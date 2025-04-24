package usecase

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type SaleOrderUseCase struct {
	SaleOrderRepo *repository.SaleOrderRepository
}

func NewSaleOrderUseCase(SaleOrderRepo *repository.SaleOrderRepository) *SaleOrderUseCase {
	return &SaleOrderUseCase{SaleOrderRepo: SaleOrderRepo}
}

// Crear una orden de venta junto con sus detalles
func (uc *SaleOrderUseCase) CreateSaleOrder(order *domain.SaleOrders, details []domain.SaleOrderDetails) error {
	// Generar referencia si no viene en la solicitud
	if order.Reference == "" {
		ref, err := uc.generateReference()
		if err != nil {
			return err
		}
		order.Reference = ref
	}
	return uc.SaleOrderRepo.CreateSaleOrder(order, details)
}

// Obtener todas las órdenes de venta
func (uc *SaleOrderUseCase) GetAllSaleOrders() ([]domain.SaleOrders, error) {
	return uc.SaleOrderRepo.GetAllSaleOrders()
}

// Obtener una orden de venta por ID junto con sus detalles
func (uc *SaleOrderUseCase) GetSaleOrderById(orderID int) (*domain.SaleOrders, error) {
	return uc.SaleOrderRepo.GetSaleOrderById(orderID)
}

// Actualizar una orden de venta junto con sus detalles
func (uc *SaleOrderUseCase) UpdateSaleOrder(order *domain.SaleOrders, details []domain.SaleOrderDetails) error {
	return uc.SaleOrderRepo.UpdateSaleOrder(order, details)
}

// Generar referencia en formato "OV-00001"
func (uc *SaleOrderUseCase) generateReference() (string, error) {
	// Obtener la última referencia almacenada en la base de datos
	lastReference, err := uc.SaleOrderRepo.GetLastSaleOrderReference()
	if err != nil {
		return "", err
	}

	// Extraer el número de la referencia actual
	var lastNumber int
	if lastReference != "" {
		_, err := fmt.Sscanf(lastReference, "OV-%05d", &lastNumber)
		if err != nil {
			return "", fmt.Errorf("error parsing last reference: %v", err)
		}
	}

	// Incrementar el número
	newNumber := lastNumber + 1

	// Formatear como "OV-00001"
	newReference := fmt.Sprintf("OV-%05d", newNumber)

	return newReference, nil
}
