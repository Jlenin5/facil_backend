package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type WarehouseHandler struct {
	WarehouseUC *usecase.WarehouseUseCase
}

func NewWarehouseHandler(WarehouseUC *usecase.WarehouseUseCase) *WarehouseHandler {
	return &WarehouseHandler{WarehouseUC: WarehouseUC}
}

func (h *WarehouseHandler) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse domain.Warehouses
	err := json.NewDecoder(r.Body).Decode(&warehouse) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo almacén
	err = h.WarehouseUC.CreateWarehouse(&warehouse)
	if err != nil {
		http.Error(w, "Failed to create warehouse", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Warehouse created successfully"})
}

func (h *WarehouseHandler) GetAllWarehouses(w http.ResponseWriter, r *http.Request) {
	warehouses, err := h.WarehouseUC.GetAllWarehouses()
	if err != nil {
		http.Error(w, "Failed to fetch warehouses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(warehouses)
}

func (h *WarehouseHandler) GetWarehouseById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	warehouse, err := h.WarehouseUC.GetWarehouseById(id)
	if err != nil {
		http.Error(w, "Warehouse not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(warehouse)
}

func (h *WarehouseHandler) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	var warehouse domain.Warehouses
	err = json.NewDecoder(r.Body).Decode(&warehouse) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	warehouse.Id = id

	// Validación básica de campos obligatorios
	if warehouse.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar el almacén
	err = h.WarehouseUC.UpdateWarehouse(&warehouse)
	if err != nil {
		http.Error(w, "Failed to update warehouse: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Warehouse updated successfully"})
}

func (h *WarehouseHandler) DeleteWarehouseById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid warehouse Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar el almacén por ID
	err = h.WarehouseUC.DeleteWarehouseById(id)
	if err != nil {
		http.Error(w, "Failed to delete warehouse", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Warehouse deleted successfully"})
}

func (h *WarehouseHandler) DeleteWarehousesByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar los almacenes por sus IDs
	err = h.WarehouseUC.DeleteWarehousesByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete warehouses", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Warehouses deleted successfully"})
}
