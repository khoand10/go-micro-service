package versions

import "gorm.io/gorm"

func Version1(tx *gorm.DB) error {
	type User struct {
		gorm.Model
		Email    string
		Name     string
		Password string
	}
	return tx.AutoMigrate(User{})
}
