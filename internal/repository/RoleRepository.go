package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// Crear un rol
func (r *RoleRepository) CreateRole(role *domain.Roles) error {
	query := `
		INSERT INTO roles (
			name, description
		) VALUES (
		 	:name, :description
		)
	`
	_, err := r.db.NamedExec(query, role)
	if err != nil {
		fmt.Printf("Error creando la rol: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los roles
func (r *RoleRepository) GetAllRoles() ([]domain.Roles, error) {
	var roles []domain.Roles
	query := "SELECT id, name, description FROM roles"
	err := r.db.Select(&roles, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return roles, nil
}

// Obtener un rol por Id
func (r *RoleRepository) GetRoleById(id int) (*domain.Roles, error) {
	query := `SELECT id, name, description FROM roles WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var role domain.Roles
	if err := row.Scan(&role.Id, &role.Name, &role.Description); err != nil {
		return nil, err
	}
	return &role, nil
}

// Actualizar un rol
func (r *RoleRepository) UpdateRole(role *domain.Roles) error {
	query := `
		UPDATE roles SET 
			name = $1, description = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(query, role.Name, role.Description, role.Id)
	if err != nil {
		fmt.Printf("Error actualizando la rol con Id %d: %v\n", role.Id, err)
		return err
	}
	return nil
}

// Eliminar un rol por Id (eliminación lógica)
func (r *RoleRepository) DeleteRoleById(id int) error {
	query := `
		UPDATE roles
		SET deleted_at = NOW()
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la rol con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples roles por Ids (eliminación lógica)
func (r *RoleRepository) DeleteRolesByIds(ids []int) error {
	query := `UPDATE roles SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
