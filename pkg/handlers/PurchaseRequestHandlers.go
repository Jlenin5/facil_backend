package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type PurchaseRequestHandler struct {
	PurchaseRequestUC *usecase.PurchaseRequestUseCase
}

func NewPurchaseRequestHandler(PurchaseRequestUC *usecase.PurchaseRequestUseCase) *PurchaseRequestHandler {
	return &PurchaseRequestHandler{PurchaseRequestUC: PurchaseRequestUC}
}

// Crear una orden de venta
func (h *PurchaseRequestHandler) CreatePurchaseRequest(w http.ResponseWriter, r *http.Request) {
	var purchaseRequest domain.PurchaseRequests
	err := json.NewDecoder(r.Body).Decode(&purchaseRequest)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.PurchaseRequestUC.CreatePurchaseRequest(&purchaseRequest, purchaseRequest.PurchaseRequestDetails)
	if err != nil {
		http.Error(w, "Failed to create sale order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sale order created successfully"})
}

// Obtener todas las órdenes de venta
func (h *PurchaseRequestHandler) GetAllPurchaseRequests(w http.ResponseWriter, r *http.Request) {
	purchaseRequests, err := h.PurchaseRequestUC.GetAllPurchaseRequests()
	if err != nil {
		http.Error(w, "Failed to fetch sale orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchaseRequests)
}

// Obtener una orden de venta por ID
func (h *PurchaseRequestHandler) GetPurchaseRequestById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid sale order Id", http.StatusBadRequest)
		return
	}

	order, err := h.PurchaseRequestUC.GetPurchaseRequestById(id)
	if err != nil {
		http.Error(w, "Sale order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Actualizar una orden de venta
func (h *PurchaseRequestHandler) UpdatePurchaseRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid sale order ID", http.StatusBadRequest)
		return
	}

	var purchaseRequest domain.PurchaseRequests
	err = json.NewDecoder(r.Body).Decode(&purchaseRequest)
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	purchaseRequest.Id = id

	err = h.PurchaseRequestUC.UpdatePurchaseRequest(&purchaseRequest, purchaseRequest.PurchaseRequestDetails)
	if err != nil {
		http.Error(w, "Failed to update sale order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sale order updated successfully"})
}

// Eliminar una orden de venta por ID
// func (h *PurchaseRequestHandler) DeletePurchaseRequestByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid sale order ID", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.PurchaseRequestUC.DeletePurchaseRequestByID(id)
// 	if err != nil {
// 		http.Error(w, "Failed to delete sale order", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Sale order deleted successfully"})
// }

// Eliminar múltiples órdenes de venta por IDs
// func (h *PurchaseRequestHandler) DeletePurchaseRequestsByIds(w http.ResponseWriter, r *http.Request) {
// 	var ids []int

// 	err := json.NewDecoder(r.Body).Decode(&ids)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.PurchaseRequestUC.DeletePurchaseRequestsByIds(ids)
// 	if err != nil {
// 		http.Error(w, "Failed to delete sale orders", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Sale orders deleted successfully"})
// }
