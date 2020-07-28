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


Table: lease_contract
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] contract_no                                    VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] lessee                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] abbreviation                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] start_date                                     DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 5] end_date                                       DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[ 6] fee                                            INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] margin                                         INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 8] contract_principal                             INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] actual_principal                               INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] term_month                                     INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[11] subject_matter                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] irr                                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[13] is_lpr                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[14] current_reprice_day                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[15] current_LPR                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[16] lpr_plus                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[17] current_rate                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[18] next_reprice_day                               DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
[19] received_principal                             INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[20] received_interest                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[21] is_finished                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[22] lessee_info_id                                 INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[23] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[24] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "received_interest": 63,    "updated_at": "2073-08-05T16:50:19.752232494Z",    "lessee": "DtBRdsNpuxuCfSCeSeZojsWko",    "fee": 43,    "is_lpr": false,    "current_reprice_day": "2053-10-01T14:35:24.24712212Z",    "next_reprice_day": "2081-12-27T06:32:29.985436509Z",    "current_lpr": 85,    "current_rate": 83,    "created_at": "2084-08-14T04:28:02.105836245Z",    "contract_no": "VHkHRSxXdVZZMBwCPjNEGYtKT",    "abbreviation": "UonadIOWRcvluLPmSQMuMrZcO",    "contract_principal": 44,    "term_month": 50,    "irr": 89,    "id": 75,    "start_date": "2248-08-31T07:45:41.919305751Z",    "actual_principal": 51,    "subject_matter": "UEVwhfhItQudcfKqvJwByiFUy",    "received_principal": 56,    "end_date": "2080-03-15T23:08:12.24363701Z",    "margin": 60,    "lpr_plus": 71,    "is_finished": true,    "lessee_info_id": 46}



*/

// LeaseContract struct is a row record of the lease_contract table in the fzzl database
type LeaseContract struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] contract_no                                    VARCHAR(255)         null: false  primary: false   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ContractNo string `gorm:"column:contract_no;type:VARCHAR;size:255;" json:"contract_no"`
	//[ 2] lessee                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Lessee string `gorm:"column:lessee;type:VARCHAR;size:255;" json:"lessee"`
	//[ 3] abbreviation                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Abbreviation null.String `gorm:"column:abbreviation;type:VARCHAR;size:255;" json:"abbreviation"`
	//[ 4] start_date                                     DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	StartDate time.Time `gorm:"column:start_date;type:DATE;" json:"start_date"`
	//[ 5] end_date                                       DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	EndDate time.Time `gorm:"column:end_date;type:DATE;" json:"end_date"`
	//[ 6] fee                                            INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Fee int64 `gorm:"column:fee;type:INT8;" json:"fee"`
	//[ 7] margin                                         INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Margin int64 `gorm:"column:margin;type:INT8;" json:"margin"`
	//[ 8] contract_principal                             INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ContractPrincipal int64 `gorm:"column:contract_principal;type:INT8;" json:"contract_principal"`
	//[ 9] actual_principal                               INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ActualPrincipal int64 `gorm:"column:actual_principal;type:INT8;" json:"actual_principal"`
	//[10] term_month                                     INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	TermMonth int32 `gorm:"column:term_month;type:INT2;" json:"term_month"`
	//[11] subject_matter                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	SubjectMatter null.String `gorm:"column:subject_matter;type:VARCHAR;size:255;" json:"subject_matter"`
	//[12] irr                                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Irr null.Int `gorm:"column:irr;type:INT4;" json:"irr"`
	//[13] is_lpr                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsLpr bool `gorm:"column:is_lpr;type:BOOL;" json:"is_lpr"`
	//[14] current_reprice_day                            DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	CurrentRepriceDay null.Time `gorm:"column:current_reprice_day;type:DATE;" json:"current_reprice_day"`
	//[15] current_LPR                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentLPR null.Int `gorm:"column:current_LPR;type:INT4;" json:"current_lpr"`
	//[16] lpr_plus                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LprPlus null.Int `gorm:"column:lpr_plus;type:INT4;" json:"lpr_plus"`
	//[17] current_rate                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentRate null.Int `gorm:"column:current_rate;type:INT4;" json:"current_rate"`
	//[18] next_reprice_day                               DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	NextRepriceDay null.Time `gorm:"column:next_reprice_day;type:DATE;" json:"next_reprice_day"`
	//[19] received_principal                             INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ReceivedPrincipal null.Int `gorm:"column:received_principal;type:INT8;" json:"received_principal"`
	//[20] received_interest                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	ReceivedInterest null.Int `gorm:"column:received_interest;type:INT8;" json:"received_interest"`
	//[21] is_finished                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsFinished bool `gorm:"column:is_finished;type:BOOL;" json:"is_finished"`
	//[22] lessee_info_id                                 INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LesseeInfoID int32 `gorm:"column:lessee_info_id;type:INT4;" json:"lessee_info_id"`
	//[23] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//[24] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
}

