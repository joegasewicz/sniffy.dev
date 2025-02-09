package migrations

import "github.com/uptrace/bun/migrate"

var migrations = migrate.NewMigrations()

func Migrations() *migrate.Migrations {
	return migrations
}
