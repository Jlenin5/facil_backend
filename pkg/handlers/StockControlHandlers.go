package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type StockControlHandler struct {
	StockControlUC *usecase.StockControlUseCase
}

func NewStockControlHandler(StockControlUC *usecase.StockControlUseCase) *StockControlHandler {
	return &StockControlHandler{StockControlUC: StockControlUC}
}

func (h *StockControlHandler) CreateStockControl(w http.ResponseWriter, r *http.Request) {
	var stockControl domain.StockControl
	err := json.NewDecoder(r.Body).Decode(&stockControl) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.StockControlUC.CreateStockControl(&stockControl)
	if err != nil {
		http.Error(w, "Failed to create stock control", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Stock control created successfully"})
}

func (h *StockControlHandler) GetAllStockControl(w http.ResponseWriter, r *http.Request) {
	stockControls, err := h.StockControlUC.GetAllStockControl()
	if err != nil {
		http.Error(w, "Failed to fetch stock control", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stockControls)
}

func (h *StockControlHandler) GetStockControlById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid stockControl ID", http.StatusBadRequest)
		return
	}

	stockControl, err := h.StockControlUC.GetStockControlById(id)
	if err != nil {
		http.Error(w, "StockControl not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stockControl)
}

func (h *StockControlHandler) UpdateStockControl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid stockControl ID", http.StatusBadRequest)
		return
	}

	var stockControl domain.StockControl
	err = json.NewDecoder(r.Body).Decode(&stockControl) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	stockControl.Id = id

	// Llama al caso de uso para actualizar la marca
	err = h.StockControlUC.UpdateStockControl(&stockControl)
	if err != nil {
		http.Error(w, "Failed to update stockControl: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "StockControl updated successfully"})
}

func (h *StockControlHandler) DeleteStockControlById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid stockControl Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.StockControlUC.DeleteStockControlById(id)
	if err != nil {
		http.Error(w, "Failed to delete stock control", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "StockControl deleted successfully"})
}

func (h *StockControlHandler) DeleteStockControlByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.StockControlUC.DeleteStockControlByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete stock control", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Stock control deleted successfully"})
}
