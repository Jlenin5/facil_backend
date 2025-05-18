package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jlenin5/facil_backend/internal/usecase"
)

type AuthHandler struct {
	AuthUC *usecase.DauthUseCase
}

func NewAuthHandler(authUC *usecase.DauthUseCase) *AuthHandler {
	return &AuthHandler{AuthUC: authUC}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	response, err := h.AuthUC.SignIn(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	response, err := h.AuthUC.SignUp(input.Username, input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) ValidateRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Obtener el token del header Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "No se proporcionó el token de autorización", http.StatusBadRequest)
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		http.Error(w, "Formato de token inválido", http.StatusBadRequest)
		return
	}
	refreshToken := authHeader[len(bearerPrefix):]

	// Validar el refresh token y obtener los datos del usuario
	userResponse, err := h.AuthUC.ValidateRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, "Refresh token inválido o expirado", http.StatusUnauthorized)
		return
	}

	// Responder con los datos del usuario
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}