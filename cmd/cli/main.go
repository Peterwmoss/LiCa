package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	_ "github.com/Peterwmoss/LiCa/internal/env"
	_ "github.com/Peterwmoss/LiCa/internal/logger"

	database "github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql/migrations"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

func main() {
	db := database.Get()
	defer db.Close()

	migrator := migrate.NewMigrator(db, migrations.Migrations)

	app := &cli.App{
		Name: "LiCa migration tool",
		Commands: []*cli.Command{
			newDatabaseCommand(migrator),
		},
	}

	app.Run(os.Args)
}

func newDatabaseCommand(migrator *migrate.Migrator) *cli.Command {
	return &cli.Command{
		Name: "migrate",
		Subcommands: []*cli.Command{
			{
				Name: "up",
				Action: func(cCtx *cli.Context) error {
					if err := migrator.Lock(cCtx.Context); err != nil {
						return err
					}
					defer migrator.Unlock(cCtx.Context)

					migrationGroup, err := migrator.Migrate(cCtx.Context)
					if err != nil {
						slog.Error("failed to migrate", "error", err)
						return err
					}
					if migrationGroup.IsZero() {
						slog.Info("no new migrations to run (database is up to date)")
						return nil
					}
					slog.Info(fmt.Sprintf("migrated to %s", migrationGroup))
					return nil
				},
			},
			{
				Name: "down",
				Action: func(cCtx *cli.Context) error {
					if err := migrator.Lock(cCtx.Context); err != nil {
						return err
					}
					defer migrator.Unlock(cCtx.Context)

					migrationGroup, err := migrator.Rollback(cCtx.Context)
					if err != nil {
						slog.Error("failed to rollback", "error", err)
						return err
					}
					if migrationGroup.IsZero() {
						slog.Info("there are no groups to roll back")
						return nil
					}
					slog.Info(fmt.Sprintf("rolled back %s", migrationGroup))
					return nil
				},
			},
			{
				Name: "init",
				Action: func(cCtx *cli.Context) error {
					return migrator.Init(cCtx.Context)
				},
			},
			{
				Name: "create",
				Args: true,
				Action: func(cCtx *cli.Context) error {
					args := cCtx.Args().Slice()
					name := strings.Join(args, "_")

					file, err := migrator.CreateGoMigration(cCtx.Context, name)
					if err != nil {
						return err
					}

					slog.Info("Migration file created", "path", file.Path)
					return nil
				},
			},
		},
	}
}
