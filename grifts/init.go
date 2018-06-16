package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/nemesisesq/groomly/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
