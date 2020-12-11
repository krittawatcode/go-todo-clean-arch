package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/domain"
)

// ToDoHandler ...
type ToDoHandler struct {
	todoUseCase domain.ToDoUseCase
}

// NewToDoHandler ...
func NewToDoHandler(usecase domain.ToDoUseCase) *ToDoHandler {
	return &ToDoHandler{
		todoUseCase: usecase,
	}
}

// GetAllToDos ...
func (t *ToDoHandler) GetAllToDos(c *gin.Context) {
	resp, err := t.todoUseCase.GetAllToDos(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateATodo ...
func (t *ToDoHandler) CreateATodo(c *gin.Context) {
	resp, err := t.todoUseCase.CreateATodo(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// GetATodo ...
func (t *ToDoHandler) GetATodo(c *gin.Context) {
	resp, err := t.todoUseCase.GetATodo(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// UpdateATodo ...
func (t *ToDoHandler) UpdateATodo(c *gin.Context) {
	resp, err := t.todoUseCase.UpdateATodo(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// DeleteATodo ...
func (t *ToDoHandler) DeleteATodo(c *gin.Context) {
	respID, err := t.todoUseCase.DeleteATodo(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + respID: "deleted"})
	}
}
