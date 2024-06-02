package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"weatherCheck/internal/usecase"
	"weatherCheck/internal/usecase/business_errors"
)

func GetTemperaturesHandler(w http.ResponseWriter, r *http.Request) {
	cep := strings.TrimPrefix(r.URL.Path, "/")

	dto, err := usecase.Execute(cep)
	if err != nil {
		if errors.Is(err, business_errors.ErrCepValidationFailed) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if errors.Is(err, business_errors.ErrCepNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if errors.Is(err, business_errors.ErrFetchTemperatureFailed) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	err = json.NewEncoder(w).Encode(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
