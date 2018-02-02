package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"strconv"
)

type User struct {
	ID           string `bson:"_id"`
	AuthToken    string `bson:"auth_token"`
	Email        string `bson:"email"`
	Password     string `bson:"-"`
	PasswordHash string `bson:"password_hash"`
	Monsters	 []Monster `bson:"monsters"`
}

type AuthCounter struct {
	AccountCount int`bson:"account_count"`
}

// Create wraps up the pattern of encrypting the password and
// running validations.
func (u *User) Create() error {
	collection := GetDBInstance().session.DB("auth").C("users")
	var testU User

	//check if email is already in DB
	u.Email = strings.ToLower(u.Email)
	collection.Find(bson.M{"email": strings.ToLower(u.Email),}).One(&testU)

	if testU.Email == u.Email {
		return errors.New("email is already in use")
	}
	// generate DB id
	id, err := generateID()
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
	collection := GetDBInstance().session.DB("auth").C("users")

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

func generateID() (string, error) {
	collection := GetDBInstance().session.DB("auth").C("counters")
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"account_count": 1}},
		ReturnNew: false,
	}

	var counterDoc AuthCounter

	_, err := collection.Find(bson.M{"_id":"0"}).Apply(change, &counterDoc)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(counterDoc.AccountCount), nil

}

func (u *User) AddMonster(no int32) error {
	db := GetDBInstance()
	c := db.session.DB("auth").C("users")

	monster, err := db.GetMonsterByNo(no)
	if err != nil {
		return errors.New("monster not found")
	}

	query := bson.M{"_id": u.ID}
	update := bson.M{"$push": bson.M{"monsters": monster}}
	return c.Update(query, update)
}

func (u *User) RenameMonster(no int32, name string) error {
	db := GetDBInstance()
	c := db.session.DB("auth").C("users")

	query := bson.M{"_id": u.ID, "monsters.no": no}
	update := bson.M{"$set": bson.M{"monsters.$.name": name}}
	return c.Update(query, update)
}

func (u *User) AddMonsterAttack(no int32, name string) error {
	db := GetDBInstance()
	c := db.session.DB("auth").C("users")

	attack, err := db.GetAttackByName(name)
	if err != nil {
		return errors.New("attack not found")
	}

	query := bson.M{"_id": u.ID, "monsters.no": no}
	update := bson.M{"$push": bson.M{"monsters.$.attacks": attack}}
	return c.Update(query, update)
}
