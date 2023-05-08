package usecase

import (
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/domain/repository"
	"go-micro-service/services/user_management/protos"
)

func GetUserProfile(userRepository repository.UserRepository, input *protos.GetUserProfileRequest, cfg *config.Config) (*protos.GetUserProfileResponse, error) {
	userFound, err := userRepository.GetByID(input.UserId)
	if err != nil {
		return nil, err
	}
	res := &protos.GetUserProfileResponse{
		Email: userFound.Email,
		Name:  userFound.Name,
		Id:    userFound.ID,
	}
	return res, nil
}
