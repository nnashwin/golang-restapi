package main

import (
	"encoding/json"
	//	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"strings"
	//"reflect"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	session := NewSession("mongodb://localhost")
	defer session.Close()

	c := session.DB("test").C("todos")

	results := Todos{}
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	session := NewSession("mongodb://localhost")
	defer session.Close()

	c := session.DB("test").C("todos")

	result := Todo{}

	err := c.Find(bson.M{"id": todoId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(&result); err != nil {
		panic(err)
	}
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {

}

func showTodoInHtml(w http.ResponseWriter, r *http.Request) {
	// session := NewSession("mongodb://localhost")

	// c := session.DB("test").C("todos")

	// result := Todo{}
	// err := c.Find(bson.M{"id": todoId}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	t := template.New("todo")
	t, _ = t.ParseFiles("static/todoList.html")
	todo := Todo{"tyler", false, "completionDate"}
	err := t.ExecuteTemplate(w, "todoList.html", todo)
	if err != nil {
		log.Fatal(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// connect with mongoDb
	session := NewSession("mongodb://localhost")

	c := session.DB("test").C("todos")

	r.ParseForm()
	stringTitle := strings.Join(r.Form["todoTitle"], "")
	completeDate := strings.Join(r.Form["completedBy"], "")
	record := Todo{stringTitle, false, completeDate}

	err := c.Insert(&record)
	if err != nil {
		log.Fatal(err)
	}
}
