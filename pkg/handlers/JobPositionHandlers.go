package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type JobPositionHandler struct {
	JobPositionUC *usecase.JobPositionUseCase
}

func NewJobPositionHandler(JobPositionUC *usecase.JobPositionUseCase) *JobPositionHandler {
	return &JobPositionHandler{JobPositionUC: JobPositionUC}
}

func (h *JobPositionHandler) CreateJobPosition(w http.ResponseWriter, r *http.Request) {
	var jobPosition domain.JobPositions
	err := json.NewDecoder(r.Body).Decode(&jobPosition) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.JobPositionUC.CreateJobPosition(&jobPosition)
	if err != nil {
		http.Error(w, "Failed to create jobPosition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area created successfully"})
}

func (h *JobPositionHandler) GetAllJobPositions(w http.ResponseWriter, r *http.Request) {
	jobPositions, err := h.JobPositionUC.GetAllJobPositions()
	if err != nil {
		http.Error(w, "Failed to fetch work areas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobPositions)
}

func (h *JobPositionHandler) GetJobPositionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid jobPosition ID", http.StatusBadRequest)
		return
	}

	jobPosition, err := h.JobPositionUC.GetJobPositionById(id)
	if err != nil {
		http.Error(w, "Work area not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobPosition)
}

func (h *JobPositionHandler) UpdateJobPosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid jobPosition ID", http.StatusBadRequest)
		return
	}

	var jobPosition domain.JobPositions
	err = json.NewDecoder(r.Body).Decode(&jobPosition) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	jobPosition.Id = id

	// Validación básica de campos obligatorios
	if jobPosition.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la marca
	err = h.JobPositionUC.UpdateJobPosition(&jobPosition)
	if err != nil {
		http.Error(w, "Failed to update jobPosition: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area updated successfully"})
}

func (h *JobPositionHandler) DeleteJobPositionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid jobPosition Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.JobPositionUC.DeleteJobPositionById(id)
	if err != nil {
		http.Error(w, "Failed to delete jobPosition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area deleted successfully"})
}

func (h *JobPositionHandler) DeleteJobPositionsByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.JobPositionUC.DeleteJobPositionsByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete work areas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "JobPositions deleted successfully"})
}