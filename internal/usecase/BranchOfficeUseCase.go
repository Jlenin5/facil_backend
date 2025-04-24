package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type BranchOfficeUseCase struct {
	BranchOfficeRepo *repository.BranchOfficeRepository
}

func NewBranchOfficeUseCase(BranchOfficeRepo *repository.BranchOfficeRepository) *BranchOfficeUseCase {
	return &BranchOfficeUseCase{BranchOfficeRepo: BranchOfficeRepo}
}

func (uc *BranchOfficeUseCase) CreateBranchOffice(branchOffice *domain.BranchOffices) error {
	return uc.BranchOfficeRepo.CreateBranchOffice(branchOffice)
}

func (uc *BranchOfficeUseCase) GetAllBranchOffices() ([]domain.BranchOffices, error) {
	return uc.BranchOfficeRepo.GetAllBranchOffices()
}

func (uc *BranchOfficeUseCase) GetBranchOfficeById(branchOfficeId int) (*domain.BranchOffices, error) {
	return uc.BranchOfficeRepo.GetBranchOfficeById(branchOfficeId)
}

func (uc *BranchOfficeUseCase) UpdateBranchOffice(branchOffice *domain.BranchOffices) error {
	return uc.BranchOfficeRepo.UpdateBranchOffice(branchOffice)
}

func (uc *BranchOfficeUseCase) DeleteBranchOfficeById(id int) error {
	return uc.BranchOfficeRepo.DeleteBranchOfficeById(id)
}

func (uc *BranchOfficeUseCase) DeleteBranchOfficesByIds(ids []int) error {
	return uc.BranchOfficeRepo.DeleteBranchOfficesByIds(ids)
}
