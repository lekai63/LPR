package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"../models"
)

// type LesseeInfo struct {
// 	//[ 0] customer_id                                    INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
// 	CustomerID int32 `gorm:"primary_key;AUTO_INCREMENT;column:customer_id;type:INT4;"`
// 	//[ 1] business_license                               VARCHAR(18)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 18      default: []
// 	BusinessLicense sql.NullString `gorm:"column:business_license;type:VARCHAR;size:18;"`
// 	//[ 2] lessee                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
// 	Lessee string `gorm:"column:lessee;type:VARCHAR;size:255;"`
// 	//[ 3] short_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
// 	ShortName sql.NullString `gorm:"column:short_name;type:VARCHAR;size:255;"`
// 	//[ 4] email                                          VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
// 	Email sql.NullString `gorm:"column:email;type:VARCHAR;size:50;"`
// 	//[ 5] contact_person                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
// 	ContactPerson sql.NullString `gorm:"column:contact_person;type:VARCHAR;size:255;"`
// 	//[ 6] contact_tel                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
// 	ContactTel sql.NullString `gorm:"column:contact_tel;type:VARCHAR;size:50;"`
// 	//[ 7] customer_manager                               VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
// 	CustomerManager sql.NullString `gorm:"column:customer_manager;type:VARCHAR;size:10;"`
// 	//[ 8] risk_manager                                   VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
// 	RiskManager sql.NullString `gorm:"column:risk_manager;type:VARCHAR;size:10;"`
// 	//[ 9] create_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
// 	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;"`
// 	//[10] modify_time                                    TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
// 	ModifyTime time.Time `gorm:"column:modify_time;type:TIMESTAMP;"`
// }

func main() {
	// db, err := gorm.Open("postgres", "host=192.168.5.11 user=remote dbname=fzzl sslmode=disable password=my032003")

	dsn := "host=192.168.5.11 user=remote password=my032003 dbname=fzzl port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{

			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		fmt.Println("fail to open postgres")
		// fmt.Errorf("数据库打开失败 %v", err)
	}

	sqlDB, err := db.DB()
	fmt.Println(sqlDB.Stats())

	defer sqlDB.Close()

	// li := new(lesseeInfo)
	li := models.LesseeInfo{}
	// db.Where("short_name = ?", "丽水南城").First(&li)

	db.Find(&li)
	fmt.Println(li)
	fmt.Print("%+v", li)
}
