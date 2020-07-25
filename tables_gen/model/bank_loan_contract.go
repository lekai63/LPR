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


Table: bank_loan_contract
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] lease_contract_id                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] bank_contract_no                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] bank_contract_name                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] bank_account                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] interest_calc_method                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] bank_name                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] bank_branch                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] loan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] loan_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] contract_start_date                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[11] contract_end_date                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[12] actual_start_date                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[13] is_lpr                                         BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[14] current_reprice_day                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[15] current_lpr                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[16] lpr_plus                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[17] current_rate                                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[18] next_reprice_day                               DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[19] all_repaid_principal                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[20] all_repaid_interest                            INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[21] is_finished                                    BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[22] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[23] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[24] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[25] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "updated_at": "2256-03-05T04:48:12.728069936Z",    "bank_branch": "nSpcnebeejEROBXQPGkGRudDW",    "loan_method": "osceXdoUpXtMDDvoZbMnBRPjv",    "contract_end_date": "2160-05-28T15:06:14.640345944Z",    "bank_account": "GULpfKwHPBPCOGGhFpXfiEJPg",    "contract_start_date": "2158-10-16T01:57:11.075148375Z",    "actual_start_date": "2248-01-16T13:01:58.291957537Z",    "is_lpr": true,    "current_reprice_day": "2070-06-08T03:29:31.098285897Z",    "current_lpr": 43,    "bank_contract_no": "NjvsSUtingPieqcliYdCvjVgH",    "bank_contract_name": "AsDYDIJHPlSFJERjcrGgncbin",    "interest_calc_method": "SDexZKvrRawdhYmTKBGkuIATP",    "contact_tel": "jQoPgKUglteidwSNYThwGOLwk",    "current_rate": 25,    "all_repaid_interest": 98,    "is_finished": true,    "loan_principal": 56,    "lpr_plus": 51,    "next_reprice_day": "2124-10-05T13:36:53.442867265Z",    "all_repaid_principal": 14,    "contact_person": "ekqcTkbYNOuDaTGUoxueCYGZD",    "id": 15,    "lease_contract_id": 35,    "bank_name": "CTgYtxVhiQCWowmsXncKKAGOw",    "created_at": "2098-08-23T10:17:04.730191942Z"}



*/

// BankLoanContract struct is a row record of the bank_loan_contract table in the fzzl database
type BankLoanContract struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] lease_contract_id                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LeaseContractID null.Int `gorm:"column:lease_contract_id;type:INT4;" json:"lease_contract_id"`
	//[ 2] bank_contract_no                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankContractNo null.String `gorm:"column:bank_contract_no;type:VARCHAR;size:255;" json:"bank_contract_no"`
	//[ 3] bank_contract_name                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankContractName null.String `gorm:"column:bank_contract_name;type:VARCHAR;size:255;" json:"bank_contract_name"`
	//[ 4] bank_account                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankAccount null.String `gorm:"column:bank_account;type:VARCHAR;size:255;" json:"bank_account"`
	//[ 5] interest_calc_method                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	InterestCalcMethod null.String `gorm:"column:interest_calc_method;type:VARCHAR;size:255;" json:"interest_calc_method"`
	//[ 6] bank_name                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankName string `gorm:"column:bank_name;type:VARCHAR;size:255;" json:"bank_name"`
	//[ 7] bank_branch                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankBranch null.String `gorm:"column:bank_branch;type:VARCHAR;size:255;" json:"bank_branch"`
	//[ 8] loan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	LoanPrincipal int64 `gorm:"column:loan_principal;type:INT8;" json:"loan_principal"`
	//[ 9] loan_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoanMethod null.String `gorm:"column:loan_method;type:VARCHAR;size:255;" json:"loan_method"`
	//[10] contract_start_date                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ContractStartDate null.Time `gorm:"column:contract_start_date;type:DATE;" json:"contract_start_date"`
	//[11] contract_end_date                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ContractEndDate null.Time `gorm:"column:contract_end_date;type:DATE;" json:"contract_end_date"`
	//[12] actual_start_date                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ActualStartDate null.Time `gorm:"column:actual_start_date;type:DATE;" json:"actual_start_date"`
	//[13] is_lpr                                         BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsLpr null.Int `gorm:"column:is_lpr;type:BOOL;" json:"is_lpr"`
	//[14] current_reprice_day                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	CurrentRepriceDay null.Time `gorm:"column:current_reprice_day;type:DATE;" json:"current_reprice_day"`
	//[15] current_lpr                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentLpr null.Int `gorm:"column:current_lpr;type:INT4;" json:"current_lpr"`
	//[16] lpr_plus                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LprPlus null.Int `gorm:"column:lpr_plus;type:INT4;" json:"lpr_plus"`
	//[17] current_rate                                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentRate int32 `gorm:"column:current_rate;type:INT4;" json:"current_rate"`
	//[18] next_reprice_day                               DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	NextRepriceDay null.Time `gorm:"column:next_reprice_day;type:DATE;" json:"next_reprice_day"`
	//[19] all_repaid_principal                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidPrincipal null.Int `gorm:"column:all_repaid_principal;type:INT8;" json:"all_repaid_principal"`
	//[20] all_repaid_interest                            INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidInterest null.Int `gorm:"column:all_repaid_interest;type:INT8;" json:"all_repaid_interest"`
	//[21] is_finished                                    BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsFinished null.Int `gorm:"column:is_finished;type:BOOL;" json:"is_finished"`
	//[22] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ContactPerson null.String `gorm:"column:contact_person;type:VARCHAR;size:255;" json:"contact_person"`
	//[23] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	ContactTel null.String `gorm:"column:contact_tel;type:VARCHAR;size:50;" json:"contact_tel"`
	//[24] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//[25] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
}

