// Package verify checks integrity after a migration (row counts, checksums, FKs).
package verify

import (
	"context"
	"errors"

	"intelligent-data-migration-system/internal/config"
	"intelligent-data-migration-system/internal/schema"
)

// Verifier compares source and target after a migration.
type Verifier struct {
	cfg config.Config
}

// New returns a Verifier for the given config.
func New(cfg config.Config) *Verifier {
	return &Verifier{cfg: cfg}
}

// Verify checks that the migration preserved the data.
func (v *Verifier) Verify(ctx context.Context, plan *schema.MappingPlan) error {
	// TODO: сверить количество строк, контрольные суммы, целостность FK.
	return errors.New("verify.Verify: не реализовано")
}
