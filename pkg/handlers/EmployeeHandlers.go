package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	EmployeeUC *usecase.EmployeeUseCase
}

func NewEmployeeHandler(employeeUC *usecase.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{EmployeeUC: employeeUC}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee domain.Employees

	err := json.NewDecoder(r.Body).Decode(&employee) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo emploeado
	err = h.EmployeeUC.CreateEmployee(&employee)
	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee created successfully"})
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.EmployeeUC.GetAllEmployees()
	if err != nil {
		http.Error(w, "Failed to fetch employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee Id", http.StatusBadRequest)
		return
	}

	employee, err := h.EmployeeUC.GetEmployeeById(id)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee domain.Employees
	err = json.NewDecoder(r.Body).Decode(&employee) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	employee.Id = id

	// Validación básica de campos obligatorios
	if employee.Document_Type == "" || employee.Document_Number == "" {
		http.Error(w, "Missing required fields (document_type, document_number, email)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar el emploeado
	err = h.EmployeeUC.UpdateEmployee(&employee)
	if err != nil {
		http.Error(w, "Failed to update employee: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated successfully"})
}

func (h *EmployeeHandler) DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid employee Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar el emploeado por ID
	err = h.EmployeeUC.DeleteEmployeeById(id)
	if err != nil {
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee deleted successfully"})
}

func (h *EmployeeHandler) DeleteEmployeesByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar los emploeados por sus IDs
	err = h.EmployeeUC.DeleteEmployeesByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete employees", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employees deleted successfully"})
}
