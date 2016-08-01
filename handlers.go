package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("todos")

	results := Todos{}
	err = c.Find(nil).All(&results)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("todos")

	result := Todo{}
	err = c.Find(bson.M{"id": todoId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(&result); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// connect with mongoDb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("todos")

	r.ParseForm()
	record := Todo{(r.Form["todoTitle"]), false, r.Form["completedBy"]}

	err = c.Insert(&record)
	if err != nil {
		log.Fatal(err)
	}
}
