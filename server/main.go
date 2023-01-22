package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Size        string
	Price       int
	Picture     string
}

var database *sql.DB

func productsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("select * from Products")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.Size, &p.Price, &p.Picture)
		if err != nil {
			continue
		}
		products = append(products, p)
	}
	resultJson, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
	}
	sendData := string(resultJson)
	fmt.Fprintf(w, "%s", sendData)
}

func main() {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err)
	}

	database = db
	defer db.Close()

	fs := http.FileServer(http.Dir("../Site"))

	http.Handle("/", fs)
	http.HandleFunc("/get-all-products", productsHandler)

	fmt.Println("Server starts...")
	http.ListenAndServe(":8080", nil)
}
