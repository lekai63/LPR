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


Table: shareholder_loan_contract
[ 0] sl_cid                                         INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
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
[11] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[12] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "abbreviation": "aNBluNpYuPSmLbwpWevVnPfPH",    "all_repaid_principal": 15,    "is_finished": false,    "create_time": "2268-04-10T16:06:53.930733287Z",    "modify_time": "2272-07-08T23:27:23.363267238Z",    "loan_start_date": "2307-06-10T08:23:36.328415109Z",    "loan_end_date": "2060-01-31T11:35:29.499468073Z",    "all_repaid_interest": 74,    "sl_cid": 84,    "creditor": "kuIuWkeBjLQLlmTNsjlEStGca",    "loan_principal": 38,    "loan_rate": 86,    "loan_contract_no": "geUFDKRqaoSVnFvdoTuuOsBOZ"}



*/

// ShareholderLoanContract struct is a row record of the shareholder_loan_contract table in the fzzl database
type ShareholderLoanContract struct {
	//[ 0] sl_cid                                         INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	SlCid int32 `gorm:"primary_key;AUTO_INCREMENT;column:sl_cid;type:INT4;" json:"sl_cid" db:"sl_cid"`
	//[ 1] creditor                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Creditor string `gorm:"column:creditor;type:VARCHAR;size:255;" json:"creditor" db:"creditor"`
	//[ 2] abbreviation                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Abbreviation null.String `gorm:"column:abbreviation;type:VARCHAR;size:255;" json:"abbreviation" db:"abbreviation"`
	//[ 3] loan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	LoanPrincipal int64 `gorm:"column:loan_principal;type:INT8;" json:"loan_principal" db:"loan_principal"`
	//[ 4] loan_rate                                      INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LoanRate int32 `gorm:"column:loan_rate;type:INT4;" json:"loan_rate" db:"loan_rate"`
	//[ 5] loan_contract_no                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoanContractNo null.String `gorm:"column:loan_contract_no;type:VARCHAR;size:255;" json:"loan_contract_no" db:"loan_contract_no"`
	//[ 6] loan_start_date                                DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	LoanStartDate null.Time `gorm:"column:loan_start_date;type:DATE;" json:"loan_start_date" db:"loan_start_date"`
	//[ 7] loan_end_date                                  DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	LoanEndDate null.Time `gorm:"column:loan_end_date;type:DATE;" json:"loan_end_date" db:"loan_end_date"`
	//[ 8] all_repaid_principal                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidPrincipal null.Int `gorm:"column:all_repaid_principal;type:INT8;" json:"all_repaid_principal" db:"all_repaid_principal"`
	//[ 9] all_repaid_interest                            INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidInterest null.Int `gorm:"column:all_repaid_interest;type:INT8;" json:"all_repaid_interest" db:"all_repaid_interest"`
	//[10] is_finished                                    BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsFinished null.Int `gorm:"column:is_finished;type:BOOL;" json:"is_finished" db:"is_finished"`
	//[11] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time" db:"create_time"`
	//[12] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	ModifyTime time.Time `gorm:"column:modify_time;type:TIMESTAMP;" json:"modify_time" db:"modify_time"`
}

var shareholder_loan_contractTableInfo = &TableInfo{
	Name: "shareholder_loan_contract",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "sl_cid",
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
			GoFieldName:        "SlCid",
			GoFieldType:        "int32",
			JSONFieldName:      "sl_cid",
			ProtobufFieldName:  "sl_cid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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
