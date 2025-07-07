package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type PerformanceReviewRepository struct {
	db *sqlx.DB
}

func NewPerformanceReviewRepository(db *sqlx.DB) *PerformanceReviewRepository {
	return &PerformanceReviewRepository{db: db}
}

const (
	performanceReviewsTable = "performance_reviews"
	employeesTable7         = "employees"
	usersTable              = "users"
)

func (r *PerformanceReviewRepository) Create(ctx context.Context, review *humanresources.PerformanceReview) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, reviewer_id, review_date, next_review_date,
			performance_score, strengths, areas_for_improvement, 
			comments, status
		) VALUES (
			:employee_id, :reviewer_id, :review_date, :next_review_date,
			:performance_score, :strengths, :areas_for_improvement, 
			:comments, :status
		) RETURNING id`, performanceReviewsTable)

	rows, err := r.db.NamedQueryContext(ctx, query, review)
	if err != nil {
		return fmt.Errorf("error creating performance review: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&review.Id); err != nil {
			return fmt.Errorf("error getting created review ID: %w", err)
		}
	}

	return nil
}

func (r *PerformanceReviewRepository) GetAll(ctx context.Context) ([]humanresources.PerformanceReview, error) {
	query := fmt.Sprintf(`
		SELECT
			pr.id, pr.employee_id, pr.reviewer_id, pr.review_date, 
			pr.next_review_date, pr.performance_score, pr.strengths, 
			pr.areas_for_improvement, pr.comments, pr.status, 
			pr.acknowledged_at, pr.created_at, pr.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s pr
		INNER JOIN %s e ON pr.employee_id = e.id
		WHERE pr.deleted_at IS NULL
		ORDER BY pr.review_date DESC`, 
		performanceReviewsTable, employeesTable7)

	var reviews []humanresources.PerformanceReview
	if err := r.db.SelectContext(ctx, &reviews, query); err != nil {
		return nil, fmt.Errorf("error getting all performance reviews: %w", err)
	}

	return reviews, nil
}

func (r *PerformanceReviewRepository) GetById(ctx context.Context, id int) (*humanresources.PerformanceReview, error) {
	query := fmt.Sprintf(`
		SELECT
			pr.id, pr.employee_id, pr.reviewer_id, pr.review_date, 
			pr.next_review_date, pr.performance_score, pr.strengths, 
			pr.areas_for_improvement, pr.comments, pr.status, 
			pr.acknowledged_at, pr.created_at, pr.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname",
			r.id AS "reviewer.id",
			r.names AS "reviewer.names",
			r.surname AS "reviewer.surname"
		FROM %s pr
		INNER JOIN %s e ON pr.employee_id = e.id
		LEFT JOIN %s r ON pr.reviewer_id = r.id
		WHERE pr.id = $1 AND pr.deleted_at IS NULL`, 
		performanceReviewsTable, employeesTable7, employeesTable7)

	var review humanresources.PerformanceReview
	if err := r.db.GetContext(ctx, &review, query, id); err != nil {
		return nil, fmt.Errorf("error getting performance review by ID: %w", err)
	}

	return &review, nil
}

func (r *PerformanceReviewRepository) Update(ctx context.Context, review *humanresources.PerformanceReview) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			employee_id = :employee_id,
			reviewer_id = :reviewer_id,
			review_date = :review_date,
			next_review_date = :next_review_date,
			performance_score = :performance_score,
			strengths = :strengths,
			areas_for_improvement = :areas_for_improvement,
			comments = :comments,
			status = :status,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, performanceReviewsTable)

	result, err := r.db.NamedExecContext(ctx, query, review)
	if err != nil {
		return fmt.Errorf("error updating performance review: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *PerformanceReviewRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, performanceReviewsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting performance review: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *PerformanceReviewRepository) CompleteReview(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'completed',
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, performanceReviewsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error completing performance review: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *PerformanceReviewRepository) Acknowledge(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'acknowledged',
			acknowledged_at = NOW(),
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, performanceReviewsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error acknowledging performance review: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}