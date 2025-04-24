package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type CurrencyHandler struct {
	CurrencyUC *usecase.CurrencyUseCase
}

func NewCurrencyHandler(currencyUC *usecase.CurrencyUseCase) *CurrencyHandler {
	return &CurrencyHandler{CurrencyUC: currencyUC}
}

func (h *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	var currency domain.Currencies
	err := json.NewDecoder(r.Body).Decode(&currency) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva moneda
	err = h.CurrencyUC.CreateCurrency(&currency)
	if err != nil {
		http.Error(w, "Failed to create currency", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Currency created successfully"})
}

func (h *CurrencyHandler) GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	categories, err := h.CurrencyUC.GetAllCurrencies()
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CurrencyHandler) GetCurrencyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid currency ID", http.StatusBadRequest)
		return
	}

	currency, err := h.CurrencyUC.GetCurrencyById(id)
	if err != nil {
		http.Error(w, "Currency not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currency)
}

func (h *CurrencyHandler) UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid currency ID", http.StatusBadRequest)
		return
	}

	var currency domain.Currencies
	err = json.NewDecoder(r.Body).Decode(&currency) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	currency.Id = id

	// Validación básica de campos obligatorios
	if currency.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la moneda
	err = h.CurrencyUC.UpdateCurrency(&currency)
	if err != nil {
		http.Error(w, "Failed to update currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Currency updated successfully"})
}

func (h *CurrencyHandler) DeleteCurrencyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid currency Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la moneda por ID
	err = h.CurrencyUC.DeleteCurrencyById(id)
	if err != nil {
		http.Error(w, "Failed to delete currency", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Currency deleted successfully"})
}

func (h *CurrencyHandler) DeleteCurrenciesByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las monedas por sus IDs
	err = h.CurrencyUC.DeleteCurrenciesByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete categories", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Currencies deleted successfully"})
}
