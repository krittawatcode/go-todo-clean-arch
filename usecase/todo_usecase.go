package usecase

import (
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

func (t *todoUseCase) GetAllToDos(c *gin.Context) (todos []models.ToDo, err error) {
	var todo []models.ToDo
	handleErr := t.todoRepo.GetAllToDos(&todo)
	return todos, handleErr
}

func (t *todoUseCase) CreateATodo(c *gin.Context) (todo models.ToDo, err error) {
	var newToDo models.ToDo
	handleErr := t.todoRepo.CreateATodo(&newToDo)
	return newToDo, handleErr
}

func (t *todoUseCase) GetATodo(c *gin.Context) (todo models.ToDo, err error) {
	id := c.Params.ByName("id")
	var respToDo models.ToDo
	handleErr := t.todoRepo.GetATodo(&respToDo, id)
	return respToDo, handleErr
}

func (t *todoUseCase) UpdateATodo(c *gin.Context) (todo models.ToDo, err error) {
	id := c.Params.ByName("id")
	var respToDo models.ToDo
	errResp := t.todoRepo.GetATodo(&respToDo, id)
	if errResp != nil {
		return respToDo, errResp
	}
	handleErr := t.todoRepo.UpdateATodo(&respToDo, id)
	return respToDo, handleErr
}

func (t *todoUseCase) DeleteATodo(c *gin.Context) (deletedID string, err error) {
	id := c.Params.ByName("id")
	var respToDo models.ToDo
	handleErr := t.todoRepo.DeleteATodo(&respToDo, id)
	return id, handleErr
}
