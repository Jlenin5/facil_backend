package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type VacationUseCase struct {
	VacationRepo *repositoryHumanresources.VacationRepository
}

func NewVacationUseCase(VacationRepo *repositoryHumanresources.VacationRepository) *VacationUseCase {
	return &VacationUseCase{VacationRepo: VacationRepo}
}

func (uc *VacationUseCase) Create(ctx context.Context, vacation *humanresources.Vacation) error {
	return uc.VacationRepo.Create(ctx, vacation)
}

func (uc *VacationUseCase) GetAll(ctx context.Context) ([]humanresources.Vacation, error) {
	return uc.VacationRepo.GetAll(ctx)
}

func (uc *VacationUseCase) GetById(ctx context.Context, id int) (*humanresources.Vacation, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.VacationRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *VacationUseCase) Update(ctx context.Context, request *humanresources.Vacation) error {
	return uc.VacationRepo.Update(ctx, request)
}

func (uc *VacationUseCase) DeleteById(ctx context.Context, id int) error {
	return uc.VacationRepo.DeleteById(ctx, id)
}

func (uc *VacationUseCase) Approve(ctx context.Context, id int, approvedBy int) error {
	return uc.VacationRepo.Approve(ctx, id, approvedBy)
}

// func (uc *VacationUseCase) Reject(id int, reason string) error {
// 	return uc.VacationRepo.Reject(id, reason)
// }

// func (uc *VacationUseCase) Cancel(id int) error {
// 	return uc.VacationRepo.Cancel(id)
// }