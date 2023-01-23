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

type Handler struct {
	DAO repository.DAO
}

// create a new instance for the Handler struct
// dependency injection - only use when there are methods on the struct
// decoupled the handler from the database using an interface
func NewHandler(dao repository.DAO) Handler {
	return Handler{DAO: dao}
}

func (h *Handler) HandlerIndexBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allBanks, err := h.DAO.SelectAllBanks()
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

		createdBank, err := h.DAO.InsertBank(bankPayload)
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

		bankByID, err = h.DAO.SelectBankByID(id)
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

		err = h.DAO.DeleteBankByID(idParam)
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

		updatedBank, err := h.DAO.UpdateBank(idParam, bank)
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
