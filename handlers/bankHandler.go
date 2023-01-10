package handlers

// TODO: imrpove error handling

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/wigwamwam/CRUD_app/initializers"
	"github.com/wigwamwam/CRUD_app/models"
)

type errorResponse struct {
	Message string
}

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	errorResponse := errorResponse{fmt.Sprintf("%v -%v", msg, err)}
	response, _ := json.Marshal(errorResponse)
	respondWithJSON(w, code, response)
}

func respondWithJSON(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HandlerIndexBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var banks []models.Bank

		initializers.DB.Find(&banks)

		response, err := json.Marshal(banks)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
			return
		}

		respondWithJSON(w, http.StatusOK, response)
	}
}

func CreateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bank := models.Bank{}

		bytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Cannot read create bank input", err)
			return
		}

		err = json.Unmarshal(bytes, &bank)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Cannot read create bank input", err)
			return
		}

		newBank := models.Bank{Name: bank.Name, IBAN: bank.IBAN}

		initializers.DB.Create(&newBank)

		response, err := json.Marshal(newBank)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
			return
		}

		respondWithJSON(w, http.StatusCreated, response)
	}
}

func ShowBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// why not use chi.URLParam?
		// url := chi.URLParam(r, "id")

		// url := r.URL.Query()["id"][0]

		// id, err := strconv.Atoi(url)

		// fmt.Println(url)

		// if err != nil {
		respondWithError(w, http.StatusNoContent, "invalid ID", errors.New(r.URL.RawQuery))
		return
		// }

		// var bank []models.Bank
		// initializers.DB.Find(&bank, id)

		// response, err := json.Marshal(bank)
		// if err != nil {
		// 	respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
		// 	return
		// }

		// respondWithJSON(w, http.StatusOK, response)
	}
}

func DeleteBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		fmt.Println(err)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "ID not found:", err)
			return
		}

		var bank []models.Bank

		initializers.DB.Find(&bank, id)

		response, err := json.Marshal(bank)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
			return
		}
		// moved permantly on deleted
		respondWithJSON(w, http.StatusMovedPermanently, response)
	}
}

func UpdateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, http.StatusNotFound, "ID not found:", err)
			// need to always add this when handling erros
			return
		}

		bank := models.Bank{}

		bytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Cannot read create bank input", err)
			return
		}

		err = json.Unmarshal(bytes, &bank)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Cannot read create bank input", err)
			return
		}

		var updatedBank []models.Bank

		initializers.DB.Find(&updatedBank, id)
		initializers.DB.Model(&updatedBank).Updates(models.Bank{Name: bank.Name, IBAN: bank.IBAN})

		response, err := json.Marshal(bank)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
			return
		}

		respondWithJSON(w, http.StatusOK, response)
	}
}
