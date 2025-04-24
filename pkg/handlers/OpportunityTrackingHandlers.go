package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type OpportunityTrackingHandler struct {
	OpportunityTrackingUC *usecase.OpportunityTrackingUseCase
}

func NewOpportunityTrackingHandler(OpportunityTrackingUC *usecase.OpportunityTrackingUseCase) *OpportunityTrackingHandler {
	return &OpportunityTrackingHandler{OpportunityTrackingUC: OpportunityTrackingUC}
}

func (h *OpportunityTrackingHandler) CreateOpportunityTracking(w http.ResponseWriter, r *http.Request) {
	var opportunityTracking domain.OpportunityTracking
	err := json.NewDecoder(r.Body).Decode(&opportunityTracking) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.OpportunityTrackingUC.CreateOpportunityTracking(&opportunityTracking)
	if err != nil {
		http.Error(w, "Failed to create opportunityTracking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "OpportunityTracking created successfully"})
}

func (h *OpportunityTrackingHandler) GetAllOpportunityTracking(w http.ResponseWriter, r *http.Request) {
	opportunityTrackings, err := h.OpportunityTrackingUC.GetAllOpportunityTracking()
	if err != nil {
		http.Error(w, "Failed to fetch opportunityTrackings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(opportunityTrackings)
}

func (h *OpportunityTrackingHandler) GetOpportunityTrackingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid opportunityTracking ID", http.StatusBadRequest)
		return
	}

	opportunityTracking, err := h.OpportunityTrackingUC.GetOpportunityTrackingById(id)
	if err != nil {
		http.Error(w, "OpportunityTracking not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(opportunityTracking)
}

func (h *OpportunityTrackingHandler) UpdateOpportunityTracking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid opportunityTracking ID", http.StatusBadRequest)
		return
	}

	var opportunityTracking domain.OpportunityTracking
	err = json.NewDecoder(r.Body).Decode(&opportunityTracking) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	opportunityTracking.Id = id

	// Llama al caso de uso para actualizar la marca
	err = h.OpportunityTrackingUC.UpdateOpportunityTracking(&opportunityTracking)
	if err != nil {
		http.Error(w, "Failed to update opportunityTracking: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "OpportunityTracking updated successfully"})
}

func (h *OpportunityTrackingHandler) DeleteOpportunityTrackingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid opportunityTracking Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.OpportunityTrackingUC.DeleteOpportunityTrackingById(id)
	if err != nil {
		http.Error(w, "Failed to delete opportunityTracking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "OpportunityTracking deleted successfully"})
}

func (h *OpportunityTrackingHandler) DeleteOpportunityTrackingByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.OpportunityTrackingUC.DeleteOpportunityTrackingByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete opportunityTrackings", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "OpportunityTracking deleted successfully"})
}
