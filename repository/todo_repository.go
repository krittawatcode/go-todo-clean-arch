package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/krittawatcode/go-todo-clean-arch/domain"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

type todoRepository struct {
	conn *gorm.DB
}

// NewToDoRepository ...
func NewToDoRepository(conn *gorm.DB) domain.ToDoRepository {
	return &todoRepository{conn}
}

func (t *todoRepository) GetAllToDos(todo *[]models.ToDo) (err error) {
	if err = t.conn.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) CreateATodo(todo *models.ToDo) (err error) {
	if err = t.conn.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) GetATodo(todo *models.ToDo, id string) (err error) {
	if err := t.conn.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) UpdateATodo(todo *models.ToDo, id string) (err error) {
	// fmt.Println(todo)
	t.conn.Save(todo) // save all field
	return nil
}

func (t *todoRepository) DeleteATodo(todo *models.ToDo, id string) (err error) {
	t.conn.Where("id = ?", id).Delete(todo)
	return nil
}
