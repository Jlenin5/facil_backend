package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type SupplierUseCase struct {
	SupplierRepo *repository.SupplierRepository
}

func NewSupplierUseCase(SupplierRepo *repository.SupplierRepository) *SupplierUseCase {
	return &SupplierUseCase{SupplierRepo: SupplierRepo}
}

func (uc *SupplierUseCase) CreateSupplier(supplier *domain.Suppliers) error {
	return uc.SupplierRepo.CreateSupplier(supplier)
}

func (uc *SupplierUseCase) GetAllSuppliers() ([]domain.Suppliers, error) {
	return uc.SupplierRepo.GetAllSuppliers()
}

func (uc *SupplierUseCase) GetSupplierById(supplierId int) (*domain.Suppliers, error) {
	return uc.SupplierRepo.GetSupplierById(supplierId)
}

func (uc *SupplierUseCase) UpdateSupplier(supplier *domain.Suppliers) error {
	return uc.SupplierRepo.UpdateSupplier(supplier)
}

func (uc *SupplierUseCase) DeleteSupplierById(id int) error {
	return uc.SupplierRepo.DeleteSupplierById(id)
}

func (uc *SupplierUseCase) DeleteSuppliersByIds(ids []int) error {
	return uc.SupplierRepo.DeleteSuppliersByIds(ids)
}
