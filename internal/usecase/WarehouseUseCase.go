package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type WarehouseUseCase struct {
	WarehouseRepo *repository.WarehouseRepository
}

func NewWarehouseUseCase(WarehouseRepo *repository.WarehouseRepository) *WarehouseUseCase {
	return &WarehouseUseCase{WarehouseRepo: WarehouseRepo}
}

func (uc *WarehouseUseCase) CreateWarehouse(warehouse *domain.Warehouses) error {
	return uc.WarehouseRepo.CreateWarehouse(warehouse)
}

func (uc *WarehouseUseCase) GetAllWarehouses() ([]domain.Warehouses, error) {
	return uc.WarehouseRepo.GetAllWarehouses()
}

func (uc *WarehouseUseCase) GetWarehouseById(warehouseId int) (*domain.Warehouses, error) {
	return uc.WarehouseRepo.GetWarehouseById(warehouseId)
}

func (uc *WarehouseUseCase) UpdateWarehouse(warehouse *domain.Warehouses) error {
	return uc.WarehouseRepo.UpdateWarehouse(warehouse)
}

func (uc *WarehouseUseCase) DeleteWarehouseById(id int) error {
	return uc.WarehouseRepo.DeleteWarehouseById(id)
}

func (uc *WarehouseUseCase) DeleteWarehousesByIds(ids []int) error {
	return uc.WarehouseRepo.DeleteWarehousesByIds(ids)
}
