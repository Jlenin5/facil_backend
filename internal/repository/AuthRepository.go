package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindUserByTypeChar(typeChar string, email string) (*domain.Users, error)
	CreateUser(user *domain.Users) (*domain.Users, error)
}

type DauthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *DauthRepository {
	return &DauthRepository{db: db}
}

func (r *DauthRepository) FindUserByTypeChar(typeChar string, email string) (*domain.Users, error) {
	var user domain.Users
	query := querySelectUser("u."+typeChar+" = $1")
	err := r.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, err
	}
	return &user, nil
}

// Obtener la consulta SQL
func querySelectUser(whereClause string) string {
	query := fmt.Sprintf(`
		SELECT 
			u.id, u.company_id, u.password, u.role_id, u.username, u.employee_id, 
			u.avatar, u.email, u.settings, u.shortcuts, u.status,
			COALESCE(c.id, 0) AS "company.id", COALESCE(c.name, '') AS "company.name",
			COALESCE(r.id, 0) AS "role.id", COALESCE(r.name, '') AS "role.name", COALESCE(r.description, '') AS "role.description",
			COALESCE(e.id, 0) AS "employee.id", COALESCE(e.first_name, '') AS "employee.first_name", COALESCE(e.second_name, '') AS "employee.second_name", COALESCE(e.third_name, '') AS "employee.third_name", COALESCE(e.surname, '') AS "employee.surname", COALESCE(e.second_surname, '') AS "employee.second_surname", COALESCE(e.document_number, '') AS "employee.document_number", COALESCE(e.job_position_id, 0) AS "employee.job_position_id", COALESCE(e.status, B'0') AS "employee.status",
			COALESCE(sub.id, 0) AS "subscription.id", COALESCE(sub.plan_id, 0) AS "subscription.plan_id", COALESCE(sub.company_id, 0) AS "subscription.company_id", COALESCE(sub.start_date, '1970-01-01') AS "subscription.start_date", COALESCE(sub.end_date, '1970-01-01') AS "subscription.end_date", COALESCE(sub.status, '') AS "subscription.status",
			COALESCE(p.id, 0) AS "subscription.plan.id", COALESCE(p.title, '') AS "subscription.plan.title", COALESCE(p.price, 0) AS "subscription.plan.price", COALESCE(p.max_branch_offices, 0) AS "subscription.plan.max_branch_offices", COALESCE(p.max_warehouses, 0) AS "subscription.plan.max_warehouses", COALESCE(p.max_purchases, 0) AS "subscription.plan.max_purchases", COALESCE(p.max_users, 0) AS "subscription.plan.max_users", COALESCE(p.max_products, 0) AS "subscription.plan.max_products", COALESCE(p.max_services, 0) AS "subscription.plan.max_services", COALESCE(p.max_documents, 0) AS "subscription.plan.max_documents", COALESCE(p.status, CAST(1 AS BIT)) AS "subscription.plan.status"
		FROM users u
		LEFT JOIN companies c ON u.company_id = c.id
		LEFT JOIN roles r ON u.role_id = r.id
		LEFT JOIN employees e ON u.employee_id = e.id
		LEFT JOIN subscriptions sub ON c.id = sub.company_id
			AND sub.status = 'active' 
			AND sub.end_date >= CURRENT_DATE
		LEFT JOIN plans p ON sub.plan_id = p.id
		WHERE %s
	`, whereClause)
	
	return query
}

func (r *DauthRepository) CreateUser(user *domain.Users) (*domain.Users, error) {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, username, email, password, status, created_at
	`
	var created domain.Users
	err := r.db.QueryRowx(query, user.Username, user.Email, user.Password).StructScan(&created)
	if err != nil {
		return nil, err
	}
	return &created, nil
}