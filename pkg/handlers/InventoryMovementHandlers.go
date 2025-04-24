package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type InventoryMovementHandler struct {
	InventoryMovementUC *usecase.InventoryMovementUseCase
}

func NewInventoryMovementHandler(InventoryMovementUC *usecase.InventoryMovementUseCase) *InventoryMovementHandler {
	return &InventoryMovementHandler{InventoryMovementUC: InventoryMovementUC}
}

func (h *InventoryMovementHandler) CreateInventoryMovement(w http.ResponseWriter, r *http.Request) {
	var inventoryMovement domain.InventoryMovements
	err := json.NewDecoder(r.Body).Decode(&inventoryMovement) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.InventoryMovementUC.CreateInventoryMovement(&inventoryMovement)
	if err != nil {
		http.Error(w, "Failed to create inventory movement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inventory movement created successfully"})
}

func (h *InventoryMovementHandler) GetAllInventoryMovements(w http.ResponseWriter, r *http.Request) {
	inventoryMovements, err := h.InventoryMovementUC.GetAllInventoryMovements()
	if err != nil {
		http.Error(w, "Failed to fetch inventory movements", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryMovements)
}

func (h *InventoryMovementHandler) GetInventoryMovementById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid inventory movement ID", http.StatusBadRequest)
		return
	}

	inventoryMovement, err := h.InventoryMovementUC.GetInventoryMovementById(id)
	if err != nil {
		http.Error(w, "InventoryMovement not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryMovement)
}

func (h *InventoryMovementHandler) UpdateInventoryMovement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid inventoryMovement ID", http.StatusBadRequest)
		return
	}

	var inventoryMovement domain.InventoryMovements
	err = json.NewDecoder(r.Body).Decode(&inventoryMovement) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	inventoryMovement.Id = id

	// Llama al caso de uso para actualizar la marca
	err = h.InventoryMovementUC.UpdateInventoryMovement(&inventoryMovement)
	if err != nil {
		http.Error(w, "Failed to update inventory movement: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inventory movement updated successfully"})
}

func (h *InventoryMovementHandler) DeleteInventoryMovementById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid inventoryMovement Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.InventoryMovementUC.DeleteInventoryMovementById(id)
	if err != nil {
		http.Error(w, "Failed to delete inventory movement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inventory movement deleted successfully"})
}

func (h *InventoryMovementHandler) DeleteInventoryMovementsByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.InventoryMovementUC.DeleteInventoryMovementsByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete inventoryMovements", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inventory movements deleted successfully"})
}
