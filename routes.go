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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ShowTodoForm",
		"GET",
		"/addTodo",
		ShowTodoForm,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/single/{todoId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"GET",
		"/todos/create",
		TodoCreate,
	},
}
