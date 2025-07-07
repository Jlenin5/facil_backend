package usecaseHumanresources

import (
	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type AttendanceTypeUseCase struct {
	AttendanceTypeRepo *repositoryHumanresources.AttendanceTypeRepository
}

func NewAttendanceTypeUseCase(AttendanceTypeRepo *repositoryHumanresources.AttendanceTypeRepository) *AttendanceTypeUseCase {
	return &AttendanceTypeUseCase{AttendanceTypeRepo: AttendanceTypeRepo}
}

func (uc *AttendanceTypeUseCase) Create(brand *humanresources.AttendanceTypes) error {
	return uc.AttendanceTypeRepo.Create(brand)
}

func (uc *AttendanceTypeUseCase) GetAll() ([]humanresources.AttendanceTypes, error) {
	return uc.AttendanceTypeRepo.GetAll()
}

func (uc *AttendanceTypeUseCase) GetById(brandId int) (*humanresources.AttendanceTypes, error) {
	return uc.AttendanceTypeRepo.GetById(brandId)
}

func (uc *AttendanceTypeUseCase) Update(AttendanceType *humanresources.AttendanceTypes) error {
	return uc.AttendanceTypeRepo.Update(AttendanceType)
}

func (uc *AttendanceTypeUseCase) DeleteById(id int) error {
	return uc.AttendanceTypeRepo.DeleteById(id)
}

func (uc *AttendanceTypeUseCase) DeleteByIds(ids []int) error {
	return uc.AttendanceTypeRepo.DeleteByIds(ids)
}