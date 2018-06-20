package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type OpportunityMetric struct {
	ID            uuid.UUID   `json:"id" db:"id"`
	CreatedAt     time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at" db:"updated_at"`
	Opportunity   Opportunity `db:"-"`
	OpportunityID uuid.UUID   `json:"opportunity_id" db:"opportunity_id"`
	Metric        Metric      `db:"-"`
	MetricID      uuid.UUID   `json:"metric_id" db:"metric_id"`
}

// String is not required by pop and may be deleted
func (o OpportunityMetric) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OpportunityMetrics is not required by pop and may be deleted
type OpportunityMetrics []OpportunityMetric

// String is not required by pop and may be deleted
func (o OpportunityMetrics) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OpportunityMetric) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OpportunityMetric) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OpportunityMetric) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
