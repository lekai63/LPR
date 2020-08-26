package models

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Gormv2 *gorm.DB

	GlobalConn db.Connection
)

func InitGormv2(c db.Connection) {
	// 赋值全局数据库连接
	GlobalConn = c
	tables = make(map[string]*TableInfo)
	tables["lease_contract"] = lease_contractTableInfo
	tables["lease_repay_plan"] = lease_repay_planTableInfo
	tables["lessee_info"] = lessee_infoTableInfo

	// 暂时注释掉，完成哪个表单就取消注释哪个
	tables["bank_loan_contract"] = bank_loan_contractTableInfo
	tables["bank_repay_plan"] = bank_repay_planTableInfo

	tables["shareholder_loan_contract"] = shareholder_loan_contractTableInfo
	tables["shareholder_loan_repaid_record"] = shareholder_loan_repaid_recordTableInfo

	dsn := "host=192.168.5.11 user=fzzl password=fzzl032003 dbname=lpr port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	sqldb := c.GetDB("default")
	var conf = postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
		Conn:                 sqldb,
	}
	var d = postgres.Dialector{
		&conf,
	}
	Gormv2, err = gorm.Open(&d, &gorm.Config{})

	if err != nil {
		fmt.Printf("%s", err)
		panic("initialize orm failed")
	}

}
