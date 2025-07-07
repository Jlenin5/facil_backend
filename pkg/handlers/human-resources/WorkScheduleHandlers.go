package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	usecaseHumanresources "github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type WorkScheduleHandler struct {
	WorkScheduleUC *usecaseHumanresources.WorkScheduleUseCase
}

func NewWorkScheduleHandler(WorkScheduleUC *usecaseHumanresources.WorkScheduleUseCase) *WorkScheduleHandler {
	return &WorkScheduleHandler{WorkScheduleUC: WorkScheduleUC}
}

func (h *WorkScheduleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var schedule humanresources.WorkSchedule
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if schedule.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if err := h.WorkScheduleUC.Create(r.Context(), &schedule); err != nil {
		http.Error(w, "Failed to create work schedule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work schedule created successfully"})
}

func (h *WorkScheduleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.WorkScheduleUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch work schedules", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

func (h *WorkScheduleHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid work schedule ID", http.StatusBadRequest)
		return
	}

	schedule, err := h.WorkScheduleUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, "Work schedule not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}

func (h *WorkScheduleHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var request humanresources.WorkSchedule
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	if err := h.WorkScheduleUC.Update(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request updated successfully"})
}

func (h *WorkScheduleHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid work schedule ID", http.StatusBadRequest)
		return
	}

	if err := h.WorkScheduleUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete work schedule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work schedule deleted successfully"})
}