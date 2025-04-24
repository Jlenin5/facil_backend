package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type CustomerUseCase struct {
	CustomerRepo *repository.CustomerRepository
}

func NewCustomerUseCase(CustomerRepo *repository.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{CustomerRepo: CustomerRepo}
}

func (uc *CustomerUseCase) CreateCustomer(customer *domain.Customers) error {
	return uc.CustomerRepo.CreateCustomer(customer)
}

func (uc *CustomerUseCase) GetAllCustomers() ([]domain.Customers, error) {
	return uc.CustomerRepo.GetAllCustomers()
}

func (uc *CustomerUseCase) GetCustomerById(cutomerId int) (*domain.Customers, error) {
	return uc.CustomerRepo.GetCustomerById(cutomerId)
}

func (uc *CustomerUseCase) UpdateCustomer(customer *domain.Customers) error {
	return uc.CustomerRepo.UpdateCustomer(customer)
}

func (uc *CustomerUseCase) DeleteCustomerById(id int) error {
	return uc.CustomerRepo.DeleteCustomerById(id)
}

func (uc *CustomerUseCase) DeleteCustomersByIds(ids []int) error {
	return uc.CustomerRepo.DeleteCustomersByIds(ids)
}
