package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func NewSession(dbAddr string) *mgo.Session {
	session, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Fatal(err)
	}

	return session
}

func getColl(dbAddr string, dbName string, collName string) *mgo.Collection {
	session := NewSession(dbAddr)

	c := session.DB(dbName).C(collName)

	return c
}

func getRecord(dbName string, collName string, recordId string) Todo {
	c := getColl("mongodb://localhost", dbName, collName)

	result := Todo{}
	err := c.Find(bson.M{"id": recordId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getAllRecords(dbName string, collName string) Todos {
	c := getColl("mongodb://localhost", dbName, collName)

	results := Todos{}
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

// type DB struct {

// }

// func SaveData ()
