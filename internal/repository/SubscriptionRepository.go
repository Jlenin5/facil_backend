package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type SubscriptionRepository struct {
	db *sqlx.DB
}

func NewSubscriptionRepository(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

// Crear una suscripción
func (r *SubscriptionRepository) CreateSubscription(subscription *domain.Subscriptions) error {
	query := `
		INSERT INTO subscriptions (
			company_id, plan_id, start_date, end_date, status
		) VALUES (
		 	:company_id, :plan_id, :start_date, :end_date, :status
		)
	`
	_, err := r.db.NamedExec(query, subscription)
	if err != nil {
		fmt.Printf("Error creando la suscripción: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las suscripciones
func (r *SubscriptionRepository) GetAllSubscriptions() ([]domain.Subscriptions, error) {
	var subscriptions []domain.Subscriptions
	query := `
		SELECT
			s.id, s.company_id, s.plan_id, s.start_date, s.end_date, s.status,
			c.id AS "company.id", c.name AS "company.name",
			p.id AS "plan.id", p.title AS "plan.title"
		FROM subscriptions s
		LEFT JOIN companies c ON s.company_id=c.id
		LEFT JOIN plans p ON s.plan_id = p.id
		WHERE s.deleted_at IS NULL
	`
	err := r.db.Select(&subscriptions, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return subscriptions, nil
}

// Obtener una suscripción por Id
func (r *SubscriptionRepository) GetSubscriptionById(id int) (*domain.Subscriptions, error) {
	var subscription domain.Subscriptions
	query := `
		SELECT
			s.id, s.company_id, s.plan_id, s.start_date, s.end_date, s.status,
			c.id AS "company.id", c.name AS "company.name",
			p.id AS "plan.id", p.title AS "plan.title"
		FROM subscriptions s
		LEFT JOIN companies c ON s.company_id=c.id
		LEFT JOIN plans p ON s.plan_id = p.id
		WHERE s.id = $1 AND s.deleted_at IS NULL
	`
	err := r.db.Get(&subscription, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Suscripción no encontrada
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &subscription, nil
}

// Actualizar una suscripción
func (r *SubscriptionRepository) UpdateSubscription(subscription *domain.Subscriptions) error {
	query := `
		UPDATE subscriptions SET 
			company_id = $1, plan_id = $2, start_date = $3, end_date = $4, status = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, subscription.Company_Id, subscription.Plan_Id, subscription.Start_Date, subscription.End_Date, subscription.Status, subscription.Id)
	if err != nil {
		fmt.Printf("Error actualizando la suscripción con Id %d: %v\n", subscription.Id, err)
		return err
	}
	return nil
}

// Eliminar una suscripción por Id (eliminación lógica)
func (r *SubscriptionRepository) DeleteSubscriptionById(id int) error {
	query := `
		UPDATE subscriptions
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la suscripción con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples suscripciones por Ids (eliminación lógica)
func (r *SubscriptionRepository) DeleteSubscriptionsByIds(ids []int) error {
	query := `UPDATE subscriptions SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}

// Obtener la suscripción de una empresa por su company_id
func (r *SubscriptionRepository) GetSubscriptionByCompanyId(companyId int) (*domain.Subscriptions, error) {
	var subscription domain.Subscriptions
	query := `
		SELECT
			s.id, s.company_id, s.plan_id, s.start_date, s.end_date, s.status
		FROM subscriptions s
		WHERE s.company_id = $1 AND s.deleted_at IS NULL
	`
	err := r.db.Get(&subscription, query, companyId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No hay suscripción para esta empresa
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &subscription, nil
}