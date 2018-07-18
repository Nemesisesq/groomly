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
		metrics := []map[string]interface{}{
			{
				"M": models.Metric{Name: "Efficiency (annual $)", Type: models.Benefit, Weight: 2},
				"V": models.Values{
					{Name: "< $5k", Score: 1},
					{Name: "$5k - 10k", Score: 2},
					{Name: "$11k - 25k", Score: 3},
					{Name: "$26k - 50k", Score: 4},
					{Name: "> $50k", Score: 5},
				}},
			{
				"M": models.Metric{Name: "Working Capital Improvement ($)", Type: models.Benefit, Weight: 2},
				"V": models.Values{
					{Name: "< $100k", Score: 1},
					{Name: "$101k - 250k", Score: 2},
					{Name: "$251k - 500k", Score: 3},
					{Name: "$501k - 1m", Score: 4},
					{Name: "> $1m", Score: 5},
				}},
			{
				"M": models.Metric{Name: "Control & Risk Avoidance", Type: models.Benefit, Weight: 3},
				"V": models.Values{
					{Name: "Lowest Impact", Score: 1},
					{Name: "Lower Impact", Score: 2},
					{Name: "Middle Impact", Score: 3},
					{Name: "Higher Impact", Score: 4},
					{Name: "Highest Impact", Score: 5},
				}},
			{
				"M": models.Metric{Name: "Stakeholder Service", Type: models.Benefit, Weight: 2},
				"V": models.Values{
					{Name: "No Improvement", Score: 1},
					{Name: "Little Improvement", Score: 2},
					{Name: "Some Improvement", Score: 3},
					{Name: "Significant Improvement", Score: 4},
					{Name: "Transformed", Score: 5},
				}},
			{
				"M": models.Metric{Name: "Time Allocation (annual hrs.)", Type: models.Benefit, Weight: 1},
				"V": models.Values{
					{Name: "< 100", Score: 1},
					{Name: "501 - 500", Score: 2},
					{Name: "501 - 1,000", Score: 3},
					{Name: "1,001 - 2,000", Score: 4},
					{Name: "> 2,000", Score: 5},
				}},

			{
				"M": models.Metric{Name: "Finance Team", Type: models.Effort, Weight: 1},
				"V": models.Values{
					{Name: "Small", Score: 1},
					{Name: "Medium", Score: 2},
					{Name: "Large", Score: 3},
					{Name: "Extra Large", Score: 4},
				}},
			{
				"M": models.Metric{Name: "Safelite Stakeholder - Excl. IT", Type: models.Effort},
				"V": models.Values{
					{Name: "Small", Score: 1},
					{Name: "Medium", Score: 2},
					{Name: "Large", Score: 3},
					{Name: "Extra Large", Score: 4},
				}},
			{
				"M": models.Metric{Name: "Safelite IT", Type: models.Effort},
				"V": models.Values{
					{Name: "Small", Score: 1},
					{Name: "Medium", Score: 2},
					{Name: "Large", Score: 3},
					{Name: "Extra Large", Score: 4},
				}},
			{
				"M": models.Metric{Name: "External Resource", Type: models.Effort},
				"V": models.Values{
					{Name: "Small", Score: 1},
					{Name: "Medium", Score: 2},
					{Name: "Large", Score: 3},
					{Name: "Extra Large", Score: 4},
				}},
			{
				"M": models.Metric{Name: "Investment $", Type: models.Effort},
				"V": models.Values{
					{Name: "Small", Score: 1},
					{Name: "Medium", Score: 2},
					{Name: "Large", Score: 3},
					{Name: "Extra Large", Score: 4},
				}},
		}

		tx, err := pop.Connect("development")
		if err != nil {
			log.Panic(err)
		}

		for _, v := range metrics {
			met := v["M"].(models.Metric)
			vals := v["V"].(models.Values)
			err := tx.Create(&met)
			if err != nil {
				log.Panic(err)
			}
			for _, v := range vals {
				tx.Create(&v)
				if err != nil {
					log.Panic(err)
				}
			}

		}

		return nil
	})

})
