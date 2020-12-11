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

func (t *todoUseCase) GetAllToDos(c *gin.Context) (todos []models.Todo, err error) {
	var todo []models.Todo
	handleErr := t.todoRepo.GetAllToDos(&todo)
	return todo, handleErr
}

func (t *todoUseCase) CreateATodo(c *gin.Context) (todo models.Todo, err error) {
	var newToDo models.Todo
	if err := c.ShouldBind(&newToDo); err != nil {
		return newToDo, err
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
	// check avaliable
	var checkingTodo models.Todo
	errResp := t.todoRepo.GetATodo(&checkingTodo, id)
	if errResp != nil {
		return checkingTodo, errResp
	}
	// update
	var reqToDo models.Todo
	if err := c.ShouldBind(&reqToDo); err != nil {
		return reqToDo, err
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
