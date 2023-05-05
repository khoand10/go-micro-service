package repository

import (
	"errors"
	"go-micro-service/services/user_management/domain/entity"
)

type fakeDBRepositoryImpl struct{}

func (f fakeDBRepositoryImpl) Create(user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (f fakeDBRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	fixUserEmail := &entity.User{
		Email:    "khoand@gmail.com",
		Name:     "khoand",
		Password: "123456",
	}

	if email == fixUserEmail.Email {
		return fixUserEmail, nil
	}
	return nil, errors.New("user not found")
}

func NewFakeDBRepositoryImpl() UserRepository {
	return &fakeDBRepositoryImpl{}
}
