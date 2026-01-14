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
	FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error)
	FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error)
	EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	DeleteByIDHeroSection(ctx context.Context, id int64) error
}

type heroSectionRepository struct {
	db *gorm.DB
}

// DeleteByIDHeroSection implements [IHeroSectionRepository].
func (h *heroSectionRepository) DeleteByIDHeroSection(ctx context.Context, id int64) error {
	modelHeroSection := model.HeroSection{}

	err = h.db.Where("id = ?", id).First(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDHeroSection - 1: %v", err)
		return err
	}

	err = h.db.Delete(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDHeroSection - 2: %v", err)
		return err
	}

	return nil
}

// EditByIDHeroSection implements [IHeroSectionRepository].
func (h *heroSectionRepository) EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	modelHeroSection := model.HeroSection{}
	err = h.db.Where("id =?", req.ID).First(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDHeroSection - 1: %v", err)
		return err
	}

	modelHeroSection.Heading = req.Heading
	modelHeroSection.SubHeading = req.SubHeading
	modelHeroSection.PathVideo = &req.PathVideo
	modelHeroSection.PathBanner = req.Banner

	err = h.db.Save(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDHeroSection - 2: %v", err)
		return err
	}

	return nil
}

// FetchByIDHeroSection implements [IHeroSectionRepository].
func (h *heroSectionRepository) FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error) {
	modelHeroSection := model.HeroSection{}
	err = h.db.Where("id = ?", id).First(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDHeroSection - 1: %v", err)
		return nil, err
	}

	return &entity.HeroSectionEntity{
		ID:         modelHeroSection.ID,
		Heading:    modelHeroSection.Heading,
		SubHeading: modelHeroSection.SubHeading,
		PathVideo:  *modelHeroSection.PathVideo,
		Banner:     modelHeroSection.PathBanner,
	}, nil
}

// FetchAllHeroSection implements [IHeroSectionRepository].
func (h *heroSectionRepository) FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error) {
	modelHeroSection := []model.HeroSection{}
	err = h.db.Select("id", "heading", "sub_heading", "path_video", "path_banner").Find(&modelHeroSection).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllHeroSection - 1: %v", err)
		return nil, err
	}
	var heroSectionEntities []entity.HeroSectionEntity
	for _, v := range modelHeroSection {
		heroSectionEntities = append(heroSectionEntities, entity.HeroSectionEntity{
			ID:         v.ID,
			Heading:    v.Heading,
			SubHeading: v.SubHeading,
			PathVideo:  *v.PathVideo,
			Banner:     v.PathBanner,
		})
	}

	return heroSectionEntities, nil
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
