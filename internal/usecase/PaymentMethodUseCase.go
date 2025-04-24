package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type PaymentMethodUseCase struct {
	PaymentMethodRepo *repository.PaymentMethodRepository
}

func NewPaymentMethodUseCase(PaymentMethodRepo *repository.PaymentMethodRepository) *PaymentMethodUseCase {
	return &PaymentMethodUseCase{PaymentMethodRepo: PaymentMethodRepo}
}

func (uc *PaymentMethodUseCase) CreatePaymentMethod(paymentMethod *domain.PaymentMethods) error {
	return uc.PaymentMethodRepo.CreatePaymentMethod(paymentMethod)
}

func (uc *PaymentMethodUseCase) GetAllPaymentMethods() ([]domain.PaymentMethods, error) {
	return uc.PaymentMethodRepo.GetAllPaymentMethods()
}

func (uc *PaymentMethodUseCase) GetPaymentMethodById(paymentMethodId int) (*domain.PaymentMethods, error) {
	return uc.PaymentMethodRepo.GetPaymentMethodById(paymentMethodId)
}

func (uc *PaymentMethodUseCase) UpdatePaymentMethod(PaymentMethod *domain.PaymentMethods) error {
	return uc.PaymentMethodRepo.UpdatePaymentMethod(PaymentMethod)
}

func (uc *PaymentMethodUseCase) DeletePaymentMethodById(id int) error {
	return uc.PaymentMethodRepo.DeletePaymentMethodById(id)
}

func (uc *PaymentMethodUseCase) DeletePaymentMethodsByIds(ids []int) error {
	return uc.PaymentMethodRepo.DeletePaymentMethodsByIds(ids)
}
