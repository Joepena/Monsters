package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID           string `bson:"_id"`
	Email        string `bson:"email"`
	Password     string `bson:"-"`
	PasswordHash string `bson:"password_hash"`
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
	return nil
}
func (u *User) Authenticate() bool {
	collection := GetDBInstance().session.DB("auth").C("users")

	passwordToAuth := u.Password
	email := strings.ToLower(u.Email)

	err := collection.Find(bson.M{"email": email,}).One(&u)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(passwordToAuth))
	if err != nil {
		return false
	}
	return true
}
