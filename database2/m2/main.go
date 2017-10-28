package main

import (
	"fmt"
	"log"
	"net/http"
	"mywriteCode/database2/models"
)

type Env struct {
	db models.Datastore  //为了替换数据源，便于测试　　技巧设计
}

func main() {
	db, err := models.NewDB("host=localhost port=5432 user=fabric_ca password=123456 dbname=fabric_ca sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := env.db.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}