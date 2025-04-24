package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type KeyHandler struct {
	KeyUC *usecase.KeyUseCase
}

func NewKeyHandler(KeyUC *usecase.KeyUseCase) *KeyHandler {
	return &KeyHandler{KeyUC: KeyUC}
}

func (h *KeyHandler) CreateKey(w http.ResponseWriter, r *http.Request) {
	var key domain.Keys
	err := json.NewDecoder(r.Body).Decode(&key) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.KeyUC.CreateKey(&key)
	if err != nil {
		http.Error(w, "Failed to create key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area created successfully"})
}

func (h *KeyHandler) GetAllKeys(w http.ResponseWriter, r *http.Request) {
	keys, err := h.KeyUC.GetAllKeys()
	if err != nil {
		http.Error(w, "Failed to fetch work areas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keys)
}

func (h *KeyHandler) GetKeyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid key ID", http.StatusBadRequest)
		return
	}

	key, err := h.KeyUC.GetKeyById(id)
	if err != nil {
		http.Error(w, "Work area not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(key)
}

func (h *KeyHandler) UpdateKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid key ID", http.StatusBadRequest)
		return
	}

	var key domain.Keys
	err = json.NewDecoder(r.Body).Decode(&key) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	key.Id = id

	// Validación básica de campos obligatorios
	if key.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la marca
	err = h.KeyUC.UpdateKey(&key)
	if err != nil {
		http.Error(w, "Failed to update key: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Work area updated successfully"})
}