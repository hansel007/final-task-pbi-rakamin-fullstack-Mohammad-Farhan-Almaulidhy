package database

import (
	"GO/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/dbfinaltaskbtpn?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{}, &models.Photo{})
	DB = database
}