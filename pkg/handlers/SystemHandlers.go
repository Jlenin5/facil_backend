package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type SystemHandler struct {
	SystemUC *usecase.SystemUseCase
}

func NewSystemHandler(SystemUC *usecase.SystemUseCase) *SystemHandler {
	return &SystemHandler{SystemUC: SystemUC}
}

func (h *SystemHandler) CreateSystem(w http.ResponseWriter, r *http.Request) {
	var system domain.Systems
	err := json.NewDecoder(r.Body).Decode(&system) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.SystemUC.CreateSystem(&system)
	if err != nil {
		http.Error(w, "Failed to create system", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area created successfully"})
}

func (h *SystemHandler) GetAllSystems(w http.ResponseWriter, r *http.Request) {
	systems, err := h.SystemUC.GetAllSystems()
	if err != nil {
		http.Error(w, "Failed to fetch work areas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(systems)
}

func (h *SystemHandler) GetSystemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid system ID", http.StatusBadRequest)
		return
	}

	system, err := h.SystemUC.GetSystemById(id)
	if err != nil {
		http.Error(w, "Work area not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(system)
}

func (h *SystemHandler) UpdateSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid system ID", http.StatusBadRequest)
		return
	}

	var system domain.Systems
	err = json.NewDecoder(r.Body).Decode(&system) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	system.Id = id

	// Validación básica de campos obligatorios
	if system.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la marca
	err = h.SystemUC.UpdateSystem(&system)
	if err != nil {
		http.Error(w, "Failed to update system: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area updated successfully"})
}