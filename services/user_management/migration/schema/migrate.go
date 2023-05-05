package schema

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"go-micro-service/services/user_management/migration/schema/versions"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	option := gormigrate.DefaultOptions
	option.TableName = "schema_migrations"

	m := gormigrate.New(db, option, []*gormigrate.Migration{
		{
			ID:      "version1",
			Migrate: versions.Version1,
		},
	})
	return m.Migrate()
}
