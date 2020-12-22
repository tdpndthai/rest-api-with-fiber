package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"rest-api-with-fiber/config"
	"rest-api-with-fiber/model"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error

	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"),config.Config("DB_NAME")))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Product{})
	fmt.Println("Database Migrated")
}
