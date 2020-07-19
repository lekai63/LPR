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


Table: shareholder_loan_repaid_record
[ 0] rid                                            INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] repaid_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 2] repaid_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 3] repaid_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] repaid_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] sl_cid                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 7] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "repaid_date": "2181-11-09T08:47:42.658699696Z",    "repaid_amount": 30,    "repaid_principal": 82,    "repaid_interest": 76,    "sl_cid": 30,    "create_time": "2226-02-07T03:25:21.201614835Z",    "modify_time": "2162-12-27T13:45:14.266006415Z",    "rid": 82}



*/

// ShareholderLoanRepaidRecord struct is a row record of the shareholder_loan_repaid_record table in the fzzl database
type ShareholderLoanRepaidRecord struct {
	//[ 0] rid                                            INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	Rid int32 `gorm:"primary_key;AUTO_INCREMENT;column:rid;type:INT4;" json:"rid" db:"rid"`
	//[ 1] repaid_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	RepaidDate null.Time `gorm:"column:repaid_date;type:DATE;" json:"repaid_date" db:"repaid_date"`
	//[ 2] repaid_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	RepaidAmount null.Int `gorm:"column:repaid_amount;type:INT8;" json:"repaid_amount" db:"repaid_amount"`
	//[ 3] repaid_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	RepaidPrincipal null.Int `gorm:"column:repaid_principal;type:INT8;" json:"repaid_principal" db:"repaid_principal"`
	//[ 4] repaid_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	RepaidInterest null.Int `gorm:"column:repaid_interest;type:INT8;" json:"repaid_interest" db:"repaid_interest"`
	//[ 5] sl_cid                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	SlCid null.Int `gorm:"column:sl_cid;type:INT4;" json:"sl_cid" db:"sl_cid"`
	//[ 6] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time" db:"create_time"`
	//[ 7] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	ModifyTime time.Time `gorm:"column:modify_time;type:TIMESTAMP;" json:"modify_time" db:"modify_time"`
}

var shareholder_loan_repaid_recordTableInfo = &TableInfo{
	Name: "shareholder_loan_repaid_record",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "rid",
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
			GoFieldName:        "Rid",
			GoFieldType:        "int32",
			JSONFieldName:      "rid",
			ProtobufFieldName:  "rid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "repaid_date",
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
			GoFieldName:        "RepaidDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "repaid_date",
			ProtobufFieldName:  "repaid_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "repaid_amount",
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
			GoFieldName:        "RepaidAmount",
			GoFieldType:        "null.Int",
			JSONFieldName:      "repaid_amount",
			ProtobufFieldName:  "repaid_amount",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "repaid_principal",
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
			GoFieldName:        "RepaidPrincipal",
			GoFieldType:        "null.Int",
			JSONFieldName:      "repaid_principal",
			ProtobufFieldName:  "repaid_principal",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "repaid_interest",
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
			GoFieldName:        "RepaidInterest",
			GoFieldType:        "null.Int",
			JSONFieldName:      "repaid_interest",
			ProtobufFieldName:  "repaid_interest",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "sl_cid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SlCid",
			GoFieldType:        "null.Int",
			JSONFieldName:      "sl_cid",
			ProtobufFieldName:  "sl_cid",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},

		{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShareholderLoanRepaidRecord) TableName() string {
	return "shareholder_loan_repaid_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShareholderLoanRepaidRecord) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShareholderLoanRepaidRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShareholderLoanRepaidRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShareholderLoanRepaidRecord) TableInfo() *TableInfo {
	return shareholder_loan_repaid_recordTableInfo
}
