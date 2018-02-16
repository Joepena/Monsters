package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/joepena/monsters/models"
	"github.com/pkg/errors"
)

func createMonsterHandler(c buffalo.Context) error {
	m := &models.Monster{}

	err := c.Bind(m)
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Create()
	if err != nil {
		return err
	}

	return c.Render(201, render.JSON(map[string]interface{}{
		"monster": m,
	}))
}

func createAttackHandler(c buffalo.Context) error {
	a := &models.Attack{}

	err := c.Bind(a)
	if err != nil {
		return errors.WithStack(err)
	}

	err = a.Create()
	if err != nil {
		return err
	}

	return c.Render(201, render.JSON(map[string]interface{}{
		"attack": a,
	}))
}
