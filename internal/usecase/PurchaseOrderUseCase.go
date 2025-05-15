package usecase

import (
	"fmt"
	
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
	if order.Reference == "" {
		ref, err := uc.generateReference()
		if err != nil {
			return err
		}
		order.Reference = ref
	}
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

// Generar referencia en formato "OC-00001"
func (uc *PurchaseOrderUseCase) generateReference() (string, error) {
	// Obtener la última referencia almacenada en la base de datos
	lastReference, err := uc.PurchaseOrderRepo.GetLastPurchaseOrderReference()
	if err != nil {
		return "", err
	}

	// Extraer el número de la referencia actual
	var lastNumber int
	if lastReference != "" {
		_, err := fmt.Sscanf(lastReference, "OC-%05d", &lastNumber)
		if err != nil {
			return "", fmt.Errorf("error parsing last reference: %v", err)
		}
	}

	// Incrementar el número
	newNumber := lastNumber + 1

	// Formatear como "OC-00001"
	newReference := fmt.Sprintf("OC-%05d", newNumber)

	return newReference, nil
}