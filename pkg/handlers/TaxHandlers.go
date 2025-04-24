package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type TaxHandler struct {
	TaxUC *usecase.TaxUseCase
}

func NewTaxHandler(TaxUC *usecase.TaxUseCase) *TaxHandler {
	return &TaxHandler{TaxUC: TaxUC}
}

func (h *TaxHandler) CreateTax(w http.ResponseWriter, r *http.Request) {
	var tax domain.Taxes
	err := json.NewDecoder(r.Body).Decode(&tax) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo impuesto
	err = h.TaxUC.CreateTax(&tax)
	if err != nil {
		http.Error(w, "Failed to create tax", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tax created successfully"})
}

func (h *TaxHandler) GetAllTaxes(w http.ResponseWriter, r *http.Request) {
	taxes, err := h.TaxUC.GetAllTaxes()
	if err != nil {
		http.Error(w, "Failed to fetch taxes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taxes)
}

func (h *TaxHandler) GetTaxById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid tax ID", http.StatusBadRequest)
		return
	}

	tax, err := h.TaxUC.GetTaxById(id)
	if err != nil {
		http.Error(w, "Tax not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tax)
}

func (h *TaxHandler) UpdateTax(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid tax ID", http.StatusBadRequest)
		return
	}

	var tax domain.Taxes
	err = json.NewDecoder(r.Body).Decode(&tax) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	tax.Id = id

	// Validación básica de campos obligatorios
	if tax.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar el impuesto
	err = h.TaxUC.UpdateTax(&tax)
	if err != nil {
		http.Error(w, "Failed to update tax: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tax updated successfully"})
}

func (h *TaxHandler) DeleteTaxById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid tax Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar el impuesto por ID
	err = h.TaxUC.DeleteTaxById(id)
	if err != nil {
		http.Error(w, "Failed to delete tax", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tax deleted successfully"})
}

func (h *TaxHandler) DeleteTaxesByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar los impuestos por sus IDs
	err = h.TaxUC.DeleteTaxesByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete taxes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Taxes deleted successfully"})
}
