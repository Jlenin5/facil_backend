package usecase

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type SubscriptionUseCase struct {
	SubscriptionRepo *repository.SubscriptionRepository
}

func NewSubscriptionUseCase(SubscriptionRepo *repository.SubscriptionRepository) *SubscriptionUseCase {
	return &SubscriptionUseCase{SubscriptionRepo: SubscriptionRepo}
}

func (uc *SubscriptionUseCase) CreateSubscription(subscription *domain.Subscriptions) error {
	return uc.SubscriptionRepo.CreateSubscription(subscription)
}

func (uc *SubscriptionUseCase) GetAllSubscriptions() ([]domain.Subscriptions, error) {
	return uc.SubscriptionRepo.GetAllSubscriptions()
}

func (uc *SubscriptionUseCase) GetSubscriptionById(subscriptionId int) (*domain.Subscriptions, error) {
	return uc.SubscriptionRepo.GetSubscriptionById(subscriptionId)
}

func (uc *SubscriptionUseCase) UpdateSubscription(Subscription *domain.Subscriptions) error {
	return uc.SubscriptionRepo.UpdateSubscription(Subscription)
}

func (uc *SubscriptionUseCase) DeleteSubscriptionById(id int) error {
	return uc.SubscriptionRepo.DeleteSubscriptionById(id)
}

func (uc *SubscriptionUseCase) DeleteSubscriptionsByIds(ids []int) error {
	return uc.SubscriptionRepo.DeleteSubscriptionsByIds(ids)
}

// Obtener el company_id de un usuario
func (uc *UserUseCase) GetCompanyIDByUserID(userId int) (int, error) {
	user, err := uc.UserRepo.GetUserById(userId)
	if err != nil {
		return 0, err
	}
	return int(*user.Company_Id.Int), nil
}

// Verificar si la suscripción es válida
func (uc *SubscriptionUseCase) IsSubscriptionValid(companyId int) (bool, error) {
	subscription, err := uc.SubscriptionRepo.GetSubscriptionByCompanyId(companyId)
	if err != nil {
		return false, err
	}

	// Si no hay suscripción o está cancelada, no es válida
	if subscription == nil || subscription.Status != "active" {
		return false, nil
	}

	// Verificar si la suscripción ha expirado
	currentTime := time.Now()
	if subscription.End_Date.Before(currentTime) {
		return false, nil
	}

	return true, nil
}