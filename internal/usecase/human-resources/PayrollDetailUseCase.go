package usecaseHumanresources

import (
	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type PayrollDetailUseCase struct {
	PayrollDetailRepo *repositoryHumanresources.PayrollDetailRepository
}

func NewPayrollDetailUseCase(PayrollDetailRepo *repositoryHumanresources.PayrollDetailRepository) *PayrollDetailUseCase {
	return &PayrollDetailUseCase{PayrollDetailRepo: PayrollDetailRepo}
}

func (uc *PayrollDetailUseCase) GetAll() ([]humanresources.PayrollDetail, error) {
	return uc.PayrollDetailRepo.GetAll()
}

func (uc *PayrollDetailUseCase) GetByPayrollID(payrollID int) ([]humanresources.PayrollDetail, error) {
	return uc.PayrollDetailRepo.GetByPayrollID(payrollID)
}

func (uc *PayrollDetailUseCase) Update(payrollDetail *humanresources.PayrollDetail) error {
	return uc.PayrollDetailRepo.Update(payrollDetail)
}

func (uc *PayrollDetailUseCase) MarkAsPaid(id int) error {
	return uc.PayrollDetailRepo.MarkAsPaid(id)
}