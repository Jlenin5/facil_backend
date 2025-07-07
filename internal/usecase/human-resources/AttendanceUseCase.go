package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type AttendanceUseCase struct {
	AttendanceRepo *repositoryHumanresources.AttendanceRepository
}

func NewAttendanceUseCase(AttendanceRepo *repositoryHumanresources.AttendanceRepository) *AttendanceUseCase {
	return &AttendanceUseCase{AttendanceRepo: AttendanceRepo}
}

func (uc *AttendanceUseCase) Create(ctx context.Context, attendance *humanresources.Attendance) error {
	return uc.AttendanceRepo.Create(ctx, attendance)
}

func (uc *AttendanceUseCase) GetAll(ctx context.Context) ([]humanresources.Attendance, error) {
	requests, err := uc.AttendanceRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get absence requests: %w", err)
	}

	if len(requests) == 0 {
		return nil, nil
	}

	return requests, nil
}

func (uc *AttendanceUseCase) GetById(ctx context.Context, id int) (*humanresources.Attendance, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.AttendanceRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *AttendanceUseCase) Update(ctx context.Context, attendance *humanresources.Attendance) error {
	return uc.AttendanceRepo.Update(ctx, attendance)
}

func (uc *AttendanceUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.AttendanceRepo.DeleteById(ctx, id)
}