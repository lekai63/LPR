package models

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
[ 0] pid                                            INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] cid                                            INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] period                                         INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 3] plan_date                                      DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 4] plan_amount                                    INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] plan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] plan_interest                                  INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] actual_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 8] actual_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] actual_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] actual_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[11] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[12] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "modify_time": "2295-05-03T09:51:57.374743635Z",    "period": 73,    "plan_date": "2057-09-26T10:02:26.68213972Z",    "plan_principal": 95,    "plan_interest": 1,    "actual_amount": 24,    "actual_principal": 47,    "pid": 94,    "cid": 60,    "plan_amount": 75,    "actual_date": "2234-01-19T20:54:37.367892328Z",    "actual_interest": 89,    "create_time": "2021-06-14T15:27:41.02052101Z"}



*/

// LeaseRepayPlan struct is a row record of the lease_repay_plan table in the fzzl database
type LeaseRepayPlan struct {
	//[ 0] pid                                            INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	Pid int32 `gorm:"primary_key;AUTO_INCREMENT;column:pid;type:INT4;" json:"pid" db:"pid"`
	//[ 1] cid                                            INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Cid int32 `gorm:"column:cid;type:INT4;" json:"cid" db:"cid"`
	//[ 2] period                                         INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Period null.Int `gorm:"column:period;type:INT2;" json:"period" db:"period"`
	//[ 3] plan_date                                      DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	PlanDate time.Time `gorm:"column:plan_date;type:DATE;" json:"plan_date" db:"plan_date"`
	//[ 4] plan_amount                                    INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanAmount int64 `gorm:"column:plan_amount;type:INT8;" json:"plan_amount" db:"plan_amount"`
	//[ 5] plan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanPrincipal int64 `gorm:"column:plan_principal;type:INT8;" json:"plan_principal" db:"plan_principal"`
	//[ 6] plan_interest                                  INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PlanInterest int64 `gorm:"column:plan_interest;type:INT8;" json:"plan_interest" db:"plan_interest"`
	//[ 7] actual_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ActualDate null.Time `gorm:"column:actual_date;type:DATE;" json:"actual_date" db:"actual_date"`
	//[ 8] actual_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualAmount null.Int `gorm:"column:actual_amount;type:INT8;" json:"actual_amount" db:"actual_amount"`
	//[ 9] actual_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualPrincipal null.Int `gorm:"column:actual_principal;type:INT8;" json:"actual_principal" db:"actual_principal"`
	//[10] actual_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualInterest null.Int `gorm:"column:actual_interest;type:INT8;" json:"actual_interest" db:"actual_interest"`
	//[11] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time" db:"create_time"`
	//[12] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	ModifyTime time.Time `gorm:"column:modify_time;type:TIMESTAMP;" json:"modify_time" db:"modify_time"`
}

var lease_repay_planTableInfo = &TableInfo{
	Name: "lease_repay_plan",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "pid",
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
			GoFieldName:        "Pid",
			GoFieldType:        "int32",
			JSONFieldName:      "pid",
			ProtobufFieldName:  "pid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "cid",
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
			GoFieldName:        "Cid",
			GoFieldType:        "int32",
			JSONFieldName:      "cid",
			ProtobufFieldName:  "cid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
			Index:              11,
			Name:               "create_time",
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
			GoFieldName:        "CreateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "uint64",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "modify_time",
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
			GoFieldName:        "ModifyTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "modify_time",
			ProtobufFieldName:  "modify_time",
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
