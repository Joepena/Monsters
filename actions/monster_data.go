package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"Monsters/models"
	"errors"
	"strconv"
)

func createMonsterHandler(c buffalo.Context) error {
	m := models.Monster{
		No:      toInt(c.Param("no")),
		Name:    c.Param("name"),
		Type:    c.Param("type"),
		Hp:      toInt(c.Param("hp")),
		Attack:  toInt(c.Param("attack")),
		Defense: toInt(c.Param("defense")),
	}

	err := m.Create(models.GetDBInstance())
	if err != nil {
		return err
	}

	return c.Render(201, render.JSON(map[string]interface{}{
		"monster": m,
	}))
}

func toInt(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		errors.New("invalid int")
	}
	return int32(i)
}
