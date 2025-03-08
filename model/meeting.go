package model

import (
	"time"

	"gorm.io/gorm"
)

// MeetingRoom 会议室
type MeetingRoom struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"` // 会议室名称
	Location    string         `gorm:"size:255" json:"location"`     // 位置
	Capacity    int            `gorm:"not null" json:"capacity"`     // 容纳人数
	Facilities  string         `gorm:"type:text" json:"facilities"`  // 设施配置，JSON数组
	Description string         `gorm:"size:500" json:"description"`  // 描述
	Status      int            `gorm:"default:1" json:"status"`      // 1:可用 2:维护中 3:停用
	Sort        int            `gorm:"default:0" json:"sort"`        // 排序
	CreatedBy   uint           `gorm:"not null" json:"created_by"`   // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// MeetingReservation 会议室预约
type MeetingReservation struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	RoomID       uint           `gorm:"not null" json:"room_id"`        // 会议室ID
	Title        string         `gorm:"size:100;not null" json:"title"` // 会议主题
	StartTime    *time.Time     `gorm:"not null" json:"start_time"`     // 开始时间
	EndTime      *time.Time     `gorm:"not null" json:"end_time"`       // 结束时间
	UserID       uint           `gorm:"not null" json:"user_id"`        // 预约人ID
	DepartmentID uint           `gorm:"not null" json:"department_id"`  // 预约部门ID
	Attendees    string         `gorm:"type:text" json:"attendees"`     // 参会人员，JSON数组
	Purpose      string         `gorm:"size:500" json:"purpose"`        // 会议用途
	Requirements string         `gorm:"type:text" json:"requirements"`  // 会议要求，JSON数组
	Status       int            `gorm:"default:1" json:"status"`        // 1:待审批 2:已通过 3:已驳回 4:已取消
	ApproverID   *uint          `json:"approver_id"`                    // 审批人ID
	ApprovalTime *time.Time     `json:"approval_time"`                  // 审批时间
	CancelReason string         `gorm:"size:500" json:"cancel_reason"`  // 取消原因
	CancelTime   *time.Time     `json:"cancel_time"`                    // 取消时间
	CheckInTime  *time.Time     `json:"check_in_time"`                  // 签到时间
	CheckOutTime *time.Time     `json:"check_out_time"`                 // 签退时间
	Files        string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	Remark       string         `gorm:"size:500" json:"remark"`         // 备注
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// MeetingMinutes 会议纪要
type MeetingMinutes struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	ReservationID uint           `gorm:"not null" json:"reservation_id"` // 会议预约ID
	Content       string         `gorm:"type:text" json:"content"`       // 会议纪要内容
	Participants  string         `gorm:"type:text" json:"participants"`  // 实际参会人员，JSON数组
	Files         string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	CreatedBy     uint           `gorm:"not null" json:"created_by"`     // 创建人ID
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// MeetingRoomMaintenance 会议室维护记录
type MeetingRoomMaintenance struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	RoomID      uint           `gorm:"not null" json:"room_id"`     // 会议室ID
	Type        int            `gorm:"not null" json:"type"`        // 1:日常维护 2:设备维修 3:环境整治
	Description string         `gorm:"size:500" json:"description"` // 维护说明
	StartTime   *time.Time     `json:"start_time"`                  // 开始时间
	EndTime     *time.Time     `json:"end_time"`                    // 结束时间
	Status      int            `gorm:"default:1" json:"status"`     // 1:待处理 2:处理中 3:已完成
	Result      string         `gorm:"size:500" json:"result"`      // 维护结果
	Files       string         `gorm:"type:text" json:"files"`      // 附件，JSON数组
	CreatedBy   uint           `gorm:"not null" json:"created_by"`  // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (MeetingRoom) TableName() string {
	return "meeting_rooms"
}

func (MeetingReservation) TableName() string {
	return "meeting_reservations"
}

func (MeetingMinutes) TableName() string {
	return "meeting_minutes"
}

func (MeetingRoomMaintenance) TableName() string {
	return "meeting_room_maintenances"
}
