package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type CurrencyUseCase struct {
	CurrencyRepo *repository.CurrencyRepository
}

func NewCurrencyUseCase(CurrencyRepo *repository.CurrencyRepository) *CurrencyUseCase {
	return &CurrencyUseCase{CurrencyRepo: CurrencyRepo}
}

func (uc *CurrencyUseCase) CreateCurrency(currency *domain.Currencies) error {
	return uc.CurrencyRepo.CreateCurrency(currency)
}

func (uc *CurrencyUseCase) GetAllCurrencies() ([]domain.Currencies, error) {
	return uc.CurrencyRepo.GetAllCurrencies()
}

func (uc *CurrencyUseCase) GetCurrencyById(currencyId int) (*domain.Currencies, error) {
	return uc.CurrencyRepo.GetCurrencyById(currencyId)
}

func (uc *CurrencyUseCase) UpdateCurrency(currency *domain.Currencies) error {
	return uc.CurrencyRepo.UpdateCurrency(currency)
}

func (uc *CurrencyUseCase) DeleteCurrencyById(id int) error {
	return uc.CurrencyRepo.DeleteCurrencyById(id)
}

func (uc *CurrencyUseCase) DeleteCurrenciesByIds(ids []int) error {
	return uc.CurrencyRepo.DeleteCurrenciesByIds(ids)
}
