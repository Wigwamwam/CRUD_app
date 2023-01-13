package handlers

// TODO: imrpove error handling

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/wigwamwam/CRUD_app/models"
	"github.com/wigwamwam/CRUD_app/repository"
)

type errorResponse struct {
	Message string
}

type Handler struct {
	db *repository.DB
}

func NewHandler(db *repository.DB) Handler {
	return Handler{db: db}
}

func (h *Handler) HandlerIndexBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allBanks, err := h.db.SelectAllBanks()
		if err != nil {
			handleAppError(w, err)
			return
		}

		js, err := json.Marshal(allBanks)
		if err != nil {
			handleAppError(w, err)
			return
		}

		respondWithJSON(w, http.StatusOK, js)
	}
}

func (h *Handler) CreateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bank := models.Bank{}

		bytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			handleAppError(w, err)
			return
		}

		err = json.Unmarshal(bytes, &bank)
		if err != nil {
			handleAppError(w, err)
			return
		}

		bankPayload := models.Bank{Name: bank.Name, IBAN: bank.IBAN}

		createdBank, err := h.db.InsertBank(bankPayload)
		if err != nil {
			handleAppError(w, err)
			return
		}

		js, err := json.Marshal(createdBank)
		if err != nil {
			handleAppError(w, err)
			return
		}

		respondWithJSON(w, http.StatusCreated, js)
	}
}

func (h *Handler) ShowBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, err)
			return
		}

		var bankByID models.Bank

		bankByID, err = h.db.SelectBankByID(id)
		if err != nil {
			handleAppError(w, err)
			return
		}

		js, err := json.Marshal(bankByID)
		if err != nil {
			handleAppError(w, err)
			return
		}

		respondWithJSON(w, http.StatusOK, js)
	}
}

func (h *Handler) DeleteBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, err := strconv.Atoi(chi.URLParam(r, "id"))
		fmt.Println(err)
		if err != nil {
			respondWithError(w, http.StatusNoContent, err)
			return
		}

		err = h.db.DeleteBankByID(idParam)
		if err != nil {
			handleAppError(w, err)
			return
		}
		// not sure what to put here:
		js, err := json.Marshal("successful deleted")
		if err != nil {
			handleAppError(w, err)
			return
		}

		respondWithJSON(w, http.StatusNoContent, js)
	}
}

func (h *Handler) UpdateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, http.StatusNotFound, err)
			return
		}

		bank := models.Bank{}

		bytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			handleAppError(w, err)
			return
		}

		err = json.Unmarshal(bytes, &bank)
		if err != nil {
			handleAppError(w, err)
			return
		}

		updatedBank, err := h.db.UpdateBank(idParam, bank)
		if err != nil {
			handleAppError(w, err)
			return
		}

		js, err := json.Marshal(updatedBank)
		if err != nil {
			handleAppError(w, err)
			return
		}

		respondWithJSON(w, http.StatusOK, js)
	}
}
