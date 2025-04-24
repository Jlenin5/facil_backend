package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type KeyUseCase struct {
	KeyRepo *repository.KeyRepository
}

func NewKeyUseCase(KeyRepo *repository.KeyRepository) *KeyUseCase {
	return &KeyUseCase{KeyRepo: KeyRepo}
}

func (uc *KeyUseCase) CreateKey(systemSetting *domain.Keys) error {
	return uc.KeyRepo.CreateKey(systemSetting)
}

func (uc *KeyUseCase) GetAllKeys() ([]domain.Keys, error) {
	return uc.KeyRepo.GetAllKeys()
}

func (uc *KeyUseCase) GetKeyById(systemSettingId int) (*domain.Keys, error) {
	return uc.KeyRepo.GetKeyById(systemSettingId)
}

func (uc *KeyUseCase) UpdateKey(systemSetting *domain.Keys) error {
	return uc.KeyRepo.UpdateKey(systemSetting)
}