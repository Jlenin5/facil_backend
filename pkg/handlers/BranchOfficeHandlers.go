package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type BranchOfficeHandler struct {
	BranchOfficeUC *usecase.BranchOfficeUseCase
}

func NewBranchOfficeHandler(BranchOfficeUC *usecase.BranchOfficeUseCase) *BranchOfficeHandler {
	return &BranchOfficeHandler{BranchOfficeUC: BranchOfficeUC}
}

func (h *BranchOfficeHandler) CreateBranchOffice(w http.ResponseWriter, r *http.Request) {
	var branchOffice domain.BranchOffices
	err := json.NewDecoder(r.Body).Decode(&branchOffice) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo sucursal
	err = h.BranchOfficeUC.CreateBranchOffice(&branchOffice)
	if err != nil {
		http.Error(w, "Failed to create branch office", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Branch office created successfully"})
}

func (h *BranchOfficeHandler) GetAllBranchOffices(w http.ResponseWriter, r *http.Request) {
	branchOffices, err := h.BranchOfficeUC.GetAllBranchOffices()
	if err != nil {
		http.Error(w, "Failed to fetch branch offices", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branchOffices)
}

func (h *BranchOfficeHandler) GetBranchOfficeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid branch office ID", http.StatusBadRequest)
		return
	}

	branchOffice, err := h.BranchOfficeUC.GetBranchOfficeById(id)
	if err != nil {
		http.Error(w, "Branch office not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branchOffice)
}

func (h *BranchOfficeHandler) UpdateBranchOffice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid branch office ID", http.StatusBadRequest)
		return
	}

	var branchOffice domain.BranchOffices
	err = json.NewDecoder(r.Body).Decode(&branchOffice) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	branchOffice.Id = id

	// Validación básica de campos obligatorios
	if branchOffice.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la sucursal
	err = h.BranchOfficeUC.UpdateBranchOffice(&branchOffice)
	if err != nil {
		http.Error(w, "Failed to update branchOffice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Branch office updated successfully"})
}

func (h *BranchOfficeHandler) DeleteBranchOfficeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid branch office Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la sucursal por ID
	err = h.BranchOfficeUC.DeleteBranchOfficeById(id)
	if err != nil {
		http.Error(w, "Failed to delete branch office", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Branch office deleted successfully"})
}

func (h *BranchOfficeHandler) DeleteBranchOfficesByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las sucursales por sus IDs
	err = h.BranchOfficeUC.DeleteBranchOfficesByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete branch offices", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Branch offices deleted successfully"})
}
