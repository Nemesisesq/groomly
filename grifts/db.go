package grifts

import (
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here

		_ = grift.Run("seed:values", c)
		_ = grift.Run("seed:metrics", c)
		_ = grift.Run("seed:fatal_attributes", c)
		return nil
	})

})
