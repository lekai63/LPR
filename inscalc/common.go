package inscalc

import (
	"fmt"
	"time"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

// slice2maps 传入两个参数，生成map[fieldname]val . 其中 "bank_loan_contract_id" 字段固定为 model.Bc.ID; fieldname 字段值为val
// 该map用于添加至 model.Brps 的dataframe中
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

// getLatestNilActualRowNum 返回第一笔实际未付的记录序号，如全部已付，则返回-1
func getLatestNilActualRowNum(df *dataframe.DataFrame) (int, error) {
	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true})

	df.Lock()
	defer df.Unlock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		if vals["actual_amount"] == nil {
			return *row, nil
		}
	}
	return -1, fmt.Errorf("无未还款记录，请检查合同是否已结束")

}

// genMonthlyInsPlanDate 生成利息还款计划，默认每月21日扣息.
// 若起始日为21日以后（含21日当日），则生成的第一个扣息日为下月21日
// 若起始日为21日以前，则生成的第一个扣息日为本月21日
// 最后一期利随本清,不额外生成一个利息还款计划
func genMonthlyInsPlanDate(start, end civil.Date) (dates []interface{}) {
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

// genSeasonlyInsPlanDate 生成利息还款计划，默认每季度（3、6、9、12月）21日扣息.
// 第一个扣息日为距离start_date 最近的一个季度末月的21日
// 最后一期利随本清,不额外生成一个利息还款计划
func genSeasonlyInsPlanDate(start, end civil.Date) (dates []interface{}) {
	//d1 为第一个扣息日
	var d1 civil.Date
	if start.Month == time.December && start.Day >= 21 {
		// start 日期处于12月21日至12月31日时
		d1 = civil.Date{
			Year:  start.Year + 1,
			Month: time.March,
			Day:   21,
		}
	} else {
		for i := time.March; i <= time.December; i = i + 3 {
			m := civil.Date{
				Year:  start.Year,
				Month: i,
				Day:   21,
			}
			if start.Before(m) {
				d1 = m
				break
			}
		}

	}

	termMonths := 12*(end.Year-start.Year) + (int(end.Month) - int(start.Month))
	for i := 0; i < termMonths; i = i + 3 {
		mInt := (int(d1.Month) + i) % 12
		var m time.Month
		var d civil.Date
		if mInt == 0 {
			d = civil.Date{
				Year:  d1.Year + (i+int(d1.Month))/12 - 1,
				Month: time.December,
				Day:   21,
			}
		} else {
			m = time.Month(mInt)
			d = civil.Date{
				Year:  d1.Year + (i+int(d1.Month))/12,
				Month: m,
				Day:   21,
			}
		}

		if d.Before(end) {
			dates = append(dates, d)
		}
	}
	// 不额外生成最后一期利随本清还息日期
	// dates = append(dates, end)

	return
}
