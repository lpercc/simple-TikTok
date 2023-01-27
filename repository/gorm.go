package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id    int64
	Code  string
	Price uint
}

func Initgorm() {
	db, err := gorm.Open(
		mysql.Open("root:123456@tcp(127.0.0.1:3306)/world?charset=utf8"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
