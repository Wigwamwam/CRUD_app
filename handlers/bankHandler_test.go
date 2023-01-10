package handlers_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/wigwamwam/CRUD_app/handlers"
)

func TestShowBank(t *testing.T) {

	// Test case 1: Successful request with a valid bank ID
	r := chi.NewRouter()

	r.Get("/banks/", handlers.ShowBank())

	url := "/banks/?id=2"

	req, err := http.NewRequest("GET", url, nil)

	fmt.Println(req.URL.RawQuery)

	// req, err := http.NewRequest("GET", "/banks/?id=1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	res := rr.Result()
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)

	fmt.Println(string(data))

	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	// }
}
