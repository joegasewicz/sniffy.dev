package migrations

import (
	"context"
	"fmt"
	"github.com/joegasewicz/sniffy.dev/api/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations().MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print("[up migrations 00001]")
		_, err := db.NewCreateTable().
			Model((*models.Domain)(nil)).
			IfNotExists().
			WithForeignKeys().
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().
			Model((*models.Path)(nil)).
			IfNotExists().
			WithForeignKeys().
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print("[down migration 00001]")
		_, err := db.NewDropTable().
			Model((*models.Path)(nil)).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewDropTable().
			Model((*models.Domain)(nil)).
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
