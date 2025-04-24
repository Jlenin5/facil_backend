package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type CashMovementUseCase struct {
	CashMovementRepo *repository.CashMovementRepository
}

func NewCashMovementUseCase(CashMovementRepo *repository.CashMovementRepository) *CashMovementUseCase {
	return &CashMovementUseCase{CashMovementRepo: CashMovementRepo}
}

func (uc *CashMovementUseCase) CreateCashMovement(cashMovement *domain.CashMovements) error {
	return uc.CashMovementRepo.CreateCashMovement(cashMovement)
}

func (uc *CashMovementUseCase) GetAllCashMovements() ([]domain.CashMovements, error) {
	return uc.CashMovementRepo.GetAllCashMovements()
}

func (uc *CashMovementUseCase) GetCashMovementById(cashMovementId int) (*domain.CashMovements, error) {
	return uc.CashMovementRepo.GetCashMovementById(cashMovementId)
}

func (uc *CashMovementUseCase) UpdateCashMovement(CashMovement *domain.CashMovements) error {
	return uc.CashMovementRepo.UpdateCashMovement(CashMovement)
}