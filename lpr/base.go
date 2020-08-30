package lpr

import (
	"context"
	"database/sql"
	"strings"
	"time"

	// "fmt"

	"cloud.google.com/go/civil"
	"github.com/guregu/null"
	"github.com/lekai63/lpr/models"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

func Test() {
	model, err := GetOneContractRepayPlan(3)
	if err != nil {
		panic(err)
	}
	icbc(model)

	// a, b := model.AddRemainPrincipalSeries(context.TODO())

}

// BankRepayPlanCalcModel  定义单个合同的计算模型
type BankRepayPlanCalcModel struct {
	Bc BankLoanContractMini
	// Brp []BankRepayPlan
	Brps *dataframe.DataFrame
}

// BankLoanContractMini 提取BankLoanContract与利息计算相关的字段
type BankLoanContractMini struct {
	// ID,InterestCalcMethod ,BankName,LoanMethod ,CurrentRate对应bank_loan_contract 中的同名字段
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 4] interest_calc_method                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	InterestCalcMethod null.String `gorm:"column:interest_calc_method;type:VARCHAR;size:255;" json:"interest_calc_method"`
	//[ 5] bank_name                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankName string `gorm:"column:bank_name;type:VARCHAR;size:255;" json:"bank_name"`
	//[ 8] loan_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoanMethod null.String `gorm:"column:loan_method;type:VARCHAR;size:255;" json:"loan_method"`
	//[16] current_rate                                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentRate int32 `gorm:"column:current_rate;type:INT4;" json:"current_rate"`
}

// GetOneContractRepayPlan 获取银行剩余本金的还本计划
func GetOneContractRepayPlan(bankLoanContractID int32) (model BankRepayPlanCalcModel, err error) {
	db := models.Gormv2

	var bc BankLoanContractMini
	bc.ID = bankLoanContractID
	db.Table("bank_loan_contract").First(&bc)
	model.Bc = bc

	var brp models.BankRepayPlan
	brpDF, err := InitDataframe(brp)

	// rows, e := db.Model(&brp).Where("bank_loan_contract_id = ? and actual_amount is null", bankLoanContractID).Rows()

	rows, e := db.Model(&brp).Where("bank_loan_contract_id = ? ", bankLoanContractID).Rows()
	if e != nil {
		err = e
		return
	}
	defer rows.Close()
	maps, e := Rows2Maps(rows)
	// fmt.Printf("\n maps: \n %+v", maps)
	if e != nil {
		err = e
		return
	}

	// 组装dataframe
	// 注意maps中字段要与series一一对应，否则报错"no. of args not equal to no. of series"
	for _, val := range maps {

		brpDF.Append(nil, val)
	}

	brpDF.Sort(context.TODO(), []dataframe.SortKey{
		{Key: "plan_date", Desc: false},
	})

	// fmt.Print(brpDF.Table())

	model.Brps = brpDF

	return
}

//Rows2Maps store rows in map
//ref https://kylewbanks.com/blog/query-result-to-map-in-golang
func Rows2Maps(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, _ := rows.Columns()
	var maps []map[string]interface{}
	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := 0; i < len(columns); i++ {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})

			// date格式转换为civil date, "_date"是为了剔除掉updated_at字段
			if strings.Contains(colName, "_date") && ((*val) != nil) {

				m[colName] = civil.DateOf((*val).(time.Time))

			} else {
				m[colName] = *val
			}

		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		// fmt.Printf("%+v \n")
		maps = append(maps, m)
	}
	// fmt.Printf("maps:%+v", maps)
	return maps, nil

}

// CalcOneInterestPlan 计算未来的还息计划
/* func CalcOneInterestPlan(model BankRepayPlanCalcModel) (plan BankRepayPlan) {
	switch model.Bc.BankName {
	case "工商银行":
		plan = calcICBC(model)
	case "建设银行":
		plan = calcCCB(model)
	case "浙商银行":
		plan = calcCZB(model)
	case "招商银行":
		plan = calcCMB(model)
	case "杭州银行":
		plan = calcHZB(model)
	case "农业银行":
		plan = calcHZB(model)
	case "浦发银行":
		plan = calcSPDB(model)
	default:

	}
	return
} */
