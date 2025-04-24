package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Crear un usuario
func (r *UserRepository) CreateUser(user *domain.Users) error {
	query := `
		INSERT INTO users (
			role, avatar, email, username, password, settings, shortcuts, status
		) VALUES (
			:role, :avatar, :email, :username, :password, :settings, :shortcuts, :status
		)
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		fmt.Printf("Error creando el usuario: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los usuarios
func (r *UserRepository) GetAllUsers() ([]domain.Users, error) {
	var users []domain.Users
	query := `
		SELECT 
			u.id, u.company_id, u.password, u.role_id, u.username, u.employee_id, u.avatar, u.email, u.settings, u.shortcuts, u.status,
			COALESCE(c.id, 0) AS "company.id", COALESCE(c.name, '') AS "company.name",
			r.id AS "role.id", r.name AS "role.name", r.description AS "role.description",
			e.id AS "employee.id", e.first_name AS "employee.first_name", e.second_name AS "employee.second_name",
			e.third_name AS "employee.third_name", e.surname AS "employee.surname", e.second_surname AS "employee.second_surname",
			e.document_number AS "employee.document_number", e.job_position_id AS "employee.job_position_id", e.status AS "employee.status"
		FROM users u
		LEFT JOIN companies c ON u.company_id = c.id
		LEFT JOIN roles r ON u.role_id = r.id
		LEFT JOIN employees e ON u.employee_id = e.id
		WHERE u.deleted_at IS NULL
	`
	err := r.db.Select(&users, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return users, nil
}

// Obtener un usuario por Id
func (r *UserRepository) GetUserById(userId int) (*domain.Users, error) {
	var user domain.Users
	query := `
		SELECT 
			u.id, u.company_id, u.password, u.role_id, u.username, u.employee_id, u.avatar, u.email, u.settings, u.shortcuts, u.status,
			COALESCE(c.id, 0) AS "company.id", COALESCE(c.name, '') AS "company.name",
			r.id AS "role.id", r.name AS "role.name", r.description AS "role.description",
			e.id AS "employee.id", e.first_name AS "employee.first_name", e.second_name AS "employee.second_name",
			e.third_name AS "employee.third_name", e.surname AS "employee.surname", e.second_surname AS "employee.second_surname",
			e.document_number AS "employee.document_number", e.job_position_id AS "employee.job_position_id", e.status AS "employee.status"
		FROM users u
		LEFT JOIN companies c ON u.company_id = c.id
		LEFT JOIN roles r ON u.role_id = r.id
		LEFT JOIN employees e ON u.employee_id = e.id
		WHERE u.id = $1 AND u.deleted_at IS NULL
	`

	// Usar Get para mapear automáticamente los resultados
	err := r.db.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &user, nil
}

// Actualizar un usuario
func (r *UserRepository) UpdateUser(user *domain.Users) error {
	query := `
		UPDATE users SET 
			company_id = :company_id, password = :password, role_id = :role_id, username = :username, 
			employee_id = :employee_id, avatar = :avatar, email = :email, settings = :settings, 
			shortcuts = :shortcuts, status = :status, updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

// Eliminar un usuario por Id (eliminación lógica)
func (r *UserRepository) DeleteUserById(id int) error {
	query := `UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando el usuario con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples usuarios por Ids (eliminación lógica)
func (r *UserRepository) DeleteUsersByIds(ids []int) error {
	query := `UPDATE users SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
