package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type EmployeeScheduleUseCase struct {
	EmployeeScheduleRepo *repositoryHumanresources.EmployeeScheduleRepository
}

func NewEmployeeScheduleUseCase(EmployeeScheduleRepo *repositoryHumanresources.EmployeeScheduleRepository) *EmployeeScheduleUseCase {
	return &EmployeeScheduleUseCase{EmployeeScheduleRepo: EmployeeScheduleRepo}
}

func (uc *EmployeeScheduleUseCase) Create(ctx context.Context, schedule *humanresources.EmployeeSchedule) error {
	return uc.EmployeeScheduleRepo.Create(ctx, schedule)
}

func (uc *EmployeeScheduleUseCase) GetAll(ctx context.Context) ([]humanresources.EmployeeSchedule, error) {
	return uc.EmployeeScheduleRepo.GetAll(ctx)
}

func (uc *EmployeeScheduleUseCase) GetById(ctx context.Context, id int) (*humanresources.EmployeeSchedule, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.EmployeeScheduleRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *EmployeeScheduleUseCase) Update(ctx context.Context, schedule *humanresources.EmployeeSchedule) error {
	return uc.EmployeeScheduleRepo.Update(ctx, schedule)
}

func (uc *EmployeeScheduleUseCase) DeleteById(ctx context.Context, id int) error {
	return uc.EmployeeScheduleRepo.DeleteById(ctx, id)
}

func (uc *EmployeeScheduleUseCase) Assign(ctx context.Context, schedule *humanresources.EmployeeSchedule) error {
	return uc.EmployeeScheduleRepo.Assign(ctx, schedule)
}

// func (uc *EmployeeScheduleUseCase) EndCurrentSchedule(employeeID int, endDate string) error {
// 	return uc.EmployeeScheduleRepo.EndCurrentSchedule(employeeID, endDate)
// }