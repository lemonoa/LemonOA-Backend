package model

import (
	"time"

	"gorm.io/gorm"
)

// Department 部门模型
type Department struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	ParentID  *uint          `gorm:"default:null" json:"parent_id"`
	Level     int            `gorm:"default:1" json:"level"`
	Sort      int            `gorm:"default:0" json:"sort"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Department) TableName() string {
	return "departments"
}
