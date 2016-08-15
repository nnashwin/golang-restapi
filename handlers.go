package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type test_struct struct {
	Test string
}

var templates = template.Must(template.ParseGlob("static/tmpl/*"))

func ShowAllTodos(w http.ResponseWriter, r *http.Request) {
	session := NewSession("mongodb://localhost")
	defer session.Close()

	results := getAllRecords("test", "todos")
	templates.ExecuteTemplate(w, "index page", results)
}

func ShowSingleTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	record := getRecord("test", "todos", todoId)
	showTodoInHtml(w, record)
}

func showTodoInHtml(w http.ResponseWriter, todo Todo) {
	t := template.New("todo")
	t, _ = t.ParseFiles("static/tmpl/singleTodoPage.html")
	err := t.ExecuteTemplate(w, "singleTodoPage.html", todo)
	if err != nil {
		log.Fatal(err)
	}
}

func ShowCreateTodoForm(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "form", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo

	r.ParseForm()

	newTodo.Id = r.PostFormValue("todoId")
	newTodo.Name = r.PostFormValue("todoTitle")
	newTodo.Due = r.PostFormValue("completedBy")
	newTodo.Desc = r.PostFormValue("description")
	newTodo.Completed = false

	CreateTodo(w, r, newTodo)

	ShowAllTodos(w, r)
}

func CreateTodo(w http.ResponseWriter, r *http.Request, todo Todo) {
	// connect with mongoDb
	session := NewSession("mongodb://localhost")
	defer session.Close()

	c := session.DB("test").C("todos")

	err := c.Insert(&todo)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	DeleteTodo(w, r, todoId)
	ShowAllTodos(w, r)
}

func HandlePut(w http.ResponseWriter, r *http.Request) {
	var updatedTodo Todo
	err := r.ParseForm()
	check(err)
	updatedTodo.Id = r.PostFormValue("todoId")
	updatedTodo.Name = r.PostFormValue("todoName")
	updatedTodo.Desc = r.PostFormValue("description")
	updatedTodo.Due = r.PostFormValue("dueDate")

	complete, err := strconv.ParseBool(r.PostFormValue("completeStatus"))
	check(err)
	updatedTodo.Completed = complete

	UpdateTodo(w, r, updatedTodo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request, todo Todo) {
	session := NewSession("mongodb://localhost")
	defer session.Close()

	c := session.DB("test").C("todos")
	log.Println("getting that session")
	log.Printf("%+v", todo)
	c.Update(bson.M{"id": todo.Id}, todo)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request, todoId string) {
	session := NewSession("mongodb://localhost")
	defer session.Close()

	c := session.DB("test").C("todos")
	c.Remove(bson.M{"id": todoId})
}
