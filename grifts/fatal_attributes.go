package grifts

import (
	"log"

	"github.com/gobuffalo/pop"
	. "github.com/markbates/grift/grift"
	"github.com/nemesisesq/groomly/models"
)

var _ = Namespace("seed", func() {

	Desc("fatal_attributes", "Seeds the database with fatal attributes")
	Add("fatal_attributes", func(c *Context) error {

		fa := models.FatalAttributes{
			{Name: "Regulatory", Summary: "The Goverment really cares about this one"},
		}

		tx, err := pop.Connect("development")
		if err != nil {
			log.Panic(err)
		}

		for _, v := range fa {
			err := tx.Create(&v)

			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})

})
