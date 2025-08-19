package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	// Parsear el formulario multipart
	err := r.ParseMultipartForm(32 << 20) // 32 MB máximo
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Obtener los datos del empleado del campo "employee"
	employeeJSON := r.FormValue("employee")
	if employeeJSON == "" {
		http.Error(w, "Employee data is required", http.StatusBadRequest)
		return
	}

	var employee domain.Employees
	err = json.Unmarshal([]byte(employeeJSON), &employee)
	if err != nil {
		http.Error(w, "Invalid employee data format", http.StatusBadRequest)
		return
	}

	// Manejar la foto subida
	file, handler, err := r.FormFile("photo")
	if err == nil {
		defer file.Close()

		// Crear directorio si no existe
		uploadDir := "uploads/images/employees"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			err = os.MkdirAll(uploadDir, 0755)
			if err != nil {
				http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
				return
			}
		}

		// Generar nombre único para el archivo
		ext := filepath.Ext(handler.Filename)
		filename := strings.ReplaceAll(strings.ToLower(employee.Names+"_"+employee.Document_Number), " ", "_") + ext
		filePath := filepath.Join(uploadDir, filename)

		dbFilePath := strings.ReplaceAll(filePath, "\\", "/")

		// Crear el archivo en el sistema
		dst, err := os.Create(dbFilePath)
		if err != nil {
			http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copiar el contenido del archivo
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		// Asignar la ruta de la foto al empleado
		employee.Photo = domain.NullString{String: &dbFilePath, Valid: true}
	} else if photoPath := r.FormValue("photo_path"); photoPath != "" {
		// Si viene una ruta de foto existente
		employee.Photo = domain.NullString{String: &photoPath, Valid: true}
	}

	// Validar campos obligatorios
	if employee.Document_Type == "" || employee.Document_Number == "" {
		http.Error(w, "Missing required fields (document_type, document_number)", http.StatusBadRequest)
		return
	}

	// Crear el empleado
	err = h.EmployeeUC.CreateEmployee(&employee)
	if err != nil {
		http.Error(w, "Failed to create employee: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	photo := ""
	if employee.Photo.Valid && employee.Photo.String != nil {
		photo = *employee.Photo.String
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Employee created successfully",
		"photo":   photo,
	})
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Parsear el formulario multipart
	err = r.ParseMultipartForm(10 << 20) // 32 MB máximo
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Obtener los datos del empleado del campo "employee"
	employeeJSON := r.FormValue("employee")
	if employeeJSON == "" {
		http.Error(w, "Employee data is required", http.StatusBadRequest)
		return
	}

	var employee domain.Employees
	err = json.Unmarshal([]byte(employeeJSON), &employee)
	if err != nil {
		http.Error(w, "Invalid employee data format", http.StatusBadRequest)
		return
	}

	employee.Id = id

	// Manejar la foto subida
	file, handler, err := r.FormFile("photo")
	if err == nil {
		defer file.Close()

		// Crear directorio si no existe
		uploadDir := "uploads/images/employees"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			err = os.MkdirAll(uploadDir, 0755)
			if err != nil {
				http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
				return
			}
		}

		// Generar nombre único para el archivo
		ext := filepath.Ext(handler.Filename)
		filename := strings.ReplaceAll(strings.ToLower(employee.Names+"_"+employee.Document_Number), " ", "_") + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ext
		filePath := filepath.Join(uploadDir, filename)

		dbFilePath := strings.ReplaceAll(filePath, "\\", "/")

		// Crear el archivo en el sistema
		dst, err := os.Create(dbFilePath)
		if err != nil {
			http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copiar el contenido del archivo
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		// Asignar la nueva ruta de la foto al empleado
		employee.Photo = domain.NullString{String: &dbFilePath, Valid: true}
	} else if photoPath := r.FormValue("photo_path"); photoPath != "" {
		// Si viene una ruta de foto existente y no se subió nueva foto
		employee.Photo = domain.NullString{String: &photoPath, Valid: true}
	}

	// Validar campos obligatorios
	if employee.Document_Type == "" || employee.Document_Number == "" {
		http.Error(w, "Missing required fields (document_type, document_number)", http.StatusBadRequest)
		return
	}

	// Actualizar el empleado
	err = h.EmployeeUC.UpdateEmployee(&employee)
	if err != nil {
		http.Error(w, "Failed to update employee: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	
	photo := ""
	if employee.Photo.Valid && employee.Photo.String != nil {
		photo = *employee.Photo.String
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Employee updated successfully",
		"photo":   photo,
	})
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
