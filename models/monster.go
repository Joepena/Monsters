package models

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Monster struct {
	No      int32  `bson:"no"`
	Name    string `bson:"name"`
	Type    string `bson:"type"`
	Hp      int32  `bson:"hp"`
	Attack  int32  `bson:"attack"`
	Defense int32  `bson:"defense"`
	//Attacks []Attack  `bson:"attacks"`
}

func (m *Monster) Create(db *DB) error {
	c := db.session.DB("monsters").C("monsters")

	err := c.Insert(m)
	if err != nil {
		return errors.New("could not create monster")
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
