package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	db *sqlx.DB
}

func NewCompanyRepository(db *sqlx.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

// Crear uns empresa
func (r *CompanyRepository) CreateCompany(company *domain.Companies) error {
	query := "INSERT INTO companies (name, ruc, email, phone, web_site, address, status) VALUES (:name, :ruc, :email, :phone, :web_site, :address, :status)"
	_, err := r.db.NamedExec(query, company)
	if err != nil {
		fmt.Printf("Error creando la empresa: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las empresas
func (r *CompanyRepository) GetAllCompanies() ([]domain.Companies, error) {
	var companies []domain.Companies
	query := `
		SELECT
			c.id, c.name, c.logo, c.ruc, c.email, c.phone, c.web_site, c.address, c.status,
			COALESCE(p.id, 0) AS "plan.id", COALESCE(p.title, '') AS "plan.title"
		FROM companies c
		LEFT JOIN subscriptions s ON c.id=s.company_id
		LEFT JOIN plans p ON s.plan_id=p.id
		WHERE c.deleted_at IS NULL
	`
	err := r.db.Select(&companies, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return companies, nil
}

// Obtener uns empresa por Id
func (r *CompanyRepository) GetCompanyById(companyId int) (*domain.Companies, error) {
	var company domain.Companies
	query := `
		SELECT
			c.id, c.name, c.logo, c.ruc, c.email, c.phone, c.web_site, c.address, c.status,
			COALESCE(p.id, 0) AS "plan.id", COALESCE(p.title, '') AS "plan.title"
		FROM companies c
		LEFT JOIN subscriptions s ON c.id=s.company_id
		LEFT JOIN plans p ON s.plan_id=p.id
		WHERE c.id = $1 AND c.deleted_at IS NULL
	`

	// Usar Get para mapear automáticamente los resultados
	err := r.db.Get(&company, query, companyId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Empresa no encontrada
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &company, nil
}

// Actualizar uns empresa
func (r *CompanyRepository) UpdateCompany(company *domain.Companies) error {
	query := `
		UPDATE companies SET 
			name = $1, ruc = $2, email = $3, phone = $4, web_site = $5, address = $6, status = $7, updated_at = NOW()
		WHERE id = $8 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, company.Name, company.Ruc, company.Email, company.Phone, company.Web_Site, company.Address, company.Status, company.Id)
	if err != nil {
		fmt.Printf("Error actualizando la empresa con Id %d: %v\n", company.Id, err)
		return err
	}
	return nil
}

// Eliminar uns empresa por Id (eliminación lógica)
func (r *CompanyRepository) DeleteCompanyById(id int) error {
	query := `
		UPDATE companies
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la empresa con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples empresas por Ids (eliminación lógica)
func (r *CompanyRepository) DeleteCompaniesByIds(ids []int) error {
	query := `UPDATE companies SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
