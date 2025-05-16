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

type SaleOrderHandler struct {
	SaleOrderUC *usecase.SaleOrderUseCase
}

func NewSaleOrderHandler(SaleOrderUC *usecase.SaleOrderUseCase) *SaleOrderHandler {
	return &SaleOrderHandler{SaleOrderUC: SaleOrderUC}
}

// Crear una orden de venta
func (h *SaleOrderHandler) CreateSaleOrder(w http.ResponseWriter, r *http.Request) {
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
	var saleOrder domain.SaleOrders
	err := json.NewDecoder(r.Body).Decode(&saleOrder)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Asignar los datos del usuario a la venta
	saleOrder.Warehouse_Id = warehouseId

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

	saleOrder.User_Id = int(userIDFloat)

	// Reiniciar los IDs de los detalles de la venta (opcional si usas auto increment en la base)
	for i := range saleOrder.SaleOrderDetails {
		saleOrder.SaleOrderDetails[i].Id = 0
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
