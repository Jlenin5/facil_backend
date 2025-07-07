package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type ScheduleDetailHandler struct {
	ScheduleDetailUC *usecaseHumanresources.ScheduleDetailUseCase
}

func NewScheduleDetailHandler(ScheduleDetailUC *usecaseHumanresources.ScheduleDetailUseCase) *ScheduleDetailHandler {
	return &ScheduleDetailHandler{ScheduleDetailUC: ScheduleDetailUC}
}

func (h *ScheduleDetailHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.ScheduleDetailUC.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *ScheduleDetailHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid schedule detail ID", http.StatusBadRequest)
		return
	}

	var detail humanresources.ScheduleDetail
	if err := json.NewDecoder(r.Body).Decode(&detail); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	detail.Id = id
	if err := h.ScheduleDetailUC.Update(&detail); err != nil {
		http.Error(w, "Failed to update schedule detail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Schedule detail updated successfully"})
}

func (h *ScheduleDetailHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid schedule detail ID", http.StatusBadRequest)
		return
	}

	if err := h.ScheduleDetailUC.Delete(id); err != nil {
		http.Error(w, "Failed to delete schedule detail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Schedule detail deleted successfully"})
}