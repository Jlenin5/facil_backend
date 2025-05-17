package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/Jlenin5/facil_backend/pkg/middleware"
	"github.com/gorilla/mux"
)

type ExchangeRateHandler struct {
	ExchangeRateUC *usecase.ExchangeRateUseCase
}

func NewExchangeRateHandler(ExchangeRateUC *usecase.ExchangeRateUseCase) *ExchangeRateHandler {
	return &ExchangeRateHandler{ExchangeRateUC: ExchangeRateUC}
}

func (h *ExchangeRateHandler) CreateExchangeRate(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos del usuario del contexto
	userData, ok := r.Context().Value(middleware.UserContextKey).(map[string]interface{})
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validar user_id del usuario autenticado
	userIDRaw, exists := userData["id"]
	if !exists {
		http.Error(w, "User ID not found", http.StatusBadRequest)
		return
	}

	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	// Struct auxiliar para capturar el JSON que viene desde el frontend
	var input struct {
		Base_Currency_Id   int    `json:"base_currency_id"`
		Target_Currency_Id int    `json:"target_currency_id"`
		Exchange_Rate      string `json:"exchange_rate"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Convertir Exchange_Rate string a float64
	exchangeRateValue, err := strconv.ParseFloat(input.Exchange_Rate, 64)
	if err != nil {
		http.Error(w, "Invalid exchange rate format", http.StatusBadRequest)
		return
	}

	// Armar la entidad real
	exchangeRate := domain.ExchangeRates{
		Base_Currency_Id:   input.Base_Currency_Id,
		Target_Currency_Id: input.Target_Currency_Id,
		Exchange_Rate:      exchangeRateValue,
		Created_By:         int(userIDFloat),
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

func ptrInt64(i int64) *int64 {
	return &i
}
func (h *ExchangeRateHandler) UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid exchangeRate ID", http.StatusBadRequest)
		return
	}

	// Obtener los datos del usuario del contexto
	userData, ok := r.Context().Value(middleware.UserContextKey).(map[string]interface{})
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validar user_id del usuario autenticado
	userIDRaw, exists := userData["id"]
	if !exists {
		http.Error(w, "User ID not found", http.StatusBadRequest)
		return
	}

	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	userID := int64(userIDFloat)

	// Struct auxiliar para capturar el JSON que viene desde el frontend
	var input struct {
		Base_Currency_Id   int    `json:"base_currency_id"`
		Target_Currency_Id int    `json:"target_currency_id"`
		Exchange_Rate      string `json:"exchange_rate"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Convertir Exchange_Rate string a float64
	exchangeRateValue, err := strconv.ParseFloat(input.Exchange_Rate, 64)
	if err != nil {
		http.Error(w, "Invalid exchange rate format", http.StatusBadRequest)
		return
	}

	// Armar la entidad real
	exchangeRate := domain.ExchangeRates{
		Base_Currency_Id:   input.Base_Currency_Id,
		Target_Currency_Id: input.Target_Currency_Id,
		Exchange_Rate:      exchangeRateValue,
		Updated_By:         domain.NullInt{Int: ptrInt64(userID), Valid: true},
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
