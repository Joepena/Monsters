package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"github.com/pkg/errors"
	"regexp"
)

type User struct {
	ID           string      `bson:"_id"`
	AuthToken    string      `bson:"auth_token"`
	Email        string      `bson:"email"          json:"email"`
	Password     string      `bson:"-"              json:"password"`
	PasswordHash string      `bson:"password_hash"`
	Monsters     []Monster   `bson:"monsters"       json:"monsters"`
	BattleStats  BattleStats `bson:"battle_stats"   json:"battleStats"`
}

type BattleStats struct {
	Wins   int32 `bson:"wins"   json:"wins"`
	Losses int32 `bson:"losses" json:"losses"`
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Create wraps up the pattern of encrypting the password and
// running validations.
func (u *User) Create() error {
	collection := GetDBInstance().Session.DB("auth").C("users")

	//init empty fields if not set
	if len(u.Monsters) < 1 {
		u.Monsters = []Monster{}
	}
	if (u.BattleStats == BattleStats{}) {
		u.BattleStats = BattleStats{
			Wins:   0,
			Losses: 0,
		}
	}

	err := u.validate()
	if err != nil {
		return err
	}

	// generate DB id
	id, err := generateAccountID()
	if err != nil {
		return err
	}
	u.ID = id
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(ph)

	return collection.Insert(u)
}

func (u *User) Authenticate() bool {
	collection := GetDBInstance().Session.DB("auth").C("users")

	passwordToAuth := u.Password
	email := strings.ToLower(u.Email)

	err := collection.Find(bson.M{"email": email}).One(&u)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(passwordToAuth))
	if err != nil {
		return false
	}

	return true
}

func (u *User) validate() error {
	u.Email = strings.ToLower(u.Email)

	if u.Email == "" {
		return errors.New("empty email was provided")
	}
	if u.Password == "" || len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if !emailRegexp.MatchString(u.Email) {
		return errors.New("provide a valid email address")
	}

	collection := GetDBInstance().Session.DB("auth").C("users")
	var testU User

	//check if email is already in DB
	u.Email = strings.ToLower(u.Email)
	collection.Find(bson.M{"email": strings.ToLower(u.Email),}).One(&testU)

	if testU.Email == u.Email {
		return errors.New("email is already in use")
	}

	return nil
}

func (u *User) AddMonster(no int32) error {
	db := GetDBInstance()
	c := db.Session.DB("auth").C("users")

	monster, err := db.GetMonsterByNo(no)
	if err != nil {
		return errors.New("monster not found")
	}

	id, err := generateMonsterID()
	if err != nil {
		return err
	}
	monster.ID = id

	query := bson.M{"_id": u.ID}
	update := bson.M{"$push": bson.M{"monsters": monster}}
	return c.Update(query, update)
}

func (u *User) RenameMonster(m *Monster) error {
	c := GetDBInstance().Session.DB("auth").C("users")

	query := bson.M{"_id": u.ID, "monsters.id": m.ID}
	update := bson.M{"$set": bson.M{"monsters.$.name": m.Name}}
	return c.Update(query, update)
}

func (u *User) UpdateMonsterStats(m *Monster) error {
	c := GetDBInstance().Session.DB("auth").C("users")

	query := bson.M{"_id": u.ID, "monsters.id": m.ID}
	update := bson.M{"$inc": bson.M{
		"monsters.$.stats.hits":             m.Stats.Hits,
		"monsters.$.stats.misses":           m.Stats.Misses,
		"monsters.$.stats.damage_dealt":     m.Stats.DamageDealt,
		"monsters.$.stats.damage_received":  m.Stats.DamageReceived,
		"monsters.$.stats.enemies_fought":   m.Stats.EnemiesFought,
		"monsters.$.stats.enemies_defeated": m.Stats.EnemiesDefeated,
		"monsters.$.stats.faints":           m.Stats.Faints,
	}}

	return c.Update(query, update)
}

func (u *User) ReplaceMonsterAttack(a *AddAttackParams) error {
	db := GetDBInstance()
	c := db.Session.DB("auth").C("users")

	attack, err := db.GetAttackByID(a.AttackID)
	if err != nil {
		return errors.New("attack not found")
	}
	attack.SlotNo = a.SlotNo

	for _, m := range u.Monsters {
		if m.ID == a.MonsterID && m.No != attack.MonsterNo {
			return errors.New("invalid attack for this monster")
		}
	}

	// Remove existing attack by slot no
	query := bson.M{"_id": u.ID, "monsters.id": a.MonsterID}
	update := bson.M{"$pull": bson.M{"monsters.$.attacks": bson.M{"slot_no": a.SlotNo}}}
	c.Update(query, update)

	 //Add new attack
	query = bson.M{"_id": u.ID, "monsters.id": a.MonsterID}
	update = bson.M{"$push": bson.M{"monsters.$.attacks": attack}}
	return c.Update(query, update)
}


func (u *User) AddBattleResult(b *BattleStats) error {
	db := GetDBInstance()
	c := db.Session.DB("auth").C("users")

	query := bson.M{"_id": u.ID}
	update := bson.M{"$inc": bson.M{
		"battle_stats.wins":   b.Wins,
		"battle_stats.losses": b.Losses,
	}}
	return c.Update(query, update)
}

