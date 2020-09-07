package insCalc

import (
	"context"
	"fmt"

	"github.com/guregu/null"
	dataframe "github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BankRepayPlanCalcModel  定义单个合同的计算模型
type BankRepayPlanCalcModel struct {
	Bc BankLoanContractMini
	// Brps存储Bc合同项下的所有还款记录, Brp []BankRepayPlan
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

// NewModel 根据bankLoanContractID 从数据库中获取数据 并生成 BankRepayPlanCalcModel
func NewBankRepayPlanCalcModel(bankLoanContractID int32) (model BankRepayPlanCalcModel, err error) {
	// conn := models.GlobalConn
	// to change back to models.GlobalConn
	db, _ := GormInitForTest()

	ctx := context.Background()
	// gen model.Bc
	var bc BankLoanContractMini
	bc.ID = bankLoanContractID
	db.Table("bank_loan_contract").First(&bc)
	// db.First(&bc)
	model.Bc = bc

	// gen model.Brps
	sqldb, _ := db.DB()
	tx, _ := sqldb.Begin()
	op := imports.SQLLoadOptions{
		// KnownRowCount: &[]int{11}[0],
		DictateDataType: map[string]interface{}{
			"id":                    int64(0),
			"bank_loan_contract_id": int64(0),
			// "plan_date":             time.Now,
			"plan_amount":    int64(0),
			"plan_principal": int64(0),
			"plan_interest":  int64(0),
			//	"actual_date":           time.Now,
			"actual_amount":    int64(0),
			"actual_principal": int64(0),
			"actual_interest":  int64(0),
			// "created_at":        time.Now,
			// "updated_at":        time.Now,
			"accrued_principal": int64(0),
		},
		Database: imports.PostgreSQL,
		Query:    `select * from "bank_repay_plan"`,
	}
	brps, err := imports.LoadFromSQL(ctx, tx, &op)
	if err != nil {
		fmt.Printf("从数据库中读取数据组装dataframe时发生错误：%+v", err)
		return model, err
	}
	// fmt.Println("brps:\n")
	fmt.Print(brps.Table())
	model.Brps = brps
	return model, nil

}

func GormInitForTest() (*gorm.DB, error) {
	dsn := "host=192.168.5.11 user=fzzl password=fzzl032003 dbname=lpr port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	gormv2, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//	TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		return nil, err
	}
	return gormv2, err

}
