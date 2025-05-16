package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/Jlenin5/facil_backend/pkg/middleware"
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
	// Obtener los datos del usuario del contexto
	userData, ok := r.Context().Value(middleware.UserContextKey).(map[string]interface{})
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verificar que employee exista en userData
	employeeData, exists := userData["employee"]
	if !exists || employeeData == nil {
		http.Error(w, "Employee not found in user data", http.StatusBadRequest)
		return
	}

	// Asegurarse que employee sea un map[string]interface{}
	employeeMap, ok := employeeData.(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid employee data", http.StatusBadRequest)
		return
	}

	// Verificar que warehouse_id exista en el map
	warehouseIDRaw, exists := employeeMap["warehouse_id"]
	if !exists || warehouseIDRaw == nil {
		http.Error(w, "Warehouse ID not found", http.StatusBadRequest)
		return
	}

	// Convertir warehouse_id a float64 y luego a int
	warehouseIDFloat, ok := warehouseIDRaw.(float64)
	if !ok {
		http.Error(w, "Invalid warehouse_id type", http.StatusBadRequest)
		return
	}

	warehouseId := int(warehouseIDFloat)

	// Decodificar el cuerpo de la solicitud (el JSON con los datos de la venta)
	var purchaseOrder domain.PurchaseOrders
	err := json.NewDecoder(r.Body).Decode(&purchaseOrder)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Asignar los datos del usuario a la venta
	purchaseOrder.Warehouse_Id = warehouseId

	// Validar user_id del usuario autenticado
	userIDRaw, exists := userData["id"]
	if !exists {
		http.Error(w, "User ID not found", http.StatusBadRequest)
		return
	}

	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	purchaseOrder.Created_By = int(userIDFloat)

	// Reiniciar los IDs de los detalles de la venta (opcional si usas auto increment en la base)
	for i := range purchaseOrder.PurchaseOrderDetails {
		purchaseOrder.PurchaseOrderDetails[i].Id = 0
	}

	err = h.PurchaseOrderUC.CreatePurchaseOrder(&purchaseOrder, purchaseOrder.PurchaseOrderDetails)
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

// Actualizar una orden de compra
func (h *PurchaseOrderHandler) UpdatePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid purchaseOrder ID", http.StatusBadRequest)
		return
	}

	var purchaseOrder domain.PurchaseOrders
	err = json.NewDecoder(r.Body).Decode(&purchaseOrder)
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	purchaseOrder.Id = id

	err = h.PurchaseOrderUC.UpdatePurchaseOrder(&purchaseOrder, purchaseOrder.PurchaseOrderDetails)
	if err != nil {
		http.Error(w, "Failed to update purchase order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Purchase updated successfully"})
}