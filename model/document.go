package model

import (
	"time"

	"gorm.io/gorm"
)

// DocumentType 公文类型
type DocumentType struct {
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

// Document 公文信息
type Document struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	Title         string         `gorm:"size:200;not null" json:"title"`  // 公文标题
	Code          string         `gorm:"size:50;unique" json:"code"`      // 公文编号
	TypeID        uint           `gorm:"not null" json:"type_id"`         // 公文类型ID
	SecurityLevel int            `gorm:"default:1" json:"security_level"` // 密级：1:普通 2:秘密 3:机密 4:绝密
	UrgencyLevel  int            `gorm:"default:1" json:"urgency_level"`  // 紧急程度：1:普通 2:紧急 3:特急
	Content       string         `gorm:"type:text" json:"content"`        // 公文内容
	Keywords      string         `gorm:"size:200" json:"keywords"`        // 关键词
	DraftUserID   uint           `gorm:"not null" json:"draft_user_id"`   // 拟稿人ID
	DraftDeptID   uint           `gorm:"not null" json:"draft_dept_id"`   // 拟稿部门ID
	DraftDate     *time.Time     `json:"draft_date"`                      // 拟稿日期
	SignDate      *time.Time     `json:"sign_date"`                       // 签发日期
	Status        int            `gorm:"default:1" json:"status"`         // 1:草稿 2:审批中 3:已签发 4:已归档 5:已作废
	Files         string         `gorm:"type:text" json:"files"`          // 附件，JSON数组
	Remark        string         `gorm:"size:500" json:"remark"`          // 备注
	CreatedBy     uint           `gorm:"not null" json:"created_by"`      // 创建人ID
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// DocumentApproval 公文审批流程
type DocumentApproval struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	DocumentID     uint           `gorm:"not null" json:"document_id"`      // 公文ID
	ApproverID     uint           `gorm:"not null" json:"approver_id"`      // 审批人ID
	ApproverDeptID uint           `gorm:"not null" json:"approver_dept_id"` // 审批部门ID
	Sort           int            `gorm:"default:0" json:"sort"`            // 审批顺序
	Status         int            `gorm:"default:1" json:"status"`          // 1:待审批 2:已通过 3:已驳回
	Comment        string         `gorm:"size:500" json:"comment"`          // 审批意见
	ApprovalTime   *time.Time     `json:"approval_time"`                    // 审批时间
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// DocumentDistribution 公文分发
type DocumentDistribution struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	DocumentID     uint           `gorm:"not null" json:"document_id"`      // 公文ID
	ReceiverID     uint           `gorm:"not null" json:"receiver_id"`      // 接收人ID
	ReceiverDeptID uint           `gorm:"not null" json:"receiver_dept_id"` // 接收部门ID
	Status         int            `gorm:"default:1" json:"status"`          // 1:未读 2:已读
	ReadTime       *time.Time     `json:"read_time"`                        // 阅读时间
	CreatedBy      uint           `gorm:"not null" json:"created_by"`       // 分发人ID
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// DocumentArchive 公文归档
type DocumentArchive struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	DocumentID  uint           `gorm:"not null" json:"document_id"`      // 公文ID
	ArchiveNo   string         `gorm:"size:50;unique" json:"archive_no"` // 归档号
	Location    string         `gorm:"size:200" json:"location"`         // 存放位置
	Status      int            `gorm:"default:1" json:"status"`          // 1:已归档 2:已借阅 3:已销毁
	ArchiveDate *time.Time     `json:"archive_date"`                     // 归档日期
	DestroyDate *time.Time     `json:"destroy_date"`                     // 销毁日期
	Files       string         `gorm:"type:text" json:"files"`           // 附件，JSON数组
	Remark      string         `gorm:"size:500" json:"remark"`           // 备注
	CreatedBy   uint           `gorm:"not null" json:"created_by"`       // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// DocumentBorrow 公文借阅
type DocumentBorrow struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	DocumentID     uint           `gorm:"not null" json:"document_id"`      // 公文ID
	BorrowerID     uint           `gorm:"not null" json:"borrower_id"`      // 借阅人ID
	BorrowerDeptID uint           `gorm:"not null" json:"borrower_dept_id"` // 借阅部门ID
	Purpose        string         `gorm:"size:500" json:"purpose"`          // 借阅用途
	BorrowDate     *time.Time     `json:"borrow_date"`                      // 借阅日期
	ReturnDate     *time.Time     `json:"return_date"`                      // 归还日期
	Status         int            `gorm:"default:1" json:"status"`          // 1:已借出 2:已归还 3:已逾期
	Files          string         `gorm:"type:text" json:"files"`           // 附件，JSON数组
	Remark         string         `gorm:"size:500" json:"remark"`           // 备注
	CreatedBy      uint           `gorm:"not null" json:"created_by"`       // 创建人ID
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (DocumentType) TableName() string {
	return "document_types"
}

func (Document) TableName() string {
	return "documents"
}

func (DocumentApproval) TableName() string {
	return "document_approvals"
}

func (DocumentDistribution) TableName() string {
	return "document_distributions"
}

func (DocumentArchive) TableName() string {
	return "document_archives"
}

func (DocumentBorrow) TableName() string {
	return "document_borrows"
}
