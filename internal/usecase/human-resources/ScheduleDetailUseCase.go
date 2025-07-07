package usecaseHumanresources

import (
	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type ScheduleDetailUseCase struct {
	ScheduleDetailRepo *repositoryHumanresources.ScheduleDetailRepository
}

func NewScheduleDetailUseCase(ScheduleDetailRepo *repositoryHumanresources.ScheduleDetailRepository) *ScheduleDetailUseCase {
	return &ScheduleDetailUseCase{ScheduleDetailRepo: ScheduleDetailRepo}
}

func (uc *ScheduleDetailUseCase) Create(detail *humanresources.ScheduleDetail) error {
	return uc.ScheduleDetailRepo.Create(detail)
}

func (uc *ScheduleDetailUseCase) GetAll() ([]humanresources.ScheduleDetail, error) {
	return uc.ScheduleDetailRepo.GetAll()
}

func (uc *ScheduleDetailUseCase) Update(detail *humanresources.ScheduleDetail) error {
	return uc.ScheduleDetailRepo.Update(detail)
}

func (uc *ScheduleDetailUseCase) Delete(id int) error {
	return uc.ScheduleDetailRepo.Delete(id)
}