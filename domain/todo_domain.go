package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoUseCase ...
type ToDoUseCase interface {
	GetAllToDos(c *gin.Context) (todos []models.ToDo, err error)
	CreateATodo(c *gin.Context) (todo models.ToDo, err error)
	GetATodo(c *gin.Context) (todo models.ToDo, err error)
	UpdateATodo(c *gin.Context) (todo models.ToDo, err error)
	DeleteATodo(c *gin.Context) (deletedID string, err error)
}

// ToDoRepository ...
type ToDoRepository interface {
	GetAllToDos(todo *[]models.ToDo) (err error)
	CreateATodo(todo *models.ToDo) (err error)
	GetATodo(todo *models.ToDo, id string) (err error)
	UpdateATodo(todo *models.ToDo, id string) (err error)
	DeleteATodo(todo *models.ToDo, id string) (err error)
}
