package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/krittawatcode/go-todo-clean-arch/database"
	"github.com/krittawatcode/go-todo-clean-arch/delivery/routes"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer database.DB.Close()
	// run the migrations: todo struct
	database.DB.AutoMigrate(&models.Todo{})
	//setup routes
	r := routes.SetupRouter()
	// running
	r.Run(":8080")
}
