package model

import (
	"time"

	"gorm.io/gorm"
)

// Seal 印章信息
type Seal struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"` // 印章名称
	Code        string         `gorm:"size:50;unique" json:"code"`   // 印章编号
	TypeID      uint           `gorm:"not null" json:"type_id"`      // 印章类型ID
	Image       string         `gorm:"size:255" json:"image"`        // 印章图片
	Status      int            `gorm:"default:1" json:"status"`      // 1:在库 2:借出 3:作废
	KeeperID    uint           `gorm:"not null" json:"keeper_id"`    // 保管人ID
	Description string         `gorm:"size:500" json:"description"`  // 印章说明
	Files       string         `gorm:"type:text" json:"files"`       // 附件，JSON数组
	Remark      string         `gorm:"size:500" json:"remark"`       // 备注
	CreatedBy   uint           `gorm:"not null" json:"created_by"`   // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// SealApplication 用印申请
type SealApplication struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	SealID           uint           `gorm:"not null" json:"seal_id"`       // 印章ID
	UserID           uint           `gorm:"not null" json:"user_id"`       // 申请人ID
	DepartmentID     uint           `gorm:"not null" json:"department_id"` // 申请部门ID
	Purpose          string         `gorm:"size:500" json:"purpose"`       // 用印事由
	Content          string         `gorm:"type:text" json:"content"`      // 用印内容
	Quantity         int            `gorm:"not null" json:"quantity"`      // 用印数量
	StartTime        *time.Time     `json:"start_time"`                    // 用印开始时间
	EndTime          *time.Time     `json:"end_time"`                      // 用印结束时间
	Status           int            `gorm:"default:1" json:"status"`       // 1:待审批 2:已通过 3:已驳回 4:已取消
	ApprovalRecordID *uint          `json:"approval_record_id"`            // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`        // 附件，JSON数组
	Remark           string         `gorm:"size:500" json:"remark"`        // 备注
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// SealRecord 用印记录
type SealRecord struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	ApplicationID uint           `gorm:"not null" json:"application_id"` // 用印申请ID
	BorrowTime    *time.Time     `json:"borrow_time"`                    // 借出时间
	ReturnTime    *time.Time     `json:"return_time"`                    // 归还时间
	Status        int            `gorm:"default:1" json:"status"`        // 1:已借出 2:已归还 3:异常
	Problem       string         `gorm:"size:500" json:"problem"`        // 问题说明
	Files         string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	Remark        string         `gorm:"size:500" json:"remark"`         // 备注
	CreatedBy     uint           `gorm:"not null" json:"created_by"`     // 创建人ID
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Seal) TableName() string {
	return "seals"
}

func (SealApplication) TableName() string {
	return "seal_applications"
}

func (SealRecord) TableName() string {
	return "seal_records"
}