var lease_contractTableInfo = &TableInfo{
	Name: "lease_contract",
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
			Name:               "contract_no",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ContractNo",
			GoFieldType:        "string",
			JSONFieldName:      "contract_no",
			ProtobufFieldName:  "contract_no",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "lessee",
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
			GoFieldName:        "Lessee",
			GoFieldType:        "string",
			JSONFieldName:      "lessee",
			ProtobufFieldName:  "lessee",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "start_date",
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
			GoFieldName:        "StartDate",
			GoFieldType:        "time.Time",
			JSONFieldName:      "start_date",
			ProtobufFieldName:  "start_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "end_date",
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
			GoFieldName:        "EndDate",
			GoFieldType:        "time.Time",
			JSONFieldName:      "end_date",
			ProtobufFieldName:  "end_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "fee",
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
			GoFieldName:        "Fee",
			GoFieldType:        "int64",
			JSONFieldName:      "fee",
			ProtobufFieldName:  "fee",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "margin",
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
			GoFieldName:        "Margin",
			GoFieldType:        "int64",
			JSONFieldName:      "margin",
			ProtobufFieldName:  "margin",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "contract_principal",
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
			GoFieldName:        "ContractPrincipal",
			GoFieldType:        "int64",
			JSONFieldName:      "contract_principal",
			ProtobufFieldName:  "contract_principal",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "actual_principal",
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
			GoFieldName:        "ActualPrincipal",
			GoFieldType:        "int64",
			JSONFieldName:      "actual_principal",
			ProtobufFieldName:  "actual_principal",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "term_month",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "TermMonth",
			GoFieldType:        "int32",
			JSONFieldName:      "term_month",
			ProtobufFieldName:  "term_month",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "subject_matter",
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
			GoFieldName:        "SubjectMatter",
			GoFieldType:        "null.String",
			JSONFieldName:      "subject_matter",
			ProtobufFieldName:  "subject_matter",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "irr",
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
			GoFieldName:        "Irr",
			GoFieldType:        "null.Int",
			JSONFieldName:      "irr",
			ProtobufFieldName:  "irr",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "is_lpr",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "IsLpr",
			GoFieldType:        "bool",
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
			Name:               "current_LPR",
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
			GoFieldName:        "CurrentLPR",
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
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "CurrentRate",
			GoFieldType:        "null.Int",
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
			Name:               "received_principal",
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
			GoFieldName:        "ReceivedPrincipal",
			GoFieldType:        "null.Int",
			JSONFieldName:      "received_principal",
			ProtobufFieldName:  "received_principal",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "received_interest",
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
			GoFieldName:        "ReceivedInterest",
			GoFieldType:        "null.Int",
			JSONFieldName:      "received_interest",
			ProtobufFieldName:  "received_interest",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "is_finished",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "IsFinished",
			GoFieldType:        "bool",
			JSONFieldName:      "is_finished",
			ProtobufFieldName:  "is_finished",
			ProtobufType:       "bool",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "lessee_info_id",
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
			GoFieldName:        "LesseeInfoID",
			GoFieldType:        "int32",
			JSONFieldName:      "lessee_info_id",
			ProtobufFieldName:  "lessee_info_id",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LeaseContract) TableName() string {
	return "lease_contract"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LeaseContract) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LeaseContract) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LeaseContract) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LeaseContract) TableInfo() *TableInfo {
	return lease_contractTableInfo
}
