package models

import "errors"

type Attack struct {
	MonsterNo   int32  `bson:"monster_no"`
	Name        string `bson:"name"`
	Type        string `bson:"type"`
	Power       int32  `bson:"power"`
	Accuracy    int32  `bson:"accuracy"`
	AnimationID int32  `bson:"animation_id"`
}

func (a *Attack) Create() error {
	c := GetDBInstance().session.DB("monsters").C("attacks")

	err := c.Insert(a)
	if err != nil {
		return errors.New("could not create attack")
	}
	return nil
}
