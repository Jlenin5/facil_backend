package repositoryHumanresources

import (
	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ScheduleDetailRepository struct {
	db *sqlx.DB
}

func NewScheduleDetailRepository(db *sqlx.DB) *ScheduleDetailRepository {
	return &ScheduleDetailRepository{db: db}
}

func (r *ScheduleDetailRepository) Create(detail *humanresources.ScheduleDetail) error {
	query := `
		INSERT INTO schedule_details (
			schedule_id, day_of_week, start_time, end_time, is_working_day
		) VALUES (
			:schedule_id, :day_of_week, :start_time, :end_time, :is_working_day
		)`
	_, err := r.db.NamedExec(query, detail)
	return err
}

func (r *ScheduleDetailRepository) GetAll() ([]humanresources.ScheduleDetail, error) {
	var scheduleDetail []humanresources.ScheduleDetail
	query := `SELECT * FROM schedule_details WHERE deleted_at IS NULL`
	err := r.db.Select(&scheduleDetail, query)
	return scheduleDetail, err
}

func (r *ScheduleDetailRepository) GetByScheduleId(scheduleID int) ([]humanresources.ScheduleDetail, error) {
	var details []humanresources.ScheduleDetail
	query := `SELECT * FROM schedule_details WHERE schedule_id = $1 AND deleted_at IS NULL ORDER BY day_of_week`
	err := r.db.Select(&details, query, scheduleID)
	return details, err
}

func (r *ScheduleDetailRepository) Update(detail *humanresources.ScheduleDetail) error {
	query := `
		UPDATE schedule_details SET
			start_time = :start_time,
			end_time = :end_time,
			is_working_day = :is_working_day,
			updated_at = NOW()
		WHERE id = :id`
	_, err := r.db.NamedExec(query, detail)
	return err
}

func (r *ScheduleDetailRepository) Delete(id int) error {
	query := `UPDATE schedule_details SET deleted_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ScheduleDetailRepository) DeleteMultiple(ids []int) error {
	query := `UPDATE schedule_details SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
