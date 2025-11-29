package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateVoteVariants, downCreateVoteVariants)
}

func upCreateVoteVariants(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE vote_variants (
			id UUID PRIMARY KEY,
			election_id UUID NOT NULL REFERENCES elections(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);

		CREATE INDEX idx_vote_variants_election_id ON vote_variants(election_id);
	`)
	return err
}

func downCreateVoteVariants(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE vote_variants;")
	return err
}
