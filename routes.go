package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ShowAllTodos",
		"GET",
		"/todos",
		ShowAllTodos,
	},
	Route{
		"ShowSingleTodo",
		"GET",
		"/todos/single/{todoId}",
		ShowSingleTodo,
	},
	Route{
		"ShowCreateTodoForm",
		"GET",
		"/todoCreateForm",
		ShowCreateTodoForm,
	},
	Route{
		"HandleCreate",
		"POST",
		"/todos",
		HandleCreate,
	},
	Route{
		"HandleDelete",
		"DELETE",
		"/todos/delete/id={todoId}",
		HandleDelete,
	},
	Route{
		"HandlePut",
		"PUT",
		"/todos/edit/id={todoId}",
		HandlePut,
	},
}
