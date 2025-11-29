package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateElections, downCreateElections)
}

func upCreateElections(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE elections (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			description VARCHAR(512),
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);

		CREATE INDEX idx_elections_user_id ON elections(user_id);
		CREATE INDEX idx_elections_name ON elections(name);
	`)
	return err
}

func downCreateElections(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE elections;")
	return err
}
