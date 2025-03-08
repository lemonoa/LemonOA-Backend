package model

import (
	"time"

	"gorm.io/gorm"
)

// ApprovalType 审批类型
type ApprovalType struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Code      string         `gorm:"size:50;not null;unique" json:"code"`
	Sort      int            `gorm:"default:0" json:"sort"`
	Status    int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ApprovalFlow 审批流程
type ApprovalFlow struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	ApprovalTypeID uint           `gorm:"not null" json:"approval_type_id"`
	Name           string         `gorm:"size:50;not null" json:"name"`
	Description    string         `gorm:"size:255" json:"description"`
	Status         int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ApprovalNode 审批节点
type ApprovalNode struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	ApprovalFlowID uint           `gorm:"not null" json:"approval_flow_id"`
	Name           string         `gorm:"size:50;not null" json:"name"`
	Type           int            `gorm:"not null" json:"type"` // 1:指定人员 2:指定角色 3:指定部门负责人
	ApproverID     *uint          `json:"approver_id"`          // 指定人员ID
	RoleID         *uint          `json:"role_id"`              // 指定角色ID
	Sort           int            `gorm:"default:0" json:"sort"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ApprovalRecord 审批记录
type ApprovalRecord struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	ApprovalFlowID uint           `gorm:"not null" json:"approval_flow_id"`
	Title          string         `gorm:"size:100;not null" json:"title"`
	Content        string         `gorm:"type:text" json:"content"`
	Status         int            `gorm:"default:1" json:"status"` // 1:待审批 2:审批中 3:已通过 4:已驳回
	ApplicantID    uint           `gorm:"not null" json:"applicant_id"`
	CurrentNodeID  uint           `gorm:"not null" json:"current_node_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ApprovalNodeRecord 审批节点记录
type ApprovalNodeRecord struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	ApprovalRecordID uint           `gorm:"not null" json:"approval_record_id"`
	ApprovalNodeID   uint           `gorm:"not null" json:"approval_node_id"`
	ApproverID       uint           `gorm:"not null" json:"approver_id"`
	Status           int            `gorm:"default:1" json:"status"` // 1:待审批 2:已通过 3:已驳回
	Comment          string         `gorm:"type:text" json:"comment"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (ApprovalType) TableName() string {
	return "approval_types"
}

func (ApprovalFlow) TableName() string {
	return "approval_flows"
}

func (ApprovalNode) TableName() string {
	return "approval_nodes"
}

func (ApprovalRecord) TableName() string {
	return "approval_records"
}

func (ApprovalNodeRecord) TableName() string {
	return "approval_node_records"
}
