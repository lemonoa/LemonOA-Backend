package model

import (
	"time"

	"gorm.io/gorm"
)

// AttendanceRule 考勤规则
type AttendanceRule struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	Name           string         `gorm:"size:50;not null" json:"name"`           // 规则名称
	WorkStartTime  string         `gorm:"size:5;not null" json:"work_start_time"` // 上班时间，格式：HH:mm
	WorkEndTime    string         `gorm:"size:5;not null" json:"work_end_time"`   // 下班时间，格式：HH:mm
	LateMinutes    int            `gorm:"default:0" json:"late_minutes"`          // 迟到判定分钟数
	EarlyMinutes   int            `gorm:"default:0" json:"early_minutes"`         // 早退判定分钟数
	RestStartTime  string         `gorm:"size:5" json:"rest_start_time"`          // 休息开始时间，格式：HH:mm
	RestEndTime    string         `gorm:"size:5" json:"rest_end_time"`            // 休息结束时间，格式：HH:mm
	WorkDays       string         `gorm:"size:20;not null" json:"work_days"`      // 工作日，例如：1,2,3,4,5
	EffectiveDate  *time.Time     `json:"effective_date"`                         // 生效日期
	ExpirationDate *time.Time     `json:"expiration_date"`                        // 失效日期
	Status         int            `gorm:"default:1" json:"status"`                // 1:启用 2:禁用
	Description    string         `gorm:"size:500" json:"description"`            // 规则说明
	CreatedBy      uint           `gorm:"not null" json:"created_by"`             // 创建人ID
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// AttendanceRecord 考勤记录
type AttendanceRecord struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	EmployeeID   uint           `gorm:"not null" json:"employee_id"`          // 员工ID
	Date         *time.Time     `gorm:"not null" json:"date"`                 // 考勤日期
	CheckInTime  *time.Time     `json:"check_in_time"`                        // 签到时间
	CheckOutTime *time.Time     `json:"check_out_time"`                       // 签退时间
	Status       int            `gorm:"default:1" json:"status"`              // 1:正常 2:迟到 3:早退 4:旷工 5:请假 6:出差
	LateMinutes  int            `gorm:"default:0" json:"late_minutes"`        // 迟到分钟数
	EarlyMinutes int            `gorm:"default:0" json:"early_minutes"`       // 早退分钟数
	WorkHours    float64        `gorm:"type:decimal(10,2)" json:"work_hours"` // 工作时长
	Location     string         `gorm:"size:255" json:"location"`             // 签到/签退地点
	Remark       string         `gorm:"size:500" json:"remark"`               // 备注
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// LeaveApplication 请假申请
type LeaveApplication struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	EmployeeID       uint           `gorm:"not null" json:"employee_id"`    // 员工ID
	Type             int            `gorm:"not null" json:"type"`           // 1:事假 2:病假 3:婚假 4:产假 5:丧假
	StartTime        *time.Time     `gorm:"not null" json:"start_time"`     // 开始时间
	EndTime          *time.Time     `gorm:"not null" json:"end_time"`       // 结束时间
	Days             float64        `gorm:"type:decimal(10,2)" json:"days"` // 请假天数
	Reason           string         `gorm:"size:500" json:"reason"`         // 请假原因
	Status           int            `gorm:"default:1" json:"status"`        // 1:待审批 2:已通过 3:已驳回 4:已取消
	ApprovalRecordID *uint          `json:"approval_record_id"`             // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	Remark           string         `gorm:"size:500" json:"remark"`         // 备注
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// OvertimeApplication 加班申请
type OvertimeApplication struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	EmployeeID       uint           `gorm:"not null" json:"employee_id"`     // 员工ID
	Type             int            `gorm:"not null" json:"type"`            // 1:工作日加班 2:休息日加班 3:节假日加班
	StartTime        *time.Time     `gorm:"not null" json:"start_time"`      // 开始时间
	EndTime          *time.Time     `gorm:"not null" json:"end_time"`        // 结束时间
	Hours            float64        `gorm:"type:decimal(10,2)" json:"hours"` // 加班小时数
	Reason           string         `gorm:"size:500" json:"reason"`          // 加班原因
	Status           int            `gorm:"default:1" json:"status"`         // 1:待审批 2:已通过 3:已驳回 4:已取消
	ApprovalRecordID *uint          `json:"approval_record_id"`              // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`          // 附件，JSON数组
	Remark           string         `gorm:"size:500" json:"remark"`          // 备注
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// BusinessTripApplication 出差申请
type BusinessTripApplication struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	EmployeeID       uint           `gorm:"not null" json:"employee_id"`    // 员工ID
	Destination      string         `gorm:"size:255" json:"destination"`    // 目的地
	StartTime        *time.Time     `gorm:"not null" json:"start_time"`     // 开始时间
	EndTime          *time.Time     `gorm:"not null" json:"end_time"`       // 结束时间
	Days             float64        `gorm:"type:decimal(10,2)" json:"days"` // 出差天数
	Purpose          string         `gorm:"size:500" json:"purpose"`        // 出差目的
	Status           int            `gorm:"default:1" json:"status"`        // 1:待审批 2:已通过 3:已驳回 4:已取消
	ApprovalRecordID *uint          `json:"approval_record_id"`             // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	Remark           string         `gorm:"size:500" json:"remark"`         // 备注
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (AttendanceRule) TableName() string {
	return "attendance_rules"
}

func (AttendanceRecord) TableName() string {
	return "attendance_records"
}

func (LeaveApplication) TableName() string {
	return "leave_applications"
}

func (OvertimeApplication) TableName() string {
	return "overtime_applications"
}

func (BusinessTripApplication) TableName() string {
	return "business_trip_applications"
}
