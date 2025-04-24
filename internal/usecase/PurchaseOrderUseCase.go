package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type PurchaseOrderUseCase struct {
	PurchaseOrderRepo *repository.PurchaseOrderRepository
}

func NewPurchaseOrderUseCase(PurchaseOrderRepo *repository.PurchaseOrderRepository) *PurchaseOrderUseCase {
	return &PurchaseOrderUseCase{PurchaseOrderRepo: PurchaseOrderRepo}
}

// Crear una orden de compra junto con sus detalles
func (uc *PurchaseOrderUseCase) CreatePurchaseOrder(order *domain.PurchaseOrders, details []domain.PurchaseOrderDetails) error {
	return uc.PurchaseOrderRepo.CreatePurchaseOrder(order, details)
}

// Obtener todas las órdenes de compra
func (uc *PurchaseOrderUseCase) GetAllPurchaseOrders() ([]domain.PurchaseOrders, error) {
	return uc.PurchaseOrderRepo.GetAllPurchaseOrders()
}

// Obtener una orden de compra por ID junto con sus detalles
func (uc *PurchaseOrderUseCase) GetPurchaseOrderByID(orderID int) (*domain.PurchaseOrders, error) {
	return uc.PurchaseOrderRepo.GetPurchaseOrderByID(orderID)
}

// Actualizar una orden de compra junto con sus detalles
func (uc *PurchaseOrderUseCase) UpdatePurchaseOrder(order *domain.PurchaseOrders, details []domain.PurchaseOrderDetails) error {
	return uc.PurchaseOrderRepo.UpdatePurchaseOrder(order, details)
}

// Eliminar (suavemente) una orden de compra
func (uc *PurchaseOrderUseCase) DeletePurchaseOrder(orderID int) error {
	return uc.PurchaseOrderRepo.DeletePurchaseOrder(orderID)
}
