package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type MetricValue struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Metric        Metric    `json:"metric" has_one:"metric"`
	MetricID      uuid.UUID `json:"metric_id" db:"metric_id"`
	Value         Value     `json:"value"  has_one:"value"`
	ValueID       uuid.UUID `json:"value_id" db:"value_id"`
	OpportunityID uuid.UUID `json:"opportunity_id" db:"opportunity_id"`
}

// String is not required by pop and may be deleted
func (m MetricValue) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// MetricValues is not required by pop and may be deleted
type MetricValues []MetricValue

// String is not required by pop and may be deleted
func (m MetricValues) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *MetricValue) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *MetricValue) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *MetricValue) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
