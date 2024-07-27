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
      ALTER TABLE products DROP COLUMN is_custom;
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
      ALTER TABLE products ADD COLUMN is_custom BOOLEAN DEFAULT FALSE;
    `).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewRaw(`
      UPDATE products p
        SET (is_custom) = (SELECT user_id IS NULL 
                           FROM products 
                           WHERE p.id = id);
    `).Exec(ctx)
		if err != nil {
			return err
		}

		return tx.Commit()
	})
}
