package models

// Todo ...
type Todo struct {
	ID          uint   `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TableName use to specific table
func (b *Todo) TableName() string {
	return "todo"
}
