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

type UserHandler struct {
	UserUC *usecase.UserUseCase
}

func NewUserHandler(UserUC *usecase.UserUseCase) *UserHandler {
	return &UserHandler{UserUC: UserUC}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.Users
	err := json.NewDecoder(r.Body).Decode(&user) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword) // Actualiza el campo password con la versión encriptada

	// Llama al caso de uso para crear un nuevo usuario
	err = h.UserUC.CreateUser(&user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserUC.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Transformar usuarios
	transformedUsers := make([]map[string]interface{}, 0, len(users))
	for _, user := range users {
		transformedUser, err := transformUser(&user)
		if err != nil {
			http.Error(w, "Failed to process users", http.StatusInternalServerError)
			return
		}
		transformedUsers = append(transformedUsers, transformedUser)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transformedUsers)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	user, err := h.UserUC.GetUserById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Transformar usuario
	transformedUsers, err := transformUser(user)
	if err != nil {
		http.Error(w, "Failed to process user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transformedUsers)
}

// UpdateUser actualiza un usuario existente
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	var user domain.Users
	err = json.NewDecoder(r.Body).Decode(&user) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword) // Actualiza el campo password con la versión encriptada

	user.Id = id // Aseguramos que el ID del usuario coincide
	err = h.UserUC.UpdateUser(&user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// DeleteUserById elimina un usuario específico por su ID
func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	err = h.UserUC.DeleteUserById(id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

// DeleteUsersByIds elimina múltiples usuarios basándose en una lista de IDs
func (h *UserHandler) DeleteUsersByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids) // Decodificar la lista de IDs desde el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.UserUC.DeleteUsersByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Users deleted successfully"})
}

// Función auxiliar para transformar un usuario en la estructura deseada
func transformUser(user *domain.Users) (map[string]interface{}, error) {
	settings := user.Settings
	shortcuts := user.Shortcuts

	// Transformar company
	var company map[string]interface{}
	if user.Employee != nil {
		company = map[string]interface{}{
			"id":              user.Employee.Id,
			"name":            user.Company.Name,
		}
	}
	
	// Transformar employee
	var employee map[string]interface{}
	if user.Employee != nil {
		employee = map[string]interface{}{
			"id":              user.Employee.Id,
			"names":      user.Employee.Names,
			"surname":         user.Employee.Surname,
			"second_surname":  user.Employee.Second_Surname,
			"document_number": user.Employee.Document_Number,
			"status":          user.Employee.Status,
		}
	}

	// Transformar role
	var role map[string]interface{}
	if user.Role != nil {
		role = map[string]interface{}{
			"id":          user.Role.Id,
			"name":        user.Role.Name,
			"description": user.Role.Description,
		}
	}

	// Estructura transformada
	return map[string]interface{}{
		"id":          user.Id,
		"company_id":  user.Company_Id,
		"company":     company,
		"role_id":     user.Role_Id,
		"role":        role,
		"username":    user.Username,
		"employee_id": user.Employee_Id,
		"employee":    employee,
		"avatar":      user.Avatar,
		"email":       user.Email,
		"settings":    settings,
		"shortcuts":   shortcuts,
		"status":      user.Status,
	}, nil
}