package actions

import (
	"github.com/gobuffalo/buffalo"
	"Monsters/models"
	"errors"
	"github.com/gobuffalo/buffalo/render"
	"strconv"
)

func userDataHandler(c buffalo.Context) error {
	db := models.GetDBInstance()

	u, err := db.GetUserById(c.Param("userID"))
	if err != nil {
		return errors.New("user not found")
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"userData": u,
	}))
}

func addMonsterHandler(c buffalo.Context) error {
	u := c.Data()["User"].(models.User)

	id, err := strconv.ParseInt(c.Param("monsterID"), 10, 32)
	if err != nil {
		return errors.New("invalid monster id")
	}

	err = u.AddMonster(int32(id))
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster added",
	}))
}
