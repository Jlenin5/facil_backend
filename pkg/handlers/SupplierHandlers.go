package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type SupplierHandler struct {
	SupplierUC *usecase.SupplierUseCase
}

func NewSupplierHandler(SupplierUC *usecase.SupplierUseCase) *SupplierHandler {
	return &SupplierHandler{SupplierUC: SupplierUC}
}

func (h *SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier domain.Suppliers

	err := json.NewDecoder(r.Body).Decode(&supplier) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo proveedor
	err = h.SupplierUC.CreateSupplier(&supplier)
	if err != nil {
		http.Error(w, "Failed to create supplier", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Supplier created successfully"})
}

func (h *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := h.SupplierUC.GetAllSuppliers()
	if err != nil {
		http.Error(w, "Failed to fetch suppliers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suppliers)
}

func (h *SupplierHandler) GetSupplierById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid supplier Id", http.StatusBadRequest)
		return
	}

	supplier, err := h.SupplierUC.GetSupplierById(id)
	if err != nil {
		http.Error(w, "Supplier not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(supplier)
}

func (h *SupplierHandler) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	var supplier domain.Suppliers
	err = json.NewDecoder(r.Body).Decode(&supplier) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	supplier.Id = id

	// Validación básica de campos obligatorios
	if supplier.Name == "" {
		http.Error(w, "Missing required fields (name, email)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar el proveedor
	err = h.SupplierUC.UpdateSupplier(&supplier)
	if err != nil {
		http.Error(w, "Failed to update supplier: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Supplier updated successfully"})
}

func (h *SupplierHandler) DeleteSupplierById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid supplier Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar el proveedor por ID
	err = h.SupplierUC.DeleteSupplierById(id)
	if err != nil {
		http.Error(w, "Failed to delete supplier", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Supplier deleted successfully"})
}

func (h *SupplierHandler) DeleteSuppliersByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar los proveedores por sus IDs
	err = h.SupplierUC.DeleteSuppliersByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete suppliers", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Suppliers deleted successfully"})
}
