package repository

import "go-micro-service/services/user_management/domain/entity"

type UserRepository interface {
	GetByEmail(email string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	GetByID(id uint64) (*entity.User, error)
}
