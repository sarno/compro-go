package repository

import (
	"compro/internal/core/domain/entity"
	"compro/internal/core/domain/model"
	"context"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type IHeroSectionRepository interface {
	CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
}

type heroSectionRepository struct {
	db *gorm.DB
}

// CreateHeroSection implements [IHeroSectionRepository].
func (h *heroSectionRepository) CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	modelHeroSection := model.HeroSection{
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  &req.PathVideo,
		PathBanner: req.Banner,
	}	

	if err = h.db.Create(&modelHeroSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateHeroSection - 1: %v", err)
		return err
	}

	return nil
}

func NewHeroSectionRepository(db *gorm.DB) IHeroSectionRepository {
	return &heroSectionRepository{db: db}
}
