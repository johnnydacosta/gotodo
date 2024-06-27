package todo

import (
	"cmp"
	"strings"
)

type Todo struct {
	ID   string `json:"id" param:"id"`
	Name string `json:"name" param:"name"`
	Done bool   `json:"done" param:"done"`
}

type CreateTodoDTO struct {
	Name string `form:"name"`
}

type UpdateTodoDTO struct {
	ID   string `param:"id"`
	Name string `form:"name"`
	Done bool   `form:"done"`
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
