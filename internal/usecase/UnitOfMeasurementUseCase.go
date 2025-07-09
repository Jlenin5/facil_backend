package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type UnitOfMeasurementUseCase struct {
	UnitOfMeasurementRepo *repository.UnitOfMeasurementRepository
}

func NewUnitOfMeasurementUseCase(UnitOfMeasurementRepo *repository.UnitOfMeasurementRepository) *UnitOfMeasurementUseCase {
	return &UnitOfMeasurementUseCase{UnitOfMeasurementRepo: UnitOfMeasurementRepo}
}

func (uc *UnitOfMeasurementUseCase) CreateUnitOfMeasurement(brand *domain.UnitsOfMeasurement) error {
	return uc.UnitOfMeasurementRepo.CreateUnitOfMeasurement(brand)
}

func (uc *UnitOfMeasurementUseCase) GetAllUnitsOfMeasurement() ([]domain.UnitsOfMeasurement, error) {
	return uc.UnitOfMeasurementRepo.GetAllUnitsOfMeasurement()
}

func (uc *UnitOfMeasurementUseCase) GetUnitOfMeasurementById(brandId int) (*domain.UnitsOfMeasurement, error) {
	return uc.UnitOfMeasurementRepo.GetUnitOfMeasurementById(brandId)
}

func (uc *UnitOfMeasurementUseCase) UpdateUnitOfMeasurement(UnitOfMeasurement *domain.UnitsOfMeasurement) error {
	return uc.UnitOfMeasurementRepo.UpdateUnitOfMeasurement(UnitOfMeasurement)
}

func (uc *UnitOfMeasurementUseCase) DeleteUnitOfMeasurementById(id int) error {
	return uc.UnitOfMeasurementRepo.DeleteUnitOfMeasurementById(id)
}

func (uc *UnitOfMeasurementUseCase) DeleteUnitsOfMeasurementByIds(ids []int) error {
	return uc.UnitOfMeasurementRepo.DeleteUnitsOfMeasurementByIds(ids)
}
