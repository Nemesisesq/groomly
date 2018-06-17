package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type FatalAttribute struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Weight    int   `json:"weight" db:"weight"`
}

// String is not required by pop and may be deleted
func (f FatalAttribute) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// FatalAttributes is not required by pop and may be deleted
type FatalAttributes []FatalAttribute

// String is not required by pop and may be deleted
func (f FatalAttributes) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *FatalAttribute) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: f.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *FatalAttribute) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *FatalAttribute) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
