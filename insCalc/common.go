package insCalc

import (
	"fmt"

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
