package repositoryHumanresources

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type AttendanceTypeRepository struct {
	db *sqlx.DB
}

func NewAttendanceTypeRepository(db *sqlx.DB) *AttendanceTypeRepository {
	return &AttendanceTypeRepository{db: db}
}

// Crear un tipo de asistencia
func (r *AttendanceTypeRepository) Create(attendanceType *humanresources.AttendanceTypes) error {
	query := `
		INSERT INTO attendance_types (
			name, code, description, status
		) VALUES (
		 	:name, :code, :description, :status
		)
	`
	_, err := r.db.NamedExec(query, attendanceType)
	if err != nil {
		fmt.Printf("Error creando la marca: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los tipos de asistencia
func (r *AttendanceTypeRepository) GetAll() ([]humanresources.AttendanceTypes, error) {
	var attendanceType []humanresources.AttendanceTypes
	query := "SELECT id, name, code, description, status FROM attendance_types WHERE deleted_at IS NULL"
	err := r.db.Select(&attendanceType, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return attendanceType, nil
}

// Obtener un tipo de asistencia por Id
func (r *AttendanceTypeRepository) GetById(id int) (*humanresources.AttendanceTypes, error) {
	var attendanceType humanresources.AttendanceTypes
	query := `
		SELECT
			id, name, code, description, status
		FROM attendance_types
		WHERE id = $1 AND deleted_at IS NULL
	`
	err := r.db.Get(&attendanceType, query, id)
	if err != nil {
		fmt.Printf("Error obteniendo cliente con ID %d: %v\n", id, err)
		return nil, err
	}
	return &attendanceType, nil
}

// Actualizar un tipo de asistencia
func (r *AttendanceTypeRepository) Update(attendanceType *humanresources.AttendanceTypes) error {
	query := `
		UPDATE attendance_types SET
			name = $name,
			code = $code,
			description = $description,
			status = $status,
			updated_at = NOW()
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, attendanceType)
	return err
}

// Eliminar un tipo de asistencia por Id (eliminación lógica)
func (r *AttendanceTypeRepository) DeleteById(id int) error {
	query := `
		UPDATE attendance_types
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la marca con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples tipos de asistencias por Ids (eliminación lógica)
func (r *AttendanceTypeRepository) DeleteByIds(ids []int) error {
	query := `UPDATE attendance_types SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
