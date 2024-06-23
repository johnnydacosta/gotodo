package todo

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
)

var DB = initDB(4)

func FindAll() Todos {
	todos := make(Todos, 0, len(DB))
	for _, v := range DB {
		todos = append(todos, *v)
	}

	slices.SortFunc(todos, TodoCmp)
	return todos
}

func FindById(id string) *Todo {
	t, ok := DB[id]

	if ok {
		return t
	}
	return nil
}

func Create(form CreateTodoDTO) *Todo {
	id := uuid.New().String()
	t := &Todo{
		ID:   id,
		Name: form.Name,
	}

	DB[id] = t
	return t
}

func Update(id string, form UpdateTodoDTO) *Todo {
	t, ok := DB[id]

	if !ok {
		return nil
	}

	t.Name = form.Name
	t.Done = form.Done

	return t
}

func Delete(id string) {
	delete(DB, id)
}

func initDB(n int) map[string]*Todo {
	db := make(map[string]*Todo, n)

	for i := 0; i < n; i++ {
		id := uuid.New().String()
		db[id] = &Todo{
			ID:   id,
			Name: fmt.Sprintf("Todo %d", i),
		}
	}
	return db
}
