package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type OpportunityTrackingUseCase struct {
	OpportunityTrackingRepo *repository.OpportunityTrackingRepository
}

func NewOpportunityTrackingUseCase(OpportunityTrackingRepo *repository.OpportunityTrackingRepository) *OpportunityTrackingUseCase {
	return &OpportunityTrackingUseCase{OpportunityTrackingRepo: OpportunityTrackingRepo}
}

func (uc *OpportunityTrackingUseCase) CreateOpportunityTracking(opportunityTracking *domain.OpportunityTracking) error {
	return uc.OpportunityTrackingRepo.CreateOpportunityTracking(opportunityTracking)
}

func (uc *OpportunityTrackingUseCase) GetAllOpportunityTracking() ([]domain.OpportunityTracking, error) {
	return uc.OpportunityTrackingRepo.GetAllOpportunityTracking()
}

func (uc *OpportunityTrackingUseCase) GetOpportunityTrackingById(opportunityTrackingId int) (*domain.OpportunityTracking, error) {
	return uc.OpportunityTrackingRepo.GetOpportunityTrackingById(opportunityTrackingId)
}

func (uc *OpportunityTrackingUseCase) UpdateOpportunityTracking(OpportunityTracking *domain.OpportunityTracking) error {
	return uc.OpportunityTrackingRepo.UpdateOpportunityTracking(OpportunityTracking)
}

func (uc *OpportunityTrackingUseCase) DeleteOpportunityTrackingById(id int) error {
	return uc.OpportunityTrackingRepo.DeleteOpportunityTrackingById(id)
}

func (uc *OpportunityTrackingUseCase) DeleteOpportunityTrackingByIds(ids []int) error {
	return uc.OpportunityTrackingRepo.DeleteOpportunityTrackingByIds(ids)
}
