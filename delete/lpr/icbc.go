package lpr

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

func icbc(model *BankRepayPlanCalcModel) {
	m, err := NewCalcModel(model)
	if err != nil {
		panic(err)
	}
	m.addIcbcInsPlanDate().AddAccruedPrincipalSeries(context.TODO())
	// fmt.Println("还款计划表:\n")
	// fmt.Print(m.Brps.Table())
	err = m.CalcIcbcFactoringIns().CreateInsToDB()
	if err != nil {
		fmt.Printf("写入错误:\n", err)
	}

}

// CalcIcbcFactoringIns 计算工行保理利息
func (model *BankRepayPlanCalcModel) CalcIcbcFactoringIns() *BankRepayPlanCalcModel {
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
			planInsterest[*row] = icbcOneInsCalc(vals, upperVals, model)

		default:

			// 因工行保理利息在每月21日扣，故本金还款日当天的利息，应加到下一个最近的21日一并扣息
			// 本row还款日非21日，将计划还款利息暂存入temp
			if d := vals["plan_date"].(civil.Date); d.Day != 21 {
				temp = icbcOneInsCalc(vals, upperVals, model)
				planInsterest[*row] = 0
			} else if x := upperVals["plan_date"].(civil.Date); x.Day != 21 {
				// 上一row为非21日，将temp提取出来，加入本row
				planInsterest[*row] = icbcOneInsCalc(vals, upperVals, model) + temp
			} else {
				// 默认planInsterest算法
				planInsterest[*row] = icbcOneInsCalc(vals, upperVals, model)
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
	df.RemoveSeries("accrued_principal")

	//添加新的plan_interest
	se := dataframe.NewSeriesInt64("plan_interest", nil, planInsterest...)
	df.AddSeries(se, nil)

	// 对原始df的直接迭代更新，只能更新部分数据，故每次都另外生成slice，再添加到df
	planAmount := make([]interface{}, nrows)
	applyDfFn := dataframe.ApplyDataFrameFn(func(val map[interface{}]interface{}, row, nRows int) map[interface{}]interface{} {
		planAmount[row] = CalcPlanAmount(val)
		return val
	})

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	dataframe.Apply(ctx, df, applyDfFn, dataframe.FilterOptions{InPlace: true})

	se2 := dataframe.NewSeriesInt64("plan_amount", nil, planAmount...)
	df.RemoveSeries("plan_amount")
	df.AddSeries(se2, nil)

	model.Brps = df
	return model
}

func (model *BankRepayPlanCalcModel) addIcbcInsPlanDate() *BankRepayPlanCalcModel {
	brps := model.Brps
	col, _ := brps.NameToColumn("plan_date")
	se := brps.Series[col]
	// fmt.Print("se:\n", se)
	startDate := se.Value(0).(civil.Date)
	endDate := se.Value(se.NRows() - 1).(civil.Date)
	planDateSlice := genIcbcInsPlanDate(startDate, endDate)
	maps := model.slice2maps("plan_date", planDateSlice...)
	for _, val := range maps {
		brps.Append(nil, val)
	}
	brps.Sort(context.TODO(), []dataframe.SortKey{
		{Key: "plan_date", Desc: false},
	})

	model.Brps = brps
	return model
}

func (model *BankRepayPlanCalcModel) slice2maps(fieldname string, vals ...interface{}) []interface{} {
	brps := model.Brps
	names := brps.Names()
	bcID := model.Bc.ID
	maps := make([]interface{}, 0)
	if vals == nil {
		panic("get no vals")
	} else {
		for _, val := range vals {
			// 初始化a
			a := make(map[string]interface{})
			for _, name := range names {
				a[name] = nil
			}

			a["bank_loan_contract_id"] = int64(bcID)
			a[fieldname] = val

			maps = append(maps, a)
		}
	}
	// fmt.Printf("maps:\n %+v", maps)
	return maps

}

// genIcbcInsPayDate 生成利息还款计划，默认每月21日扣息.
// 若起始日为21日以后（含21日当日），则生成的第一个扣息日为下月21日
// 若起始日为21日以前，则生成的第一个扣息日为本月21日
// 最后一期利随本清,不额外生成一个利息还款计划
func genIcbcInsPlanDate(start, end civil.Date) (dates []interface{}) {
	termMonths := 12*(end.Year-start.Year) + (int(end.Month) - int(start.Month))
	plus := 0
	if start.Day >= 21 {
		plus = 1
	}

	for i := 0; i < termMonths; i++ {
		mInt := (int(start.Month) + i + plus) % 12
		var m time.Month
		if mInt == 0 {
			m = time.December
		} else {
			m = time.Month(mInt)
		}
		d := civil.Date{
			Year:  start.Year + (i+int(start.Month))/12,
			Month: m,
			Day:   21,
		}

		if d.Before(end) {
			dates = append(dates, d)
		}
	}
	// 不额外生成最后一期利随本清还息日期
	// dates = append(dates, end)

	return
}

// 工行单个保理利息计算
func icbcOneInsCalc(vals map[interface{}]interface{}, upperVals map[interface{}]interface{}, model *BankRepayPlanCalcModel) int64 {
	calcDays := vals["plan_date"].(civil.Date).DaysSince(upperVals["plan_date"].(civil.Date))
	planInsB := big.NewInt(0)
	accruedB := big.NewInt(vals["accrued_principal"].(int64))
	RateB := big.NewInt(int64(model.Bc.CurrentRate))
	// 计划利息 = 应计本金×年利率×期间天数/360 （因利率单位为0.01%，所以再除以10000）
	planInsB = planInsB.Mul(accruedB, RateB).Mul(planInsB, big.NewInt(int64(calcDays))).Div(planInsB, big.NewInt(3600000))
	return planInsB.Int64()
}
