package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type CashRegisterUseCase struct {
	CashRegisterRepo *repository.CashRegisterRepository
}

func NewCashRegisterUseCase(CashRegisterRepo *repository.CashRegisterRepository) *CashRegisterUseCase {
	return &CashRegisterUseCase{CashRegisterRepo: CashRegisterRepo}
}

func (uc *CashRegisterUseCase) CreateCashRegister(cashRegister *domain.CashRegisters) error {
	return uc.CashRegisterRepo.CreateCashRegister(cashRegister)
}

func (uc *CashRegisterUseCase) GetAllCashRegisters() ([]domain.CashRegisters, error) {
	return uc.CashRegisterRepo.GetAllCashRegisters()
}

func (uc *CashRegisterUseCase) GetCashRegisterById(cashRegisterId int) (*domain.CashRegisters, error) {
	return uc.CashRegisterRepo.GetCashRegisterById(cashRegisterId)
}

func (uc *CashRegisterUseCase) UpdateCashRegister(CashRegister *domain.CashRegisters) error {
	return uc.CashRegisterRepo.UpdateCashRegister(CashRegister)
}

func (uc *CashRegisterUseCase) DeleteCashRegisterById(id int) error {
	return uc.CashRegisterRepo.DeleteCashRegisterById(id)
}

func (uc *CashRegisterUseCase) DeleteCashRegistersByIds(ids []int) error {
	return uc.CashRegisterRepo.DeleteCashRegistersByIds(ids)
}