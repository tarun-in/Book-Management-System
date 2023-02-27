package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	connectionString := "root:tnluser123@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connectionString)

	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
