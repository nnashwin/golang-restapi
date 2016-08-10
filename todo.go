package main

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Due         string `json:"due"`
}

type Todos []Todo
