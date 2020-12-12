package usecase

import (
	"github.com/krittawatcode/go-todo-clean-arch/domain"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// UseCase don't give a shit about the world!!
type todoUseCase struct {
	todoRepo domain.ToDoRepository
}

// NewToDoUseCase ...
func NewToDoUseCase(repo domain.ToDoRepository) domain.ToDoUseCase {
	return &todoUseCase{
		todoRepo: repo,
	}
}

func (t *todoUseCase) GetAllToDos() (todos []models.Todo, err error) {
	var todo []models.Todo
	handleErr := t.todoRepo.GetAllToDos(&todo)
	return todo, handleErr
}

func (t *todoUseCase) CreateATodo(input *models.Todo) (err error) {
	handleErr := t.todoRepo.CreateATodo(input)
	return handleErr
}

func (t *todoUseCase) GetATodo(input *models.Todo, id string) (err error) {
	handleErr := t.todoRepo.GetATodo(input, id)
	return handleErr
}

func (t *todoUseCase) UpdateATodo(input *models.Todo, id string) (err error) {
	// check avaliable
	var checkingTodo models.Todo
	errResp := t.todoRepo.GetATodo(&checkingTodo, id)
	if errResp != nil {
		return errResp
	}
	// update
	handleErr := t.todoRepo.UpdateATodo(input, id)
	return handleErr
}

func (t *todoUseCase) DeleteATodo(input *models.Todo, id string) (err error) {
	handleErr := t.todoRepo.DeleteATodo(input, id)
	return handleErr
}
