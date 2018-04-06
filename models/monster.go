package models

import (
	"github.com/pkg/errors"
)

type Monster struct {
	ID		string   `bson:"id" json:"monsterID"` //not _id, set when added to a user
	No      int32    `bson:"no" json:"monsterNo"`
	Assets  AssetIDSet `bson:"asset_id_set" json:"assetIDSet"`
	Name    string   `bson:"name" json:"name"`
	Type    string   `bson:"type" json:"type"`
	Hp      int32    `bson:"hp" json:"hp"`
	Attack  int32    `bson:"attack" json:"attack"`
	Defense int32    `bson:"defense" json:"defense"`
	Attacks []Attack `bson:"attacks" json:"attacks"`
	Stats   Stats    `bson:"stats"   json:"stats"`
}

type Stats struct {
	Hits            int32 `bson:"hits"             json:"hits"`
	Misses          int32 `bson:"misses"           json:"misses"`
	DamageDealt     int32 `bson:"damage_dealt"     json:"damageDealt"`
	DamageReceived  int32 `bson:"damage_received"  json:"damageReceived"`
	EnemiesFought   int32 `bson:"enemies_fought"   json:"enemiesFought"`
	EnemiesDefeated int32 `bson:"enemies_defeated" json:"enemiesDefeated"`
	Faints          int32 `bson:"faints"           json:"faints"`

}

type AssetIDSet struct {
	Texture1ID int `bson:"texture_1_id" json:"texture1ID"`
	Texture2ID int `bson:"texture_2_id" json:"texture2ID"`
	IOS struct{
		DaeID int `bson:"dae_id" json:"daeID"`
		AnimationSet []AnimationPair `bson:"animation_set" json:"animationSet"`
	} `bson:"ios" json:"ios"`
	Android struct{
		AssetID int `bson:"asset_id" json:"assetID"`
	} `bson:"android" json:"android"`
}

type AnimationPair struct {
	Name string `bson:"name" json:"name"`
	AssetID int `bson:"asset_id" json:"assetID"`
}


func (m *Monster) Create() error {
	c := GetDBInstance().Session.DB("dex").C("monsters")

	err := c.Insert(m)
	if err != nil {
		return errors.New("could not create monster")
	}
	return nil
}
