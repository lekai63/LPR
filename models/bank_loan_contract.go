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


Table: bank_loan_contract
[ 0] bl_cid                                         INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] cid                                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
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
[24] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[25] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "loan_method": "ZAmlexYdjVhoMFaOaryBpQGmf",    "contract_start_date": "2295-01-06T20:51:04.531030878Z",    "contract_end_date": "2117-11-29T01:32:43.470544219Z",    "lpr_plus": 85,    "current_rate": 30,    "bank_contract_name": "sgslNaPcQqJEXJxQmWEdhPkKE",    "interest_calc_method": "DRskXWCdXfPqDDHovHFlFxPBh",    "current_reprice_day": "2236-07-13T20:40:45.652126614Z",    "all_repaid_principal": 71,    "all_repaid_interest": 6,    "create_time": "2246-04-02T10:41:12.101041613Z",    "bl_cid": 41,    "cid": 9,    "bank_name": "McLpsvCqLbWDTLkIyyDCThThg",    "loan_principal": 96,    "next_reprice_day": "2100-05-12T15:06:40.858306442Z",    "is_finished": false,    "contact_person": "lHBxgSKYJocZqiAcpnMEepcPB",    "contact_tel": "CayGFjfXGSVYOHMYhbsNnISJC",    "bank_contract_no": "ekfDCyykxwxkFrQaoTjWRcHKe",    "bank_account": "xitBJcCdLvnLpChMBsgARUxRV",    "modify_time": "2140-03-17T22:13:59.422772725Z",    "is_lpr": true,    "current_lpr": 32,    "bank_branch": "nEcOiecJitNAImOkATXWpEbFb",    "actual_start_date": "2133-07-27T04:40:41.654567581Z"}



*/

// BankLoanContract struct is a row record of the bank_loan_contract table in the fzzl database
type BankLoanContract struct {
	//[ 0] bl_cid                                         INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	BlCid int32 `gorm:"primary_key;AUTO_INCREMENT;column:bl_cid;type:INT4;" json:"bl_cid" db:"bl_cid"`
	//[ 1] cid                                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Cid null.Int `gorm:"column:cid;type:INT4;" json:"cid" db:"cid"`
	//[ 2] bank_contract_no                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankContractNo null.String `gorm:"column:bank_contract_no;type:VARCHAR;size:255;" json:"bank_contract_no" db:"bank_contract_no"`
	//[ 3] bank_contract_name                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankContractName null.String `gorm:"column:bank_contract_name;type:VARCHAR;size:255;" json:"bank_contract_name" db:"bank_contract_name"`
	//[ 4] bank_account                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankAccount null.String `gorm:"column:bank_account;type:VARCHAR;size:255;" json:"bank_account" db:"bank_account"`
	//[ 5] interest_calc_method                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	InterestCalcMethod null.String `gorm:"column:interest_calc_method;type:VARCHAR;size:255;" json:"interest_calc_method" db:"interest_calc_method"`
	//[ 6] bank_name                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankName string `gorm:"column:bank_name;type:VARCHAR;size:255;" json:"bank_name" db:"bank_name"`
	//[ 7] bank_branch                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankBranch null.String `gorm:"column:bank_branch;type:VARCHAR;size:255;" json:"bank_branch" db:"bank_branch"`
	//[ 8] loan_principal                                 INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	LoanPrincipal int64 `gorm:"column:loan_principal;type:INT8;" json:"loan_principal" db:"loan_principal"`
	//[ 9] loan_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoanMethod null.String `gorm:"column:loan_method;type:VARCHAR;size:255;" json:"loan_method" db:"loan_method"`
	//[10] contract_start_date                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ContractStartDate null.Time `gorm:"column:contract_start_date;type:DATE;" json:"contract_start_date" db:"contract_start_date"`
	//[11] contract_end_date                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ContractEndDate null.Time `gorm:"column:contract_end_date;type:DATE;" json:"contract_end_date" db:"contract_end_date"`
	//[12] actual_start_date                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ActualStartDate null.Time `gorm:"column:actual_start_date;type:DATE;" json:"actual_start_date" db:"actual_start_date"`
	//[13] is_lpr                                         BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsLpr null.Int `gorm:"column:is_lpr;type:BOOL;" json:"is_lpr" db:"is_lpr"`
	//[14] current_reprice_day                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	CurrentRepriceDay null.Time `gorm:"column:current_reprice_day;type:DATE;" json:"current_reprice_day" db:"current_reprice_day"`
	//[15] current_lpr                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentLpr null.Int `gorm:"column:current_lpr;type:INT4;" json:"current_lpr" db:"current_lpr"`
	//[16] lpr_plus                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LprPlus null.Int `gorm:"column:lpr_plus;type:INT4;" json:"lpr_plus" db:"lpr_plus"`
	//[17] current_rate                                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentRate int32 `gorm:"column:current_rate;type:INT4;" json:"current_rate" db:"current_rate"`
	//[18] next_reprice_day                               DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	NextRepriceDay null.Time `gorm:"column:next_reprice_day;type:DATE;" json:"next_reprice_day" db:"next_reprice_day"`
	//[19] all_repaid_principal                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidPrincipal null.Int `gorm:"column:all_repaid_principal;type:INT8;" json:"all_repaid_principal" db:"all_repaid_principal"`
	//[20] all_repaid_interest                            INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AllRepaidInterest null.Int `gorm:"column:all_repaid_interest;type:INT8;" json:"all_repaid_interest" db:"all_repaid_interest"`
	//[21] is_finished                                    BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsFinished null.Int `gorm:"column:is_finished;type:BOOL;" json:"is_finished" db:"is_finished"`
	//[22] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ContactPerson null.String `gorm:"column:contact_person;type:VARCHAR;size:255;" json:"contact_person" db:"contact_person"`
	//[23] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	ContactTel null.String `gorm:"column:contact_tel;type:VARCHAR;size:50;" json:"contact_tel" db:"contact_tel"`
	//[24] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time" db:"create_time"`
	//[25] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	ModifyTime time.Time `gorm:"column:modify_time;type:TIMESTAMP;" json:"modify_time" db:"modify_time"`
}

var bank_loan_contractTableInfo = &TableInfo{
	Name: "bank_loan_contract",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "bl_cid",
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
			GoFieldName:        "BlCid",
			GoFieldType:        "int32",
			JSONFieldName:      "bl_cid",
			ProtobufFieldName:  "bl_cid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "cid",
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
			GoFieldName:        "Cid",
			GoFieldType:        "null.Int",
			JSONFieldName:      "cid",
			ProtobufFieldName:  "cid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
			Index:              24,
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
			ProtobufPos:        25,
		},

		{
			Index:              25,
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
