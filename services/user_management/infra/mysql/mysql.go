package mysql

import (
	"go-micro-service/services/user_management/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.MysqlDSN
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
