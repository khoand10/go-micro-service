package usecase

import (
	"errors"
	"go-micro-service/services/user_management/domain/entity"
	"go-micro-service/services/user_management/domain/repository"
	"go-micro-service/services/user_management/pkg/utils"
)

type (
	RegisterRequest struct {
		Name     string
		Email    string
		Password string
	}

	RegisterResponse struct {
		Id    uint64
		Email string
	}
)

func Register(userRepository repository.UserRepository, input *RegisterRequest) (*RegisterResponse, error) {
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

	res := &RegisterResponse{
		Id:    user.ID,
		Email: user.Email,
	}
	return res, nil
}
