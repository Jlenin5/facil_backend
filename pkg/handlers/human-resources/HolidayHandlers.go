package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type HolidayHandler struct {
	HolidayUC *usecaseHumanresources.HolidayUseCase
}

func NewHolidayHandler(HolidayUC *usecaseHumanresources.HolidayUseCase) *HolidayHandler {
	return &HolidayHandler{HolidayUC: HolidayUC}
}

func (h *HolidayHandler) Create(w http.ResponseWriter, r *http.Request) {
	var holiday humanresources.Holiday
	if err := json.NewDecoder(r.Body).Decode(&holiday); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if holiday.Name == "" || holiday.Date.IsZero() {
		http.Error(w, "Name and date are required", http.StatusBadRequest)
		return
	}

	if holiday.Date.Before(time.Now().AddDate(0, 0, -1)) {
		http.Error(w, "Date cannot be in the past", http.StatusBadRequest)
		return
	}

	if err := h.HolidayUC.Create(r.Context(), &holiday); err != nil {
		http.Error(w, "Failed to create holiday", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Holiday created successfully"})
}

func (h *HolidayHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	holidays, err := h.HolidayUC.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch holidays", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(holidays)
}

func (h *HolidayHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	request, err := h.HolidayUC.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}

func (h *HolidayHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence request ID", http.StatusBadRequest)
		return
	}

	var request humanresources.Holiday
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	if err := h.HolidayUC.Update(r.Context(), &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absence request updated successfully"})
}

func (h *HolidayHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid holiday ID", http.StatusBadRequest)
		return
	}

	if err := h.HolidayUC.DeleteById(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete holiday", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Holiday deleted successfully"})
}

func (h *HolidayHandler) GetByYear(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	if yearStr == "" {
		http.Error(w, "Year parameter is required", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	holidays, err := h.HolidayUC.GetByYear(r.Context(), year)
	if err != nil {
		http.Error(w, "Failed to fetch holidays", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(holidays)
}