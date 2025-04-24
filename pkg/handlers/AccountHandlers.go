package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type AccountHandler struct {
	AccountUC *usecase.AccountUseCase
}

func NewAccountHandler(AccountUC *usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{AccountUC: AccountUC}
}

func (h *AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	user, err := h.AccountUC.GetAccountById(id)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var account domain.Account
	err = json.NewDecoder(r.Body).Decode(&account) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	account.Id = id

	// Llama al caso de uso para actualizar la marca
	err = h.AccountUC.UpdateAccount(&account)
	if err != nil {
		http.Error(w, "Failed to update account: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Brand updated successfully"})
}

func (h *AccountHandler) GetChangePassowrdById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	user, err := h.AccountUC.GetChangePassowrdById(id)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *AccountHandler) UpdateChangePassowrd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}
	
	var userDomain domain.ChangePassword
	err = json.NewDecoder(r.Body).Decode(&userDomain) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.AccountUC.GetChangePassowrdById(id)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Current_Password), []byte(userDomain.Current_Password)); err != nil {
		return
	}

	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDomain.New_Password), 12)
	if err != nil {
		http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
		return
	}
	userDomain.New_Password = string(hashedPassword) // Actualiza el campo password con la versión encriptada

	userDomain.Id = id // Aseguramos que el ID del usuario coincide
	err = h.AccountUC.UpdateChangePassowrd(&userDomain)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func (h *AccountHandler) GetPlanBillingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	user, err := h.AccountUC.GetPlanBillingById(id)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}