package models

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/config"

	_ "github.com/lib/pq"
)

type Product struct {
	ID           int64  `json:"id"`
	Product_name string `json:"product_name"`
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
}

func AddProduct(product Product) int64 {

	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO product (product_name, price, quantity) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, product.Product_name, product.Price, product.Quantity).Scan(&id)

	if err != nil {
		log.Fatalf("Can't execute the query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	return id

}

func TakeAllProducts() ([]Product, error) {

	db := config.CreateConnection()

	defer db.Close()

	var products []Product

	sqlStatement := `SELECT * FROM product`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Can't execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product Product

		err := rows.Scan(&product.ID, &product.Product_name, &product.Price, &product.Quantity)

		if err != nil {
			log.Fatalf("Can't take the data. %v", err)
		}

		products = append(products, product)
	}

	return products, err

}

func TakeProduct(id int64) (Product, error) {

	db := config.CreateConnection()

	defer db.Close()

	var product Product

	sqlStatement := `SELECT * FROM product WHERE id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&product.ID, &product.Product_name, &product.Price, &product.Quantity)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("There is no data to find!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("Can't take the data. %v", err)
	}

	return product, err

}

func UpdateProduct(id int64, product Product) int64 {

	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE product SET product_name=$2, price=$3, quantity=$4 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, product.Product_name, product.Price, product.Quantity)

	if err != nil {
		log.Fatalf("Can't execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error when checking the rows/data to be updated. %v", err)
	}

	fmt.Printf("Total rows/record to be updated %v \n", rowsAffected)

	return rowsAffected

}

func DeleteProduct(id int64) int64 {

	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM product WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Can't execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Can't find the data. %v", err)
	}

	fmt.Printf("Total data to be deleted %v", rowsAffected)

	return rowsAffected

}
