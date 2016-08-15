package main

type Todo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"description"`
	Completed bool   `json:"completed"`
	Due       string `json:"due"`
}

type Todos []Todo
