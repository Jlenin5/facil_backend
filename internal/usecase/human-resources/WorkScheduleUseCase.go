package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type WorkScheduleUseCase struct {
	WorkScheduleRepo *repositoryHumanresources.WorkScheduleRepository
}

func NewWorkScheduleUseCase(WorkScheduleRepo *repositoryHumanresources.WorkScheduleRepository) *WorkScheduleUseCase {
	return &WorkScheduleUseCase{WorkScheduleRepo: WorkScheduleRepo}
}

func (uc *WorkScheduleUseCase) Create(ctx context.Context, schedule *humanresources.WorkSchedule) error {
	return uc.WorkScheduleRepo.Create(ctx, schedule)
}

func (uc *WorkScheduleUseCase) GetAll(ctx context.Context) ([]humanresources.WorkSchedule, error) {
	return uc.WorkScheduleRepo.GetAll(ctx)
}

func (uc *WorkScheduleUseCase) GetById(ctx context.Context, id int) (*humanresources.WorkSchedule, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.WorkScheduleRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *WorkScheduleUseCase) Update(ctx context.Context, request *humanresources.WorkSchedule) error {
	return uc.WorkScheduleRepo.Update(ctx, request)
}

func (uc *WorkScheduleUseCase) DeleteById(ctx context.Context, id int) error {
	return uc.WorkScheduleRepo.DeleteById(ctx, id)
}