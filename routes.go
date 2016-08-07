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
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"UpdateTodo",
		"POST",
		"/todos/update/{todoId}",
		UpdateTodo,
	},
	Route{
		"HandleDelete",
		"DELETE",
		"/todos/delete/id={todoId}",
		HandleDelete,
	},
}
