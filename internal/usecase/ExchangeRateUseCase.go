package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type ExchangeRateUseCase struct {
	ExchangeRateRepo *repository.ExchangeRateRepository
}

func NewExchangeRateUseCase(ExchangeRateRepo *repository.ExchangeRateRepository) *ExchangeRateUseCase {
	return &ExchangeRateUseCase{ExchangeRateRepo: ExchangeRateRepo}
}

func (uc *ExchangeRateUseCase) CreateExchangeRate(exchangeRate *domain.ExchangeRates) error {
	return uc.ExchangeRateRepo.CreateExchangeRate(exchangeRate)
}

func (uc *ExchangeRateUseCase) GetAllExchangeRates() ([]domain.ExchangeRates, error) {
	return uc.ExchangeRateRepo.GetAllExchangeRates()
}

func (uc *ExchangeRateUseCase) GetExchangeRateById(exchangeRateId int) (*domain.ExchangeRates, error) {
	return uc.ExchangeRateRepo.GetExchangeRateById(exchangeRateId)
}

func (uc *ExchangeRateUseCase) UpdateExchangeRate(ExchangeRate *domain.ExchangeRates) error {
	return uc.ExchangeRateRepo.UpdateExchangeRate(ExchangeRate)
}

func (uc *ExchangeRateUseCase) DeleteExchangeRateById(id int) error {
	return uc.ExchangeRateRepo.DeleteExchangeRateById(id)
}

func (uc *ExchangeRateUseCase) DeleteExchangeRatesByIds(ids []int) error {
	return uc.ExchangeRateRepo.DeleteExchangeRatesByIds(ids)
}
