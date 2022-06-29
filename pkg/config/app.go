package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	// Use your password instead of "password" in next line
	d, err := gorm.Open("mysql", "arun:password@/BookStore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
