package grifts

import (
	"log"

	"github.com/gobuffalo/pop"
	. "github.com/markbates/grift/grift"
	"github.com/nemesisesq/groomly/models"
	"github.com/gobuffalo/envy"
)

var _ = Namespace("seed", func() {

	Desc("fatal_attributes", "Seeds the database with fatal attributes")
	Add("fatal_attributes", func(c *Context) error {

		fa := models.FatalAttributes{
			{Name: "Regulatory", Summary: "The Goverment really cares about this one"},
		}

		env := envy.Get("GO_ENV", "development")
		tx, err := pop.Connect(env)
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
