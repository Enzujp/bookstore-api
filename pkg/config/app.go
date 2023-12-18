package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// connect to database
func Connect() {
	d, err := gorm.Open("mysql", "enzu: jessedavid/bookstore?charset=utf-8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}


func GetDB() *gorm.DB {
	return db
}
