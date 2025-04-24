package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type PlanRepository struct {
	db *sqlx.DB
}

func NewPlanRepository(db *sqlx.DB) *PlanRepository {
	return &PlanRepository{db: db}
}

// Crear un plan
func (r *PlanRepository) CreatePlan(plan *domain.Plans) error {
	query := `
		INSERT INTO plans (
			title, subtitle, price, max_branch_offices, max_warehouses, max_purchases, max_users, max_products, max_services, max_documents, status
		) VALUES (
		 	:title, :subtitle, :price, :max_branch_offices, :max_warehouses, :max_purchases, :max_users, :max_products, :max_services, :max_documents, :status
		)
	`
	_, err := r.db.NamedExec(query, plan)
	if err != nil {
		fmt.Printf("Error creando el plan: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las plans
func (r *PlanRepository) GetAllPlans() ([]domain.Plans, error) {
	var plans []domain.Plans
	query := `
		SELECT
			id, title, subtitle, price, max_branch_offices, max_warehouses, max_purchases, max_users, max_products, max_services, max_documents, status
		FROM plans
		WHERE deleted_at IS NULL
		ORDER BY id DESC
	`
	err := r.db.Select(&plans, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return plans, nil
}

// Obtener un plan por Id
func (r *PlanRepository) GetPlanById(id int) (*domain.Plans, error) {
	var plan domain.Plans
	query := `
		SELECT
			id, title, subtitle, price, max_branch_offices, max_warehouses, max_purchases, max_users, max_products, max_services, max_documents, status
		FROM plans
		WHERE id = $1 AND deleted_at IS NULL
	`
	err := r.db.Get(&plan, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Plan no encontrado
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return &plan, nil
}

// Actualizar un plan
func (r *PlanRepository) UpdatePlan(plan *domain.Plans) error {
	query := `
		UPDATE plans SET 
			title = :title,
			subtitle = :subtitle,
			price = :price,
			max_branch_offices = :max_branch_offices,
			max_warehouses = :max_warehouses,
			max_purchases = :max_purchases,
			max_users = :max_users,
			max_products = :max_products,
			max_services = :max_services,
			max_documents = :max_documents,
			status = :status,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err := r.db.NamedExec(query, plan)
	if err != nil {
		fmt.Printf("Error actualizando el plan con Id %d: %v\n", plan.Id, err)
		return err
	}
	return nil
}

// Eliminar un plan por Id (eliminación lógica)
func (r *PlanRepository) DeletePlanById(id int) error {
	query := `
		UPDATE plans
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la plan con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples plans por Ids (eliminación lógica)
func (r *PlanRepository) DeletePlansByIds(ids []int) error {
	query := `UPDATE plans SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
