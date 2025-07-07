package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type AbsenceTypeUseCase struct {
	AbsenceTypeRepo *repositoryHumanresources.AbsenceTypeRepository
}

func NewAbsenceTypeUseCase(AbsenceTypeRepo *repositoryHumanresources.AbsenceTypeRepository) *AbsenceTypeUseCase {
	return &AbsenceTypeUseCase{AbsenceTypeRepo: AbsenceTypeRepo}
}

func (uc *AbsenceTypeUseCase) Create(ctx context.Context, request *humanresources.AbsenceType) error {
	return uc.AbsenceTypeRepo.Create(ctx, request)
}

func (uc *AbsenceTypeUseCase) GetAll(ctx context.Context) ([]humanresources.AbsenceType, error) {
	requests, err := uc.AbsenceTypeRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get absence requests: %w", err)
	}

	if len(requests) == 0 {
		return nil, nil
	}

	return requests, nil
}

func (uc *AbsenceTypeUseCase) GetById(ctx context.Context, id int) (*humanresources.AbsenceType, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.AbsenceTypeRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *AbsenceTypeUseCase) Update(ctx context.Context, request *humanresources.AbsenceType) error {
	return uc.AbsenceTypeRepo.Update(ctx, request)
}

func (uc *AbsenceTypeUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.AbsenceTypeRepo.DeleteById(ctx, id)
}