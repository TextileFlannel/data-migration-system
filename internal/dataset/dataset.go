// Package dataset builds the training CSV (features + label) from generator
// output: raw documents (JSONL) + a labels file. Признаки считает пакет
// features — тем же кодом, что и на инференсе.
package dataset

import (
	"context"
	"errors"
)

// Options configures a build run.
type Options struct {
	DocsPath   string // JSONL с документами от генератора
	LabelsPath string // метки: путь поля -> решение
	OutCSV     string // train.csv на выходе
}

// Build reads documents + labels, computes features, and writes train.csv.
func Build(ctx context.Context, opts Options) error {
	// TODO: профилировать документы, Extract признаки, присоединить метки, записать CSV.
	return errors.New("dataset.Build: не реализовано")
}
