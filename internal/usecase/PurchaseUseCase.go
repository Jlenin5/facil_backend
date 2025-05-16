package usecase

import (
	"fmt"
	_"strconv"
	_"strings"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type PurchaseUseCase struct {
	PurchaseRepo *repository.PurchaseRepository
}

func NewPurchaseUseCase(PurchaseRepo *repository.PurchaseRepository) *PurchaseUseCase {
	return &PurchaseUseCase{PurchaseRepo: PurchaseRepo}
}

// Crear una venta junto con sus detalles
func (uc *PurchaseUseCase) CreatePurchase(purchase *domain.Purchases, details []domain.PurchaseDetails) error {
	if purchase.Reference == "" {
		ref, err := uc.generateReference()
		if err != nil {
			return err
		}
		purchase.Reference = ref
	}
	return uc.PurchaseRepo.CreatePurchase(purchase, details)
}

// Obtener todas las ventas
func (uc *PurchaseUseCase) GetAllPurchases() ([]domain.Purchases, error) {
	return uc.PurchaseRepo.GetAllPurchases()
}

// Obtener una venta por ID junto con sus detalles
func (uc *PurchaseUseCase) GetPurchaseById(purchaseID int) (*domain.Purchases, error) {
	return uc.PurchaseRepo.GetPurchaseById(purchaseID)
}

// Obtener una venta por IDs
func (uc *PurchaseUseCase) GetPurchasesByIds(ids []int) ([]domain.Purchases, error) {
	return uc.PurchaseRepo.GetPurchasesByIds(ids)
}

// Actualizar una venta junto con sus detalles
func (uc *PurchaseUseCase) UpdatePurchase(purchase *domain.Purchases, details []domain.PurchaseDetails) error {
	return uc.PurchaseRepo.UpdatePurchase(purchase, details)
}

// Generar referencia en formato "CP-00001"
func (uc *PurchaseUseCase) generateReference() (string, error) {
	// Obtener la última referencia almacenada en la base de datos
	lastReference, err := uc.PurchaseRepo.GetLastPurchaseReference()
	if err != nil {
		return "", err
	}

	// Extraer el número de la referencia actual
	var lastNumber int
	if lastReference != "" {
		_, err := fmt.Sscanf(lastReference, "CP-%05d", &lastNumber)
		if err != nil {
			return "", fmt.Errorf("error parsing last reference: %v", err)
		}
	}

	// Incrementar el número
	newNumber := lastNumber + 1

	// Formatear como "CP-00001"
	newReference := fmt.Sprintf("CP-%05d", newNumber)

	return newReference, nil
}