package routesHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Size        string
	Price       int
	Picture     string
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	pageNumber, _ := strconv.ParseInt(r.URL.Query().Get("pageNumber"), 10, 32)

	rows, err := Database.Query("select * from Products")
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

	totalCount := strconv.Itoa(len(products))
	resultJson, err := json.Marshal(products[(pageNumber-1)*limit : limit*pageNumber])

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Add("total-count", totalCount)
	w.Header().Add("Access-Control-Expose-Headers", "total-count")

	sendData := string(resultJson)
	fmt.Fprintf(w, "%s", sendData)
}

func OneProductHandler(w http.ResponseWriter, r *http.Request) {
	reg, _ := regexp.Compile("/get-product/([0-9]+)")
	if len(reg.FindStringSubmatch(r.URL.Path)) > 0 {
		id, err := strconv.ParseInt(reg.FindStringSubmatch(r.URL.Path)[1], 0, 64)
		if err != nil {
			panic(err)
		}
		row := Database.QueryRow("select * from Products where id=$1", id)
		p := Product{}
		err = row.Scan(&p.Id, &p.Title, &p.Description, &p.Size, &p.Price, &p.Picture)
		if err != nil {
			panic(err)
		}
		resultJson, err := json.Marshal(p)
		if err != nil {
			fmt.Println(err)
		}
		sendData := string(resultJson)
		fmt.Fprintf(w, "%s", sendData)
	}
}
