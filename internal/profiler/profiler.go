// Package profiler scans MongoDB collections (or JSONL document dumps) and
// aggregates per-field statistics into a schema.Profile.
package profiler

import (
	"context"
	"errors"

	"intelligent-data-migration-system/internal/config"
	"intelligent-data-migration-system/internal/schema"
)

// Profiler builds a schema.Profile from a data source.
type Profiler struct {
	cfg config.Config
}

// New returns a Profiler for the given config.
func New(cfg config.Config) *Profiler {
	return &Profiler{cfg: cfg}
}

// Profile scans the configured MongoDB database and returns its profile.
func (p *Profiler) Profile(ctx context.Context) (*schema.Profile, error) {
	// TODO: подключиться к MongoDB, пройти по коллекциям, собрать FieldProfile.
	return nil, errors.New("profiler.Profile: не реализовано")
}

// ProfileJSONL profiles documents from a JSONL dump — тот же код профилирования,
// что и для живой Mongo, используется при сборке обучающего набора.
func (p *Profiler) ProfileJSONL(ctx context.Context, path string) (*schema.Profile, error) {
	// TODO: прочитать документы из JSONL и собрать профиль.
	return nil, errors.New("profiler.ProfileJSONL: не реализовано")
}
