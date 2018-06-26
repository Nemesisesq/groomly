package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Type int

const (
	Benefit Type = iota + 1
	Effort
)

type Metric struct {
	ID          uuid.UUID   `json:"id" db:"id"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	Name        string      `json:"name" db:"name"`
	Weight      int         `json:"weight" db:"weight"`
	Value       Value       `json:"value" has_one:"value"`
	Type        Type        `json:"type" db:"type"`
	Opportunity Opportunity `belongs_to:"opportunity"`
	OpportunityID uuid.UUID `db:"opportunity_id"`
}

// String is not required by pop and may be deleted
func (m Metric) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Metrics is not required by pop and may be deleted
type Metrics []Metric

// String is not required by pop and may be deleted
func (m Metrics) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *Metric) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: m.Name, Name: "Name"},
		&validators.IntIsPresent{Field: m.Weight, Name: "Weight"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *Metric) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *Metric) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
