package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/monsters/models"
	"github.com/pkg/errors"
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
		"token": user.AuthToken,
		"email": user.Email,
		"monsters": user.Monsters,
	}))
}

func userAnimationsHandler(c buffalo.Context) error {
	db := models.GetDBInstance()

	user, err := db.GetUserById(c.Param("userID"))
	if err != nil {
		return errors.New("user not found")
	}

	type animation struct {
		MonsterNo    int32
		AnimationIDs []int32
	}
	var animations []animation

	for _, monster := range user.Monsters {
		var ids []int32
		for _, attack := range monster.Attacks {
			ids = append(ids, attack.AnimationID)
		}
		a := animation{
			MonsterNo:    monster.No,
			AnimationIDs: ids,
		}
		animations = append(animations, a)
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"animations": animations,
	}))
}

func addMonsterHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)
	m := &models.Monster{}

	err := c.Bind(m)
	if err != nil {
		return errors.WithStack(err)
	}

	err = user.AddMonster(m.No)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster added",
	}))
}

func renameMonsterHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)
	m := &models.Monster{}

	err := c.Bind(m)
	if err != nil {
		return errors.WithStack(err)
	}

	err = user.RenameMonster(m)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster renamed to " + m.Name,
	}))
}

func addMonsterAttackHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)
	a := &models.AddAttackParams{}

	err := c.Bind(a)
	if err != nil {
		return errors.WithStack(err)
	}

	err = user.ReplaceMonsterAttack(a)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "attack added",
	}))
}
