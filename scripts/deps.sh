#!/bin/bash

echo "Importing Deps"
go get -u -v "github.com/dgrijalva/jwt-go"
go get -u -v "gopkg.in/mgo.v2"
go get -u -v "golang.org/x/crypto/bcrypt"