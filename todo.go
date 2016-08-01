package main

type Todo struct {
	Name      []string `json:"name"`
	Completed bool     `json:"completed"`
	Due       []string `json:"due"`
}

type Todos []Todo
