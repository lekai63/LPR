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


Table: lessee_info
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] business_license                               VARCHAR(18)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 18      default: []
[ 2] lessee                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] short_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] email                                          VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 5] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 7] customer_manager                               VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
[ 8] risk_manager                                   VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
[ 9] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[10] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "business_license": "ihwaAiJCTJWVjMbPCekNJuKQX",    "lessee": "QfsIAhXYaegccuQDSUAVQUHvW",    "short_name": "cMKbLWDTLDFutYXyIqDabRESC",    "email": "mDleomEDHngLEJdLqjxODCjVN",    "contact_person": "PoMNgJBmHJQcBomaCFsIYTtOT",    "contact_tel": "LBXucnnUBGfLZSrjquJnqLCRi",    "risk_manager": "FvjnCGuYAPnGrbdBGvBEeOfID",    "id": 39,    "customer_manager": "USjonDdCTXwSRIIZFqxqxoiJI",    "created_at": "2106-02-28T16:54:26.994308429Z",    "updated_at": "2275-08-12T05:23:30.253323103Z"}



*/

// LesseeInfo struct is a row record of the lessee_info table in the fzzl database
type LesseeInfo struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] business_license                               VARCHAR(18)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 18      default: []
	BusinessLicense null.String `gorm:"column:business_license;type:VARCHAR;size:18;" json:"business_license"`
	//[ 2] lessee                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Lessee string `gorm:"column:lessee;type:VARCHAR;size:255;" json:"lessee"`
	//[ 3] short_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ShortName null.String `gorm:"column:short_name;type:VARCHAR;size:255;" json:"short_name"`
	//[ 4] email                                          VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	Email null.String `gorm:"column:email;type:VARCHAR;size:50;" json:"email"`
	//[ 5] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ContactPerson null.String `gorm:"column:contact_person;type:VARCHAR;size:255;" json:"contact_person"`
	//[ 6] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	ContactTel null.String `gorm:"column:contact_tel;type:VARCHAR;size:50;" json:"contact_tel"`
	//[ 7] customer_manager                               VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
	CustomerManager null.String `gorm:"column:customer_manager;type:VARCHAR;size:10;" json:"customer_manager"`
	//[ 8] risk_manager                                   VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
	RiskManager null.String `gorm:"column:risk_manager;type:VARCHAR;size:10;" json:"risk_manager"`
	//[ 9] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;" json:"created_at"`
	//[10] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
}

var lessee_infoTableInfo = &TableInfo{
	Name: "lessee_info",
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
			Name:               "business_license",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(18)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       18,
			GoFieldName:        "BusinessLicense",
			GoFieldType:        "null.String",
			JSONFieldName:      "business_license",
			ProtobufFieldName:  "business_license",
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
			Name:               "short_name",
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
			GoFieldName:        "ShortName",
			GoFieldType:        "null.String",
			JSONFieldName:      "short_name",
			ProtobufFieldName:  "short_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "null.String",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "customer_manager",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       10,
			GoFieldName:        "CustomerManager",
			GoFieldType:        "null.String",
			JSONFieldName:      "customer_manager",
			ProtobufFieldName:  "customer_manager",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "risk_manager",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       10,
			GoFieldName:        "RiskManager",
			GoFieldType:        "null.String",
			JSONFieldName:      "risk_manager",
			ProtobufFieldName:  "risk_manager",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LesseeInfo) TableName() string {
	return "lessee_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LesseeInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LesseeInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LesseeInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LesseeInfo) TableInfo() *TableInfo {
	return lessee_infoTableInfo
}
