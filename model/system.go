package model

import (
	"time"

	"gorm.io/gorm"
)

// SystemConfig 系统配置
type SystemConfig struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Key       string         `gorm:"size:50;not null;unique" json:"key"`
	Value     string         `gorm:"type:text" json:"value"`
	Desc      string         `gorm:"size:255" json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Module 功能模块
type Module struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Code      string         `gorm:"size:50;not null;unique" json:"code"`
	Icon      string         `gorm:"size:50" json:"icon"`
	Sort      int            `gorm:"default:0" json:"sort"`
	Status    int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	ParentID  *uint          `json:"parent_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ModuleConfig 模块配置
type ModuleConfig struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	ModuleID  uint           `gorm:"not null" json:"module_id"`
	Key       string         `gorm:"size:50;not null" json:"key"`
	Value     string         `gorm:"type:text" json:"value"`
	Desc      string         `gorm:"size:255" json:"desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// FunctionNode 功能节点
type FunctionNode struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	ModuleID  uint           `gorm:"not null" json:"module_id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Code      string         `gorm:"size:50;not null;unique" json:"code"`
	Type      int            `gorm:"default:1" json:"type"` // 1:菜单 2:按钮 3:接口
	Path      string         `gorm:"size:255" json:"path"`
	Method    string         `gorm:"size:10" json:"method"`
	Sort      int            `gorm:"default:0" json:"sort"`
	Status    int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// RoleFunction 角色功能权限
type RoleFunction struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	RoleID         uint           `gorm:"not null" json:"role_id"`
	FunctionNodeID uint           `gorm:"not null" json:"function_node_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// OperationLog 操作日志
type OperationLog struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Module    string         `gorm:"size:50" json:"module"`
	Action    string         `gorm:"size:50" json:"action"`
	Method    string         `gorm:"size:10" json:"method"`
	Path      string         `gorm:"size:255" json:"path"`
	Params    string         `gorm:"type:text" json:"params"`
	Response  string         `gorm:"type:text" json:"response"`
	IP        string         `gorm:"size:50" json:"ip"`
	UserAgent string         `gorm:"size:255" json:"user_agent"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Attachment 附件
type Attachment struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	Path        string         `gorm:"size:255;not null" json:"path"`
	Size        int64          `gorm:"not null" json:"size"`
	Type        string         `gorm:"size:50" json:"type"`
	UploadedBy  uint           `gorm:"not null" json:"uploaded_by"`
	Module      string         `gorm:"size:50" json:"module"`
	RelatedID   uint           `json:"related_id"`
	RelatedType string         `gorm:"size:50" json:"related_type"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// BackupRecord 备份记录
type BackupRecord struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Path      string         `gorm:"size:255;not null" json:"path"`
	Size      int64          `gorm:"not null" json:"size"`
	Type      int            `gorm:"default:1" json:"type"`   // 1:全量备份 2:增量备份
	Status    int            `gorm:"default:1" json:"status"` // 1:备份中 2:备份成功 3:备份失败
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ScheduledTask 定时任务
type ScheduledTask struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	Name        string     `gorm:"size:50;not null" json:"name"`
	Description string     `gorm:"size:255" json:"description"`
	Cron        string     `gorm:"size:50;not null" json:"cron"`
	Command     string     `gorm:"type:text;not null" json:"command"`
	Status      int        `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	LastRunAt   *time.Time `json:"last_run_at"`
}
