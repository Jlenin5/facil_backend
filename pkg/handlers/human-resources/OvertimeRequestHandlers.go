package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type OvertimeRequestHandler struct {
	OvertimeRequestUC *usecaseHumanresources.OvertimeRequestUseCase
}

func NewOvertimeRequestHandler(OvertimeRequestUC *usecaseHumanresources.OvertimeRequestUseCase) *OvertimeRequestHandler {
	return &OvertimeRequestHandler{OvertimeRequestUC: OvertimeRequestUC}
}

func (h *OvertimeRequestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request humanresources.OvertimeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if request.EmployeeId == 0 || request.Date.IsZero() || request.Reason == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if request.EndTime <= request.StartTime {
		http.Error(w, "End time must be after start time", http.StatusBadRequest)
		return
	}

	if err := h.OvertimeRequestUC.Create(r.Context(), &request); err != nil {
		http.Error(w, "Failed to create overtime request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Overtime request created successfully"})
}

func (h *OvertimeRequestHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.OvertimeRequestUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *OvertimeRequestHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	request, err := h.OvertimeRequestUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}

func (h *OvertimeRequestHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var request humanresources.OvertimeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	if err := h.OvertimeRequestUC.Update(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request updated successfully"})
}

func (h *OvertimeRequestHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
		return
	}

	if err := h.OvertimeRequestUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete holiday", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "OvertimeRequest deleted successfully"})
}

func (h *OvertimeRequestHandler) Approve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid overtime request ID", http.StatusBadRequest)
		return
	}

	var body struct {
		ApprovedBy int `json:"approvedBy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.OvertimeRequestUC.Approve(r.Context(), id, body.ApprovedBy); err != nil {
		http.Error(w, "Failed to approve overtime request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Overtime request approved successfully"})
}