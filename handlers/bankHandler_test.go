package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/wigwamwam/CRUD_app/handlers"
)

func TestShowBank(t *testing.T) {

	// Test case 1: Successful request with a valid bank ID
	bankID := 2

	url := fmt.Sprintf("/banks/%v", strconv.Itoa(bankID))

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ShowBank())
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}


}
