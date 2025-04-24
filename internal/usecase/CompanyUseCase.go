package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type CompanyUseCase struct {
	CompanyRepo *repository.CompanyRepository
}

func NewCompanyUseCase(CompanyRepo *repository.CompanyRepository) *CompanyUseCase {
	return &CompanyUseCase{CompanyRepo: CompanyRepo}
}

func (uc *CompanyUseCase) CreateCompany(company *domain.Companies) error {
	return uc.CompanyRepo.CreateCompany(company)
}

func (uc *CompanyUseCase) GetAllCompanies() ([]domain.Companies, error) {
	return uc.CompanyRepo.GetAllCompanies()
}

func (uc *CompanyUseCase) GetCompanyById(companyId int) (*domain.Companies, error) {
	return uc.CompanyRepo.GetCompanyById(companyId)
}

func (uc *CompanyUseCase) UpdateCompany(company *domain.Companies) error {
	return uc.CompanyRepo.UpdateCompany(company)
}

func (uc *CompanyUseCase) DeleteCompanyById(id int) error {
	return uc.CompanyRepo.DeleteCompanyById(id)
}

func (uc *CompanyUseCase) DeleteCompaniesByIds(ids []int) error {
	return uc.CompanyRepo.DeleteCompaniesByIds(ids)
}
