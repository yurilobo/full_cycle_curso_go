package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	db.Create(&Product{
		Name:  "Celular",
		Price: 1800.12,
	})
	//create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1500.0},
	// 	{Name: "Tablet", Price: 800.50},
	// 	{Name: "Monitor", Price: 1000.0},
	// 	{Name: "Mouse", Price: 50.0},
	// 	{Name: "Teclado", Price: 120.0},
	// }
	// db.Create(&products)

	//select one
	// var product Product
	// db.First(&product, 2)
	// db.First(&product, "name = 7", "Mouse")
	// fmt.Println(product)

	//select all
	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	//where
	// var products []Product
	// db.Where("price >?", 1).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }
	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)

	//
}
