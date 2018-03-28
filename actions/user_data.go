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
		"email": user.Email,
		"monsters": user.Monsters,
		"battleStats": user.BattleStats,
	}))
}

func userAssetDataHandler(c buffalo.Context) error {
	db := models.GetDBInstance()
	userID := c.Param("userID")

	user, err := db.GetUserById(userID)
	if err != nil {
		return errors.New("user not found")
	}

	type MonsterAsset struct {
		Name         string
		MonsterNo    int32
		AssetSet     models.AssetIDSet
	}
	var assets []MonsterAsset

	for _, monster := range user.Monsters {

		mA := MonsterAsset{
			Name: monster.Name,
			MonsterNo: monster.No,
			AssetSet: monster.Assets,
		}
		assets = append(assets, mA)
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"userID": userID,
		"assets": assets,
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

func updateMonsterStatsHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)
	m := &models.Monster{}

	err := c.Bind(m)
	if err != nil {
		return errors.WithStack(err)
	}

	err = user.UpdateMonsterStats(m)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "monster stats updated",
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

func addBattleResultHandler(c buffalo.Context) error {
	user := c.Data()["User"].(models.User)
	b := &models.BattleStats{}

	err := c.Bind(b)
	if err != nil {
		return errors.WithStack(err)
	}

	err = user.AddBattleResult(b)
	if err != nil {
		return err
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"status": "battle stats updated",
	}))
}
