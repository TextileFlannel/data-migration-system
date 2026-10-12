// Package planner turns a profile plus model predictions into a MappingPlan.
package planner

import (
	"errors"

	"intelligent-data-migration-system/internal/model"
	"intelligent-data-migration-system/internal/schema"
)

// Planner builds a relational mapping plan from a MongoDB profile.
type Planner struct {
	model *model.Model
}

// New returns a Planner backed by the given model.
func New(m *model.Model) *Planner {
	return &Planner{model: m}
}

// Plan predicts a decision per field and assembles the relational plan.
func (p *Planner) Plan(prof *schema.Profile) (*schema.MappingPlan, error) {
	// TODO: для каждого поля features.Extract -> model.Predict -> собрать таблицы/связи.
	return nil, errors.New("planner.Plan: не реализовано")
}
