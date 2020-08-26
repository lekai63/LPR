package lpr

import (
	"time"

	"github.com/guregu/null"
	"github.com/lekai63/lpr/models"
)

// BankRepayPlanCalcModel  定义单个合同的计算模型
type BankRepayPlanCalcModel struct {
	Bc BankLoanContractMini
	// Brp []BankRepayPlan
	Brp map[string]interface{}
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

// BankRepayPlan 源于models中BankRepayPlan，但去掉了CreatedAt UpdatedAt两个无关的计算项目
type BankRepayPlan struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] bank_loan_contract_id                          INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	BankLoanContractID int32 `gorm:"column:bank_loan_contract_id;type:INT4;" json:"bank_loan_contract_id"`
	//[ 2] plan_date                                      DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	PlanDate time.Time `gorm:"column:plan_date;type:DATE;" json:"plan_date"`
	//[ 3] plan_amount                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanAmount null.Int `gorm:"column:plan_amount;type:INT8;" json:"plan_amount"`
	//[ 4] plan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanPrincipal int64 `gorm:"column:plan_principal;type:INT8;" json:"plan_principal"`
	//[ 5] plan_interest                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanInterest null.Int `gorm:"column:plan_interest;type:INT8;" json:"plan_interest"`
	//[ 6] actual_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ActualDate null.Time `gorm:"column:actual_date;type:DATE;" json:"actual_date"`
	//[ 7] actual_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualAmount null.Int `gorm:"column:actual_amount;type:INT8;" json:"actual_amount"`
	//[ 8] actual_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualPrincipal null.Int `gorm:"column:actual_principal;type:INT8;" json:"actual_principal"`
	//[ 9] actual_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualInterest null.Int `gorm:"column:actual_interest;type:INT8;" json:"actual_interest"`
}

// GetOneContractRepayPlan 获取银行剩余本金的还本计划
func GetOneContractRepayPlan(bankLoanContractID int32) (model BankRepayPlanCalcModel) {
	dbGorm := models.Gormv2
	var bc BankLoanContractMini
	bc.ID = bankLoanContractID
	dbGorm.Table("bank_loan_contract").Debug().First(&bc)
	// fmt.Printf("bc: %+v \n", bc)
	/* var brp []map[string]interface{}
	dbGorm.Table("bank_repay_plan").Debug().Where("bank_loan_contract_id = ? AND (actual_amount is null) ", bankLoanContractID).Scan(&brp) */
	// fmt.Printf("brp: %+v \n", brp)
	model.Bc = bc

	// model.Brp = brp

	return

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
