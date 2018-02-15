package models

import "errors"

type Monster struct {
	ID		string   `bson:"id" json:"monsterID"` //not _id, set when added to a user
	No      int32    `bson:"no" json:"monsterNo"`
	Name    string   `bson:"name" json:"name"`
	Type    string   `bson:"type" json:"type"`
	Hp      int32    `bson:"hp" json:"hp"`
	Attack  int32    `bson:"attack" json:"attack"`
	Defense int32    `bson:"defense" json:"defense"`
	Attacks []Attack `bson:"attacks"`
}

func (m *Monster) Create() error {
	c := GetDBInstance().session.DB("dex").C("monsters")

	err := c.Insert(m)
	if err != nil {
		return errors.New("could not create monster")
	}
	return nil
}
