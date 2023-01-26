package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db gorm.DB

func ConnectAndCheck() {
	db, err := gorm.Open(
		mysql.Open("root:123456@tcp(127.0.0.1:3306)/world?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	// Read
	var product Product
	//result := db.First(&product, 1) // find product with integer primary key
	// find product with code D42
	if db.Last(&product, 1).Error != nil {
		fmt.Println("Err")
	}

}
