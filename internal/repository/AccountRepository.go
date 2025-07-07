package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetAccountById(userId int) (*domain.Account, error) {
	var user domain.Account
	query := `
		SELECT 
			u.id, u.username, u.employee_id, u.avatar, u.email,
			COALESCE(c.name, '') AS "company_name",
			CONCAT_WS(' ', e.names, e.surname, e.second_surname) AS name, e.phone,
			COALESCE(sub.id, 0) AS "subscription.id", COALESCE(sub.plan_id, 0) AS "subscription.plan_id", COALESCE(sub.company_id, 0) AS "subscription.company_id", COALESCE(sub.start_date, '1970-01-01') AS "subscription.start_date", COALESCE(sub.end_date, '1970-01-01') AS "subscription.end_date", COALESCE(sub.status, '') AS "subscription.status",
			COALESCE(p.id, 0) AS "subscription.plan.id", COALESCE(p.title, '') AS "subscription.plan.title", COALESCE(p.price, 0) AS "subscription.plan.price", COALESCE(p.max_branch_offices, 0) AS "subscription.plan.max_branch_offices", COALESCE(p.max_warehouses, 0) AS "subscription.plan.max_warehouses", COALESCE(p.max_purchases, 0) AS "subscription.plan.max_purchases", COALESCE(p.max_users, 0) AS "subscription.plan.max_users", COALESCE(p.max_products, 0) AS "subscription.plan.max_products", COALESCE(p.max_services, 0) AS "subscription.plan.max_services", COALESCE(p.max_documents, 0) AS "subscription.plan.max_documents", COALESCE(p.status, CAST(1 AS BIT)) AS "subscription.plan.status"
		FROM users u
		LEFT JOIN companies c ON u.company_id = c.id
		LEFT JOIN employees e ON u.employee_id = e.id
		LEFT JOIN subscriptions sub ON c.id = sub.company_id
		LEFT JOIN plans p ON sub.plan_id = p.id
		WHERE u.id = $1
	`
	err := r.db.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, err
	}

	return &user, nil
}

func (r *AccountRepository) UpdateAccount(user *domain.Account) error {
	query := `
		UPDATE users SET 
			username = :username, updated_at = NOW()
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (r *AccountRepository) GetChangePassowrdById(userId int) (*domain.ChangePassword, error) {
	var user domain.ChangePassword
	query := `
		SELECT 
			id, password AS "current_password"
		FROM users WHERE id = $1
	`
	err := r.db.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, err
	}

	return &user, nil
}

func (r *AccountRepository) UpdateChangePassowrd(user *domain.ChangePassword) error {
	query := `
		UPDATE users SET 
			password = :new_password, updated_at = NOW()
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (r *AccountRepository) GetPlanBillingById(userId int) (*domain.PlanBilling, error) {
	var user domain.PlanBilling
	query := `
		SELECT 
			u.id, COALESCE(p.id, 0) AS "plan_id",
			COALESCE(p.id, 0) AS "plan.id", COALESCE(p.title, '') AS "plan.title", COALESCE(p.price, 0) AS "plan.price", COALESCE(p.max_branch_offices, 0) AS "plan.max_branch_offices", COALESCE(p.max_warehouses, 0) AS "plan.max_warehouses", COALESCE(p.max_purchases, 0) AS "plan.max_purchases", COALESCE(p.max_users, 0) AS "plan.max_users", COALESCE(p.max_products, 0) AS "plan.max_products", COALESCE(p.max_services, 0) AS "plan.max_services", COALESCE(p.max_documents, 0) AS "plan.max_documents", COALESCE(p.status, CAST(1 AS BIT)) AS "plan.status"
		FROM users u
		LEFT JOIN companies c ON u.company_id = c.id
		LEFT JOIN employees e ON u.employee_id = e.id
		LEFT JOIN subscriptions sub ON c.id = sub.company_id
		LEFT JOIN plans p ON sub.plan_id = p.id
		WHERE u.id = $1
	`
	err := r.db.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, err
	}

	return &user, nil
}
