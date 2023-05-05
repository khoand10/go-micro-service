package repository

import "go-micro-service/services/user_management/domain/entity"

type UserRepository interface {
	GetByEmail(email string) (*entity.User, error)
}
