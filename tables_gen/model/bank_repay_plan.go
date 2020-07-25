package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: bank_repay_plan
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] bank_loan_contract_id                          INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] plan_date                                      DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 3] plan_amount                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] plan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] plan_interest                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] actual_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 7] actual_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 8] actual_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] actual_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[11] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "plan_amount": 91,    "plan_interest": 66,    "actual_amount": 26,    "actual_principal": 26,    "updated_at": "2300-01-18T08:02:37.147327156Z",    "id": 67,    "bank_loan_contract_id": 16,    "plan_date": "2159-08-03T14:14:14.922878857Z",    "plan_principal": 90,    "actual_date": "2084-06-07T08:57:10.753963539Z",    "actual_interest": 35,    "created_at": "2219-01-24T11:56:18.51500847Z"}



*/

// BankRepayPlan struct is a row record of the bank_repay_plan table in the fzzl database
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
	//[10] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//[11] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
}

var bank_repay_planTableInfo = &TableInfo{
	Name: "bank_repay_plan",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "bank_loan_contract_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "BankLoanContractID",
			GoFieldType:        "int32",
			JSONFieldName:      "bank_loan_contract_id",
			ProtobufFieldName:  "bank_loan_contract_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "plan_date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "DATE",
			DatabaseTypePretty: "DATE",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATE",
			ColumnLength:       -1,
			GoFieldName:        "PlanDate",
			GoFieldType:        "time.Time",
			JSONFieldName:      "plan_date",
			ProtobufFieldName:  "plan_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "plan_amount",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "PlanAmount",
			GoFieldType:        "null.Int",
			JSONFieldName:      "plan_amount",
			ProtobufFieldName:  "plan_amount",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "plan_principal",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "PlanPrincipal",
			GoFieldType:        "int64",
			JSONFieldName:      "plan_principal",
			ProtobufFieldName:  "plan_principal",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "plan_interest",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "PlanInterest",
			GoFieldType:        "null.Int",
			JSONFieldName:      "plan_interest",
			ProtobufFieldName:  "plan_interest",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "actual_date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "DATE",
			DatabaseTypePretty: "DATE",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATE",
			ColumnLength:       -1,
			GoFieldName:        "ActualDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "actual_date",
			ProtobufFieldName:  "actual_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "actual_amount",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ActualAmount",
			GoFieldType:        "null.Int",
			JSONFieldName:      "actual_amount",
			ProtobufFieldName:  "actual_amount",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "actual_principal",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ActualPrincipal",
			GoFieldType:        "null.Int",
			JSONFieldName:      "actual_principal",
			ProtobufFieldName:  "actual_principal",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "actual_interest",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ActualInterest",
			GoFieldType:        "null.Int",
			JSONFieldName:      "actual_interest",
			ProtobufFieldName:  "actual_interest",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BankRepayPlan) TableName() string {
	return "bank_repay_plan"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BankRepayPlan) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BankRepayPlan) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BankRepayPlan) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BankRepayPlan) TableInfo() *TableInfo {
	return bank_repay_planTableInfo
}
