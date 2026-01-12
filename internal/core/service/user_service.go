package service

import (
	"compro/config"
	"compro/internal/adapter/repository"
	"compro/internal/core/domain/entity"
	"compro/utils/auth"
	"compro/utils/conv"
	"context"
	"errors"

	"github.com/rs/zerolog/log"
)

var (
	err  error
	code string
)

type IUserService interface {
	LoginAdmin(ctx context.Context, req entity.UserEntity) (string, error)
}

type UserService struct {
	userRepo repository.IUserRepository
	cfg      *config.Config
	jwtAuth  auth.IJwt
}

// LoginAdmin implements [IUserService].
func (u *UserService) LoginAdmin(ctx context.Context, req entity.UserEntity) (string, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		code = "[SERVICE] LoginAdmin - 1"
		log.Err(err).Msg(code)
		return "", err
	}

	if checkPass := conv.CheckPasswordHash(req.Password, user.Password); !checkPass {
		code = "[SERVICE] LoginAdmin - 2"
		err = errors.New("invalid password")
		log.Err(err).Msg(code)
		return "", err
	}

	jwtData := &entity.JwtData{
		UserID: float64(user.ID),
	}

	token, _, err := u.jwtAuth.GenerateToken(jwtData)
	if err != nil {
		code = "[SERVICE] LoginAdmin - 3"
		log.Err(err).Msg(code)
		return "", err
	}

	return token, nil
}

func NewUserService(userRepository repository.IUserRepository, cfg *config.Config, jwtAuth auth.IJwt) IUserService {
	return &UserService{
		userRepo: userRepository, 
		cfg: cfg, 
		jwtAuth: jwtAuth,
	}
}
