package main

import (
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	go log.Fatal(http.ListenAndServe(":8080", router))
}
