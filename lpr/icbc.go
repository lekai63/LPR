package lpr

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

func icbc(model BankRepayPlanCalcModel) {
	m, err := NewCalcModel(model)
	if err != nil {
		panic(err)
	}
	m.AddAccruedPrincipalSeries(context.TODO()).addIcbcInsPlanDate()

	// fmt.Print(df.Table())

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

	fmt.Print(brps.Table())

	return nil

}

func (model *BankRepayPlanCalcModel) slice2maps(fieldname string, vals ...interface{}) []interface{} {
	brps := model.Brps
	names := brps.Names()

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
			a[fieldname] = val
			maps = append(maps, a)
		}
	}
	fmt.Printf("maps:\n %+v", maps)
	return maps

}

// genIcbcInsPayDate 生成利息还款计划，默认每月21日扣息.
// 若起始日为21日以后（含21日当日），则生成的第一个扣息日为下月21日
// 若起始日为21日以前，则生成的第一个扣息日为本月21日
// 最后一期利随本清,额外生成一个利息还款计划
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
	// 生成最后一期利随本清还息日期
	dates = append(dates, end)

	return
}
