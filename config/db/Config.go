package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
var err error

func InitDB(){
	dbDns := "root:password@tcp(127.0.0.1:3306)/golang_books?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: dbDns,
		},
	), &gorm.Config{})

	if err != nil {
		panic("Can  not connect to DB")
	}
	
	fmt.Print("Success Connet to Database")
	
}