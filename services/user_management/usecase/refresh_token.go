package usecase

import (
	"errors"
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/domain/repository"
	"go-micro-service/services/user_management/pkg/jwt"
)

type (
	RefreshTokenRequest struct {
		RefreshToken string
	}

	RefreshTokenResponse struct {
		Token        string
		RefreshToken string
	}
)

func RefreshToken(userRepository repository.UserRepository, input *RefreshTokenRequest, cfg *config.Config) (*RefreshTokenResponse, error) {
	claims, err := jwt.VerifyToken(input.RefreshToken, cfg.JwtSecretKey)
	if err != nil {
		return nil, err
	}

	userFound, err := userRepository.GetByID(claims.UserId)
	if err != nil {
		return nil, errors.New("refresh token fail")
	}

	newToken, err := jwt.CreateJWT(userFound, cfg.JwtSecretKey, cfg.TokenExpirationHour)
	if err != nil {
		return nil, err
	}

	res := &RefreshTokenResponse{
		Token:        newToken,
		RefreshToken: input.RefreshToken,
	}
	return res, nil
}
