package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoUseCase ...
type ToDoUseCase interface {
	GetAllToDos(c *gin.Context) (todos []models.Todo, err error)
	CreateATodo(c *gin.Context) (todo models.Todo, err error)
	GetATodo(c *gin.Context) (todo models.Todo, err error)
	UpdateATodo(c *gin.Context) (todo models.Todo, err error)
	DeleteATodo(c *gin.Context) (deletedID string, err error)
}

// ToDoRepository ...
type ToDoRepository interface {
	GetAllToDos(todo *[]models.Todo) (err error)
	CreateATodo(todo *models.Todo) (err error)
	GetATodo(todo *models.Todo, id string) (err error)
	UpdateATodo(todo *models.Todo, id string) (err error)
	DeleteATodo(todo *models.Todo, id string) (err error)
}
