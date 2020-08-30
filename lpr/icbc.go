package lpr

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
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
	fmt.Print("se:\n", se)
	startDate := se.Value(0).(civil.Date)
	endDate := se.Value(se.NRows() - 1).(civil.Date)
	dates := genIcbcInsPayDate(startDate, endDate)

	fmt.Printf("\n dates:\n %s", dates)

	return nil

}

// genIcbcInsPayDate 生成利息还款计划，默认每月21日扣息.
// 若起始日为21日以后（含21日当日），则生成的第一个扣息日为下月21日
// 若起始日为21日以前，则生成的第一个扣息日为本月21日
// 最后一期利随本清,额外生成一个利息还款计划
func genIcbcInsPayDate(start, end civil.Date) (dates []civil.Date) {
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

/* func genIcbcInsPayDate(start, end civil.Date) (dates []civil.Date) {
	// year,month 初始化
	year := start.Year
	month := start.Month
	if start.Day >= 21 {
		if month == time.December {
			year = year + 1
			month = time.January
		} else {
			month = month + 1
		}
	}

	for y := year; y <= end.Year; y++ {
		for m := time.January; m <= 12; m++ {
			if y == start.Year {
				m = month
			}
			// 最后一期本金还款日小于21日时，不生成该月21日的利息还款计划
			if (y == end.Year) && (m == end.Month) && (end.Day < 21) {
				break
			}
			d := civil.Date{Year: y, Month: m, Day: 21}
			dates = append(dates, d)
		}
	}
	// 最后一期利随本清，单独生成一条利息还款计划
	dlast := civil.Date{Year: end.Year, Month: end.Month, Day: end.Day}
	dates = append(dates, dlast)
	return

} */
