package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type AttendanceHandler struct {
	AttendanceUC *usecaseHumanresources.AttendanceUseCase
}

func NewAttendanceHandler(AttendanceUC *usecaseHumanresources.AttendanceUseCase) *AttendanceHandler {
	return &AttendanceHandler{AttendanceUC: AttendanceUC}
}

func (h *AttendanceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var attendance humanresources.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.AttendanceUC.Create(r.Context(), &attendance); err != nil {
		http.Error(w, "Failed to create attendance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Attendance created successfully"})
}

func (h *AttendanceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.AttendanceUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *AttendanceHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence type ID", http.StatusBadRequest)
		return
	}

	absenceType, err := h.AttendanceUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, "Absence type not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(absenceType)
}

func (h *AttendanceHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid attendance ID", http.StatusBadRequest)
		return
	}

	var attendance humanresources.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	attendance.Id = id
	if err := h.AttendanceUC.Update(r.Context(), &attendance); err != nil {
		http.Error(w, "Failed to update attendance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Attendance updated successfully"})
}

func (h *AttendanceHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid attendance ID", http.StatusBadRequest)
		return
	}

	if err := h.AttendanceUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete attendance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Attendance deleted successfully"})
}
