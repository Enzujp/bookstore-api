package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// connect to database
func connect() {
	d, err := gorm.Open("mysql", "enzu: jessedavid/bookstore?charset=utf-8&parseTime=True&loc=local")
	if err != null {
		panic(err)
	}
	db = d
}


func getDB() *gorm.DB {
	return db
}
