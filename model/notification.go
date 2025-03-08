package model

import (
	"time"

	"gorm.io/gorm"
)

// Notification 消息通知模型
type Notification struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"size:100;not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	Type      int            `gorm:"not null" json:"type"`    // 1:系统消息 2:审批通知
	Status    int            `gorm:"default:1" json:"status"` // 1:未读 2:已读
	UserID    uint           `gorm:"not null" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Notification) TableName() string {
	return "notifications"
}
