package grifts

import (
	log "github.com/sirupsen/logrus"
	"github.com/markbates/grift/grift"
	"github.com/joepena/monsters/models"
	"github.com/joepena/monsters/actions"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"encoding/json"
)

var _ = grift.Namespace("seed", func() {

	grift.Desc("init", "Initializes necessary documents to operate the application")
	grift.Add("init", func(c *grift.Context) error {
		db := models.GetDBInstance()
		err := InsertCounters(db)
		if err != nil {
			log.Error(err)
		}
		return nil
	})

	grift.Desc("dex", "Seeds basic pokemon and attacks to the dex")
	grift.Add("dex", func(c *grift.Context) error {
		db := models.GetDBInstance()
		err := InsertToDex(db, "./grifts/seeds/monsters_seed.json", "./grifts/seeds/attacks_seed.json")
		if err != nil {
			log.Error(err)
		}
		return nil
	})

	grift.Desc("users", "Seed fake users for leaderboard")
	grift.Add("users", func(c *grift.Context) error {
		db := models.GetDBInstance()
		err := InsertToUsers(db, "./grifts/seeds/users_seed.json")
		if err != nil {
			log.Error(err)
		}
		return nil
	})

})

func InsertCounters(db *models.DB) error {
	// Auth account counter
	c := db.Session.DB("auth").C("counters")
	cDoc := bson.M{"_id":"0", "account_count":1}
	err := c.Insert(cDoc)
	if err != nil {
		return err
	}

	// Asset counter
	c = db.Session.DB("dex").C("counters")
	cDoc = bson.M{"_id":"asset_counter", "asset_count": 1}
	err = c.Insert(cDoc)
	if err != nil {
		return err
	}

	// Monster count counter
	cDoc = bson.M{"_id":"monster_counter", "asset_count": 1}
	err = c.Insert(cDoc)
	if err != nil {
		return err
	}

	log.Info("Successfully seeded counters into DB")
	return nil
}

func InsertToDex(db *models.DB, monsterSeed string, attackSeed string) interface{} {
	// seed monsters
	f, err := ioutil.ReadFile(monsterSeed)
	if err != nil {
		return err
	}

	var monsters[] models.Monster
	err = json.Unmarshal(f, &monsters)
	if err != nil {
		return err
	}

	c := db.Session.DB("dex").C("monsters")
	for _, m := range monsters {
		err = c.Insert(m)
		if err != nil {
			return err
		}
	}
	log.Info("Successfully seeded monsters into DB")

	//seed attacks
	f, err = ioutil.ReadFile(attackSeed)
	if err != nil {
		return err
	}

	var attacks[] models.Attack
	err = json.Unmarshal(f, &attacks)
	if err != nil {
		return err
	}

	c = db.Session.DB("dex").C("attacks")
	for _, a := range attacks {
		err = c.Insert(a)
		if err != nil {
			return err
		}
	}

	log.Info("Successfully seeded attacks into DB")
	return nil
}

func InsertToUsers(db *models.DB, userSeed string) interface{} {
	f, err := ioutil.ReadFile(userSeed)
	if err != nil {
		return err
	}

	var users[] models.User
	err = json.Unmarshal(f, &users)
	if err != nil {
		return err
	}

	for _, u := range users {
		token, err := actions.GetAuthToken(&u)
		if err != nil {
			return err
		}
		u.AuthToken = token

		err = u.Create()
		if err != nil {
			return err
		}
	}
	log.Info("Successfully seeded test users into DB")

	return nil
}
