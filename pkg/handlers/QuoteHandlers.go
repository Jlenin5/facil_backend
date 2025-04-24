package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type QuoteHandler struct {
	QuoteUC *usecase.QuoteUseCase
}

func NewQuoteHandler(QuoteUC *usecase.QuoteUseCase) *QuoteHandler {
	return &QuoteHandler{QuoteUC: QuoteUC}
}

// Crear una cotización
func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote domain.Quotes
	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validar fechas vacías
	if quote.Expiration_Date.Valid && quote.Expiration_Date.Time.IsZero() {
		quote.Expiration_Date.Valid = false
	}
	if quote.Approved_At.Valid && quote.Approved_At.Time.IsZero() {
		quote.Approved_At.Valid = false
	}
	if quote.Canceled_At.Valid && quote.Canceled_At.Time.IsZero() {
		quote.Canceled_At.Valid = false
	}

	for i := range quote.QuoteDetails {
		quote.QuoteDetails[i].Id = 0
	}

	err = h.QuoteUC.CreateQuote(&quote, quote.QuoteDetails)
	if err != nil {
		http.Error(w, "Failed to create quote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Quote created successfully"})
}

// Obtener todas las órdenes de venta
func (h *QuoteHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := h.QuoteUC.GetAllQuotes()
	if err != nil {
		http.Error(w, "Failed to fetch quotes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

// Obtener una cotización por ID
func (h *QuoteHandler) GetQuoteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid quote Id", http.StatusBadRequest)
		return
	}

	order, err := h.QuoteUC.GetQuoteById(id)
	if err != nil {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Actualizar una cotización
func (h *QuoteHandler) UpdateQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid quote ID", http.StatusBadRequest)
		return
	}

	var quote domain.Quotes
	err = json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	quote.Id = id

	err = h.QuoteUC.UpdateQuote(&quote, quote.QuoteDetails)
	if err != nil {
		http.Error(w, "Failed to update quote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Quote updated successfully"})
}
