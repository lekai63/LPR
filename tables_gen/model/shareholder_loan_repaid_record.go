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


Table: shareholder_loan_repaid_record
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] repaid_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 2] repaid_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 3] repaid_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] repaid_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] shareholder_loan_contract_id                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 7] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "created_at": "2260-09-27T05:42:50.384349782+08:00",    "updated_at": "2025-02-10T10:24:54.026179139+08:00",    "id": 83,    "repaid_date": "2059-07-14T11:05:52.232246349+08:00",    "repaid_amount": 91,    "repaid_principal": 43,    "repaid_interest": 52,    "shareholder_loan_contract_id": 15}



*/

// ShareholderLoanRepaidRecord struct is a row record of the shareholder_loan_repaid_record table in the fzzl database
type ShareholderLoanRepaidRecord struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] repaid_date                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	RepaidDate null.Time `gorm:"column:repaid_date;type:DATE;" json:"repaid_date"`
	//[ 2] repaid_amount                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	RepaidAmount null.Int `gorm:"column:repaid_amount;type:INT8;" json:"repaid_amount"`
	//[ 3] repaid_principal                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	RepaidPrincipal null.Int `gorm:"column:repaid_principal;type:INT8;" json:"repaid_principal"`
	//[ 4] repaid_interest                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	RepaidInterest null.Int `gorm:"column:repaid_interest;type:INT8;" json:"repaid_interest"`
	//[ 5] shareholder_loan_contract_id                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	ShareholderLoanContractID null.Int `gorm:"column:shareholder_loan_contract_id;type:INT4;" json:"shareholder_loan_contract_id"`
	//[ 6] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//[ 7] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
}

var shareholder_loan_repaid_recordTableInfo = &TableInfo{
	Name: "shareholder_loan_repaid_record",
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

		&ColumnInfo{
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

		&ColumnInfo{
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

		&ColumnInfo{
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

		&ColumnInfo{
			Index:              5,
			Name:               "shareholder_loan_contract_id",
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
			GoFieldName:        "ShareholderLoanContractID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "shareholder_loan_contract_id",
			ProtobufFieldName:  "shareholder_loan_contract_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
