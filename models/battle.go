package models

import "time"

type Battle struct {
	VictorID string    `bson:"victor_id" json:"victorID"`
	LoserID  string    `bson:"loser_id"  json:"loserID"`
	Date     time.Time `bson:"date"      json:"date"`
	Location Location  `bson:"location"  json:"location"`
}

type Location struct {
	X float32 `bson:"x" json:"x"`
	Y float32 `bson:"y" json:"y"`
}
