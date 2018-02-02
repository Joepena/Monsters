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

	user, err := db.GetUserById(c.Param("userID"))
	if err != nil {
		return errors.New("user not found")
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"userData": user,
	}))
}

func addMonsterHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)

	id, err := strconv.ParseInt(c.Param("monsterID"), 10, 32)
	if err != nil {
		return errors.New("invalid monster id")
	}

	err = user.AddMonster(int32(id))
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster added",
	}))
}

func renameMonsterHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)

	no := c.Param("no")
	name := c.Param("name")

	i, err := strconv.ParseInt(no, 10, 32)
	if err != nil {
		return err
	}

	err = user.RenameMonster(int32(i), name)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster renamed to " + name,
	}))
}
