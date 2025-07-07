package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type EmployeeBenefitHandler struct {
	EmployeeBenefitUC *usecaseHumanresources.EmployeeBenefitUseCase
}

func NewEmployeeBenefitHandler(EmployeeBenefitUC *usecaseHumanresources.EmployeeBenefitUseCase) *EmployeeBenefitHandler {
	return &EmployeeBenefitHandler{EmployeeBenefitUC: EmployeeBenefitUC}
}

func (h *EmployeeBenefitHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request humanresources.EmployeeBenefit
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.EmployeeBenefitUC.Create(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Absence request created successfully",
		"id":      request.Id,
	})
}

func (h *EmployeeBenefitHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.EmployeeBenefitUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *EmployeeBenefitHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	request, err := h.EmployeeBenefitUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}

func (h *EmployeeBenefitHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee benefit ID", http.StatusBadRequest)
		return
	}

	var benefit humanresources.EmployeeBenefit
	if err := json.NewDecoder(r.Body).Decode(&benefit); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	benefit.Id = id
	if err := h.EmployeeBenefitUC.Update(r.Context(), &benefit); err != nil {
		http.Error(w, "Failed to update employee benefit", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee benefit updated successfully"})
}

func (h *EmployeeBenefitHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	if err := h.EmployeeBenefitUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request deleted successfully"})
}