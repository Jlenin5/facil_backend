package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type PaymentMethodHandler struct {
	PaymentMethodUC *usecase.PaymentMethodUseCase
}

func NewPaymentMethodHandler(PaymentMethodUC *usecase.PaymentMethodUseCase) *PaymentMethodHandler {
	return &PaymentMethodHandler{PaymentMethodUC: PaymentMethodUC}
}

func (h *PaymentMethodHandler) CreatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	var paymentMethod domain.PaymentMethods
	err := json.NewDecoder(r.Body).Decode(&paymentMethod) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo método de pago
	err = h.PaymentMethodUC.CreatePaymentMethod(&paymentMethod)
	if err != nil {
		http.Error(w, "Failed to create payment method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment method created successfully"})
}

func (h *PaymentMethodHandler) GetAllPaymentMethods(w http.ResponseWriter, r *http.Request) {
	paymentMethods, err := h.PaymentMethodUC.GetAllPaymentMethods()
	if err != nil {
		http.Error(w, "Failed to fetch payment methods", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentMethods)
}

func (h *PaymentMethodHandler) GetPaymentMethodById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid paymentMethod ID", http.StatusBadRequest)
		return
	}

	paymentMethod, err := h.PaymentMethodUC.GetPaymentMethodById(id)
	if err != nil {
		http.Error(w, "PaymentMethod not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentMethod)
}

func (h *PaymentMethodHandler) UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid payment method ID", http.StatusBadRequest)
		return
	}

	var paymentMethod domain.PaymentMethods
	err = json.NewDecoder(r.Body).Decode(&paymentMethod) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	paymentMethod.Id = id

	// Validación básica de campos obligatorios
	if paymentMethod.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar el método de pago
	err = h.PaymentMethodUC.UpdatePaymentMethod(&paymentMethod)
	if err != nil {
		http.Error(w, "Failed to update payment method: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment method updated successfully"})
}

func (h *PaymentMethodHandler) DeletePaymentMethodById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid paymentMethod Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar el método de pago por ID
	err = h.PaymentMethodUC.DeletePaymentMethodById(id)
	if err != nil {
		http.Error(w, "Failed to delete payment method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment method deleted successfully"})
}

func (h *PaymentMethodHandler) DeletePaymentMethodsByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar los métodos de pago por sus IDs
	err = h.PaymentMethodUC.DeletePaymentMethodsByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete paymentMethods", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment methods deleted successfully"})
}
