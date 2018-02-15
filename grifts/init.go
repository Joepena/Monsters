package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/monsters/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
