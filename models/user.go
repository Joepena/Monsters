package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	log "github.com/sirupsen/logrus"
	"errors"
)

type User struct {
	ID           string    `bson:"_id"`
	Email        string    `bson:"email"`
	Password     string    `bson:"-"`
	PasswordHash string    `bson:"password_hash"`
	Monsters	 []Monster `bson:"monsters"`
}

// Create wraps up the pattern of encrypting the password and
// running validations.
func (u *User) Create(db *DB) error {
	u.Email = strings.ToLower(u.Email)
	log.Info(u.Password)
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	log.Info(string(ph))
	u.PasswordHash = string(ph)
	return db.session.DB("auth").C("users").Insert(u)
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

func (u *User) AddMonster(id int32) error {
	db := GetDBInstance()
	c := db.session.DB("auth").C("users")

	monster, err := db.GetMonsterByNo(id)
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
