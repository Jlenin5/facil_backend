package repositoryHumanresources

import (
	"context"
	"fmt"

	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AbsenceTypeRepository struct {
	db *sqlx.DB
}

func NewAbsenceTypeRepository(db *sqlx.DB) *AbsenceTypeRepository {
	return &AbsenceTypeRepository{db: db}
}

const (
	absenceTypesTable = "absence_types"
)

func (r *AbsenceTypeRepository) Create(ctx context.Context, absenceType *humanresources.AbsenceType) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			name, code, description, requires_approval, 
			is_paid, deducts_vacation, status
		) VALUES (
			:name, :code, :description, :requires_approval, 
			:is_paid, :deducts_vacation, :status
		) RETURNING id, created_at, updated_at`, absenceTypesTable)

	rows, err := r.db.NamedQueryContext(ctx, query, absenceType)
	if err != nil {
		return fmt.Errorf("error creating absence type: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&absenceType.Id,
			&absenceType.Created_at,
			&absenceType.Updated_at,
		); err != nil {
			return fmt.Errorf("error getting created absence type data: %w", err)
		}
	}

	return nil
}

func (r *AbsenceTypeRepository) GetAll(ctx context.Context) ([]humanresources.AbsenceType, error) {
	query := fmt.Sprintf(`
		SELECT id, name, code, description, requires_approval, is_paid, deducts_vacation, status, created_at, updated_at FROM %s 
		WHERE deleted_at IS NULL`, absenceTypesTable)

	var absenceTypes []humanresources.AbsenceType
	if err := r.db.SelectContext(ctx, &absenceTypes, query); err != nil {
		return nil, fmt.Errorf("error getting absence types: %w", err)
	}

	return absenceTypes, nil
}

func (r *AbsenceTypeRepository) GetById(ctx context.Context, id int) (*humanresources.AbsenceType, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	query := fmt.Sprintf(`
		SELECT * FROM %s 
		WHERE id = $1 AND deleted_at IS NULL`, absenceTypesTable)

	var absenceType humanresources.AbsenceType
	if err := r.db.GetContext(ctx, &absenceType, query, id); err != nil {
		return nil, fmt.Errorf("error getting absence type by ID: %w", err)
	}

	return &absenceType, nil
}

func (r *AbsenceTypeRepository) Update(ctx context.Context, absenceType *humanresources.AbsenceType) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = :name,
			code = :code,
			description = :description,
			requires_approval = :requires_approval,
			is_paid = :is_paid,
			deducts_vacation = :deducts_vacation,
			status = :status,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, absenceTypesTable)

	result, err := r.db.NamedExecContext(ctx, query, absenceType)
	if err != nil {
		return fmt.Errorf("error updating absence type: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *AbsenceTypeRepository) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, absenceTypesTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting absence type: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *AbsenceTypeRepository) DeleteByIds(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return fmt.Errorf("no IDs provided")
	}

	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = ANY($1) AND deleted_at IS NULL`, absenceTypesTable)

	result, err := r.db.ExecContext(ctx, query, pq.Array(ids))
	if err != nil {
		return fmt.Errorf("error deleting multiple absence types: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, records may not exist")
	}

	return nil
}
