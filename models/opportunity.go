package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
)

type Opportunity struct {
	ID               uuid.UUID       `json:"id" db:"id"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" db:"updated_at"`
	Name             string          `json:"name" db:"name"`
	Summary          string          `json:"summary" db:"summary"`
	BusinessCategory string          `json:"business_category" db:"business_category"`
	MetricValues     MetricValues    `json:"metric_values" has_many:"metric_values"`
	FatalAttributes  FatalAttributes `json:"fatal_attributes" many_to_many:"opportunity_fatal_attributes"`
	Score            int             `json:"score" db:"-"`
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

func (o *Opportunity) ComputeScore() error {
    var total  = 0
    var fa = 1
	for _,  v := range o.MetricValues {
      x := v.Metric.Weight * v.Value.Score

      total = total + x

	}

	for _, v := range o.FatalAttributes {
		fa = fa + v.Weight
	}

	o.Score = total * fa

	return nil
}
func (os Opportunities) ComputeScore()  {

	for key, v := range os {
		v.ComputeScore()
		os[key] = v
		print(key)
	}

	println(os)
}

func (opportunity *Opportunity)PopulateMetricValues(tx *pop.Connection) {
	for i, mv := range opportunity.MetricValues {
		m := &Metric{}
		v := &Value{}
		tx.Eager().Find(m, mv.MetricID)
		tx.Eager().Find(v, mv.ValueID)
		mv.Metric = *m
		mv.Value = *v
		opportunity.MetricValues[i] = mv

	}
}

func(os *Opportunities)PopulateMetricValues(tx *pop.Connection){
	for _, v := range *os {
		v.PopulateMetricValues(tx)
	}
}

func (o *Opportunity)CreateMetricValues(tx *pop.Connection) (verrs *validate.Errors, err error){
	for _, v := range o.MetricValues {

		mv := MetricValue{}
		mv.OpportunityID = o.ID
		mv.MetricID = v.Metric.ID
		mv.ValueID = v.Value.ID

		// link Metrics with The Opportunity

		verrs, err = tx.ValidateAndCreate(mv)

		if err != nil || verrs.HasAny() {
			return verrs, errors.WithStack(err)
		}
	}

	return verrs, err

}

func (o *Opportunity)UpdateMetricValues(tx *pop.Connection) (verrs *validate.Errors, err error){
	for _, v := range o.MetricValues {

		mv := &MetricValue{}
		mv.ID = v.ID
		mv.OpportunityID = o.ID
		mv.MetricID = v.Metric.ID
		mv.ValueID = v.Value.ID

		// link Metrics with The Opportunity

		verrs, err = tx.ValidateAndUpdate(mv)

		if err != nil || verrs.HasAny() {
			return verrs, errors.WithStack(err)
		}
	}

	return verrs, err

}