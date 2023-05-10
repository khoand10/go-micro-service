package usecase

import (
	"go-micro-service/services/user_management/domain/repository"
)

type (
	GetUserRequest struct {
		UserID uint64
	}

	GetUserResponse struct {
		Id    uint64
		Email string
		Name  string
	}
)

func GetUserDetail(userRepository repository.UserRepository, req GetUserRequest) (*GetUserResponse, error) {
	userFound, err := userRepository.GetByID(req.UserID)
	if err != nil {
		return nil, err
	}
	res := &GetUserResponse{
		Email: userFound.Email,
		Name:  userFound.Name,
		Id:    userFound.ID,
	}
	return res, nil
}
