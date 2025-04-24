package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type InventoryMovementUseCase struct {
	InventoryMovementRepo *repository.InventoryMovementRepository
}

func NewInventoryMovementUseCase(InventoryMovementRepo *repository.InventoryMovementRepository) *InventoryMovementUseCase {
	return &InventoryMovementUseCase{InventoryMovementRepo: InventoryMovementRepo}
}

func (uc *InventoryMovementUseCase) CreateInventoryMovement(inventoryMovement *domain.InventoryMovements) error {
	return uc.InventoryMovementRepo.CreateInventoryMovement(inventoryMovement)
}

func (uc *InventoryMovementUseCase) GetAllInventoryMovements() ([]domain.InventoryMovements, error) {
	return uc.InventoryMovementRepo.GetAllInventoryMovements()
}

func (uc *InventoryMovementUseCase) GetInventoryMovementById(inventoryMovementId int) (*domain.InventoryMovements, error) {
	return uc.InventoryMovementRepo.GetInventoryMovementById(inventoryMovementId)
}

func (uc *InventoryMovementUseCase) UpdateInventoryMovement(InventoryMovement *domain.InventoryMovements) error {
	return uc.InventoryMovementRepo.UpdateInventoryMovement(InventoryMovement)
}

func (uc *InventoryMovementUseCase) DeleteInventoryMovementById(id int) error {
	return uc.InventoryMovementRepo.DeleteInventoryMovementById(id)
}

func (uc *InventoryMovementUseCase) DeleteInventoryMovementsByIds(ids []int) error {
	return uc.InventoryMovementRepo.DeleteInventoryMovementsByIds(ids)
}
