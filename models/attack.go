package models

import "github.com/pkg/errors"

type Attack struct {
	SlotNo		int32  `bson:"slot_no" json:"slotNo"`
	MonsterNo   int32  `bson:"monster_no" json:"monsterNo"`
	AssetID     int32    `bson:"asset_id" json:"assetID"`
	Name        string `bson:"name" json:"name"`
	Type        string `bson:"type" json:"type"`
	Power       int32  `bson:"power" json:"power"`
	Accuracy    int32  `bson:"accuracy" json:"accuracy"`
}

type AddAttackParams struct {
	AttackID  string `json:"attackID"`
	MonsterID string `json:"monsterID"`
	SlotNo    int32  `json:"slotNo"`
}

func (a *Attack) Create() error {
	c := GetDBInstance().Session.DB("dex").C("attacks")

	err := c.Insert(a)
	if err != nil {
		return errors.New("could not create attack")
	}
	return nil
}
