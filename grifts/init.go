package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/invitation/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
