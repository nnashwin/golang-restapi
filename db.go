package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

func NewSession(dbAddr string) *mgo.Session {
	session, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatal(err)
	}

	return session
}

// type DB struct {

// }

// func SaveData ()
