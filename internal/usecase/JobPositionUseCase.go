package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type JobPositionUseCase struct {
	JobPositionRepo *repository.JobPositionRepository
}

func NewJobPositionUseCase(JobPositionRepo *repository.JobPositionRepository) *JobPositionUseCase {
	return &JobPositionUseCase{JobPositionRepo: JobPositionRepo}
}

func (uc *JobPositionUseCase) CreateJobPosition(jobPosition *domain.JobPositions) error {
	return uc.JobPositionRepo.CreateJobPosition(jobPosition)
}

func (uc *JobPositionUseCase) GetAllJobPositions() ([]domain.JobPositions, error) {
	return uc.JobPositionRepo.GetAllJobPositions()
}

func (uc *JobPositionUseCase) GetJobPositionById(jobPositionId int) (*domain.JobPositions, error) {
	return uc.JobPositionRepo.GetJobPositionById(jobPositionId)
}

func (uc *JobPositionUseCase) UpdateJobPosition(jobPosition *domain.JobPositions) error {
	return uc.JobPositionRepo.UpdateJobPosition(jobPosition)
}

func (uc *JobPositionUseCase) DeleteJobPositionById(id int) error {
	return uc.JobPositionRepo.DeleteJobPositionById(id)
}

func (uc *JobPositionUseCase) DeleteJobPositionsByIds(ids []int) error {
	return uc.JobPositionRepo.DeleteJobPositionsByIds(ids)
}
