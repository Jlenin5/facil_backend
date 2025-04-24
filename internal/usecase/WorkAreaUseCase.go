package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type WorkAreaUseCase struct {
	WorkAreaRepo *repository.WorkAreaRepository
}

func NewWorkAreaUseCase(WorkAreaRepo *repository.WorkAreaRepository) *WorkAreaUseCase {
	return &WorkAreaUseCase{WorkAreaRepo: WorkAreaRepo}
}

func (uc *WorkAreaUseCase) CreateWorkArea(workArea *domain.WorkAreas) error {
	return uc.WorkAreaRepo.CreateWorkArea(workArea)
}

func (uc *WorkAreaUseCase) GetAllWorkAreas() ([]domain.WorkAreas, error) {
	return uc.WorkAreaRepo.GetAllWorkAreas()
}

func (uc *WorkAreaUseCase) GetWorkAreaById(workAreaId int) (*domain.WorkAreas, error) {
	return uc.WorkAreaRepo.GetWorkAreaById(workAreaId)
}

func (uc *WorkAreaUseCase) UpdateWorkArea(workArea *domain.WorkAreas) error {
	return uc.WorkAreaRepo.UpdateWorkArea(workArea)
}

func (uc *WorkAreaUseCase) DeleteWorkAreaById(id int) error {
	return uc.WorkAreaRepo.DeleteWorkAreaById(id)
}

func (uc *WorkAreaUseCase) DeleteWorkAreasByIds(ids []int) error {
	return uc.WorkAreaRepo.DeleteWorkAreasByIds(ids)
}