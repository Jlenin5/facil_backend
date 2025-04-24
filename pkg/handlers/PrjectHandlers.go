package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jlenin5/facil_backend/internal/usecase"
)

type ProjectHandler struct {
	ProjectUC *usecase.ProjectUseCase
}

func NewProjectHandler(ProjectUC *usecase.ProjectUseCase) *ProjectHandler {
	return &ProjectHandler{ProjectUC: ProjectUC}
}

func (h *ProjectHandler) DashboardSummary(w http.ResponseWriter, r *http.Request) {
	data, err := h.ProjectUC.DashboardSummary()
	if err != nil {
		http.Error(w, "Failed to fetch dashboard data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *ProjectHandler) GetProjectData(w http.ResponseWriter, r *http.Request) {
	data, err := h.ProjectUC.GetProjectData()
	if err != nil {
		http.Error(w, "Failed to fetch dashboard data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}