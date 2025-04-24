package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type UserUseCase struct {
	UserRepo *repository.UserRepository
}

func NewUserUseCase(UserRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: UserRepo}
}

func (uc *UserUseCase) CreateUser(user *domain.Users) error {
	return uc.UserRepo.CreateUser(user)
}

func (uc *UserUseCase) GetAllUsers() ([]domain.Users, error) {
	return uc.UserRepo.GetAllUsers()
}

func (uc *UserUseCase) GetUserById(userId int) (*domain.Users, error) {
	return uc.UserRepo.GetUserById(userId)
}

func (uc *UserUseCase) UpdateUser(user *domain.Users) error {
	return uc.UserRepo.UpdateUser(user)
}

func (uc *UserUseCase) DeleteUserById(id int) error {
	return uc.UserRepo.DeleteUserById(id)
}

func (uc *UserUseCase) DeleteUsersByIds(ids []int) error {
	return uc.UserRepo.DeleteUsersByIds(ids)
}
