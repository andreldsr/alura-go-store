package models

import (
	conn "alura-go-store/db"
	"database/sql"
	"strconv"
)

type Product struct {
	Id                int
	Name, Description string
	Quantity          int
	Price             float64
}

func FindAllProducts() (products []Product) {
	db := conn.Connect()
	p := Product{}
	allProducts, err := db.Query("SELECT * FROM products ORDER BY 1 DESC")
	if err != nil {
		panic(err.Error())
		return nil
	}
	for allProducts.Next() {
		p = buildProductFromDb(allProducts)
		products = append(products, p)
	}
	defer db.Close()
	return
}

func buildProductFromDb(rows *sql.Rows) (p Product) {
	var id, quantity int
	var name, description string
	var price float64
	err := rows.Scan(&id, &name, &description, &price, &quantity)
	if err != nil {
		panic(err.Error())
	}
	p.Id = id
	p.Name = name
	p.Description = description
	p.Price = price
	p.Quantity = quantity
	return
}

func FindById(id int) (p Product, err error) {
	db := conn.Connect()
	findById, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return
	}
	findById.Next()
	p = buildProductFromDb(findById)
	return
}

func AddProduct(name, description, priceString, quantityString string) (err error) {
	p, err := buildProduct(name, description, priceString, quantityString)
	if err != nil {
		return
	}
	db := conn.Connect()
	insert, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return
	}
	_, err = insert.Exec(p.Name, p.Description, p.Price, p.Quantity)
	if err != nil {
		return
	}
	defer db.Close()
	return
}

func buildProduct(name string, description string, priceString string, quantityString string) (p Product, err error) {
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		return
	}
	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		return
	}
	p = Product{0, name, description, quantity, price}
	return
}

func DeleteProduct(id int) {
	db := conn.Connect()
	deleteQuery, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	_, err = deleteQuery.Exec(id)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateProduct(idString, name, description, priceString, quantityString string) (p Product, err error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return
	}
	p, err = buildProduct(name, description, priceString, quantityString)
	if err != nil {
		return
	}
	p.Id = id
	db := conn.Connect()
	update, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		return
	}
	_, err = update.Exec(p.Name, p.Description, p.Price, p.Quantity, p.Id)
	return
}
