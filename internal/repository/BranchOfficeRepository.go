package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type BranchOfficeRepository struct {
	db *sqlx.DB
}

func NewBranchOfficeRepository(db *sqlx.DB) *BranchOfficeRepository {
	return &BranchOfficeRepository{db: db}
}

// Crear una sucursal
func (r *BranchOfficeRepository) CreateBranchOffice(branchOffice *domain.BranchOffices) error {
	query := "INSERT INTO branch_offices (company_id, name, description, address, phone, status) VALUES (:company_id, :name, :description, :address, :phone, :status)"
	_, err := r.db.NamedExec(query, branchOffice)
	if err != nil {
		fmt.Printf("Error creando la sucursal: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las sucursales
func (r *BranchOfficeRepository) GetAllBranchOffices() ([]domain.BranchOffices, error) {
	var branchOffices []domain.BranchOffices
	query := "SELECT id, company_id, name, description, address, phone, status FROM branch_offices WHERE deleted_at IS NULL"
	err := r.db.Select(&branchOffices, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	// Iterar sobre las sucursales
	for i := range branchOffices {
		// Obtener empresa de la sucursal
		var company domain.Companies
		err = r.db.Get(&company, `
			SELECT id, name, ruc, email, phone, web_site, address, status
			FROM companies
			WHERE id = $1 AND deleted_at IS NULL
		`, branchOffices[i].Company_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay empresa asociada
			fmt.Printf("Error obteniendo empresa para la sucursal %d: %v\n", branchOffices[i].Id, err)
			return nil, err
		}
		branchOffices[i].Company = &company
	}

	return branchOffices, nil
}

// Obtener un sucursal por Id
func (r *BranchOfficeRepository) GetBranchOfficeById(branchOfficeId int) (*domain.BranchOffices, error) {
	query := `SELECT id, company_id, name, description, address, phone, status FROM branch_offices WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, branchOfficeId)

	var branchOffice domain.BranchOffices
	if err := row.Scan(&branchOffice.Id, &branchOffice.Company_Id, &branchOffice.Name, &branchOffice.Description, &branchOffice.Address, &branchOffice.Phone, &branchOffice.Status); err != nil {
		return nil, err
	}

	// Obtener empresa de la sucursal
	var company domain.Companies
	err := r.db.Get(&company, `
		SELECT id, name, ruc, email, phone, web_site, address, status
		FROM companies
		WHERE id = $1 AND deleted_at IS NULL
	`, branchOffice.Company_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay empresa asociada
		fmt.Printf("Error obteniendo empresa para la sucursal %d: %v\n", branchOfficeId, err)
		return nil, err
	}
	branchOffice.Company = &company

	return &branchOffice, nil
}

// Actualizar una sucursal
func (r *BranchOfficeRepository) UpdateBranchOffice(branchOffice *domain.BranchOffices) error {
	query := `
		UPDATE branch_offices SET 
			company_id = $1, name = $2, description = $3, address = $4, phone = $5, status = $6, updated_at = NOW()
		WHERE id = $7 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, branchOffice.Company_Id, branchOffice.Name, branchOffice.Description, branchOffice.Address, branchOffice.Phone, branchOffice.Status, branchOffice.Id)
	if err != nil {
		fmt.Printf("Error actualizando la sucursal con Id %d: %v\n", branchOffice.Id, err)
		return err
	}
	return nil
}

// Eliminar una sucursal por Id (eliminación lógica)
func (r *BranchOfficeRepository) DeleteBranchOfficeById(id int) error {
	query := `
		UPDATE branch_offices
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la sucursal con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples sucursales por Ids (eliminación lógica)
func (r *BranchOfficeRepository) DeleteBranchOfficesByIds(ids []int) error {
	query := `UPDATE branch_offices SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
