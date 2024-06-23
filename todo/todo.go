package todo

import (
	"cmp"
	"strings"
)

type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type CreateTodoDTO struct {
	Name string `json:"name"`
}

type UpdateTodoDTO struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type Todos []Todo

func TodoCmp(a, b Todo) int {
	return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
}

func New(name string) Todo {
	return Todo{
		Name: name,
	}
}
