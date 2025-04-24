package usecase

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type QuoteUseCase struct {
	QuoteRepo *repository.QuoteRepository
}

func NewQuoteUseCase(QuoteRepo *repository.QuoteRepository) *QuoteUseCase {
	return &QuoteUseCase{QuoteRepo: QuoteRepo}
}

// Crear una cotización junto con sus detalles
func (uc *QuoteUseCase) CreateQuote(quote *domain.Quotes, details []domain.QuoteDetails) error {
	// Generar referencia si no viene en la solicitud
	if quote.Reference == "" {
		ref, err := uc.generateReference()
		if err != nil {
			return err
		}
		quote.Reference = ref
	}
	return uc.QuoteRepo.CreateQuote(quote, details)
}

// Obtener todas las cotizaciones
func (uc *QuoteUseCase) GetAllQuotes() ([]domain.Quotes, error) {
	return uc.QuoteRepo.GetAllQuotes()
}

// Obtener una cotización por ID junto con sus detalles
func (uc *QuoteUseCase) GetQuoteById(quoteID int) (*domain.Quotes, error) {
	return uc.QuoteRepo.GetQuoteById(quoteID)
}

// Actualizar una cotización junto con sus detalles
func (uc *QuoteUseCase) UpdateQuote(quote *domain.Quotes, details []domain.QuoteDetails) error {
	return uc.QuoteRepo.UpdateQuote(quote, details)
}

// Generar referencia en formato "CT-00001"
func (uc *QuoteUseCase) generateReference() (string, error) {
	// Obtener la última referencia almacenada en la base de datos
	lastReference, err := uc.QuoteRepo.GetLastQuoteReference()
	if err != nil {
		return "", err
	}

	// Extraer el número de la referencia actual
	var lastNumber int
	if lastReference != "" {
		_, err := fmt.Sscanf(lastReference, "CT-%05d", &lastNumber)
		if err != nil {
			return "", fmt.Errorf("error parsing last reference: %v", err)
		}
	}

	// Incrementar el número
	newNumber := lastNumber + 1

	// Formatear como "CT-00001"
	newReference := fmt.Sprintf("CT-%05d", newNumber)

	return newReference, nil
}
