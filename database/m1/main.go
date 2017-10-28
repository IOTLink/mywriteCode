package main

import (
	"mywriteCode/database/models"
	"fmt"
	"net/http"
)

func main() {
	//models.InitDB("postgres://fabric_ca:123456@127.0.0.1/fabric_ca sslmode=disable")
	models.InitDB("host=localhost port=5432 user=fabric_ca password=123456 dbname=fabric_ca sslmode=disable")
	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}