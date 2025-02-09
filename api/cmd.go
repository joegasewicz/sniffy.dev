package main

import (
	"context"
	"github.com/joegasewicz/sniffy.dev/api/migrations"
	"github.com/joegasewicz/sniffy.dev/api/utils"
	"github.com/uptrace/bun/migrate"
	"log"
)

func main() {
	db := utils.Database()
	migrator := migrate.NewMigrator(db, migrations.Migrations())

	ctx := context.Background()
	if err := migrator.Init(ctx); err != nil {
		panic(err)
	}
	group, err := migrator.Migrate(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Migrations applied: %v\n", group)
}
