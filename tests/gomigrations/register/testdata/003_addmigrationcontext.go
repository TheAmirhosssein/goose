package register

import (
	"context"
	"database/sql"

	"github.com/TheAmirhosssein/goose/v3"
)

func init() {
	goose.AddMigrationContext(
		func(_ context.Context, _ *sql.Tx) error { return nil },
		func(_ context.Context, _ *sql.Tx) error { return nil },
	)
}
