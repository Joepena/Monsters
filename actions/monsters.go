package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	//"log"
	"Monsters/models"
	"strconv"
	"errors"
)

func userMonstersHandler (c buffalo.Context) error {
	// Get user data once logged in
	u := models.User{
		Email: "email",
		Password: "pass",
	}

	validUser := u.Authenticate()
	if !validUser {
		return errors.New("invalid credentials")
	}

	dbInstance := models.GetDBInstance()

	monsters := []models.Monster{}
	for _, no := range u.Monsters {
		monster, err := dbInstance.GetMonsterByNo(no)
		if err != nil {
			return err
		}
		monsters = append(monsters, monster)
	}

	return c.Render(201, render.JSON(map[string]interface{}{
		"monsters": monsters,
	}))

}

func createMonsterHandler (c buffalo.Context) error {
	m := models.Monster{}
	err := m.Create(models.GetDBInstance())
	if err != nil {
		return err
	}
	return c.Render(201, render.JSON(map[string]string{
		"status": "monster created",
	}))
}

func monsterDataHandler (c buffalo.Context) error {

	no, err := strconv.ParseInt(c.Param("monsterID"), 10, 32)
	if err != nil {
		return err
	}

	m := models.Monster{
		No: int32(no),
	}

	err = m.Find(models.GetDBInstance())
	if err != nil {
		return err
	}

	return c.Render(201, render.JSON(map[string]interface{}{
		"no": m.No,
		"name": m.Name,
		"type": m.Type,
		"model": m.Model,
		"hp": m.Hp,
		"attack": m.Attack,
		"defense": m.Defense,
	}))
}

