package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	customErrors "github.com/wigwamwam/CRUD_app/repository/errors"
)

// to do:
// - build out more handle app error

func internalServerError(w http.ResponseWriter, err error) {
	errorResponse := errorResponse{fmt.Sprintf("%v", err)}
	response, _ := json.Marshal(errorResponse)
	respondWithJSON(w, http.StatusInternalServerError, response)
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	errorResponse := errorResponse{fmt.Sprintf("%v", err)}
	response, _ := json.Marshal(errorResponse)
	respondWithJSON(w, code, response)
}

func respondWithJSON(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		internalServerError(w, err)
	}
}

func handleAppError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *customErrors.NotFoundError:
		respondWithError(w, http.StatusNotFound, err)
		return
	case *customErrors.IdNotFoundError:
		respondWithError(w, http.StatusNotFound, err)
		return
	case *customErrors.DeletingBankError:
		respondWithError(w, http.StatusNotFound, err)
		return
	case *customErrors.CreatingBankError:
		respondWithError(w, http.StatusInternalServerError, err)
		return
	case *customErrors.ScanningIdError:
		respondWithError(w, http.StatusInternalServerError, err)
		return
	default:
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
}
