package usecase

import (
	"errors"
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/domain/repository"
	"go-micro-service/services/user_management/pkg/jwt"
	"go-micro-service/services/user_management/pkg/utils"
)

type (
	LoginRequest struct {
		Email    string
		Password string
	}

	LoginResponse struct {
		Token        string
		RefreshToken string
	}
)

func Login(userRepository repository.UserRepository, input *LoginRequest, cfg *config.Config) (*LoginResponse, error) {
	userFound, err := userRepository.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if err := utils.ComparePassword(userFound.Password, input.Password); err != nil {
		return nil, errors.New("wrong email or password")
	}

	token, err := jwt.CreateJWT(userFound, cfg.JwtSecretKey, cfg.TokenExpirationHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateJWT(userFound, cfg.JwtSecretKey, cfg.RefreshTokenExpirationHour)
	if err != nil {
		return nil, err
	}

	res := &LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
	return res, nil
}
