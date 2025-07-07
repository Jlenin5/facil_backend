package handlersHumanresources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/usecase/human-resources"
	"github.com/gorilla/mux"
)

type AttendanceTypeHandler struct {
	AttendanceTypeUC *usecaseHumanresources.AttendanceTypeUseCase
}

func NewAttendanceTypeHandler(AttendanceTypeUC *usecaseHumanresources.AttendanceTypeUseCase) *AttendanceTypeHandler {
	return &AttendanceTypeHandler{AttendanceTypeUC: AttendanceTypeUC}
}

func (h *AttendanceTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var attendanceType humanresources.AttendanceTypes
	err := json.NewDecoder(r.Body).Decode(&attendanceType) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear una nueva marca
	err = h.AttendanceTypeUC.Create(&attendanceType)
	if err != nil {
		http.Error(w, "Failed to create attendance type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "AttendanceType created successfully"})
}

func (h *AttendanceTypeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	attendanceTypes, err := h.AttendanceTypeUC.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch attendance types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendanceTypes)
}

func (h *AttendanceTypeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid attendance type ID", http.StatusBadRequest)
		return
	}

	attendanceType, err := h.AttendanceTypeUC.GetById(id)
	if err != nil {
		http.Error(w, "AttendanceType not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendanceType)
}

func (h *AttendanceTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid attendance type ID", http.StatusBadRequest)
		return
	}

	var attendanceType humanresources.AttendanceTypes
	err = json.NewDecoder(r.Body).Decode(&attendanceType) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	attendanceType.Id = id

	// Validación básica de campos obligatorios
	if attendanceType.Name == "" {
		http.Error(w, "Missing required fields (name)", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para actualizar la marca
	err = h.AttendanceTypeUC.Update(&attendanceType)
	if err != nil {
		http.Error(w, "Failed to update attendance type: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "AttendanceType updated successfully"})
}

func (h *AttendanceTypeHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid attendance type Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar la marca por ID
	err = h.AttendanceTypeUC.DeleteById(id)
	if err != nil {
		http.Error(w, "Failed to delete attendance type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Attendance type deleted successfully"})
}

func (h *AttendanceTypeHandler) DeleteByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar las marcas por sus IDs
	err = h.AttendanceTypeUC.DeleteByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete attendance types", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Attendance types deleted successfully"})
}