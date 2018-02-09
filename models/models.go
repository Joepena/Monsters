package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"github.com/gobuffalo/envy"
	"crypto/tls"
	"net"
	"strings"
)

var MONGOD_URL = envy.Get("MONGOD_URL","localhost:27017")
var ON_ATLAS = envy.Get("ON_ATLAS", "false")

var (
	once       sync.Once
	dbInstance DB
)

type DB struct {
	session *mgo.Session
	logger  *log.Entry
}

type ReadRequest struct {
	dBName         string
	collectionName string
	logger         *log.Entry
}

type WriteRequest struct {
	DbName         string
	CollectionName string
	Data           interface{}
}

func (db *DB) Write(wR *WriteRequest) {
	c := db.session.DB(wR.DbName).C(wR.CollectionName) // maybe can cache this later?
	err := c.Insert(wR.Data)
	if err != nil {
		log.Fatal("failed to insert data into db. dbName: %v, collectionName: %v, data: %v, error: %v", wR.DbName, wR.CollectionName, wR.Data, err)
	}
}

func (db *DB) CreateCappedCollection(dbName string, collectionName string, capacity int) {
	collectionInfo := &mgo.CollectionInfo{
		Capped:   true,
		MaxBytes: capacity,
	}
	err := db.session.DB(dbName).C(collectionName).Create(collectionInfo)
	if err != nil {
		log.WithField("Function", "CreateCappedCollection").Error(err)
	}
}

/* User */
func (db *DB) GetUserById(id string) (User, error) {
	collection := db.session.DB("auth").C("users")

	var user User
	err := collection.Find(bson.M{"_id": id}).One(&user)

	return user, err
}

func (db *DB) GetUserByAuthToken(token string) (User, error) {
	c := db.session.DB("auth").C("users")

	var user User
	err := c.Find(bson.M{"auth_token": token}).One(&user)

	return user, err
}

/* Monster */
func (db *DB) GetMonsterByNo(no int32) (Monster, error) {
	collection := db.session.DB("monsters").C("monsters")

	var monster Monster
	err := collection.Find(bson.M{"no": no}).One(&monster)

	return monster, err
}

/* Attack */
func (db *DB) GetAttackById(id string) (Attack, error) {
	collection := db.session.DB("monsters").C("attacks")

	var attack Attack
	err := collection.Find(bson.M{"_id": id}).One(&attack)

	return attack, err
}

/* DB */
func GetDBInstance() *DB {
	once.Do(func() {
		dbInstance = DB{
			session: establishMongoDBSession(),
			logger:  log.WithField("Component", "DB"),
		}
	})

	return &dbInstance
}

func establishMongoDBSession() *mgo.Session {
	var session *mgo.Session
	var err error

	if strings.EqualFold(ON_ATLAS, "true") { // if we are on MongoDB Atlas dial with TLS connec info
		dialInfo, err := mgo.ParseURL(MONGOD_URL)
		if err != nil {
			panic(fmt.Sprintf("failed to establish session to %v. error: %v", MONGOD_URL, err.Error()))
		}

		tlsConfig := &tls.Config{}
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
		session, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			panic(fmt.Sprintf("failed to establish session to %v. error: %v", MONGOD_URL, err.Error()))
		}

	} else {
		session, err = mgo.Dial(MONGOD_URL)
		if err != nil {
			panic(fmt.Sprintf("failed to establish session to %v. error: %v", MONGOD_URL, err.Error()))
		}
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session

}

func (rR *ReadRequest) SetDB(dBName string) {
	rR.dBName = dBName
}

func (rR *ReadRequest) DBName() string {
	return rR.dBName
}

func (rR *ReadRequest) SetCollection(collectionName string) {
	rR.collectionName = collectionName
}

func (rR *ReadRequest) CollectionName() string {
	return rR.collectionName
}
