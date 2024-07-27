package postgresql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DbConfig struct {
	Addr     string
	Database string
	User     string
	Password string
}

func Get() *bun.DB {
	config := &DbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	}

	return connect(config)
}

func connect(config *DbConfig) *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.User, config.Password, config.Addr, config.Database)
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDb, pgdialect.New())

  db.AddQueryHook(bundebug.NewQueryHook(
    bundebug.WithVerbose(true),
    bundebug.FromEnv("BUNDEBUG"),
  ))

	return db
}
