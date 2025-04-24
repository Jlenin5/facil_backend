package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type PlanUseCase struct {
	PlanRepo *repository.PlanRepository
}

func NewPlanUseCase(PlanRepo *repository.PlanRepository) *PlanUseCase {
	return &PlanUseCase{PlanRepo: PlanRepo}
}

func (uc *PlanUseCase) CreatePlan(plan *domain.Plans) error {
	return uc.PlanRepo.CreatePlan(plan)
}

func (uc *PlanUseCase) GetAllPlans() ([]domain.Plans, error) {
	return uc.PlanRepo.GetAllPlans()
}

func (uc *PlanUseCase) GetPlanById(planId int) (*domain.Plans, error) {
	return uc.PlanRepo.GetPlanById(planId)
}

func (uc *PlanUseCase) UpdatePlan(Plan *domain.Plans) error {
	return uc.PlanRepo.UpdatePlan(Plan)
}

func (uc *PlanUseCase) DeletePlanById(id int) error {
	return uc.PlanRepo.DeletePlanById(id)
}

func (uc *PlanUseCase) DeletePlansByIds(ids []int) error {
	return uc.PlanRepo.DeletePlansByIds(ids)
}
