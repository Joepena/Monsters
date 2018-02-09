package actions

import (
	"github.com/gobuffalo/buffalo"
	"Monsters/models"
	"errors"
	"github.com/gobuffalo/buffalo/render"
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

	err := user.AddMonster(toInt(c.Param("monsterNo")))
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster added",
	}))
}

func renameMonsterHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)

	no := toInt(c.Param("no"))
	name := c.Param("name")

	err := user.RenameMonster(no, name)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster renamed to " + name,
	}))
}

func addMonsterAttackHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)

	no := c.Param("monsterNo")
	id := c.Param("attackID")

	err := user.AddMonsterAttack(toInt(no), id)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "attack added",
	}))
}
