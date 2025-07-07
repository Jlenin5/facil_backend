package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type HolidayUseCase struct {
	HolidayRepo *repositoryHumanresources.HolidayRepository
}

func NewHolidayUseCase(HolidayRepo *repositoryHumanresources.HolidayRepository) *HolidayUseCase {
	return &HolidayUseCase{HolidayRepo: HolidayRepo}
}

func (uc *HolidayUseCase) Create(ctx context.Context, holiday *humanresources.Holiday) error {
	return uc.HolidayRepo.Create(ctx, holiday)
}

func (uc *HolidayUseCase) GetAll(ctx context.Context) ([]humanresources.Holiday, error) {
	return uc.HolidayRepo.GetAll(ctx)
}

func (uc *HolidayUseCase) GetById(ctx context.Context, id int) (*humanresources.Holiday, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.HolidayRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *HolidayUseCase) Update(ctx context.Context, request *humanresources.Holiday) error {
	return uc.HolidayRepo.Update(ctx, request)
}

func (uc *HolidayUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.HolidayRepo.DeleteById(ctx, id)
}

func (uc *HolidayUseCase) GetByYear(ctx context.Context, year int) ([]humanresources.Holiday, error) {
	return uc.HolidayRepo.GetByYear(ctx, year)
}