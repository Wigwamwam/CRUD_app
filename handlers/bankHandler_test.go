package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wigwamwam/CRUD_app/handlers"
	"github.com/wigwamwam/CRUD_app/models"
	"github.com/wigwamwam/CRUD_app/repository"
	customErrors "github.com/wigwamwam/CRUD_app/repository/errors"
)

// Questions:
// How to refactor

func Test_HandlerIndexBanks(t *testing.T) {
	mockController := gomock.NewController(t)
	mockDao := repository.NewMockDAO(mockController)
	handler := handlers.NewHandler(mockDao)
	r := chi.NewRouter()

	r.Get("/banks", handler.HandlerIndexBanks())

	type tc struct {
		expectedError        error
		responseCode         int
		expectedResponseBody string
	}

	cases := map[string]*tc{
		"returns not found error": {
			expectedError:        &customErrors.NotFoundError{},
			responseCode:         http.StatusNotFound,
			expectedResponseBody: `{"Message":"no entries found"}`,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			mockDao.EXPECT().SelectAllBanks().Return(nil, tc.expectedError)

			req, err := http.NewRequest("GET", "/banks", nil)
			assert.Equal(t, nil, err)

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			assert.Equal(t, tc.responseCode, rr.Code)
			assert.JSONEq(t, tc.expectedResponseBody, rr.Body.String())
		})
	}

	t.Run("Valid Index Route", func(t *testing.T) {

		mockDao.EXPECT().SelectAllBanks().Return([]models.Bank{}, nil)

		req, err := http.NewRequest("GET", "/banks", nil)
		if err != nil {
			t.Fatal(err)
		}
		// Why does this through an error:
		// assert.Equal(t, req, nil)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("NotFoundError", func(t *testing.T) {
		mockDao.EXPECT().SelectAllBanks().Return(nil, &customErrors.NotFoundError{})

		req, err := http.NewRequest("GET", "/banks", nil)
		assert.Equal(t, nil, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("Unhandled error", func(t *testing.T) {
		mockDao.EXPECT().SelectAllBanks().Return(nil, errors.New("unhandled error"))
		req, err := http.NewRequest("GET", "/banks", nil)
		assert.Equal(t, nil, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func Test_CreateBank(t *testing.T) {
	mockController := gomock.NewController(t)
	mockDao := repository.NewMockDAO(mockController)
	handler := handlers.NewHandler(mockDao)
	r := chi.NewRouter()
	r.Post("/banks", handler.CreateBank())

	t.Run("Valid Create Request", func(t *testing.T) {
		incomingPayload := models.Bank{Name: "Test Ban", IBAN: "12345"}
		returnedBankFromDB := models.Bank{ID: 1, Name: "test bank", IBAN: "1234567890"}

		mockDao.EXPECT().InsertBank(incomingPayload).Return(returnedBankFromDB, nil).Times(1)

		body, _ := json.Marshal(incomingPayload)
		req, _ := http.NewRequest("POST", "/banks", bytes.NewReader(body))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		expectedBankResponse := models.Bank{}
		json.Unmarshal(rr.Body.Bytes(), &expectedBankResponse)

		assert.Equal(t, expectedBankResponse, returnedBankFromDB)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("Returns Creating Bank Error", func(t *testing.T) {
		incomingPayload := models.Bank{Name: "Test Ban", IBAN: "12345"}

		mockDao.EXPECT().InsertBank(incomingPayload).Return(models.Bank{}, &customErrors.CreatingBankError{}).Times(1)

		body, _ := json.Marshal(incomingPayload)
		req, err := http.NewRequest("POST", "/banks", bytes.NewReader(body))
		assert.Equal(t, nil, err)

		rr := httptest.NewRecorder()
		handler.CreateBank()(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("Unhandled Error", func(t *testing.T) {
		incomingPayload := models.Bank{Name: "Test Ban", IBAN: "12345"}

		mockDao.EXPECT().InsertBank(incomingPayload).Return(models.Bank{}, errors.New("unhandled error")).Times(1)

		body, _ := json.Marshal(incomingPayload)
		req, err := http.NewRequest("POST", "/banks", bytes.NewReader(body))
		assert.Equal(t, nil, err)

		rr := httptest.NewRecorder()
		handler.CreateBank()(rr, req)

		// assert.JSONEq(t, `{"ID":1,"name":"test bank","iban":"1234567890","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}`, rr.Body.String())
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func Test_ShowBank(t *testing.T) {
	mockController := gomock.NewController(t)
	mockDao := repository.NewMockDAO(mockController)
	handler := handlers.NewHandler(mockDao)
	r := chi.NewRouter()
	r.Get("/banks/{id}", handler.ShowBank())

	t.Run("Valid Show", func(t *testing.T) {
		id := 2
		req, _ := http.NewRequest("GET", "/banks/2", nil)
		returnedBankFromDB := models.Bank{ID: 2, Name: "test bank", IBAN: "1234567890"}
		mockDao.EXPECT().SelectBankByID(id).Return(returnedBankFromDB, nil)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		expectedBankResponse := models.Bank{}
		json.Unmarshal(rr.Body.Bytes(), &expectedBankResponse)

		assert.Equal(t, expectedBankResponse, returnedBankFromDB)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("return id not found error", func(t *testing.T) {
		id := 2
		mockDao.EXPECT().SelectBankByID(id).Return(models.Bank{}, &customErrors.IdNotFoundError{})
		req, err := http.NewRequest("GET", "/banks/2", nil)
		assert.Equal(t, nil, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// how to get the 0 to be a 2 - how to parse the id into the customer error handler?
		assert.JSONEq(t, `{"Message":"id: 0 not found"}`, rr.Body.String())

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("return scanning id error", func(t *testing.T) {
		id := 2
		mockDao.EXPECT().SelectBankByID(id).Return(models.Bank{}, &customErrors.ScanningIdError{})
		req, err := http.NewRequest("GET", "/banks/2", nil)
		assert.Equal(t, nil, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// how to get the 0 to be a 2 - how to parse the id into the customer error handler?
		assert.JSONEq(t, `{"Message":"error scanning bank with ID 0"}`, rr.Body.String())

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func Test_UpdateBank(t *testing.T) {
	mockController := gomock.NewController(t)
	mockDao := repository.NewMockDAO(mockController)
	handler := handlers.NewHandler(mockDao)
	r := chi.NewRouter()
	r.Put("/banks/{id}", handler.UpdateBank())

	t.Run("Valid update", func(t *testing.T) {
		id := 1
		incomingPayload := models.Bank{ID: 1, Name: "Test Bank", IBAN: "1234567890"}
		returnedBankFromDB := models.Bank{ID: 2, Name: "test bank", IBAN: "1234567890"}

		body, _ := json.Marshal(incomingPayload)
		req, _ := http.NewRequest("PUT", "/banks/1", bytes.NewReader(body))

		mockDao.EXPECT().UpdateBank(id, incomingPayload).Return(returnedBankFromDB, nil).Times(1)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		expectedBankResponse := models.Bank{}
		json.Unmarshal(rr.Body.Bytes(), &expectedBankResponse)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedBankResponse, returnedBankFromDB)
	})

	t.Run("Return Id not found", func(t *testing.T) {
		id := 1
		incomingPayload := models.Bank{ID: 1, Name: "Test Bank", IBAN: "1234567890"}
		// returnedBankFromDB := models.Bank{ID: 2, Name: "test bank", IBAN: "1234567890"}

		body, _ := json.Marshal(incomingPayload)
		req, _ := http.NewRequest("PUT", "/banks/1", bytes.NewReader(body))

		mockDao.EXPECT().UpdateBank(id,  incomingPayload).Return(models.Bank{}, &customErrors.IdNotFoundError{}).Times(1)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		expectedBankResponse := models.Bank{}
		json.Unmarshal(rr.Body.Bytes(), &expectedBankResponse)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
func Test_DeleteBank(t *testing.T) {
	mockController := gomock.NewController(t)
	mockDao := repository.NewMockDAO(mockController)
	handler := handlers.NewHandler(mockDao)
	r := chi.NewRouter()
	r.Delete("/banks/{id}", handler.DeleteBank())

	t.Run("Valid Delete", func(t *testing.T) {
		id := 2
		req, _ := http.NewRequest("DELETE", "/banks/2", nil)

		mockDao.EXPECT().DeleteBankByID(id).Return(nil)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("Returns DeletingBankError", func(t *testing.T) {
		id := 2
		req, _ := http.NewRequest("DELETE", "/banks/2", nil)

		mockDao.EXPECT().DeleteBankByID(id).Return(&customErrors.DeletingBankError{})

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("Returns NotFoundError", func(t *testing.T) {
		id := 2
		req, _ := http.NewRequest("DELETE", "/banks/2", nil)

		mockDao.EXPECT().DeleteBankByID(id).Return(&customErrors.NotFoundError{})

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}




// func TestDeleteBank(t *testing.T) {
// 	r := chi.NewRouter()
// 	r.Delete("/banks/{id}", handlers.ShowBank())

// 	t.Run("Invalid ID", func(t *testing.T) {
// 		url := "/banks/invalid"
// 		req, err := http.NewRequest("DELETE", url, nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		r.ServeHTTP(rr, req)

// 		if status := rr.Code; status != http.StatusNoContent {
// 			t.Errorf("expected %d but got %d", http.StatusNoContent, status)
// 		}
// 	})

// 	t.Run("Valid ID", func(t *testing.T) {
// 		url := "/banks/2"
// 		req, err := http.NewRequest("DELETE", url, nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		r.ServeHTTP(rr, req)

// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("expected %d but got %d", http.StatusOK, status)
// 		}
// 	})
// }
