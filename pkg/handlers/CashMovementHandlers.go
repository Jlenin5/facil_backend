package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type CashMovementHandler struct {
	CashMovementUC *usecase.CashMovementUseCase
}

func NewCashMovementHandler(CashMovementUC *usecase.CashMovementUseCase) *CashMovementHandler {
	return &CashMovementHandler{CashMovementUC: CashMovementUC}
}

func (h *CashMovementHandler) CreateCashMovement(w http.ResponseWriter, r *http.Request) {
	var cashMovement domain.CashMovements
	err := json.NewDecoder(r.Body).Decode(&cashMovement) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva movimiento de caja
	err = h.CashMovementUC.CreateCashMovement(&cashMovement)
	if err != nil {
		http.Error(w, "Failed to create cashMovement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "CashMovement created successfully"})
}

func (h *CashMovementHandler) GetAllCashMovements(w http.ResponseWriter, r *http.Request) {
	cashMovements, err := h.CashMovementUC.GetAllCashMovements()
	if err != nil {
		http.Error(w, "Failed to fetch cashMovements", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cashMovements)
}

func (h *CashMovementHandler) GetCashMovementById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid cashMovement ID", http.StatusBadRequest)
		return
	}

	cashMovement, err := h.CashMovementUC.GetCashMovementById(id)
	if err != nil {
		http.Error(w, "CashMovement not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cashMovement)
}

func (h *CashMovementHandler) UpdateCashMovement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid cashMovement ID", http.StatusBadRequest)
		return
	}

	var cashMovement domain.CashMovements
	err = json.NewDecoder(r.Body).Decode(&cashMovement) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	cashMovement.Id = id

	// Llama al caso de uso para actualizar la movimiento de caja
	err = h.CashMovementUC.UpdateCashMovement(&cashMovement)
	if err != nil {
		http.Error(w, "Failed to update cashMovement: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "CashMovement updated successfully"})
}