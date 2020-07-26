package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Gorm *gorm.DB
	err  error
)

func Init(c db.Connection) {
	tables = make(map[string]*TableInfo)
	// 暂时注释掉
	// tables["bank_loan_contract"] = bank_loan_contractTableInfo
	// tables["bank_repay_plan"] = bank_repay_planTableInfo
	// tables["lease_contract"] = lease_contractTableInfo
	// tables["lease_repay_plan"] = lease_repay_planTableInfo
	tables["lessee_info"] = lessee_infoTableInfo
	// tables["shareholder_loan_contract"] = shareholder_loan_contractTableInfo
	// tables["shareholder_loan_repaid_record"] = shareholder_loan_repaid_recordTableInfo

	Gorm, err = gorm.Open("postgresql", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}

	Gorm.LogMode(true)
}
