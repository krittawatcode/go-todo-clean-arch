package domains

import (
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoUseCase ...
type ToDoUseCase interface {
	GetAllTodo() (t []models.Todo, err error)
	CreateATodo(t *models.Todo) (err error)
	GetATodo(t *models.Todo, id string) (err error)
	UpdateATodo(t *models.Todo, id string) (err error)
	DeleteATodo(t *models.Todo, id string) (err error)
}

// ToDoRepository ...
type ToDoRepository interface {
	GetAllTodo(t *[]models.Todo) (err error)
	CreateATodo(t *models.Todo) (err error)
	GetATodo(t *models.Todo, id string) (err error)
	UpdateATodo(t *models.Todo, id string) (err error)
	DeleteATodo(t *models.Todo, id string) (err error)
}
