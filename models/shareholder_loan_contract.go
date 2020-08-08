package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
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


Table: shareholder_loan_contract
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] creditor                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] abbreviation                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] loan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] loan_rate                                      INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] loan_contract_no                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] loan_start_date                                DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 7] loan_end_date                                  DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 8] all_repaid_principal                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] all_repaid_interest                            INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] is_finished                                    BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[11] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[12] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "loan_rate": 74,    "loan_contract_no": "FapEgsCAxOJqSSOoSbVEkRkbQ",    "loan_start_date": "2092-01-08T05:22:46.577368818Z",    "is_finished": true,    "abbreviation": "CGciHHgunjEEahQTQoCuqjsKu",    "creditor": "XwjsrCMTjUnOnxxGjufFYIJUi",    "loan_principal": 29,    "loan_end_date": "2145-11-09T06:45:39.167096431Z",    "all_repaid_principal": 61,    "all_repaid_interest": 90,    "created_at": "2191-01-07T16:50:27.147232762Z",    "updated_at": "2292-06-16T15:34:16.377053985Z",    "id": 67}



*/

// ShareholderLoanContract struct is a row record of the shareholder_loan_contract table in the fzzl database
type ShareholderLoanContract struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] creditor                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Creditor string `gorm:"column:creditor;type:VARCHAR;size:255;" json:"creditor"`
	//[ 2] abbreviation                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Abbreviation null.String `gorm:"column:abbreviation;type:VARCHAR;size:255;" json:"abbreviation"`
	//[ 3] loan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	LoanPrincipal int64 `gorm:"column:loan_principal;type:INT8;" json:"loan_principal"`
	//[ 4] loan_rate                                      INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LoanRate int32 `gorm:"column:loan_rate;type:INT4;" json:"loan_rate"`
	//[ 5] loan_contract_no                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoanContractNo null.String `gorm:"column:loan_contract_no;type:VARCHAR;size:255;" json:"loan_contract_no"`
	//[ 6] loan_start_date                                DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	LoanStartDate null.Time `gorm:"column:loan_start_date;type:DATE;" json:"loan_start_date"`
	//[ 7] loan_end_date                                  DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	LoanEndDate null.Time `gorm:"column:loan_end_date;type:DATE;" json:"loan_end_date"`
	//[ 8] all_repaid_principal                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidPrincipal null.Int `gorm:"column:all_repaid_principal;type:INT8;" json:"all_repaid_principal"`
	//[ 9] all_repaid_interest                            INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidInterest null.Int `gorm:"column:all_repaid_interest;type:INT8;" json:"all_repaid_interest"`
	//[10] is_finished                                    BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsFinished null.Int `gorm:"column:is_finished;type:BOOL;" json:"is_finished"`
	//[11] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//[12] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
}

var shareholder_loan_contractTableInfo = &TableInfo{
	Name: "shareholder_loan_contract",
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
			Name:               "creditor",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Creditor",
			GoFieldType:        "string",
			JSONFieldName:      "creditor",
			ProtobufFieldName:  "creditor",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "abbreviation",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Abbreviation",
			GoFieldType:        "null.String",
			JSONFieldName:      "abbreviation",
			ProtobufFieldName:  "abbreviation",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "loan_principal",
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
			GoFieldName:        "LoanPrincipal",
			GoFieldType:        "int64",
			JSONFieldName:      "loan_principal",
			ProtobufFieldName:  "loan_principal",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "loan_rate",
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
			GoFieldName:        "LoanRate",
			GoFieldType:        "int32",
			JSONFieldName:      "loan_rate",
			ProtobufFieldName:  "loan_rate",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "loan_contract_no",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "LoanContractNo",
			GoFieldType:        "null.String",
			JSONFieldName:      "loan_contract_no",
			ProtobufFieldName:  "loan_contract_no",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "loan_start_date",
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
			GoFieldName:        "LoanStartDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "loan_start_date",
			ProtobufFieldName:  "loan_start_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "loan_end_date",
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
			GoFieldName:        "LoanEndDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "loan_end_date",
			ProtobufFieldName:  "loan_end_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "all_repaid_principal",
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
			GoFieldName:        "AllRepaidPrincipal",
			GoFieldType:        "null.Int",
			JSONFieldName:      "all_repaid_principal",
			ProtobufFieldName:  "all_repaid_principal",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "all_repaid_interest",
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
			GoFieldName:        "AllRepaidInterest",
			GoFieldType:        "null.Int",
			JSONFieldName:      "all_repaid_interest",
			ProtobufFieldName:  "all_repaid_interest",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "is_finished",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "IsFinished",
			GoFieldType:        "null.Int",
			JSONFieldName:      "is_finished",
			ProtobufFieldName:  "is_finished",
			ProtobufType:       "bool",
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
func (s *ShareholderLoanContract) TableName() string {
	return "shareholder_loan_contract"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShareholderLoanContract) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShareholderLoanContract) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShareholderLoanContract) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShareholderLoanContract) TableInfo() *TableInfo {
	return shareholder_loan_contractTableInfo
}
