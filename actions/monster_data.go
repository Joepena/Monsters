package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	//"log"
	"Monsters/models"
	"strconv"
	"errors"
)

//func getMonstersHandler (c buffalo.Context) error {
//	db := models.GetDBInstance()
//
//	user, err := db.GetUserById(c.Param("userID"))
//	if err != nil {
//		return errors.New("User not found")
//	}
//
//	monsters := []models.Monster{{}}
//	for _, no := range user.Monsters {
//		monster, err := db.GetMonsterByNo(no)
//		if err != nil {
//			return err
//		}
//		monsters = append(monsters, monster)
//	}
//
//	if len(monsters) > 0 {
//		return c.Render(200, render.JSON(map[string]interface{}{
//			"monsters": monsters,
//		}))
//	}
//
//	c.Render(204, render.JSON(map[string]string{
//		"message": "User has no monsters",
//	}))
//	return nil
//}

func createMonsterHandler(c buffalo.Context) error {
	m := models.Monster{}
	err := m.Create(models.GetDBInstance())
	if err != nil {
		return errors.New("Could not create monster")
	}
	return c.Render(201, render.JSON(map[string]string{
		"status": "monster created",
	}))
}

func monsterDataHandler(c buffalo.Context) error {
	no, err := strconv.ParseInt(c.Param("monsterID"), 10, 32)
	if err != nil {
		return errors.New("Invalid monster ID")
	}

	m := models.Monster{
		No: int32(no),
	}

	err = m.Find(models.GetDBInstance())
	if err != nil {
		return errors.New("Monster not found")
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"no": m.No,
		"name": m.Name,
		"type": m.Type,
		"hp": m.Hp,
		"attack": m.Attack,
		"defense": m.Defense,
	}))
}

