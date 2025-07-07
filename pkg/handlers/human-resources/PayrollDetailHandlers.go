package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type PayrollDetailHandler struct {
	PayrollDetailUC *usecaseHumanresources.PayrollDetailUseCase
}

func NewPayrollDetailHandler(PayrollDetailUC *usecaseHumanresources.PayrollDetailUseCase) *PayrollDetailHandler {
	return &PayrollDetailHandler{PayrollDetailUC: PayrollDetailUC}
}

func (h *PayrollDetailHandler) GetByPayrollID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	payrollID, err := strconv.Atoi(vars["payrollId"])
	if err != nil {
		http.Error(w, "Invalid payroll ID", http.StatusBadRequest)
		return
	}

	details, err := h.PayrollDetailUC.GetByPayrollID(payrollID)
	if err != nil {
		http.Error(w, "Failed to fetch payroll details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}

func (h *PayrollDetailHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.PayrollDetailUC.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *PayrollDetailHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid payroll detail ID", http.StatusBadRequest)
		return
	}

	var detail humanresources.PayrollDetail
	if err := json.NewDecoder(r.Body).Decode(&detail); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	detail.Id = id
	if err := h.PayrollDetailUC.Update(&detail); err != nil {
		http.Error(w, "Failed to update payroll detail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payroll detail updated successfully"})
}

func (h *PayrollDetailHandler) MarkAsPaid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid payroll detail ID", http.StatusBadRequest)
		return
	}

	if err := h.PayrollDetailUC.MarkAsPaid(id); err != nil {
		http.Error(w, "Failed to mark payroll detail as paid", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payroll detail marked as paid successfully"})
}