package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type SaleOrderHandler struct {
	SaleOrderUC *usecase.SaleOrderUseCase
}

func NewSaleOrderHandler(SaleOrderUC *usecase.SaleOrderUseCase) *SaleOrderHandler {
	return &SaleOrderHandler{SaleOrderUC: SaleOrderUC}
}

// Crear una orden de venta
func (h *SaleOrderHandler) CreateSaleOrder(w http.ResponseWriter, r *http.Request) {
	var saleOrder domain.SaleOrders
	err := json.NewDecoder(r.Body).Decode(&saleOrder)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.SaleOrderUC.CreateSaleOrder(&saleOrder, saleOrder.SaleOrderDetails)
	if err != nil {
		http.Error(w, "Failed to create sale order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sale order created successfully"})
}

// Obtener todas las órdenes de venta
func (h *SaleOrderHandler) GetAllSaleOrders(w http.ResponseWriter, r *http.Request) {
	saleOrders, err := h.SaleOrderUC.GetAllSaleOrders()
	if err != nil {
		http.Error(w, "Failed to fetch sale orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saleOrders)
}

// Obtener una orden de venta por ID
func (h *SaleOrderHandler) GetSaleOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid sale order Id", http.StatusBadRequest)
		return
	}

	order, err := h.SaleOrderUC.GetSaleOrderById(id)
	if err != nil {
		http.Error(w, "Sale order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Actualizar una orden de venta
func (h *SaleOrderHandler) UpdateSaleOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid sale order ID", http.StatusBadRequest)
		return
	}

	var saleOrder domain.SaleOrders
	err = json.NewDecoder(r.Body).Decode(&saleOrder)
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	saleOrder.Id = id

	err = h.SaleOrderUC.UpdateSaleOrder(&saleOrder, saleOrder.SaleOrderDetails)
	if err != nil {
		http.Error(w, "Failed to update sale order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sale order updated successfully"})
}

// Eliminar una orden de venta por ID
// func (h *SaleOrderHandler) DeleteSaleOrderByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid sale order ID", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.SaleOrderUC.DeleteSaleOrderByID(id)
// 	if err != nil {
// 		http.Error(w, "Failed to delete sale order", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Sale order deleted successfully"})
// }

// Eliminar múltiples órdenes de venta por IDs
// func (h *SaleOrderHandler) DeleteSaleOrdersByIds(w http.ResponseWriter, r *http.Request) {
// 	var ids []int

// 	err := json.NewDecoder(r.Body).Decode(&ids)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.SaleOrderUC.DeleteSaleOrdersByIds(ids)
// 	if err != nil {
// 		http.Error(w, "Failed to delete sale orders", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Sale orders deleted successfully"})
// }
