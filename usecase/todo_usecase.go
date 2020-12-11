package usecase

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/domain"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

type todoUseCase struct {
	todoRepo domain.ToDoRepository
}

// NewToDoUseCase ...
func NewToDoUseCase(repo domain.ToDoRepository) domain.ToDoUseCase {
	return &todoUseCase{
		todoRepo: repo,
	}
}

func (t *todoUseCase) GetAllToDos(c *gin.Context) (todos []models.Todo, err error) {
	var todo []models.Todo
	handleErr := t.todoRepo.GetAllToDos(&todo)
	fmt.Println(handleErr)
	fmt.Println(todo)
	return todo, handleErr
}

func (t *todoUseCase) CreateATodo(c *gin.Context) (todo models.Todo, err error) {
	fmt.Println("Usecase : CreateATodo")
	var newToDo models.Todo
	if err := c.ShouldBind(&newToDo); err != nil {
		log.Fatal("cannot bind")
	}
	handleErr := t.todoRepo.CreateATodo(&newToDo)
	return newToDo, handleErr
}

func (t *todoUseCase) GetATodo(c *gin.Context) (todo models.Todo, err error) {
	id := c.Params.ByName("id")
	var respToDo models.Todo
	handleErr := t.todoRepo.GetATodo(&respToDo, id)
	return respToDo, handleErr
}

func (t *todoUseCase) UpdateATodo(c *gin.Context) (todo models.Todo, err error) {
	id := c.Params.ByName("id")
	var reqToDo models.Todo
	if err := c.ShouldBind(&reqToDo); err != nil {
		log.Fatal("cannot bind")
	}
	var requestToDo models.Todo
	errResp := t.todoRepo.GetATodo(&requestToDo, id)
	if errResp != nil {
		return requestToDo, errResp
	}
	handleErr := t.todoRepo.UpdateATodo(&reqToDo, id)
	return reqToDo, handleErr
}

func (t *todoUseCase) DeleteATodo(c *gin.Context) (deletedID string, err error) {
	id := c.Params.ByName("id")
	var respToDo models.Todo
	handleErr := t.todoRepo.DeleteATodo(&respToDo, id)
	return id, handleErr
}
