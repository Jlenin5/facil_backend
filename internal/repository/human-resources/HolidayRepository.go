package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type HolidayRepository struct {
	db *sqlx.DB
}

func NewHolidayRepository(db *sqlx.DB) *HolidayRepository {
	return &HolidayRepository{db: db}
}

const (
	holidaysTable = "holidays"
)

func (r *HolidayRepository) Create(ctx context.Context, holiday *humanresources.Holiday) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			name, description, date, recurring, status
		) VALUES (
			:name, :description, :date, :recurring, :status
		) RETURNING id`, holidaysTable)

	rows, err := r.db.NamedQueryContext(ctx, query, holiday)
	if err != nil {
		return fmt.Errorf("error creating holiday: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&holiday.Id); err != nil {
			return fmt.Errorf("error getting created ID: %w", err)
		}
	}

	return nil
}

func (r *HolidayRepository) GetAll(ctx context.Context) ([]humanresources.Holiday, error) {
	query := fmt.Sprintf(`
		SELECT 
			id, name, description, date, recurring, status, created_at, updated_at
		FROM %s 
		WHERE deleted_at IS NULL 
		ORDER BY date`, holidaysTable)

	var holidays []humanresources.Holiday
	if err := r.db.SelectContext(ctx, &holidays, query); err != nil {
		return nil, fmt.Errorf("error getting holidays: %w", err)
	}

	return holidays, nil
}

func (r *HolidayRepository) GetById(ctx context.Context, id int) (*humanresources.Holiday, error) {
	query := fmt.Sprintf(`
		SELECT 
			id, name, description, date, recurring, status, created_at, updated_at
		FROM %s 
		WHERE id = $1 AND deleted_at IS NULL`, holidaysTable)

	var holiday humanresources.Holiday
	if err := r.db.GetContext(ctx, &holiday, query, id); err != nil {
		return nil, fmt.Errorf("error getting holiday by ID: %w", err)
	}

	return &holiday, nil
}

func (r *HolidayRepository) Update(ctx context.Context, holiday *humanresources.Holiday) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = :name,
			description = :description,
			date = :date,
			recurring = :recurring,
			status = :status,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, holidaysTable)

	result, err := r.db.NamedExecContext(ctx, query, holiday)
	if err != nil {
		return fmt.Errorf("error updating holiday: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *HolidayRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, holidaysTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting holiday: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *HolidayRepository) GetByYear(ctx context.Context, year int) ([]humanresources.Holiday, error) {
	query := fmt.Sprintf(`
		SELECT 
			id, name, description, date, recurring, status, created_at, updated_at
		FROM %s 
		WHERE (EXTRACT(YEAR FROM date) = $1 OR recurring = true)
		AND deleted_at IS NULL
		ORDER BY date`, holidaysTable)

	var holidays []humanresources.Holiday
	if err := r.db.SelectContext(ctx, &holidays, query, year); err != nil {
		return nil, fmt.Errorf("error getting holidays by year: %w", err)
	}

	return holidays, nil
}

func (r *HolidayRepository) Activate(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = true,
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, holidaysTable)

	return r.executeStatusUpdate(ctx, query, id)
}

func (r *HolidayRepository) Deactivate(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = false,
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, holidaysTable)

	return r.executeStatusUpdate(ctx, query, id)
}

func (r *HolidayRepository) executeStatusUpdate(ctx context.Context, query string, args ...interface{}) error {
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error executing status update: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}