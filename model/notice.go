package model

import (
	"time"

	"gorm.io/gorm"
)

// Notice 公告信息
type Notice struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title"` // 公告标题
	TypeID    uint           `gorm:"not null" json:"type_id"`        // 公告类型ID
	Content   string         `gorm:"type:text" json:"content"`       // 公告内容
	StartTime *time.Time     `json:"start_time"`                     // 生效时间
	EndTime   *time.Time     `json:"end_time"`                       // 失效时间
	Priority  int            `gorm:"default:1" json:"priority"`      // 优先级：1:普通 2:重要 3:紧急
	Status    int            `gorm:"default:1" json:"status"`        // 1:草稿 2:已发布 3:已撤回
	Files     string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	Remark    string         `gorm:"size:500" json:"remark"`         // 备注
	CreatedBy uint           `gorm:"not null" json:"created_by"`     // 创建人ID
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// NoticeRead 公告阅读记录
type NoticeRead struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	NoticeID  uint           `gorm:"not null" json:"notice_id"` // 公告ID
	UserID    uint           `gorm:"not null" json:"user_id"`   // 用户ID
	ReadTime  *time.Time     `json:"read_time"`                 // 阅读时间
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Notice) TableName() string {
	return "notices"
}

func (NoticeRead) TableName() string {
	return "notice_reads"
}
