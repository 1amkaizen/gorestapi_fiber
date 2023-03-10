package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/go_restapi_fiber"))
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&Book{})
	DB = db
}
