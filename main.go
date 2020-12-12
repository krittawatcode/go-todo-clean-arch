package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/krittawatcode/go-todo-clean-arch/databases"
	"github.com/krittawatcode/go-todo-clean-arch/deliveries/routes"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

var err error

func main() {
	databases.DB, err = gorm.Open("mysql", databases.DbURL(databases.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer databases.DB.Close()
	// run the migrations: todo struct
	databases.DB.AutoMigrate(&models.Todo{})
	//setup routes
	r := routes.SetupRouter()
	// running
	r.Run(":8080")
}
