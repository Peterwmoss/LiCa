package database

import (
	"database/sql"
	"fmt"
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DbConfig struct {
	Addr     string
	Database string
	User     string
	Password string
}

func Get() *bun.DB {
	config := &DbConfig{
		User:     "postgres",
		Password: "postgres",
		Database: "lica",
		Addr:     "localhost:5433",
	}

	return connect(config)
}

func CreateSchema(ctx context.Context) error {
	models := []interface{}{
		(*User)(nil),
		(*Category)(nil),
		(*Item)(nil),
	}

	db := Get()
	defer db.Close()

	for _, model := range models {
    _, err := db.
      NewCreateTable().
      Model(model).
      IfNotExists().
      WithForeignKeys().
      Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func connect(config *DbConfig) *bun.DB {
  dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.User, config.Password, config.Addr, config.Database)
  sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDb, pgdialect.New())
	return db
}
