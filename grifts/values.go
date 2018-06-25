package grifts

import (
	"log"

	"github.com/gobuffalo/pop"
	. "github.com/markbates/grift/grift"
	"github.com/nemesisesq/groomly/models"
)

var _ error = Namespace("seed", func() {

	Desc("values", "This task seed values in the database")
	Add("values", func(c *Context) error {
		values := models.Values{
			{Name: "< $5k", Score: 1},
			{Name: "$5k - 10k", Score: 2},
			{Name: "$11k - 25k", Score: 3},
			{Name: "$26k - 50k", Score: 4},
			{Name: "> $50k", Score: 5},

			{Name: "< $100k", Score: 1},
			{Name: "$101k - 250k", Score: 2},
			{Name: "$251k - 500k", Score: 3},
			{Name: "$501k - 1m", Score: 4},
			{Name: "> $1m", Score: 5},

			{Name: "Lowest Impact", Score: 1},
			{Name: "Lower Impact", Score: 2},
			{Name: "Middle Impact", Score: 3},
			{Name: "Higher Impact", Score: 4},
			{Name: "Highest Impact", Score: 5},

			{Name: "No Improvement", Score: 1},
			{Name: "Little Improvement", Score: 2},
			{Name: "Some Improvement", Score: 3},
			{Name: "Significant Improvement", Score: 4},
			{Name: "Transformed", Score: 5},

			{Name: "< 100", Score: 1},
			{Name: "501 - 500", Score: 2},
			{Name: "501 - 1,000", Score: 3},
			{Name: "1,001 - 2,000", Score: 4},
			{Name: "> 2,000", Score: 5},
		}

		tx, err := pop.Connect("development")
		if err != nil {
			log.Panic(err)
		}

		for _, v := range values {
			err := tx.Create(&v)

			if err != nil {
				log.Panic(err)
			}
		}

		return err
	})

})
