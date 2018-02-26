package models

import "time"

type Battle struct {
	VictorID string    `bson:"victor_id"`
	LoserID  string    `bson:"loser_id"`
	Date     time.Time `bson:"date"`
	Location Location  `bson:"location"`
}

type Location struct {
	X float32 `bson:"x"`
	Y float32 `bson:"y"`
}
