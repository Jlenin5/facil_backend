package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type SystemUseCase struct {
	SystemRepo *repository.SystemRepository
}

func NewSystemUseCase(SystemRepo *repository.SystemRepository) *SystemUseCase {
	return &SystemUseCase{SystemRepo: SystemRepo}
}

func (uc *SystemUseCase) CreateSystem(system *domain.Systems) error {
	return uc.SystemRepo.CreateSystem(system)
}

func (uc *SystemUseCase) GetAllSystems() ([]domain.Systems, error) {
	return uc.SystemRepo.GetAllSystems()
}

func (uc *SystemUseCase) GetSystemById(systemId int) (*domain.Systems, error) {
	return uc.SystemRepo.GetSystemById(systemId)
}

func (uc *SystemUseCase) UpdateSystem(system *domain.Systems) error {
	return uc.SystemRepo.UpdateSystem(system)
}