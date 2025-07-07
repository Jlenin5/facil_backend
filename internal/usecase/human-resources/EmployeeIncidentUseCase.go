package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type EmployeeIncidentUseCase struct {
	EmployeeIncidentRepo *repositoryHumanresources.EmployeeIncidentRepository
}

func NewEmployeeIncidentUseCase(EmployeeIncidentRepo *repositoryHumanresources.EmployeeIncidentRepository) *EmployeeIncidentUseCase {
	return &EmployeeIncidentUseCase{EmployeeIncidentRepo: EmployeeIncidentRepo}
}

func (uc *EmployeeIncidentUseCase) Create(ctx context.Context, incident *humanresources.EmployeeIncident) error {
	return uc.EmployeeIncidentRepo.Create(ctx, incident)
}

func (uc *EmployeeIncidentUseCase) GetAll(ctx context.Context) ([]humanresources.EmployeeIncident, error) {
	return uc.EmployeeIncidentRepo.GetAll(ctx)
}

func (uc *EmployeeIncidentUseCase) GetById(ctx context.Context, id int) (*humanresources.EmployeeIncident, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.EmployeeIncidentRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *EmployeeIncidentUseCase) Update(ctx context.Context, incident *humanresources.EmployeeIncident) error {
	return uc.EmployeeIncidentRepo.Update(ctx, incident)
}

func (uc *EmployeeIncidentUseCase) DeleteById(ctx context.Context, id int) error {
	return uc.EmployeeIncidentRepo.DeleteById(ctx, id)
}

func (uc *EmployeeIncidentUseCase) GetByStatus(ctx context.Context, status string) ([]humanresources.EmployeeIncident, error) {
	return uc.EmployeeIncidentRepo.GetByStatus(ctx, status)
}

func (uc *EmployeeIncidentUseCase) Resolve(ctx context.Context, id int, actionTaken string) error {
	return uc.EmployeeIncidentRepo.Resolve(ctx, id, actionTaken)
}

func (uc *EmployeeIncidentUseCase) Close(ctx context.Context, id int) error {
	return uc.EmployeeIncidentRepo.Close(ctx, id)
}