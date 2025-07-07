package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type VacationBalanceUseCase struct {
	VacationBalanceRepo *repositoryHumanresources.VacationBalanceRepository
}

func NewVacationBalanceUseCase(VacationBalanceRepo *repositoryHumanresources.VacationBalanceRepository) *VacationBalanceUseCase {
	return &VacationBalanceUseCase{VacationBalanceRepo: VacationBalanceRepo}
}

func (uc *VacationBalanceUseCase) Create(ctx context.Context, balance *humanresources.VacationBalance) error {
	return uc.VacationBalanceRepo.Create(ctx, balance)
}

func (uc *VacationBalanceUseCase) GetAll(ctx context.Context) ([]humanresources.VacationBalance, error) {
	return uc.VacationBalanceRepo.GetAll(ctx)
}

func (uc *VacationBalanceUseCase) GetById(ctx context.Context, id int) (*humanresources.VacationBalance, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.VacationBalanceRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *VacationBalanceUseCase) Update(ctx context.Context, request *humanresources.VacationBalance) error {
	return uc.VacationBalanceRepo.Update(ctx, request)
}

func (uc *VacationBalanceUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.VacationBalanceRepo.DeleteById(ctx, id)
}

func (uc *VacationBalanceUseCase) GetByEmployeeAndYear(ctx context.Context, employeeID int, year int) (*humanresources.VacationBalance, error) {
	return uc.VacationBalanceRepo.GetByEmployeeAndYear(ctx, employeeID, year)
}