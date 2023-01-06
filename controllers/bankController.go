package controllers

// TODO: imrpove error handling


import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/wigwamwam/CRUD_app/initializers"
	"github.com/wigwamwam/CRUD_app/models"
)

func IndexBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var banks []models.Bank
		initializers.DB.Find(&banks)

		response, err := json.Marshal(banks)

		if err != nil {
			fmt.Println("Error", err)
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func CreateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := models.Bank{}

		// marshall vs unmarchall

		json.NewDecoder(r.Body).Decode(&body)

		newBank := models.Bank{Name: body.Name, IBAN: body.IBAN}

		initializers.DB.Create(&newBank)

		response, err := json.Marshal(newBank)

		if err != nil {
			fmt.Println("Error", err)
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

func ShowBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		var bank []models.Bank
		initializers.DB.Find(&bank, id)

		response, err := json.Marshal(bank)

		if err != nil {
			fmt.Println("Error", err)
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

func DeleteBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		var bank []models.Bank
		initializers.DB.Find(&bank, id)

		response, err := json.Marshal(bank)

		if err != nil {
			fmt.Println("Error", err)
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

func UpdateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		body := models.Bank{}

		json.NewDecoder(r.Body).Decode(&body)

		var bank []models.Bank

		initializers.DB.Find(&bank, id)

		initializers.DB.Model(&bank).Updates(models.Bank{Name: body.Name, IBAN: body.IBAN})

		response, err := json.Marshal(bank)

		if err != nil {
			fmt.Println("Error", err)
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
