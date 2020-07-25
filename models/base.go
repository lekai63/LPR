package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	Gorm *gorm.DB
	err  error
)

func Init(c db.Connection) {
	Gorm, err = gorm.Open("postgresql", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}
}
