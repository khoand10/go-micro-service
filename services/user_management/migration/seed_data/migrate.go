package seed_data

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"go-micro-service/services/user_management/migration/seed_data/versions"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	option := gormigrate.DefaultOptions
	option.TableName = "data_migrations"
	m := gormigrate.New(db, option, []*gormigrate.Migration{
		{
			ID:      "default user",
			Migrate: versions.DefaultUser,
		},
	})
	return m.Migrate()
}
