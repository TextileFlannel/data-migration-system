// Package migrator executes a MappingPlan: creates the PostgreSQL schema and
// copies data from MongoDB in batches.
package migrator

import (
	"context"
	"errors"

	"intelligent-data-migration-system/internal/config"
	"intelligent-data-migration-system/internal/schema"
)

// Migrator applies a plan against source MongoDB and target PostgreSQL.
type Migrator struct {
	cfg config.Config
}

// New returns a Migrator for the given config.
func New(cfg config.Config) *Migrator {
	return &Migrator{cfg: cfg}
}

// Migrate creates the schema and transfers the data per plan.
func (m *Migrator) Migrate(ctx context.Context, plan *schema.MappingPlan) error {
	// TODO: сгенерировать DDL, создать таблицы, батчами перенести документы.
	return errors.New("migrator.Migrate: не реализовано")
}
