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


Table: lessee_info
[ 0] customer_id                                    INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] business_license                               VARCHAR(18)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 18      default: []
[ 2] lessee                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] short_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] email                                          VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 5] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 7] customer_manager                               VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
[ 8] risk_manager                                   VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
[ 9] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[10] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "create_time": "2224-11-03T06:01:33.47894552Z",    "modify_time": "2101-05-08T20:44:08.158349487Z",    "business_license": "uaPtOWCVrlVLwVRYSEfNjWSBt",    "lessee": "DhBsBEGyAHLxhhwinknDPyuie",    "short_name": "iNlVZDwKSoTVWFLvCFqCqXOfK",    "contact_person": "UjiEhIoMgergihfOPRIdDkDUr",    "risk_manager": "EXsUptgdToIMfeboanajumTHn",    "customer_id": 93,    "email": "vwpRnCjTpaEQqXLTxqtTOBral",    "contact_tel": "auqQFuAcFQOcdHAQqZrFRUOZv",    "customer_manager": "DCPfAsppHVSnoGrpQutbaGhup"}



*/

// LesseeInfo struct is a row record of the lessee_info table in the fzzl database
type LesseeInfo struct {
	//[ 0] customer_id                                    INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	CustomerID int32 `gorm:"primary_key;AUTO_INCREMENT;column:customer_id;type:INT4;" json:"customer_id" db:"customer_id"`
	//[ 1] business_license                               VARCHAR(18)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 18      default: []
	BusinessLicense null.String `gorm:"column:business_license;type:VARCHAR;size:18;" json:"business_license" db:"business_license"`
	//[ 2] lessee                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Lessee string `gorm:"column:lessee;type:VARCHAR;size:255;" json:"lessee" db:"lessee"`
	//[ 3] short_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ShortName null.String `gorm:"column:short_name;type:VARCHAR;size:255;" json:"short_name" db:"short_name"`
	//[ 4] email                                          VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	Email null.String `gorm:"column:email;type:VARCHAR;size:50;" json:"email" db:"email"`
	//[ 5] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ContactPerson null.String `gorm:"column:contact_person;type:VARCHAR;size:255;" json:"contact_person" db:"contact_person"`
	//[ 6] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	ContactTel null.String `gorm:"column:contact_tel;type:VARCHAR;size:50;" json:"contact_tel" db:"contact_tel"`
	//[ 7] customer_manager                               VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
	CustomerManager null.String `gorm:"column:customer_manager;type:VARCHAR;size:10;" json:"customer_manager" db:"customer_manager"`
	//[ 8] risk_manager                                   VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
	RiskManager null.String `gorm:"column:risk_manager;type:VARCHAR;size:10;" json:"risk_manager" db:"risk_manager"`
	//[ 9] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time" db:"create_time"`
	//[10] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	ModifyTime time.Time `gorm:"column:modify_time;type:TIMESTAMP;" json:"modify_time" db:"modify_time"`
}

var lessee_infoTableInfo = &TableInfo{
	Name: "lessee_info",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "customer_id",
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
			GoFieldName:        "CustomerID",
			GoFieldType:        "int32",
			JSONFieldName:      "customer_id",
			ProtobufFieldName:  "customer_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
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

		{
			Index:              9,
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
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
