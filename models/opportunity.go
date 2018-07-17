package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Opportunity struct {
	ID               uuid.UUID       `json:"id" db:"id"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" db:"updated_at"`
	Name             string          `json:"name" db:"name"`
	Summary          string          `json:"summary" db:"summary"`
	BusinessCategory string          `json:"business_category" db:"business_category"`
	Metrics          Metrics         `json:"metrics" many_to_many:"opportunity_metrics"`
	FatalAttributes  FatalAttributes `json:"fatal_attributes" many_to_many:"opportunity_fatal_attributes"`
}

// String is not required by pop and may be deleted
func (o Opportunity) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Opportunities is not required by pop and may be deleted
type Opportunities []Opportunity

// String is not required by pop and may be deleted
func (o Opportunities) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *Opportunity) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: o.Name, Name: "Name"},
		&validators.StringIsPresent{Field: o.Summary, Name: "Summary"},
		&validators.StringIsPresent{Field: o.BusinessCategory, Name: "BusinessCategory"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *Opportunity) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *Opportunity) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
