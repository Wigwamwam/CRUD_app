package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/wigwamwam/CRUD_app/handlers"
	"github.com/wigwamwam/CRUD_app/models"
)

// Questions:
// How to refactor

func TestHandlerIndexBanks(t *testing.T) {
	t.Run("Valid Index Route", func(t *testing.T) {
		r := chi.NewRouter()
		r.Get("/banks", handlers.HandlerIndexBanks())

		url := "/banks"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		res := rr.Result()
		defer res.Body.Close()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}

func TestCreateBank(t *testing.T) {
	t.Run("Valid Create Request", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/banks", handlers.HandlerIndexBanks())

		// setup test data
		bank := models.Bank{Name: "Test Ban", IBAN: "12345"}
		payload, _ := json.Marshal(bank)
		url := "/banks"

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new bank
		w := httptest.NewRecorder()
		handler.CreateBank()(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code to be 201, got %d", resp.StatusCode)
		}

		var newBank models.Bank
		b, _ := io.ReadAll(resp.Body)
		json.Unmarshal(b, &newBank)

		if newBank.Name != bank.Name {
			t.Errorf("Expected bank name to be %s, got %s", bank.Name, newBank.Name)
		}

		if newBank.IBAN != bank.IBAN {
			t.Errorf("Expected bank IBAN to be %s, got %s", bank.IBAN, newBank.IBAN)
		}
	})

	t.Run("Test error when reading request body with invalid JSON", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/banks", handlers.HandlerIndexBanks())
		url := "/banks"

		// testing invalid request body:
		bank := "invalid json"
		payload, _ := json.Marshal(bank)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handlers.CreateBank()(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code to be 400, got %d", resp.StatusCode)
		}
	})
}

func TestShowBank(t *testing.T) {
	// Test case 1: Successful request with a valid bank ID

	r := chi.NewRouter()
	r.Get("/banks/{id}", handlers.ShowBank())

	t.Run("Valid ID - with id=2", func(t *testing.T) {

		url := "/banks/2"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		res := rr.Result()
		defer res.Body.Close()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	// Test case 2: Unsuccessful request with a invalid bank ID
	t.Run("Invalid ID - with id=invalid", func(t *testing.T) {

		url := "/banks/invalid"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		res := rr.Result()
		defer res.Body.Close()

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
		}
	})
}

func TestDeleteBank(t *testing.T) {
	r := chi.NewRouter()
	r.Delete("/banks/{id}", handlers.ShowBank())

	t.Run("Invalid ID", func(t *testing.T) {
		url := "/banks/invalid"
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("expected %d but got %d", http.StatusNoContent, status)
		}
	})

	t.Run("Valid ID", func(t *testing.T) {
		url := "/banks/2"
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected %d but got %d", http.StatusOK, status)
		}
	})
}
