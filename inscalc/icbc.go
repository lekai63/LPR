package inscalc

import (
	"errors"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

// ToICBC 根据model 生成工行还本付息计划
// 首次生成，则isFirst 为true；
func (model *BankRepayPlanCalcModel) ToICBC(isFirst bool) (*BankRepayPlanCalcModel, error) {
	if model.Bc.BankName != "工商银行" {
		return nil, errors.New("输入模型的银行名称不是工商银行，请检查")
	}

	// 注意使用括号决定计算优先级，不要直接链式调用
	if isFirst {
		model.FillInsPlanDate()
	}
	model.AddAccruedPrincipal()
	model.AddIcbcFactoringIns()
	model.AddPlanAmount()
	return model, nil
}

// AddIcbcFactoringIns 计算工行保理利息并添加到列，本函数将df.lock 注意避免与其他函数形成死锁
func (model *BankRepayPlanCalcModel) AddIcbcFactoringIns() *BankRepayPlanCalcModel {
	model.Sort("plan_date")
	df := model.Brps
	// upperVal 用于存储上一个row的信息
	var upperVals map[interface{}]interface{}
	var temp int64
	// planInsterest Slice用于存储计算后的计划利息,因最后要将 planInsterest 传入NewSeries，故类型直接选择[]interface{} ,如定义为[]int,还需再循环转换为[]interface{}
	// https://golang.org/doc/faq#convert_slice_of_interface
	nrows := df.NRows()
	planInsterest := make([]interface{}, nrows)
	// fmt.Println("完成planIns初始化")

	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	df.Lock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		switch {
		// 第0行为最后一笔实际还款记录，计划利息已算好，无需重新测算
		case (*row) == 0:
			if vals["plan_interest"] != nil {
				planInsterest[0] = vals["plan_interest"].(int64)
			} else {
				planInsterest[0] = 0
			}
			// 最后一行利随本清
		case (*row) == nrows-1: // 如使用(*row) == df.NRows() 游标直接到最后，从而无法执行
			planInsterest[*row] = model.rowInsCalc(vals, upperVals)
		default:
			// 因工行保理利息在每月21日扣，故本金还款日当天的利息，应加到下一个最近的21日一并扣息
			// 本row还款日非21日，将计划还款利息暂存入temp
			if d := vals["plan_date"].(civil.Date); d.Day != 21 {
				temp = model.rowInsCalc(vals, upperVals)
				planInsterest[*row] = 0
			} else if x := upperVals["plan_date"].(civil.Date); x.Day != 21 {
				// 上一row为非21日，将temp提取出来，加入本row
				planInsterest[*row] = model.rowInsCalc(vals, upperVals) + temp
			} else {
				// 默认planInsterest算法
				planInsterest[*row] = model.rowInsCalc(vals, upperVals)
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
