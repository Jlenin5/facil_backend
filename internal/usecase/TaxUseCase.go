package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type TaxUseCase struct {
	TaxRepo *repository.TaxRepository
}

func NewTaxUseCase(TaxRepo *repository.TaxRepository) *TaxUseCase {
	return &TaxUseCase{TaxRepo: TaxRepo}
}

func (uc *TaxUseCase) CreateTax(tax *domain.Taxes) error {
	return uc.TaxRepo.CreateTax(tax)
}

func (uc *TaxUseCase) GetAllTaxes() ([]domain.Taxes, error) {
	return uc.TaxRepo.GetAllTaxes()
}

func (uc *TaxUseCase) GetTaxById(taxId int) (*domain.Taxes, error) {
	return uc.TaxRepo.GetTaxById(taxId)
}

func (uc *TaxUseCase) UpdateTax(Tax *domain.Taxes) error {
	return uc.TaxRepo.UpdateTax(Tax)
}

func (uc *TaxUseCase) DeleteTaxById(id int) error {
	return uc.TaxRepo.DeleteTaxById(id)
}

func (uc *TaxUseCase) DeleteTaxesByIds(ids []int) error {
	return uc.TaxRepo.DeleteTaxesByIds(ids)
}
