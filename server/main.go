package main

import (
	"fmt"
	"net/http"
	"server/routesHandlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	routesHandlers.SqlConnect()
	defer routesHandlers.Database.Close()

	fs := http.FileServer(http.Dir("../Site"))

	http.Handle("/", fs)
	http.HandleFunc("/get-all-products", routesHandlers.ProductsHandler)
	http.HandleFunc("/get-product/", routesHandlers.OneProductHandler)

	fmt.Println("Server starts...")
	http.ListenAndServe(":8080", nil)
}
