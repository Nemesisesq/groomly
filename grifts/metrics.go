package grifts

import (
	"log"

	"github.com/gobuffalo/pop"
	. "github.com/markbates/grift/grift"
	"github.com/nemesisesq/groomly/models"
)

var _ = Namespace("seed", func() {

	Desc("metrics", "Seeds the Database with Metrics")
	Add("metrics", func(c *Context) error {
		metrics := models.Metrics{
			{Name: "Efficiency (annual $)", Type: models.Benefit, Weight: 2, Value:models.Value{}},
			{Name: "Working Capital Improvement ($)", Type: models.Benefit, Weight: 2, Value:models.Value{}},
			{Name: "Control & Risk Avoidance", Type: models.Benefit, Weight: 3, Value:models.Value{}},
			{Name: "Stakeholder Service", Type: models.Benefit, Weight: 2, Value:models.Value{}},
			{Name: "Time Allocation (annual hrs.)", Type: models.Benefit, Weight: 1, Value:models.Value{}},

			{Name: "Finance Team", Type: models.Effort, Value:models.Value{}},
			{Name: "Safelite Stakeholder - Excl. IT", Type: models.Effort, Value:models.Value{}},
			{Name: "Safelite IT", Type: models.Effort, Value:models.Value{}},
			{Name: "External Resource", Type: models.Effort, Value:models.Value{}},
			{Name: "Investment $", Type: models.Effort, Value:models.Value{}},
		}

		tx, err := pop.Connect("development")
		if err != nil {
			log.Panic(err)
		}

		for _, v := range metrics {
			err := tx.Create(&v)

			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})

})
