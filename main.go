package main

import (
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Person struct {
	Name  string
	Phone string
}

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	// session, err := mgo.Dial("mongodb://localhost")

	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// session.SetMode(mgo.Monotonic, true)

	// c := session.DB("test").C("people")

	// err = c.Insert(&Person{"Ale", "+886 0906343210"},
	// 	&Person{"Cla", "+886 0902343291"})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Phone:", result.Phone)
}
