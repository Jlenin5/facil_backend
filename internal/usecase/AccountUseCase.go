package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type AccountUseCase struct {
	AccountRepo *repository.AccountRepository
}

func NewAccountUseCase(AccountRepo *repository.AccountRepository) *AccountUseCase {
	return &AccountUseCase{AccountRepo: AccountRepo}
}

func (uc *AccountUseCase) GetAccountById(userId int) (*domain.Account, error) {
	return uc.AccountRepo.GetAccountById(userId)
}

func (uc *AccountUseCase) UpdateAccount(user *domain.Account) error {
	return uc.AccountRepo.UpdateAccount(user)
}

func (uc *AccountUseCase) GetChangePassowrdById(userId int) (*domain.ChangePassword, error) {
	return uc.AccountRepo.GetChangePassowrdById(userId)
}

func (uc *AccountUseCase) UpdateChangePassowrd(user *domain.ChangePassword) error {
	return uc.AccountRepo.UpdateChangePassowrd(user)
}

func (uc *AccountUseCase) GetPlanBillingById(userId int) (*domain.PlanBilling, error) {
	return uc.AccountRepo.GetPlanBillingById(userId)
}