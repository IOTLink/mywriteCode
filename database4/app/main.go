package main

import (
	"mywriteCode/database4/models"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	h(ctx, rw, req)
}

type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}

func main() {
	db, err := models.NewDB("host=localhost port=5432 user=fabric_ca password=123456 dbname=fabric_ca sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	ctx := context.WithValue(context.Background(), "db", db)

	http.Handle("/books", &ContextAdapter{ctx, ContextHandlerFunc(booksIndex)})
	http.ListenAndServe(":3000", nil)
}

func booksIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("start http")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks(ctx)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

