package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type PerformanceReviewHandler struct {
	PerformanceReviewUC *usecaseHumanresources.PerformanceReviewUseCase
}

func NewPerformanceReviewHandler(PerformanceReviewUC *usecaseHumanresources.PerformanceReviewUseCase) *PerformanceReviewHandler {
	return &PerformanceReviewHandler{PerformanceReviewUC: PerformanceReviewUC}
}

func (h *PerformanceReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	var performanceReview humanresources.PerformanceReview
	if err := json.NewDecoder(r.Body).Decode(&performanceReview); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.PerformanceReviewUC.Create(r.Context(), &performanceReview); err != nil {
		http.Error(w, "Failed to create performanceReview", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Holiday created successfully"})
}

func (h *PerformanceReviewHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.PerformanceReviewUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *PerformanceReviewHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	request, err := h.PerformanceReviewUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}

func (h *PerformanceReviewHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var request humanresources.PerformanceReview
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	if err := h.PerformanceReviewUC.Update(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request updated successfully"})
}

func (h *PerformanceReviewHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
		return
	}

	if err := h.PerformanceReviewUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete holiday", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Holiday deleted successfully"})
}