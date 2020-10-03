package inscalc

import (
	"errors"
	"math/big"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

func (model *BankRepayPlanCalcModel) ToHZBank(isFirst bool) (*BankRepayPlanCalcModel, error) {
	if model.Bc.BankName != "杭州银行" {
		return nil, errors.New("输入模型的银行名称不是杭州银行，请检查")
	}
	if isFirst {
		model.FillInsPlanDate()
	}

	model.AddAccruedPrincipal()
	model.AddHZBankFactoringIns()
	model.AddPlanAmount()
	return model, nil

}

// AddHZBankFactoringIns 计算杭州银行保理利息并添加到列，本函数将df.lock 注意避免与其他函数形成死锁
// 杭州银行保理利息方案为：利息在每季度21日偿还，本金在其他日期偿还。
// 本金偿还时需配套付息，付息金额=本次偿还本金*（偿还日-最近一期季度付息日）/30 * 月利率
// 注意杭州银行采用月利率，即年利率除以12；日利率=月利率/30
func (model *BankRepayPlanCalcModel) AddHZBankFactoringIns() *BankRepayPlanCalcModel {
	model.Sort("plan_date")
	df := model.Brps
	// upperVal 用于存储上一个row的信息
	upperVals := make(map[interface{}]interface{})
	var temp int64
	nrows := df.NRows()
	// planInsterest Slice用于存储计算后的计划利息,因最后要将 planInsterest 传入NewSeries，故类型直接选择[]interface{} ,如定义为[]int,还需再循环转换为[]interface{}
	// https://golang.org/doc/faq#convert_slice_of_interface
	planInsterest := make([]interface{}, nrows)
	var e error
	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	df.Lock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		switch {
		// 第0行为第一笔利息还款计划
		case (*row) == 0:
			upperVals["plan_date"] = civil.DateOf(model.Bc.ActualStartDate.ValueOrZero())
			planInsterest[*row], e = model.rowInsCalc(vals, upperVals, "monthly")
			check(e)
			// 最后一行利随本清
		case (*row) == nrows-1: // 如使用(*row) == df.NRows() 游标直接到最后，从而无法执行
			planInsterest[*row], e = model.rowInsCalc(vals, upperVals, "monthly")
			check(e)
		default:
			// 杭州银行保理利息在每季末21日扣
			// 本金偿还时需配套付息，付息金额=本次偿还本金*（偿还日-最近一期季度付息日）/30 * 月利率
			// 其余本金应加到下一季末的21日一并扣息
			// 本row还款日非21日，将计划还款利息暂存入temp
			if d := vals["plan_date"].(civil.Date); d.Day != 21 {
				// 与默认保理利息计算方式不同的地方，本金还款时还会有本笔本金对应的利息
				rowins := model.hzRowInsCalc(vals, upperVals)
				planInsterest[*row] = rowins
				m, e := model.rowInsCalc(vals, upperVals, "monthly")
				check(e)
				temp = m - rowins
			} else if x := upperVals["plan_date"].(civil.Date); x.Day != 21 {
				// 上一row为非21日，将temp提取出来，加入本row
				m, e := model.rowInsCalc(vals, upperVals, "monthly")
				check(e)
				planInsterest[*row] = m + temp
			} else {
				// 默认planInsterest算法，即本row 上一row 均为21日
				planInsterest[*row], e = model.rowInsCalc(vals, upperVals, "monthly")
				check(e)
			}
		}

		// 存在数据竞争，如直接更新row，最后的df数据不全，故删去，转而在前面数据处理时，组装一个series，再将series整体加入到df
		// df.UpdateRow(*row, &dataframe.DontLock, vals)

		// 本次循环结束时，将本行赋值给upperVal用于下次循环
		upperVals = vals
	}
	df.Unlock()
	// 移除旧的plan_interest 以及不需要用的几个字段
	df.RemoveSeries("plan_interest")
	df.RemoveSeries("created_at")
	df.RemoveSeries("updated_at")
	// df.RemoveSeries("accrued_principal")

	//添加新的plan_interest
	se := dataframe.NewSeriesInt64("plan_interest", nil, planInsterest...)
	df.AddSeries(se, nil)

	model.Brps = df
	return model

}

func (model *BankRepayPlanCalcModel) hzRowInsCalc(vals map[interface{}]interface{}, upperVals map[interface{}]interface{}) int64 {
	upperDay := upperVals["plan_date"].(civil.Date)
	if upperVals["actual_date"] != nil {
		upperDay = (upperVals["actual_date"].(civil.Date))
	}

	rowDay := vals["plan_date"].(civil.Date)
	if vals["actual_date"] != nil {
		rowDay = (vals["actual_date"].(civil.Date))
	}

	calcDays := (rowDay).DaysSince(upperDay)

	planInsB := big.NewInt(0)
	accruedB := big.NewInt(vals["plan_principal"].(int64))                                                                        //此行与默认计息方式不同
	rateMonthB := big.NewInt(int64(model.Bc.CurrentRate * 10 / 12))                                                               //为了月利率的精度，小数点右移一位后再除以12
	planInsB = planInsB.Mul(accruedB, rateMonthB).Mul(planInsB, big.NewInt(int64(calcDays))).Div(planInsB, big.NewInt(300000000)) // 注意最后要多除以10，即把上面移动的小数点移回去
	// 四舍五入
	p := planInsB.Int64()
	return rounding(p)
}
