package usecaseHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/Jlenin5/facil_backend/internal/repository/human-resources"
)

type PerformanceReviewUseCase struct {
	PerformanceReviewRepo *repositoryHumanresources.PerformanceReviewRepository
}

func NewPerformanceReviewUseCase(PerformanceReviewRepo *repositoryHumanresources.PerformanceReviewRepository) *PerformanceReviewUseCase {
	return &PerformanceReviewUseCase{PerformanceReviewRepo: PerformanceReviewRepo}
}

func (uc *PerformanceReviewUseCase) Create(ctx context.Context, review *humanresources.PerformanceReview) error {
	return uc.PerformanceReviewRepo.Create(ctx, review)
}

func (uc *PerformanceReviewUseCase) GetAll(ctx context.Context) ([]humanresources.PerformanceReview, error) {
	return uc.PerformanceReviewRepo.GetAll(ctx)
}

func (uc *PerformanceReviewUseCase) GetById(ctx context.Context, id int) (*humanresources.PerformanceReview, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	request, err := uc.PerformanceReviewRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get absence request: %w", err)
	}

	return request, nil
}

func (uc *PerformanceReviewUseCase) Update(ctx context.Context, request *humanresources.PerformanceReview) error {
	return uc.PerformanceReviewRepo.Update(ctx, request)
}

func (uc *PerformanceReviewUseCase) DeleteById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	return uc.PerformanceReviewRepo.DeleteById(ctx, id)
}

func (uc *PerformanceReviewUseCase) CompleteReview(ctx context.Context, id int) error {
	return uc.PerformanceReviewRepo.CompleteReview(ctx, id)
}

func (uc *PerformanceReviewUseCase) Acknowledge(ctx context.Context, id int) error {
	return uc.PerformanceReviewRepo.Acknowledge(ctx, id)
}