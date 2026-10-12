// Package model loads a trained LightGBM model (via leaves) and predicts the
// mapping decision for a feature vector.
package model

import (
	"errors"

	"intelligent-data-migration-system/internal/schema"
)

// Model wraps a loaded gradient-boosted model for inference.
type Model struct {
	// TODO: *leaves.Ensemble (github.com/dmitryikh/leaves) или аналог.
}

// Load reads a model file (models/model.txt) produced by training.
func Load(path string) (*Model, error) {
	// TODO: загрузить модель через github.com/dmitryikh/leaves.
	return nil, errors.New("model.Load: не реализовано")
}

// Predict returns the structural decision for one feature vector.
func (m *Model) Predict(features []float64) (schema.Decision, error) {
	// TODO: прогнать features через модель, вернуть класс с макс. вероятностью.
	return "", errors.New("model.Predict: не реализовано")
}
