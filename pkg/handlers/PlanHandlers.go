package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type PlanHandler struct {
	PlanUC *usecase.PlanUseCase
}

func NewPlanHandler(PlanUC *usecase.PlanUseCase) *PlanHandler {
	return &PlanHandler{PlanUC: PlanUC}
}

func (h *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {
	var plan domain.Plans
	err := json.NewDecoder(r.Body).Decode(&plan) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo plan
	err = h.PlanUC.CreatePlan(&plan)
	if err != nil {
		http.Error(w, "Failed to create plan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Plan created successfully"})
}

func (h *PlanHandler) GetAllPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := h.PlanUC.GetAllPlans()
	if err != nil {
		http.Error(w, "Failed to fetch plans", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

func (h *PlanHandler) GetPlanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid plan ID", http.StatusBadRequest)
		return
	}

	plan, err := h.PlanUC.GetPlanById(id)
	if err != nil {
		http.Error(w, "Plan not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}

func (h *PlanHandler) UpdatePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid plan ID", http.StatusBadRequest)
		return
	}

	var plan domain.Plans
	err = json.NewDecoder(r.Body).Decode(&plan) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	plan.Id = id

	// Llama al caso de uso para actualizar la plan
	err = h.PlanUC.UpdatePlan(&plan)
	if err != nil {
		http.Error(w, "Failed to update plan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Plan updated successfully"})
}

func (h *PlanHandler) DeletePlanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid plan Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la plan por ID
	err = h.PlanUC.DeletePlanById(id)
	if err != nil {
		http.Error(w, "Failed to delete plan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Plan deleted successfully"})
}

func (h *PlanHandler) DeletePlansByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las plans por sus IDs
	err = h.PlanUC.DeletePlansByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete plans", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Plans deleted successfully"})
}
