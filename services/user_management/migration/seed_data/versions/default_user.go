package versions

import (
	"go-micro-service/services/user_management/pkg/utils"
	"gorm.io/gorm"
)

func DefaultUser(tx *gorm.DB) error {
	type User struct {
		gorm.Model
		Email    string
		Name     string
		Password string
	}
	hashPassword := utils.HashPassword("123456")
	user := User{Name: "Khoand", Email: "khoand@gmail.com", Password: hashPassword}
	err := tx.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
