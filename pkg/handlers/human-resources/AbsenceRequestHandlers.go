package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type AbsenceRequestHandler struct {
	AbsenceRequestUC *usecaseHumanresources.AbsenceRequestUseCase
}

func NewAbsenceRequestHandler(AbsenceRequestUC *usecaseHumanresources.AbsenceRequestUseCase) *AbsenceRequestHandler {
	return &AbsenceRequestHandler{AbsenceRequestUC: AbsenceRequestUC}
}

func (h *AbsenceRequestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request humanresources.AbsenceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.AbsenceRequestUC.Create(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Absence request created successfully",
		"id":      request.Id,
	})
}

func (h *AbsenceRequestHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	requests, err := h.AbsenceRequestUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if requests == nil {
		requests = []humanresources.AbsenceRequest{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

func (h *AbsenceRequestHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	request, err := h.AbsenceRequestUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}

func (h *AbsenceRequestHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var request humanresources.AbsenceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	if err := h.AbsenceRequestUC.Update(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request updated successfully"})
}

func (h *AbsenceRequestHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	if err := h.AbsenceRequestUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request deleted successfully"})
}

func (h *AbsenceRequestHandler) Approve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var body struct {
		ApprovedBy int `json:"approvedBy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.AbsenceRequestUC.Approve(r.Context(), id, body.ApprovedBy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request approved successfully"})
}

func (h *AbsenceRequestHandler) Reject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var body struct {
		ApprovedBy int    `json:"approvedBy"`
		Reason     string `json:"reason"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.AbsenceRequestUC.Reject(r.Context(), id, body.ApprovedBy, body.Reason); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request rejected successfully"})
}