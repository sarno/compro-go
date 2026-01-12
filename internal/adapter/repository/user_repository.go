package repository

import (
	"compro/internal/core/domain/entity"
	"compro/internal/core/domain/model"
	"context"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	err  error
	code string
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type UserRepository struct {
	db *gorm.DB
}

// GetUserByEmail implements [IUserRepository].
func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var modelUser model.User
	err = u.db.Select("email", "password", "name", "id").Where("email = ?", email).First(&modelUser).Error
	if err != nil {
		code = "[REPOSITORY] GetUserByEmail - 1"
		log.Err(err).Msg(code)
		return nil, err
	}

	return &entity.UserEntity{
		ID:       modelUser.ID,
		Name:     modelUser.Name,
		Email:    modelUser.Email,
		Password: modelUser.Password,
	}, nil
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}
