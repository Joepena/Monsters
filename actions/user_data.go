package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/monsters/models"
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
		"id": user.ID,
		"email": user.Email,
		"monsters": user.Monsters,
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

	id := c.Param("monsterID")
	name := c.Param("name")

	err := user.RenameMonster(id, name)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster renamed to " + name,
	}))
}

func addMonsterAttackHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)

	monsterID := c.Param("monsterID")
	attackID := c.Param("attackID")
	slotNo := toInt(c.Param("slot_no"))

	err := user.ReplaceMonsterAttack(monsterID, attackID, slotNo)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "attack added",
	}))
}
