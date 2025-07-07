package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type AbsenceRequestUseCase struct {
	absenceRequestRepo *repositoryHumanresources.AbsenceRequestRepository
}

func NewAbsenceRequestUseCase(repo *repositoryHumanresources.AbsenceRequestRepository) *AbsenceRequestUseCase {
	return &AbsenceRequestUseCase{absenceRequestRepo: repo}
}

func (uc *AbsenceRequestUseCase) Create(ctx context.Context, request *humanresources.AbsenceRequest) error {
	if err := validateAbsenceRequest(request); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return uc.absenceRequestRepo.Create(ctx, request)
}

func (uc *AbsenceRequestUseCase) GetAll(ctx context.Context) ([]humanresources.AbsenceRequest, error) {
	requests, err := uc.absenceRequestRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get absence requests: %w", err)
	}

	if len(requests) == 0 {
		return nil, nil
	}

	return requests, nil
}

func (uc *AbsenceRequestUseCase) GetById(ctx context.Context, id int) (*humanresources.AbsenceRequest, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.absenceRequestRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *AbsenceRequestUseCase) Update(ctx context.Context, request *humanresources.AbsenceRequest) error {
	if err := validateAbsenceRequest(request); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return uc.absenceRequestRepo.Update(ctx, request)
}

func (uc *AbsenceRequestUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.absenceRequestRepo.DeleteById(ctx, id)
}

func (uc *AbsenceRequestUseCase) Approve(ctx context.Context, id int, approvedBy int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	if approvedBy <= 0 {
		return fmt.Errorf("invalid approver ID")
	}

	return uc.absenceRequestRepo.Approve(ctx, id, approvedBy)
}

func (uc *AbsenceRequestUseCase) Reject(ctx context.Context, id int, approvedBy int, reason string) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	if approvedBy <= 0 {
		return fmt.Errorf("invalid approver ID")
	}

	if reason == "" {
		return fmt.Errorf("rejection reason cannot be empty")
	}

	return uc.absenceRequestRepo.Reject(ctx, id, approvedBy, reason)
}

// validateAbsenceRequest realiza validaciones básicas de la estructura
func validateAbsenceRequest(request *humanresources.AbsenceRequest) error {
	if request.EmployeeId <= 0 {
		return fmt.Errorf("employee ID is required")
	}

	if request.AbsenceTypeId <= 0 {
		return fmt.Errorf("absence type ID is required")
	}

	if request.StartDate.IsZero() {
		return fmt.Errorf("start date is required")
	}

	if request.EndDate.IsZero() {
		return fmt.Errorf("end date is required")
	}

	if request.EndDate.Before(request.StartDate) {
		return fmt.Errorf("end date cannot be before start date")
	}

	return nil
}
