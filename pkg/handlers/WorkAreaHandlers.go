package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type WorkAreaHandler struct {
	WorkAreaUC *usecase.WorkAreaUseCase
}

func NewWorkAreaHandler(WorkAreaUC *usecase.WorkAreaUseCase) *WorkAreaHandler {
	return &WorkAreaHandler{WorkAreaUC: WorkAreaUC}
}

func (h *WorkAreaHandler) CreateWorkArea(w http.ResponseWriter, r *http.Request) {
	var workArea domain.WorkAreas
	err := json.NewDecoder(r.Body).Decode(&workArea) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.WorkAreaUC.CreateWorkArea(&workArea)
	if err != nil {
		http.Error(w, "Failed to create workArea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area created successfully"})
}

func (h *WorkAreaHandler) GetAllWorkAreas(w http.ResponseWriter, r *http.Request) {
	workAreas, err := h.WorkAreaUC.GetAllWorkAreas()
	if err != nil {
		http.Error(w, "Failed to fetch work areas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workAreas)
}

func (h *WorkAreaHandler) GetWorkAreaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid workArea ID", http.StatusBadRequest)
		return
	}

	workArea, err := h.WorkAreaUC.GetWorkAreaById(id)
	if err != nil {
		http.Error(w, "Work area not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workArea)
}

func (h *WorkAreaHandler) UpdateWorkArea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid workArea ID", http.StatusBadRequest)
		return
	}

	var workArea domain.WorkAreas
	err = json.NewDecoder(r.Body).Decode(&workArea) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	workArea.Id = id

	// Validación básica de campos obligatorios
	if workArea.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la marca
	err = h.WorkAreaUC.UpdateWorkArea(&workArea)
	if err != nil {
		http.Error(w, "Failed to update workArea: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area updated successfully"})
}

func (h *WorkAreaHandler) DeleteWorkAreaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid workArea Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.WorkAreaUC.DeleteWorkAreaById(id)
	if err != nil {
		http.Error(w, "Failed to delete workArea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area deleted successfully"})
}

func (h *WorkAreaHandler) DeleteWorkAreasByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.WorkAreaUC.DeleteWorkAreasByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete work areas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "WorkAreas deleted successfully"})
}