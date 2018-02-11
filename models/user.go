package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"strconv"
	"regexp"
)

type User struct {
	ID           string `bson:"_id"`
	AuthToken    string `bson:"auth_token"`
	Email        string `bson:"email" json:"email"`
	Password     string `bson:"-" json:"password"`
	PasswordHash string `bson:"password_hash"`
}

type AuthCounter struct {
	AccountCount int`bson:"account_count"`
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Create wraps up the pattern of encrypting the password and
// running validations.
func (u *User) Create() error {
	collection := GetDBInstance().session.DB("auth").C("users")

	err := u.validate()
	if err != nil {
		return err
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

	collection := GetDBInstance().session.DB("auth").C("users")
	var testU User

	//check if email is already in DB
	u.Email = strings.ToLower(u.Email)
	collection.Find(bson.M{"email": strings.ToLower(u.Email),}).One(&testU)

	if testU.Email == u.Email {
		return errors.New("email is already in use")
	}

	return nil
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

