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


Table: lease_repay_plan
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] lease_contract_id                              INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] period                                         INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 3] plan_date                                      DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 4] plan_amount                                    INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] plan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] plan_interest                                  INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] actual_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 8] actual_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] actual_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] actual_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[11] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[12] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 61,    "plan_interest": 37,    "updated_at": "2204-05-14T22:40:04.418355921Z",    "actual_principal": 70,    "lease_contract_id": 39,    "period": 91,    "plan_date": "2155-05-26T21:04:56.758641363Z",    "plan_amount": 63,    "plan_principal": 12,    "actual_date": "2150-12-24T22:55:11.841700994Z",    "actual_amount": 57,    "actual_interest": 61,    "created_at": "2148-05-10T08:25:28.04157684Z"}



*/

// LeaseRepayPlan struct is a row record of the lease_repay_plan table in the fzzl database
type LeaseRepayPlan struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] lease_contract_id                              INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LeaseContractID int32 `gorm:"column:lease_contract_id;type:INT4;" json:"lease_contract_id"`
	//[ 2] period                                         INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Period null.Int `gorm:"column:period;type:INT2;" json:"period"`
	//[ 3] plan_date                                      DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	PlanDate time.Time `gorm:"column:plan_date;type:DATE;" json:"plan_date"`
	//[ 4] plan_amount                                    INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanAmount int64 `gorm:"column:plan_amount;type:INT8;" json:"plan_amount"`
	//[ 5] plan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanPrincipal int64 `gorm:"column:plan_principal;type:INT8;" json:"plan_principal"`
	//[ 6] plan_interest                                  INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanInterest int64 `gorm:"column:plan_interest;type:INT8;" json:"plan_interest"`
	//[ 7] actual_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ActualDate null.Time `gorm:"column:actual_date;type:DATE;" json:"actual_date"`
	//[ 8] actual_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualAmount null.Int `gorm:"column:actual_amount;type:INT8;" json:"actual_amount"`
	//[ 9] actual_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualPrincipal null.Int `gorm:"column:actual_principal;type:INT8;" json:"actual_principal"`
	//[10] actual_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualInterest null.Int `gorm:"column:actual_interest;type:INT8;" json:"actual_interest"`
	//[11] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//[12] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
}

var lease_repay_planTableInfo = &TableInfo{
	Name: "lease_repay_plan",
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
			Name:               "lease_contract_id",
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
			GoFieldName:        "LeaseContractID",
			GoFieldType:        "int32",
			JSONFieldName:      "lease_contract_id",
			ProtobufFieldName:  "lease_contract_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "period",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "Period",
			GoFieldType:        "null.Int",
			JSONFieldName:      "period",
			ProtobufFieldName:  "period",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "plan_amount",
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
			GoFieldName:        "PlanAmount",
			GoFieldType:        "int64",
			JSONFieldName:      "plan_amount",
			ProtobufFieldName:  "plan_amount",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "plan_interest",
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
			GoFieldName:        "PlanInterest",
			GoFieldType:        "int64",
			JSONFieldName:      "plan_interest",
			ProtobufFieldName:  "plan_interest",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LeaseRepayPlan) TableName() string {
	return "lease_repay_plan"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LeaseRepayPlan) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LeaseRepayPlan) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LeaseRepayPlan) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LeaseRepayPlan) TableInfo() *TableInfo {
	return lease_repay_planTableInfo
}
