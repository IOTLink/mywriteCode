package main

import (
	"mywriteCode/database2/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

type mockDB struct{}

type Env struct {
	db models.Datastore
}

func (mdb *mockDB) AllBooks() ([]*models.Book, error) {
	bks := make([]*models.Book, 0)
	bks = append(bks, &models.Book{"978-1503261969", "Emma", "Jayne Austen", 9.44})
	bks = append(bks, &models.Book{"978-1505255607", "The Time Machine", "H. G. Wells", 5.99})
	return bks, nil
}

func TestBooksIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.booksIndex).ServeHTTP(rec, req)

	expected := "978-1503261969, Emma, Jayne Austen, £9.44\n978-1505255607, The Time Machine, H. G. Wells, £5.99\n"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

/////////////////add my
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