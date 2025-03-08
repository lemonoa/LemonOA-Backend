package model

import (
	"time"

	"gorm.io/gorm"
)

// Todo 待办事项模型
type Todo struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"size:100;not null" json:"title"`
	Content     string         `gorm:"type:text" json:"content"`
	Type        int            `gorm:"not null" json:"type"`    // 1:审批任务 2:工作任务
	Status      int            `gorm:"default:1" json:"status"` // 1:待完成 2:已完成
	UserID      uint           `gorm:"not null" json:"user_id"`
	DueDate     *time.Time     `json:"due_date"`
	CompletedAt *time.Time     `json:"completed_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Todo) TableName() string {
	return "todos"
}
