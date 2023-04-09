package config

import (
	"github.com/bitwisemj/api-market/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() {

	db, err := GetConnection()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Product{})
}

func GetConnection() (*gorm.DB, error) {

	dsn := "root:root@tcp(localhost:3306)/market?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
