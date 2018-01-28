package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Monster struct {
	No      int32  `bson:"no"`
	Name    string `bson:"name"`
	Type    string `bson:"type"`
	Model   string `bson:"model"`
	Hp      int32  `bson:"hp"`
	Attack  int32  `bson:"attack"`
	Defense int32  `bson:"defense"`
	//Attacks []Attack  `bson:"attacks"`
}

func (m *Monster) Create(db *DB) error {
	c := db.session.DB("monsters").C("monsters")

	monster := Monster{
		No:      25,
		Name:    "Pikachu",
		Type:    "Electric",
		Model:   "uri.go/yvisalxjeelf",
		Hp:      35,
		Attack:  55,
		Defense: 40,
	}

	if err := c.Insert(monster); err != nil {
		return err
	}
	return nil
}

func (m *Monster) Find(db *DB) error {
	c := db.session.DB("monsters").C("monsters")

	err := c.Find(bson.M{"no": m.No}).One(&m)
	if err != nil {
		return err
	}
	return nil
}
