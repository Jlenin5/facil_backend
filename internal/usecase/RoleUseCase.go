package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type RoleUseCase struct {
	RoleRepo *repository.RoleRepository
}

func NewRoleUseCase(RoleRepo *repository.RoleRepository) *RoleUseCase {
	return &RoleUseCase{RoleRepo: RoleRepo}
}

func (uc *RoleUseCase) CreateRole(role *domain.Roles) error {
	return uc.RoleRepo.CreateRole(role)
}

func (uc *RoleUseCase) GetAllRoles() ([]domain.Roles, error) {
	return uc.RoleRepo.GetAllRoles()
}

func (uc *RoleUseCase) GetRoleById(roleId int) (*domain.Roles, error) {
	return uc.RoleRepo.GetRoleById(roleId)
}

func (uc *RoleUseCase) UpdateRole(Role *domain.Roles) error {
	return uc.RoleRepo.UpdateRole(Role)
}

func (uc *RoleUseCase) DeleteRoleById(id int) error {
	return uc.RoleRepo.DeleteRoleById(id)
}

func (uc *RoleUseCase) DeleteRolesByIds(ids []int) error {
	return uc.RoleRepo.DeleteRolesByIds(ids)
}
