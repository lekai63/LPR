package inscalc

import (
	// "context"
	"math/big"
	"time"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

// FillPlanDateMonthly 对还未还款的记录（actual_date 为nil），生成每月21日的还息日期plan_date，并按planDate升序排序。
// 其他字段以nil进行填充 （bank_loan_contract_id 用 bc.id填充）
func (model *BankRepayPlanCalcModel) FillPlanDateMonthly() *BankRepayPlanCalcModel {
	brps := model.Brps
	col, _ := brps.NameToColumn("plan_date")
	se := brps.Series[col]
	se.Sort(ctx, dataframe.SortOptions{Desc: false})

	n, err := getLatestNilActualRowNum(brps)
	nrow := se.NRows()
	check(err)

	startDate := se.Value(n).(civil.Date)
	endDate := se.Value(nrow - 1).(civil.Date)
	planDateSlice := genIcbcInsPlanDate(startDate, endDate)

	// 填充生成的planDate，并对其他字段进行填充
	model.Brps = brps
	maps := model.slice2maps("plan_date", planDateSlice...)

	// 组装dataframe
	// 注意maps中字段要与series一一对应，否则报错"no. of args not equal to no. of series"
	for _, val := range maps {
		brps.Append(nil, val)
	}

	brps.Sort(ctx, []dataframe.SortKey{
		{Key: "plan_date", Desc: false},
	})

	model.Brps = brps
	return model
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
// todo:改造为通用计算
func icbcOneInsCalc(vals map[interface{}]interface{}, upperVals map[interface{}]interface{}, model *BankRepayPlanCalcModel) int64 {
	calcDays := vals["plan_date"].(civil.Date).DaysSince(upperVals["plan_date"].(civil.Date))
	planInsB := big.NewInt(0)
	accruedB := big.NewInt(vals["accrued_principal"].(int64))
	RateB := big.NewInt(int64(model.Bc.CurrentRate))
	// 计划利息 = 应计本金×年利率×期间天数/360 （因利率单位为0.01%，所以再除以10000）
	planInsB = planInsB.Mul(accruedB, RateB).Mul(planInsB, big.NewInt(int64(calcDays))).Div(planInsB, big.NewInt(3600000))
	return planInsB.Int64()
}
