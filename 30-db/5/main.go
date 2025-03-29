package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

// type SerialNumber struct {
// 	ID        int `gorm:"primaryKey"`
// 	Number    string
// 	ProductID int
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/goex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// //create
	category := Category{Name: "Eletrodomesticos"}
	db.Create(&category)

	// //create products
	db.Create(&Product{
		Name:       "Celular",
		Price:      1800.12,
		CategoryID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ": ")
		for _, product := range category.Products {
			fmt.Println("-", product.Name)
		}

	}

}
