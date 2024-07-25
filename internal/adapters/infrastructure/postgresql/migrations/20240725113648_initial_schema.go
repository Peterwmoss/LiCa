package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      CREATE TABLE users (
        id UUID PRIMARY KEY,
        email TEXT NOT NULL UNIQUE
      );
    `).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      CREATE TABLE products (
        id UUID PRIMARY KEY,
        name TEXT NOT NULL,
        user_id UUID REFERENCES users(id),
        is_custom boolean NOT NULL DEFAULT false,
        UNIQUE(name, user_id)
      );
    `).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      CREATE TABLE categories (
        id UUID PRIMARY KEY,
        name TEXT NOT NULL,
        user_id UUID REFERENCES users(id),
        UNIQUE(name, user_id)
      );
    `).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      CREATE TABLE product_categories (
        product_id UUID REFERENCES products(id),
        user_id UUID REFERENCES users(id),
        category_id UUID REFERENCES categories(id),
        PRIMARY KEY (product_id, user_id, category_id)
      );
    `).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      CREATE TABLE lists (
        id UUID PRIMARY KEY,
        name TEXT NOT NULL,
        user_id UUID NOT NULL REFERENCES users(id),
        UNIQUE(name, user_id)
      );
    `).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      CREATE TABLE list_items (
        id UUID PRIMARY KEY,
        unit TEXT,
        amount DECIMAL NOT NULL DEFAULT 1.0,
        list_id UUID NOT NULL REFERENCES lists(id),
        product_id UUID NOT NULL REFERENCES products(id),
        category_id UUID NOT NULL REFERENCES categories(id),
        UNIQUE(list_id, product_id, category_id)
      );
    `).Exec(ctx)
		if err != nil {
			return err
		}

    return tx.Commit()
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
    
		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      DROP TABLE list_items;
      DROP TABLE lists;
      DROP TABLE product_categories;
      DROP TABLE categories;
      DROP TABLE products;
      DROP TABLE users;
    `).Exec(ctx)
		if err != nil {
			return err
		}

    return tx.Commit()
	})
}
