package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type UnitOfMeasurementHandler struct {
	UnitOfMeasurementUC *usecase.UnitOfMeasurementUseCase
}

func NewUnitOfMeasurementHandler(UnitOfMeasurementUC *usecase.UnitOfMeasurementUseCase) *UnitOfMeasurementHandler {
	return &UnitOfMeasurementHandler{UnitOfMeasurementUC: UnitOfMeasurementUC}
}

func (h *UnitOfMeasurementHandler) CreateUnitOfMeasurement(w http.ResponseWriter, r *http.Request) {
	var brand domain.UnitsOfMeasurement
	err := json.NewDecoder(r.Body).Decode(&brand) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.UnitOfMeasurementUC.CreateUnitOfMeasurement(&brand)
	if err != nil {
		http.Error(w, "Failed to create brand", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "UnitOfMeasurement created successfully"})
}

func (h *UnitOfMeasurementHandler) GetAllUnitsOfMeasurement(w http.ResponseWriter, r *http.Request) {
	brands, err := h.UnitOfMeasurementUC.GetAllUnitsOfMeasurement()
	if err != nil {
		http.Error(w, "Failed to fetch brands", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brands)
}

func (h *UnitOfMeasurementHandler) GetUnitOfMeasurementById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid brand ID", http.StatusBadRequest)
		return
	}

	brand, err := h.UnitOfMeasurementUC.GetUnitOfMeasurementById(id)
	if err != nil {
		http.Error(w, "UnitOfMeasurement not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brand)
}

func (h *UnitOfMeasurementHandler) UpdateUnitOfMeasurement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid brand ID", http.StatusBadRequest)
		return
	}

	var brand domain.UnitsOfMeasurement
	err = json.NewDecoder(r.Body).Decode(&brand) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	brand.Id = id

	// Validación básica de campos obligatorios
	if brand.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la marca
	err = h.UnitOfMeasurementUC.UpdateUnitOfMeasurement(&brand)
	if err != nil {
		http.Error(w, "Failed to update brand: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "UnitOfMeasurement updated successfully"})
}

func (h *UnitOfMeasurementHandler) DeleteUnitOfMeasurementById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid brand Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.UnitOfMeasurementUC.DeleteUnitOfMeasurementById(id)
	if err != nil {
		http.Error(w, "Failed to delete brand", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "UnitOfMeasurement deleted successfully"})
}

func (h *UnitOfMeasurementHandler) DeleteUnitsOfMeasurementByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.UnitOfMeasurementUC.DeleteUnitsOfMeasurementByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete brands", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "UnitsOfMeasurement deleted successfully"})
}
