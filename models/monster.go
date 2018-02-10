package models

import "errors"

type Monster struct {
	ID		string   `bson:"id"` //not _id, set when added to a user
	No      int32    `bson:"no"`
	Name    string   `bson:"name"`
	Type    string   `bson:"type"`
	Hp      int32    `bson:"hp"`
	Attack  int32    `bson:"attack"`
	Defense int32    `bson:"defense"`
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
