package models

// Todo ...
type Todo struct {
	ID          int    `gorm:"column:id;primary_key" json:"id"`
	Title       string `gorm:"column:title" json:"title" `
	Description string `gorm:"column:description" json:"description"`
}

// TableName use to specific table
func (b *Todo) TableName() string {
	return "todo"
}
