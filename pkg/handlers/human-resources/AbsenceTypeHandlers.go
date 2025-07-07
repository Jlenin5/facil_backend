package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type AbsenceTypeHandler struct {
	AbsenceTypeUC *usecaseHumanresources.AbsenceTypeUseCase
}

func NewAbsenceTypeHandler(AbsenceTypeUC *usecaseHumanresources.AbsenceTypeUseCase) *AbsenceTypeHandler {
	return &AbsenceTypeHandler{AbsenceTypeUC: AbsenceTypeUC}
}

func (h *AbsenceTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var absenceType humanresources.AbsenceType
	if err := json.NewDecoder(r.Body).Decode(&absenceType); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.AbsenceTypeUC.Create(r.Context(), &absenceType); err != nil {
		http.Error(w, "Failed to create absence type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence type created successfully"})
}

func (h *AbsenceTypeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	absenceTypes, err := h.AbsenceTypeUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch absence types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(absenceTypes)
}

func (h *AbsenceTypeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence type ID", http.StatusBadRequest)
		return
	}

	absenceType, err := h.AbsenceTypeUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, "Absence type not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(absenceType)
}

func (h *AbsenceTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence type ID", http.StatusBadRequest)
		return
	}

	var absenceType humanresources.AbsenceType
	if err := json.NewDecoder(r.Body).Decode(&absenceType); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	absenceType.Id = id
	if err := h.AbsenceTypeUC.Update(r.Context(), &absenceType); err != nil {
		http.Error(w, "Failed to update absence type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence type updated successfully"})
}

func (h *AbsenceTypeHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence type ID", http.StatusBadRequest)
		return
	}

	if err := h.AbsenceTypeUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete absence type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence type deleted successfully"})
}
