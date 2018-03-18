package models

import "github.com/pkg/errors"

type Attack struct {
	SlotNo		int32  `bson:"slot_no"`
	MonsterNo   int32  `bson:"monster_no"   json:"monsterNo"`
	Name        string `bson:"name"         json:"name"`
	Type        string `bson:"type"         json:"type"`
	Power       int32  `bson:"power"        json:"power"`
	Accuracy    int32  `bson:"accuracy"     json:"accuracy"`
	AnimationID int32  `bson:"animation_id" json:"animationID"`
}

func (a *Attack) Create() error {
	c := GetDBInstance().session.DB("dex").C("attacks")

	err := c.Insert(a)
	if err != nil {
		return errors.New("could not create attack")
	}
	return nil
}
