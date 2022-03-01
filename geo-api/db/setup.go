package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	USERNAME := "root"
	PASSWORD := "666666"
	HOST := "localhost"
	PORT := "3306"
	DB_NAME := "geo_db"
	CONN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DB_NAME)
	fmt.Printf("CONN: %v\n", CONN)
	db, err := gorm.Open("mysql", CONN)
	if err != nil {
		panic(err.Error())
	}
	return db
}
