package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type WorkScheduleRepository struct {
	db *sqlx.DB
}

func NewWorkScheduleRepository(db *sqlx.DB) *WorkScheduleRepository {
	return &WorkScheduleRepository{db: db}
}

const (
	workSchedulesTable = "work_schedules"
)

func (r *WorkScheduleRepository) Create(ctx context.Context, schedule *humanresources.WorkSchedule) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			name, description, is_default, status
		) VALUES (
			:name, :description, :is_default, :status
		) RETURNING id`, workSchedulesTable)

	rows, err := r.db.NamedQueryContext(ctx, query, schedule)
	if err != nil {
		return fmt.Errorf("error creating work schedule: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&schedule.Id); err != nil {
			return fmt.Errorf("error getting created ID: %w", err)
		}
	}

	return nil
}

func (r *WorkScheduleRepository) GetAll(ctx context.Context) ([]humanresources.WorkSchedule, error) {
	query := fmt.Sprintf(`
		SELECT 
			id, name, description, is_default, status,
			created_at, updated_at
		FROM %s 
		WHERE deleted_at IS NULL
		ORDER BY name ASC`, workSchedulesTable)

	var schedules []humanresources.WorkSchedule
	if err := r.db.SelectContext(ctx, &schedules, query); err != nil {
		return nil, fmt.Errorf("error getting work schedules: %w", err)
	}

	return schedules, nil
}

func (r *WorkScheduleRepository) GetById(ctx context.Context, id int) (*humanresources.WorkSchedule, error) {
	query := fmt.Sprintf(`
		SELECT 
			id, name, description, is_default, status,
			created_at, updated_at
		FROM %s 
		WHERE id = $1 AND deleted_at IS NULL`, workSchedulesTable)

	var schedule humanresources.WorkSchedule
	if err := r.db.GetContext(ctx, &schedule, query, id); err != nil {
		return nil, fmt.Errorf("error getting work schedule by ID: %w", err)
	}

	return &schedule, nil
}

func (r *WorkScheduleRepository) Update(ctx context.Context, schedule *humanresources.WorkSchedule) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = :name,
			description = :description,
			is_default = :is_default,
			status = :status,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, workSchedulesTable)

	result, err := r.db.NamedExecContext(ctx, query, schedule)
	if err != nil {
		return fmt.Errorf("error updating work schedule: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *WorkScheduleRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, workSchedulesTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting work schedule: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *WorkScheduleRepository) SetAsDefault(ctx context.Context, id int) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	// Primero quitamos el default de todos
	clearDefaultQuery := fmt.Sprintf(`
		UPDATE %s 
		SET is_default = false, updated_at = NOW() 
		WHERE is_default = true AND deleted_at IS NULL`, workSchedulesTable)

	result, err := tx.ExecContext(ctx, clearDefaultQuery)
	if err != nil {
		return fmt.Errorf("error clearing default work schedules: %w", err)
	}

	// Luego establecemos el nuevo default
	setDefaultQuery := fmt.Sprintf(`
		UPDATE %s 
		SET is_default = true, updated_at = NOW() 
		WHERE id = $1 AND deleted_at IS NULL`, workSchedulesTable)

	result, err = tx.ExecContext(ctx, setDefaultQuery, id)
	if err != nil {
		return fmt.Errorf("error setting default work schedule: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}