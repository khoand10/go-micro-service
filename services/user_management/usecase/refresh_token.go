package usecase

import (
	"errors"
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/domain/repository"
	"go-micro-service/services/user_management/pkg/jwt"
	"go-micro-service/services/user_management/protos"
)

func RefreshToken(userRepository repository.UserRepository, input *protos.RefreshTokenRequest, cfg *config.Config) (*protos.RefreshTokenResponse, error) {
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

	res := &protos.RefreshTokenResponse{
		Token:        newToken,
		RefreshToken: input.RefreshToken,
	}
	return res, nil
}
