package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateVotes, downCreateVotes)
}

func upCreateVotes(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE votes (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			variant_id UUID NOT NULL REFERENCES vote_variants(id) ON DELETE CASCADE,
			UNIQUE(user_id, variant_id),
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);

		CREATE INDEX idx_votes_user_id ON votes(user_id);
		CREATE INDEX idx_votes_variant_id ON votes(variant_id);
	`)
	return err
}

func downCreateVotes(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE votes;")
	return err
}
