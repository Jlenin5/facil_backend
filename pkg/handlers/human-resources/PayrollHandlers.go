package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	usecaseHumanresources "github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type PayrollHandler struct {
	PayrollUC *usecaseHumanresources.PayrollUseCase
}

func NewPayrollHandler(PayrollUC *usecaseHumanresources.PayrollUseCase) *PayrollHandler {
	return &PayrollHandler{PayrollUC: PayrollUC}
}

func (h *PayrollHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payroll humanresources.Payroll
	if err := json.NewDecoder(r.Body).Decode(&payroll); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if payroll.Reference == "" || payroll.PeriodStart.IsZero() || payroll.PeriodEnd.IsZero() || payroll.PaymentDate.IsZero() || payroll.CreatedBy == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if payroll.PeriodEnd.Before(payroll.PeriodStart) {
		http.Error(w, "Period end must be after period start", http.StatusBadRequest)
		return
	}

	if payroll.PaymentDate.Before(payroll.PeriodEnd) {
		http.Error(w, "Payment date must be after period end", http.StatusBadRequest)
		return
	}

	if err := h.PayrollUC.Create(&payroll); err != nil {
		http.Error(w, "Failed to create payroll", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payroll created successfully"})
}

func (h *PayrollHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendances, err := h.PayrollUC.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (h *PayrollHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid payroll ID", http.StatusBadRequest)
		return
	}

	payroll, err := h.PayrollUC.GetById(id)
	if err != nil {
		http.Error(w, "Payroll not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payroll)
}

func (h *PayrollHandler) Approve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid payroll ID", http.StatusBadRequest)
		return
	}

	var body struct {
		ApprovedBy int `json:"approvedBy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.PayrollUC.Approve(id, body.ApprovedBy); err != nil {
		http.Error(w, "Failed to approve payroll", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payroll approved successfully"})
}
