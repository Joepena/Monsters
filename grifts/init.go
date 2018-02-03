package grifts

import (
	"github.com/gobuffalo/buffalo"
	"Monsters/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
