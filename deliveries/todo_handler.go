package deliveries

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoHandler use for handle framwork here and present as a controller
type ToDoHandler struct {
	todoUseCase domains.ToDoUseCase
}

// NewToDoHandler ...
func NewToDoHandler(usecase domains.ToDoUseCase) *ToDoHandler {
	return &ToDoHandler{
		todoUseCase: usecase,
	}
}

// GetAllTodo ...
func (t *ToDoHandler) GetAllTodo(c *gin.Context) {
	resp, err := t.todoUseCase.GetAllTodo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateATodo ...
func (t *ToDoHandler) CreateATodo(c *gin.Context) {
	var newToDo models.Todo
	if err := c.ShouldBind(&newToDo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := t.todoUseCase.CreateATodo(&newToDo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newToDo)
	}
}

// GetATodo ...
func (t *ToDoHandler) GetATodo(c *gin.Context) {
	var newToDo models.Todo
	id := c.Params.ByName("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err = t.todoUseCase.GetATodo(&newToDo, intID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newToDo)
	}
}

// UpdateATodo ...
func (t *ToDoHandler) UpdateATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var reqToDo models.Todo
	if err := c.ShouldBind(&reqToDo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err = t.todoUseCase.UpdateATodo(&reqToDo, intID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reqToDo)
	}
}

// DeleteATodo ...
func (t *ToDoHandler) DeleteATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	var respToDo models.Todo
	err = t.todoUseCase.DeleteATodo(&respToDo, intID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
