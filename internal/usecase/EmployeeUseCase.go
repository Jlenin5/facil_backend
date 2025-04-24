package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type EmployeeUseCase struct {
	EmployeeRepo *repository.EmployeeRepository
}

func NewEmployeeUseCase(EmployeeRepo *repository.EmployeeRepository) *EmployeeUseCase {
	return &EmployeeUseCase{EmployeeRepo: EmployeeRepo}
}

func (uc *EmployeeUseCase) CreateEmployee(employee *domain.Employees) error {
	return uc.EmployeeRepo.CreateEmployee(employee)
}

func (uc *EmployeeUseCase) GetAllEmployees() ([]domain.Employees, error) {
	return uc.EmployeeRepo.GetAllEmployees()
}

func (uc *EmployeeUseCase) GetEmployeeById(employeeId int) (*domain.Employees, error) {
	return uc.EmployeeRepo.GetEmployeeById(employeeId)
}

func (uc *EmployeeUseCase) UpdateEmployee(employee *domain.Employees) error {
	return uc.EmployeeRepo.UpdateEmployee(employee)
}

func (uc *EmployeeUseCase) DeleteEmployeeById(id int) error {
	return uc.EmployeeRepo.DeleteEmployeeById(id)
}

func (uc *EmployeeUseCase) DeleteEmployeesByIds(ids []int) error {
	return uc.EmployeeRepo.DeleteEmployeesByIds(ids)
}
