package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type VacationHandler struct {
	VacationUC *usecaseHumanresources.VacationUseCase
}

func NewVacationHandler(VacationUC *usecaseHumanresources.VacationUseCase) *VacationHandler {
	return &VacationHandler{VacationUC: VacationUC}
}

func (h *VacationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var vacation humanresources.Vacation
	if err := json.NewDecoder(r.Body).Decode(&vacation); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if vacation.EndDate.Before(vacation.StartDate) {
		http.Error(w, "End date must be after start date", http.StatusBadRequest)
		return
	}

	if err := h.VacationUC.Create(r.Context(), &vacation); err != nil {
		http.Error(w, "Failed to create vacation request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vacation request created successfully"})
}

func (h *VacationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.VacationUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *VacationHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	request, err := h.VacationUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}

func (h *VacationHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var request humanresources.Vacation
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	if err := h.VacationUC.Update(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request updated successfully"})
}

func (h *VacationHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid vacation request ID", http.StatusBadRequest)
		return
	}

	if err := h.VacationUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete vacation request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vacation request deleted successfully"})
}

func (h *VacationHandler) Approve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid vacation request ID", http.StatusBadRequest)
		return
	}

	var body struct {
		ApprovedBy int `json:"approvedBy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.VacationUC.Approve(r.Context(), id, body.ApprovedBy); err != nil {
		http.Error(w, "Failed to approve vacation request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vacation request approved successfully"})
}