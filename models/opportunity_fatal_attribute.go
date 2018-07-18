package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type OpportunityFatalAttribute struct {
	ID               uuid.UUID      `json:"id" db:"id"`
	CreatedAt        time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at" db:"updated_at"`
	OpportunityID    uuid.UUID      `json:"opportunity_id" db:"opportunity_id"`
	Opportunity      Opportunity    `db:"-"`
	FatalAttributeID uuid.UUID      `json:"fatal_attribute_id" db:"fatal_attribute_id"`
	FatalAttribute   FatalAttribute `db:"-"`
}

// String is not required by pop and may be deleted
func (o OpportunityFatalAttribute) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OpportunityFatalAttributes is not required by pop and may be deleted
type OpportunityFatalAttributes []OpportunityFatalAttribute

// String is not required by pop and may be deleted
func (o OpportunityFatalAttributes) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OpportunityFatalAttribute) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OpportunityFatalAttribute) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OpportunityFatalAttribute) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
