package models

import "errors"

type Attack struct {
	Name     string `bson:"name"`
	Type     string `bson:"type"`
	Power    int32  `bson:"power"`
	Accuracy int32  `bson:"accuracy"`
}

func (a *Attack) Create(db *DB) error {
	c := db.session.DB("monsters").C("attacks")

	err := c.Insert(a)
	if err != nil {
		return errors.New("could not create attack")
	}
	return nil
}
