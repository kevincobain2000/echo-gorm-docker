package db

import (
	"fmt"

	"github.com/kevincobain2000/echo-gorm-docker/config"
	"github.com/kevincobain2000/echo-gorm-docker/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connect_string)
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	db.LogMode(true)
	db.AutoMigrate(&model.User{})

}

func DbManager() *gorm.DB {
	return db
}
