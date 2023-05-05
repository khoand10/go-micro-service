package mysql

import (
	"go-micro-service/services/user_management/domain/entity"
	"go-micro-service/services/user_management/domain/repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u userRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
