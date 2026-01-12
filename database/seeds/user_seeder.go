package seeds

import (
	"compro/internal/core/domain/model"
	"compro/utils/conv"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	// Hapus data admin yang ada untuk memastikan data baru yang di-seed
	if err := db.Unscoped().Where("email = ?", "admin@mail.com").Delete(&model.User{}).Error; err != nil {
		log.Fatal().Err(err).Msg("Gagal menghapus admin user yang ada")
	}

	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	admin := model.User{
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: bytes,
	}

	if err = db.Create(&admin).Error; err != nil {
		log.Fatal().Err(err).Msg("Gagal membuat admin user")
	} else {
		log.Info().Msg("Admin user berhasil di-seed")
	}
}