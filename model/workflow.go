package model

import (
	"time"

	"gorm.io/gorm"
)

// WorkflowType 流程类型
type WorkflowType struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"` // 类型名称
	Code        string         `gorm:"size:50;unique" json:"code"`   // 类型编码
	Description string         `gorm:"size:500" json:"description"`  // 类型说明
	Sort        int            `gorm:"default:0" json:"sort"`        // 排序
	Status      int            `gorm:"default:1" json:"status"`      // 1:启用 2:禁用
	CreatedBy   uint           `gorm:"not null" json:"created_by"`   // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// WorkflowDefinition 流程定义
type WorkflowDefinition struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name"` // 流程名称
	TypeID      uint           `gorm:"not null" json:"type_id"`       // 流程类型ID
	Description string         `gorm:"size:500" json:"description"`   // 流程说明
	Form        string         `gorm:"type:text" json:"form"`         // 表单配置，JSON格式
	Status      int            `gorm:"default:1" json:"status"`       // 1:草稿 2:已发布 3:已停用
	Version     int            `gorm:"default:1" json:"version"`      // 版本号
	CreatedBy   uint           `gorm:"not null" json:"created_by"`    // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// WorkflowNode 流程节点
type WorkflowNode struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	DefinitionID uint           `gorm:"not null" json:"definition_id"` // 流程定义ID
	Name         string         `gorm:"size:50;not null" json:"name"`  // 节点名称
	Type         int            `gorm:"not null" json:"type"`          // 1:开始 2:审批 3:抄送 4:条件 5:并行 6:结束
	Config       string         `gorm:"type:text" json:"config"`       // 节点配置，JSON格式
	Sort         int            `gorm:"default:0" json:"sort"`         // 排序
	CreatedBy    uint           `gorm:"not null" json:"created_by"`    // 创建人ID
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// WorkflowInstance 流程实例
type WorkflowInstance struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	DefinitionID uint           `gorm:"not null" json:"definition_id"`  // 流程定义ID
	Title        string         `gorm:"size:200;not null" json:"title"` // 流程标题
	Content      string         `gorm:"type:text" json:"content"`       // 流程内容
	FormData     string         `gorm:"type:text" json:"form_data"`     // 表单数据，JSON格式
	Status       int            `gorm:"default:1" json:"status"`        // 1:进行中 2:已完成 3:已取消
	StartTime    *time.Time     `json:"start_time"`                     // 开始时间
	EndTime      *time.Time     `json:"end_time"`                       // 结束时间
	Files        string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	CreatedBy    uint           `gorm:"not null" json:"created_by"`     // 创建人ID
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// WorkflowTask 流程任务
type WorkflowTask struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	InstanceID uint           `gorm:"not null" json:"instance_id"` // 流程实例ID
	NodeID     uint           `gorm:"not null" json:"node_id"`     // 流程节点ID
	AssigneeID uint           `gorm:"not null" json:"assignee_id"` // 处理人ID
	Action     int            `gorm:"default:0" json:"action"`     // 0:未处理 1:同意 2:驳回 3:转办 4:已阅
	Comment    string         `gorm:"size:500" json:"comment"`     // 处理意见
	HandleTime *time.Time     `json:"handle_time"`                 // 处理时间
	Status     int            `gorm:"default:1" json:"status"`     // 1:待处理 2:已处理 3:已转办 4:已取消
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (WorkflowType) TableName() string {
	return "workflow_types"
}

func (WorkflowDefinition) TableName() string {
	return "workflow_definitions"
}

func (WorkflowNode) TableName() string {
	return "workflow_nodes"
}

func (WorkflowInstance) TableName() string {
	return "workflow_instances"
}

func (WorkflowTask) TableName() string {
	return "workflow_tasks"
}
