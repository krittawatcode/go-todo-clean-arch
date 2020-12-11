package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/database"
	"github.com/krittawatcode/go-todo-clean-arch/delivery"
	"github.com/krittawatcode/go-todo-clean-arch/repository"
	"github.com/krittawatcode/go-todo-clean-arch/usecase"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {

	todoRepo := repository.NewToDoRepository(database.DB)
	todoUseCase := usecase.NewToDoUseCase(todoRepo)
	todoHandler := delivery.NewToDoHandler(todoUseCase)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", todoHandler.GetAllToDos)
		v1.POST("todo", todoHandler.CreateATodo)
		v1.GET("todo/:id", todoHandler.GetATodo)
		v1.PUT("todo/:id", todoHandler.UpdateATodo)
		v1.DELETE("todo/:id", todoHandler.DeleteATodo)
	}
	return r
}
