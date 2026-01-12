package service

import (
	"compro/internal/adapter/repository"
	"compro/internal/core/domain/entity"
	"context"
)

type IHeroSectionService interface {
	CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
}

type HeroSectionService struct {
	heroSectionRepo repository.IHeroSectionRepository
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
