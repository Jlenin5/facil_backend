package usecase

import (
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
func (uc *PurchaseUseCase) CreatePurchase(sale *domain.Purchases, details []domain.PurchaseDetails) error {
	if sale.Purchase_Status == "" {
		sale.Purchase_Status = "issued"
	}

	return uc.PurchaseRepo.CreatePurchase(sale, details)
}

// Obtener todas las ventas
func (uc *PurchaseUseCase) GetAllPurchases() ([]domain.Purchases, error) {
	return uc.PurchaseRepo.GetAllPurchases()
}

// Obtener una venta por ID junto con sus detalles
func (uc *PurchaseUseCase) GetPurchaseById(saleID int) (*domain.Purchases, error) {
	return uc.PurchaseRepo.GetPurchaseById(saleID)
}

// Obtener una venta por IDs
func (uc *PurchaseUseCase) GetPurchasesByIds(ids []int) ([]domain.Purchases, error) {
	return uc.PurchaseRepo.GetPurchasesByIds(ids)
}

// Actualizar una venta junto con sus detalles
func (uc *PurchaseUseCase) UpdatePurchase(sale *domain.Purchases, details []domain.PurchaseDetails) error {
	return uc.PurchaseRepo.UpdatePurchase(sale, details)
}

func (uc *PurchaseUseCase) GetPurchaseByBill(bill string) (*domain.Purchases, error) {
	// Buscar la venta en la base de datos por el número de factura
	sale, err := uc.PurchaseRepo.GetPurchaseByBill(bill)
	if err != nil {
		return nil, err
	}
	return sale, nil
}