var bank_loan_contractTableInfo = &TableInfo{
	Name: "bank_loan_contract",
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
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "LeaseContractID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "lease_contract_id",
			ProtobufFieldName:  "lease_contract_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "bank_contract_no",
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
			GoFieldName:        "BankContractNo",
			GoFieldType:        "null.String",
			JSONFieldName:      "bank_contract_no",
			ProtobufFieldName:  "bank_contract_no",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "bank_contract_name",
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
			GoFieldName:        "BankContractName",
			GoFieldType:        "null.String",
			JSONFieldName:      "bank_contract_name",
			ProtobufFieldName:  "bank_contract_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "bank_account",
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
			GoFieldName:        "BankAccount",
			GoFieldType:        "null.String",
			JSONFieldName:      "bank_account",
			ProtobufFieldName:  "bank_account",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "interest_calc_method",
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
			GoFieldName:        "InterestCalcMethod",
			GoFieldType:        "null.String",
			JSONFieldName:      "interest_calc_method",
			ProtobufFieldName:  "interest_calc_method",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "bank_name",
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
			GoFieldName:        "BankName",
			GoFieldType:        "string",
			JSONFieldName:      "bank_name",
			ProtobufFieldName:  "bank_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "bank_branch",
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
			GoFieldName:        "BankBranch",
			GoFieldType:        "null.String",
			JSONFieldName:      "bank_branch",
			ProtobufFieldName:  "bank_branch",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "loan_method",
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
			GoFieldName:        "LoanMethod",
			GoFieldType:        "null.String",
			JSONFieldName:      "loan_method",
			ProtobufFieldName:  "loan_method",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "contract_start_date",
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
			GoFieldName:        "ContractStartDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "contract_start_date",
			ProtobufFieldName:  "contract_start_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "contract_end_date",
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
			GoFieldName:        "ContractEndDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "contract_end_date",
			ProtobufFieldName:  "contract_end_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "actual_start_date",
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
			GoFieldName:        "ActualStartDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "actual_start_date",
			ProtobufFieldName:  "actual_start_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "is_lpr",
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
			GoFieldName:        "IsLpr",
			GoFieldType:        "null.Int",
			JSONFieldName:      "is_lpr",
			ProtobufFieldName:  "is_lpr",
			ProtobufType:       "bool",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "current_reprice_day",
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
			GoFieldName:        "CurrentRepriceDay",
			GoFieldType:        "null.Time",
			JSONFieldName:      "current_reprice_day",
			ProtobufFieldName:  "current_reprice_day",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "current_lpr",
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
			GoFieldName:        "CurrentLpr",
			GoFieldType:        "null.Int",
			JSONFieldName:      "current_lpr",
			ProtobufFieldName:  "current_lpr",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "lpr_plus",
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
			GoFieldName:        "LprPlus",
			GoFieldType:        "null.Int",
			JSONFieldName:      "lpr_plus",
			ProtobufFieldName:  "lpr_plus",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "current_rate",
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
			GoFieldName:        "CurrentRate",
			GoFieldType:        "int32",
			JSONFieldName:      "current_rate",
			ProtobufFieldName:  "current_rate",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "next_reprice_day",
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
			GoFieldName:        "NextRepriceDay",
			GoFieldType:        "null.Time",
			JSONFieldName:      "next_reprice_day",
			ProtobufFieldName:  "next_reprice_day",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "contact_person",
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
			GoFieldName:        "ContactPerson",
			GoFieldType:        "null.String",
			JSONFieldName:      "contact_person",
			ProtobufFieldName:  "contact_person",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "contact_tel",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       50,
			GoFieldName:        "ContactTel",
			GoFieldType:        "null.String",
			JSONFieldName:      "contact_tel",
			ProtobufFieldName:  "contact_tel",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
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
			ProtobufPos:        26,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BankLoanContract) TableName() string {
	return "bank_loan_contract"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BankLoanContract) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BankLoanContract) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BankLoanContract) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BankLoanContract) TableInfo() *TableInfo {
	return bank_loan_contractTableInfo
}
