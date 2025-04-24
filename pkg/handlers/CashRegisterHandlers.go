package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type CashRegisterHandler struct {
	CashRegisterUC *usecase.CashRegisterUseCase
}

func NewCashRegisterHandler(CashRegisterUC *usecase.CashRegisterUseCase) *CashRegisterHandler {
	return &CashRegisterHandler{CashRegisterUC: CashRegisterUC}
}

func (h *CashRegisterHandler) CreateCashRegister(w http.ResponseWriter, r *http.Request) {
	var cashRegister domain.CashRegisters
	err := json.NewDecoder(r.Body).Decode(&cashRegister) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva registro de caja
	err = h.CashRegisterUC.CreateCashRegister(&cashRegister)
	if err != nil {
		http.Error(w, "Failed to create cashRegister", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "CashRegister created successfully"})
}

func (h *CashRegisterHandler) GetAllCashRegisters(w http.ResponseWriter, r *http.Request) {
	cashRegisters, err := h.CashRegisterUC.GetAllCashRegisters()
	if err != nil {
		http.Error(w, "Failed to fetch cashRegisters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cashRegisters)
}

func (h *CashRegisterHandler) GetCashRegisterById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid cashRegister ID", http.StatusBadRequest)
		return
	}

	cashRegister, err := h.CashRegisterUC.GetCashRegisterById(id)
	if err != nil {
		http.Error(w, "CashRegister not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cashRegister)
}

func (h *CashRegisterHandler) UpdateCashRegister(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid cashRegister ID", http.StatusBadRequest)
		return
	}

	var cashRegister domain.CashRegisters
	err = json.NewDecoder(r.Body).Decode(&cashRegister) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	cashRegister.Id = id

	// Llama al caso de uso para actualizar la registro de caja
	err = h.CashRegisterUC.UpdateCashRegister(&cashRegister)
	if err != nil {
		http.Error(w, "Failed to update cashRegister: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "CashRegister updated successfully"})
}

func (h *CashRegisterHandler) DeleteCashRegisterById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid cashRegister Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la registro de caja por ID
	err = h.CashRegisterUC.DeleteCashRegisterById(id)
	if err != nil {
		http.Error(w, "Failed to delete cashRegister", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "CashRegister deleted successfully"})
}

func (h *CashRegisterHandler) DeleteCashRegistersByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las registro de cajas por sus IDs
	err = h.CashRegisterUC.DeleteCashRegistersByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete cashRegisters", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "CashRegisters deleted successfully"})
}