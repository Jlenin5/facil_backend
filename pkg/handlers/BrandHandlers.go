package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type BrandHandler struct {
	BrandUC *usecase.BrandUseCase
}

func NewBrandHandler(BrandUC *usecase.BrandUseCase) *BrandHandler {
	return &BrandHandler{BrandUC: BrandUC}
}

func (h *BrandHandler) CreateBrand(w http.ResponseWriter, r *http.Request) {
	var brand domain.Brands
	err := json.NewDecoder(r.Body).Decode(&brand) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.BrandUC.CreateBrand(&brand)
	if err != nil {
		http.Error(w, "Failed to create brand", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Brand created successfully"})
}

func (h *BrandHandler) GetAllBrands(w http.ResponseWriter, r *http.Request) {
	brands, err := h.BrandUC.GetAllBrands()
	if err != nil {
		http.Error(w, "Failed to fetch brands", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brands)
}

func (h *BrandHandler) GetBrandById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid brand ID", http.StatusBadRequest)
		return
	}

	brand, err := h.BrandUC.GetBrandById(id)
	if err != nil {
		http.Error(w, "Brand not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brand)
}

func (h *BrandHandler) UpdateBrand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid brand ID", http.StatusBadRequest)
		return
	}

	var brand domain.Brands
	err = json.NewDecoder(r.Body).Decode(&brand) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	brand.Id = id

	// Validación básica de campos obligatorios
	if brand.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la marca
	err = h.BrandUC.UpdateBrand(&brand)
	if err != nil {
		http.Error(w, "Failed to update brand: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Brand updated successfully"})
}

func (h *BrandHandler) DeleteBrandById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid brand Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.BrandUC.DeleteBrandById(id)
	if err != nil {
		http.Error(w, "Failed to delete brand", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Brand deleted successfully"})
}

func (h *BrandHandler) DeleteBrandsByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.BrandUC.DeleteBrandsByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete brands", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Brands deleted successfully"})
}