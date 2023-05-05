package usecase

import (
	"errors"
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/domain/entity"
	"go-micro-service/services/user_management/domain/repository"
	"go-micro-service/services/user_management/pkg/utils"
	"go-micro-service/services/user_management/protos"
)

func Register(userRepository repository.UserRepository, input *protos.RegisterRequest, cfg *config.Config) (*protos.RegisterResponse, error) {
	userFound, _ := userRepository.GetByEmail(input.Email)
	if userFound.ID != 0 {
		return nil, errors.New("email already used")
	}

	user := &entity.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: utils.HashPassword(input.Password),
	}
	_, err := userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	res := &protos.RegisterResponse{
		Id:    user.ID,
		Email: user.Email,
	}
	return res, nil
}
