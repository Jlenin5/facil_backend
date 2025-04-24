package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type ExchangeRateHandler struct {
	ExchangeRateUC *usecase.ExchangeRateUseCase
}

func NewExchangeRateHandler(ExchangeRateUC *usecase.ExchangeRateUseCase) *ExchangeRateHandler {
	return &ExchangeRateHandler{ExchangeRateUC: ExchangeRateUC}
}

func (h *ExchangeRateHandler) CreateExchangeRate(w http.ResponseWriter, r *http.Request) {
	var exchangeRate domain.ExchangeRates
	err := json.NewDecoder(r.Body).Decode(&exchangeRate) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.ExchangeRateUC.CreateExchangeRate(&exchangeRate)
	if err != nil {
		http.Error(w, "Failed to create exchangeRate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "ExchangeRate created successfully"})
}

func (h *ExchangeRateHandler) GetAllExchangeRates(w http.ResponseWriter, r *http.Request) {
	exchangeRates, err := h.ExchangeRateUC.GetAllExchangeRates()
	if err != nil {
		http.Error(w, "Failed to fetch exchangeRates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchangeRates)
}

func (h *ExchangeRateHandler) GetExchangeRateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid exchangeRate ID", http.StatusBadRequest)
		return
	}

	exchangeRate, err := h.ExchangeRateUC.GetExchangeRateById(id)
	if err != nil {
		http.Error(w, "ExchangeRate not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchangeRate)
}

func (h *ExchangeRateHandler) UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid exchangeRate ID", http.StatusBadRequest)
		return
	}

	var exchangeRate domain.ExchangeRates
	err = json.NewDecoder(r.Body).Decode(&exchangeRate) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	exchangeRate.Id = id

	// Llama al caso de uso para actualizar la marca
	err = h.ExchangeRateUC.UpdateExchangeRate(&exchangeRate)
	if err != nil {
		http.Error(w, "Failed to update exchangeRate: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ExchangeRate updated successfully"})
}

func (h *ExchangeRateHandler) DeleteExchangeRateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid exchangeRate Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.ExchangeRateUC.DeleteExchangeRateById(id)
	if err != nil {
		http.Error(w, "Failed to delete exchangeRate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ExchangeRate deleted successfully"})
}

func (h *ExchangeRateHandler) DeleteExchangeRatesByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.ExchangeRateUC.DeleteExchangeRatesByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete exchangeRates", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ExchangeRates deleted successfully"})
}
