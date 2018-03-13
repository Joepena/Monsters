package grifts

import (
	log "github.com/sirupsen/logrus"
	"github.com/markbates/grift/grift"
	"github.com/joepena/monsters/models"
	"gopkg.in/mgo.v2/bson"
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

