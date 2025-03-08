package model

import (
	"time"

	"gorm.io/gorm"
)

// Employee 员工模型
type Employee struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	Name         string         `gorm:"size:50;not null" json:"name"`
	Email        string         `gorm:"size:100;unique" json:"email"`
	Phone        string         `gorm:"size:20" json:"phone"`
	Avatar       string         `gorm:"size:255" json:"avatar"`
	DepartmentID uint           `gorm:"not null" json:"department_id"`
	Position     string         `gorm:"size:50" json:"position"`
	Status       int            `gorm:"default:1" json:"status"` // 1:在职 2:离职
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Employee) TableName() string {
	return "employees"
}
