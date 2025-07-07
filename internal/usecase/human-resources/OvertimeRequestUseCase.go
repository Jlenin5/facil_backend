package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type OvertimeRequestUseCase struct {
	OvertimeRequestRepo *repositoryHumanresources.OvertimeRequestRepository
}

func NewOvertimeRequestUseCase(OvertimeRequestRepo *repositoryHumanresources.OvertimeRequestRepository) *OvertimeRequestUseCase {
	return &OvertimeRequestUseCase{OvertimeRequestRepo: OvertimeRequestRepo}
}

func (uc *OvertimeRequestUseCase) Create(ctx context.Context, request *humanresources.OvertimeRequest) error {
	return uc.OvertimeRequestRepo.Create(ctx, request)
}

func (uc *OvertimeRequestUseCase) GetAll(ctx context.Context) ([]humanresources.OvertimeRequest, error) {
	return uc.OvertimeRequestRepo.GetAll(ctx)
}

func (uc *OvertimeRequestUseCase) GetById(ctx context.Context, id int) (*humanresources.OvertimeRequest, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.OvertimeRequestRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *OvertimeRequestUseCase) Update(ctx context.Context, request *humanresources.OvertimeRequest) error {
	return uc.OvertimeRequestRepo.Update(ctx, request)
}

func (uc *OvertimeRequestUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.OvertimeRequestRepo.DeleteById(ctx, id)
}

func (uc *OvertimeRequestUseCase) Approve(ctx context.Context, id int, approvedBy int) error {
	return uc.OvertimeRequestRepo.Approve(ctx, id, approvedBy)
}

// func (uc *OvertimeRequestUseCase) Reject(ctx context.Context, id int, reason string) error {
// 	return uc.OvertimeRequestRepo.Reject(ctx, id, reason)
// }