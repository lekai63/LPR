package models

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Gormv2 *gorm.DB

	GlobalConn db.Connection
)

func InitGormv2(c db.Connection) {

	tables = make(map[string]*TableInfo)
	tables["lease_contract"] = lease_contractTableInfo
	tables["lease_repay_plan"] = lease_repay_planTableInfo
	tables["lessee_info"] = lessee_infoTableInfo

	// 暂时注释掉，完成哪个表单就取消注释哪个
	tables["bank_loan_contract"] = bank_loan_contractTableInfo
	tables["bank_repay_plan"] = bank_repay_planTableInfo

	tables["shareholder_loan_contract"] = shareholder_loan_contractTableInfo
	tables["shareholder_loan_repaid_record"] = shareholder_loan_repaid_recordTableInfo

	var err error
	if c == nil {
		dsn := "user=fzzl password=fzzl032003 dbname=lpr port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		Gormv2, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//	TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})
		// sqlDB, _ := Gormv2.DB()

	} else {
		// 赋值全局数据库连接
		GlobalConn = c
		sqlDB := c.GetDB("default")

		Gormv2, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	}

	if err != nil {
		fmt.Printf("%s", err)
		panic("initialize orm failed")
	}

}
