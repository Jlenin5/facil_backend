package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type PurchaseOrderHandler struct {
	PurchaseOrderUC *usecase.PurchaseOrderUseCase
}

func NewPurchaseOrderHandler(PurchaseOrderUC *usecase.PurchaseOrderUseCase) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{PurchaseOrderUC: PurchaseOrderUC}
}

// Crear una orden de compra
func (h *PurchaseOrderHandler) PostPurchaseOrder(w http.ResponseWriter, r *http.Request) {
	var purchaseOrder struct {
		Order   domain.PurchaseOrders         `json:"order"`
		Details []domain.PurchaseOrderDetails `json:"details"`
	}

	err := json.NewDecoder(r.Body).Decode(&purchaseOrder)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.PurchaseOrderUC.CreatePurchaseOrder(&purchaseOrder.Order, purchaseOrder.Details)
	if err != nil {
		http.Error(w, "Failed to create purchase order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Purchase order created successfully"})
}

// Obtener todas las órdenes de compra
func (h *PurchaseOrderHandler) GetAllPurchaseOrders(w http.ResponseWriter, r *http.Request) {
	purchaseOrders, err := h.PurchaseOrderUC.GetAllPurchaseOrders()
	if err != nil {
		http.Error(w, "Failed to fetch purchase orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchaseOrders)
}

// Obtener una orden de compra por ID
func (h *PurchaseOrderHandler) GetPurchaseOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid purchase order Id", http.StatusBadRequest)
		return
	}

	order, err := h.PurchaseOrderUC.GetPurchaseOrderByID(id)
	if err != nil {
		http.Error(w, "Purchase order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
