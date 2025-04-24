package usecase

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type PurchaseRequestUseCase struct {
	PurchaseRequestRepo *repository.PurchaseRequestRepository
}

func NewPurchaseRequestUseCase(PurchaseRequestRepo *repository.PurchaseRequestRepository) *PurchaseRequestUseCase {
	return &PurchaseRequestUseCase{PurchaseRequestRepo: PurchaseRequestRepo}
}

// Crear una orden de venta junto con sus detalles
func (uc *PurchaseRequestUseCase) CreatePurchaseRequest(purchaseRequest *domain.PurchaseRequests, details []domain.PurchaseRequestDetails) error {
	// Generar referencia si no viene en la solicitud
	if purchaseRequest.Reference == "" {
		ref, err := uc.generateReference()
		if err != nil {
			return err
		}
		purchaseRequest.Reference = ref
	}
	return uc.PurchaseRequestRepo.CreatePurchaseRequest(purchaseRequest, details)
}

// Obtener todas las órdenes de venta
func (uc *PurchaseRequestUseCase) GetAllPurchaseRequests() ([]domain.PurchaseRequests, error) {
	return uc.PurchaseRequestRepo.GetAllPurchaseRequests()
}

// Obtener una orden de venta por ID junto con sus detalles
func (uc *PurchaseRequestUseCase) GetPurchaseRequestById(purchaseRequestID int) (*domain.PurchaseRequests, error) {
	return uc.PurchaseRequestRepo.GetPurchaseRequestById(purchaseRequestID)
}

// Actualizar una orden de venta junto con sus detalles
func (uc *PurchaseRequestUseCase) UpdatePurchaseRequest(purchaseRequest *domain.PurchaseRequests, details []domain.PurchaseRequestDetails) error {
	return uc.PurchaseRequestRepo.UpdatePurchaseRequest(purchaseRequest, details)
}

// Eliminar (suavemente) una orden de venta
func (uc *PurchaseRequestUseCase) DeletePurchaseRequest(purchaseRequestID int) error {
	return uc.PurchaseRequestRepo.DeletePurchaseRequest(purchaseRequestID)
}

// Generar referencia en formato "PR000001"
func (uc *PurchaseRequestUseCase) generateReference() (string, error) {
	// Obtener la última referencia almacenada en la base de datos
	lastReference, err := uc.PurchaseRequestRepo.GetLastPurchaseRequestReference()
	if err != nil {
		return "", err
	}

	// Extraer el número de la referencia actual
	var lastNumber int
	if lastReference != "" {
		_, err := fmt.Sscanf(lastReference, "PR%06d", &lastNumber)
		if err != nil {
			return "", fmt.Errorf("error parsing last reference: %v", err)
		}
	}

	// Incrementar el número
	newNumber := lastNumber + 1

	// Formatear como "PR000001"
	newReference := fmt.Sprintf("PR%06d", newNumber)

	return newReference, nil
}
