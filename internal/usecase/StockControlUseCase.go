package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type StockControlUseCase struct {
	StockControlRepo *repository.StockControlRepository
}

func NewStockControlUseCase(StockControlRepo *repository.StockControlRepository) *StockControlUseCase {
	return &StockControlUseCase{StockControlRepo: StockControlRepo}
}

func (uc *StockControlUseCase) CreateStockControl(stockControl *domain.StockControl) error {
	return uc.StockControlRepo.CreateStockControl(stockControl)
}

func (uc *StockControlUseCase) GetAllStockControl() ([]domain.StockControl, error) {
	return uc.StockControlRepo.GetAllStockControl()
}

func (uc *StockControlUseCase) GetStockControlById(stockControlId int) (*domain.StockControl, error) {
	return uc.StockControlRepo.GetStockControlById(stockControlId)
}

func (uc *StockControlUseCase) UpdateStockControl(StockControl *domain.StockControl) error {
	return uc.StockControlRepo.UpdateStockControl(StockControl)
}

func (uc *StockControlUseCase) DeleteStockControlById(id int) error {
	return uc.StockControlRepo.DeleteStockControlById(id)
}

func (uc *StockControlUseCase) DeleteStockControlByIds(ids []int) error {
	return uc.StockControlRepo.DeleteStockControlByIds(ids)
}
