package main

import (
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	// session, err := mgo.Dial("mongodb://localhost")

	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// session.SetMode(mgo.Monotonic, true)

}
