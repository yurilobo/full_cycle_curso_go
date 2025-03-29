package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() { //cria conexão com banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goex")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 1800.12)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 1500.0
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Mudança: Renomeando a variável 'product' para 'products' para armazenar o slice de produtos
	products, err := selecAllProduct(db)
	if err != nil {
		panic(err)
	}

	// Mudança: Iterando sobre 'products'
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

// Função para inserir um novo dado
func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO product (id, name, price) VALUES(?, ?,?)") //usamos ?
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE product SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selecAllProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product // Mudança: Renomeando a variável de 'product' para 'products'
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price) //adiciona os resultados na variável p
		if err != nil {
			return nil, err
		}
		products = append(products, p) // Adiciona o produto ao slice 'products'
	}
	return products, nil // retorna o slice de produtos
}
func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM product WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
