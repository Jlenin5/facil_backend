package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type VacationBalanceHandler struct {
	VacationBalanceUC *usecaseHumanresources.VacationBalanceUseCase
}

func NewVacationBalanceHandler(VacationBalanceUC *usecaseHumanresources.VacationBalanceUseCase) *VacationBalanceHandler {
	return &VacationBalanceHandler{VacationBalanceUC: VacationBalanceUC}
}

func (h *VacationBalanceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var balance humanresources.VacationBalance
	if err := json.NewDecoder(r.Body).Decode(&balance); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.VacationBalanceUC.Create(r.Context(), &balance); err != nil {
		http.Error(w, "Failed to create vacation balance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vacation balance created successfully"})
}

func (h *VacationBalanceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.VacationBalanceUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *VacationBalanceHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	request, err := h.VacationBalanceUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}

func (h *VacationBalanceHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var request humanresources.VacationBalance
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	if err := h.VacationBalanceUC.Update(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request updated successfully"})
}

func (h *VacationBalanceHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
		return
	}

	if err := h.VacationBalanceUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete holiday", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Holiday deleted successfully"})
}

func (h *VacationBalanceHandler) GetByEmployeeAndYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeID, err := strconv.Atoi(vars["employeeId"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	balance, err := h.VacationBalanceUC.GetByEmployeeAndYear(r.Context(), employeeID, year)
	if err != nil {
		http.Error(w, "Failed to fetch vacation balance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}