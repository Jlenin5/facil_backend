package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type SaleUseCase struct {
	SaleRepo *repository.SaleRepository
}

func NewSaleUseCase(SaleRepo *repository.SaleRepository) *SaleUseCase {
	return &SaleUseCase{SaleRepo: SaleRepo}
}

func ptrInt64(i int64) *int64 {
	return &i
}
// Crear una venta junto con sus detalles
func (uc *SaleUseCase) CreateSale(sale *domain.Sales, details []domain.SaleDetails) error {
	// Generar comprobante si no viene en la solicitud
	if *sale.Bill.String == "" {
		ref, err := uc.generateBill(sale.Document_Type)
		if err != nil {
			return err
		}

		var document string
		if sale.Document_Type == "ticket" {
			document = "B001"
		} else {
			document = "F001"
		}

		numStr := strings.TrimPrefix(ref, fmt.Sprintf("%s-", document))
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error al convertir a número:", err)
			return nil
		}

		*sale.Series.String = document
		sale.Number = domain.NullInt{
			Int:   ptrInt64(int64(num)),
			Valid: true,
		}
		*sale.Bill.String = ref
	}

	if sale.Sale_Status == "" {
		sale.Sale_Status = "issued"
	}

	return uc.SaleRepo.CreateSale(sale, details)
}

// Obtener todas las ventas
func (uc *SaleUseCase) GetAllSales() ([]domain.Sales, error) {
	return uc.SaleRepo.GetAllSales()
}

// Obtener una venta por ID junto con sus detalles
func (uc *SaleUseCase) GetSaleById(saleID int) (*domain.Sales, error) {
	return uc.SaleRepo.GetSaleById(saleID)
}

// Obtener una venta por IDs
func (uc *SaleUseCase) GetSalesByIds(ids []int) ([]domain.Sales, error) {
	return uc.SaleRepo.GetSalesByIds(ids)
}

// Actualizar una venta junto con sus detalles
func (uc *SaleUseCase) UpdateSale(sale *domain.Sales, details []domain.SaleDetails) error {
	return uc.SaleRepo.UpdateSale(sale, details)
}

func (uc *SaleUseCase) GetSaleByBill(bill string) (*domain.Sales, error) {
	// Buscar la venta en la base de datos por el número de factura
	sale, err := uc.SaleRepo.GetSaleByBill(bill)
	if err != nil {
		return nil, err
	}
	return sale, nil
}

// Generar factura en formato "B001-0000001"
func (uc *SaleUseCase) generateBill(document_type string) (string, error) {
	// Obtener la última factura almacenada en la base de datos
	lastBill, err := uc.SaleRepo.GetLastSaleBill(document_type)
	if err != nil {
		return "", err
	}

	var document string
	if document_type == "ticket" {
		document = "B001"
	} else {
		document = "F001"
	}

	// Extraer el número de la factura actual
	var lastNumber int
	if lastBill != "" {
		_, err := fmt.Sscanf(lastBill, fmt.Sprintf("%s-%%07d", document), &lastNumber)
		if err != nil {
			return "", fmt.Errorf("error parsing last Bill: %v", err)
		}
	}

	// Incrementar el número
	newNumber := lastNumber + 1

	// Formatear como "document-0000001"
	newBill := fmt.Sprintf("%s-%07d", document, newNumber)

	return newBill, nil
}
