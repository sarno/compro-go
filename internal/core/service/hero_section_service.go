package service

import (
	"compro/internal/adapter/repository"
	"compro/internal/core/domain/entity"
	"context"
)

type IHeroSectionService interface {
	CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error)
	FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error)
	EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	DeleteByIDHeroSection(ctx context.Context, id int64) error
}

type HeroSectionService struct {
	heroSectionRepo repository.IHeroSectionRepository
}

// DeleteByIDHeroSection implements [IHeroSectionService].
func (h *HeroSectionService) DeleteByIDHeroSection(ctx context.Context, id int64) error {
	return h.heroSectionRepo.DeleteByIDHeroSection(ctx, id)
}

// EditByIDHeroSection implements [IHeroSectionService].
func (h *HeroSectionService) EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	return h.heroSectionRepo.EditByIDHeroSection(ctx, req)
}

// FetchByIDHeroSection implements [IHeroSectionService].
func (h *HeroSectionService) FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error) {
	return h.heroSectionRepo.FetchByIDHeroSection(ctx, id)
}

// FetchAllHeroSection implements [IHeroSectionService].
func (h *HeroSectionService) FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error) {
	return h.heroSectionRepo.FetchAllHeroSection(ctx)
}

// CreateHeroSection implements [IHeroSectionService].
func (h *HeroSectionService) CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	return h.heroSectionRepo.CreateHeroSection(ctx, req)
}

func NewHeroSectionService(heroSectionRepo repository.IHeroSectionRepository) IHeroSectionService {
	return &HeroSectionService{
		heroSectionRepo: heroSectionRepo,
	}
}
