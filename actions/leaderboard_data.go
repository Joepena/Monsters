package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/joepena/monsters/models"
	"github.com/pkg/errors"
	"strings"
)

func leaderboardDataHandler(c buffalo.Context) error {
	db := models.GetDBInstance()

	users, err := db.GetLeaderboardData()
	if err != nil {
		return errors.New("user data not available")
	}

	type userMetadata struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Wins   int32  `json:"wins"`
		Losses int32  `json:"losses"`
	}
	var leaderboard []userMetadata

	for _, u := range users {
		user := userMetadata{
			ID:     u.ID,
			Name:   strings.Split(u.Email, "@")[0],
			Wins:   u.BattleStats.Wins,
			Losses: u.BattleStats.Losses,
		}
		leaderboard = append(leaderboard, user)
	}

	return c.Render(200, render.JSON(map[string]interface{}{
		"users": leaderboard,
	}))
}

