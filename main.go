package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"  // Use "localhost" or "127.0.0.1" for connecting from host to container
	port     = 5432         // Port as defined in docker-compose.yml
	user     = "myuser"     // Should match POSTGRES_USER in docker-compose.yml
	password = "mypassword" // Should match POSTGRES_PASSWORD in docker-compose.yml
	dbname   = "mydatabase" // Should match POSTGRES_DB in docker-compose.yml
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {

psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)


	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	// Connection Database Successful
	fmt.Println("Connection Database Successful")

	// ##### SELECT
	product, err := getProduct(3)

	fmt.Println("Get Successful!", product)

	// ##### Create
	// err = createProduct(&Product{Name: "Go product 2", Price: 444})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Create Successful")

}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id, name, price FROM products where id=$1;",id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}
