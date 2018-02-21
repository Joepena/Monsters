package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type AuthCounter struct {
	AccountCount int` bson:"account_count"`
}

type MonsterCounter struct {
	MonsterCount int `bson:"monster_count"`
}

type AssetCounter struct {
	AssetCount int `bson:"asset_count"`
}

func generateMonsterID() (string, error) {
	c := GetDBInstance().session.DB("dex").C("counters")
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"monster_count": 1}},
		ReturnNew: false,
	}

	var counterDoc MonsterCounter

	_, err := c.Find(bson.M{"_id": "monster_counter"}).Apply(change, &counterDoc)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(counterDoc.MonsterCount), nil
}

func generateAssetId() (string, error) {
	collection := GetDBInstance().session.DB("dex").C("counters")
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"asset_count": 1}},
		ReturnNew: false,
	}

	var counterDoc AssetCounter

	_, err := collection.Find(bson.M{"_id":"asset_counter"}).Apply(change, &counterDoc)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(counterDoc.AssetCount), nil
}

func generateAccountID() (string, error) {
	collection := GetDBInstance().session.DB("auth").C("counters")
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"account_count": 1}},
		ReturnNew: false,
	}

	var counterDoc AuthCounter

	_, err := collection.Find(bson.M{"_id":"0"}).Apply(change, &counterDoc)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(counterDoc.AccountCount), nil
}

