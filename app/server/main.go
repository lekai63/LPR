package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lekai63/lpr/models"
)

func main() {
	// db, err := gorm.Open("postgres", "host=192.168.5.11 user=remote dbname=fzzl sslmode=disable password=my032003")

	dsn := "host=192.168.5.11 user=remote password=my032003 dbname=fzzl port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("fail to open postgres")
	}

	sqlDB, err := db.DB()
	fmt.Println(sqlDB.Stats())

	defer sqlDB.Close()

	li := models.LesseeInfo{}

	db.Find(&li)

	fmt.Print(li)
}
