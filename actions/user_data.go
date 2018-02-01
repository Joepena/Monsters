package actions

import (
	"github.com/gobuffalo/buffalo"
	"Monsters/models"
	"errors"
	"github.com/gobuffalo/buffalo/render"
)

func userDataHandler (c buffalo.Context) error {
	db := models.GetDBInstance()

	u, err := db.GetUserById(c.Param("userID"))
	if err != nil {
		return errors.New("User not found")
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"userData": u,
	}))
}
