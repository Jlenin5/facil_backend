package usecaseHumanresources

import (
	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	repositoryHumanresources "github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type PayrollUseCase struct {
	PayrollRepo *repositoryHumanresources.PayrollRepository
}

func NewPayrollUseCase(PayrollRepo *repositoryHumanresources.PayrollRepository) *PayrollUseCase {
	return &PayrollUseCase{PayrollRepo: PayrollRepo}
}

func (uc *PayrollUseCase) Create(payroll *humanresources.Payroll) error {
	return uc.PayrollRepo.Create(payroll)
}

func (uc *PayrollUseCase) GetAll() ([]humanresources.Payroll, error) {
	return uc.PayrollRepo.GetAll()
}

func (uc *PayrollUseCase) GetById(id int) (*humanresources.Payroll, error) {
	return uc.PayrollRepo.GetById(id)
}

func (uc *PayrollUseCase) Approve(id int, approvedBy int) error {
	return uc.PayrollRepo.Approve(id, approvedBy)
}
