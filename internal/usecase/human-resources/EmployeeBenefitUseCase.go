package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type EmployeeBenefitUseCase struct {
	EmployeeBenefitRepo *repositoryHumanresources.EmployeeBenefitRepository
}

func NewEmployeeBenefitUseCase(EmployeeBenefitRepo *repositoryHumanresources.EmployeeBenefitRepository) *EmployeeBenefitUseCase {
	return &EmployeeBenefitUseCase{EmployeeBenefitRepo: EmployeeBenefitRepo}
}

func (uc *EmployeeBenefitUseCase) Create(ctx context.Context, benefit *humanresources.EmployeeBenefit) error {
	return uc.EmployeeBenefitRepo.Create(ctx, benefit)
}

func (uc *EmployeeBenefitUseCase) GetAll(ctx context.Context) ([]humanresources.EmployeeBenefit, error) {
	return uc.EmployeeBenefitRepo.GetAll(ctx)
}

func (uc *EmployeeBenefitUseCase) GetById(ctx context.Context, id int) (*humanresources.EmployeeBenefit, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.EmployeeBenefitRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *EmployeeBenefitUseCase) Update(ctx context.Context, benefit *humanresources.EmployeeBenefit) error {
	return uc.EmployeeBenefitRepo.Update(ctx, benefit)
}

func (uc *EmployeeBenefitUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.EmployeeBenefitRepo.DeleteById(ctx, id)
}

func (uc *EmployeeBenefitUseCase) Deactivate(ctx context.Context, id int) error {
	return uc.EmployeeBenefitRepo.Deactivate(ctx, id)
}