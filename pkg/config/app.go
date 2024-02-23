package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func init(){
	godotenv.Load(".env")
	
}

// connect to database
func Connect() {
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DBURL not found in env")
	}
	
	d, err := gorm.Open("mysql", DB_URL)
	if err != nil {
		panic(err)
	}
	db = d
}


func GetDB() *gorm.DB {
	return db
}
