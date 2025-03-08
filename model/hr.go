package model

import (
	"time"

	"gorm.io/gorm"
)

// Position 岗位职称
type Position struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// EmployeeArchive 员工档案
type EmployeeArchive struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	EmployeeID    uint           `gorm:"not null" json:"employee_id"`
	Education     string         `gorm:"size:50" json:"education"`      // 学历
	School        string         `gorm:"size:100" json:"school"`        // 毕业院校
	Major         string         `gorm:"size:100" json:"major"`         // 专业
	GraduationAt  *time.Time     `json:"graduation_at"`                 // 毕业时间
	WorkStartAt   *time.Time     `json:"work_start_at"`                 // 参加工作时间
	MaritalStatus string         `gorm:"size:20" json:"marital_status"` // 婚姻状况
	Political     string         `gorm:"size:50" json:"political"`      // 政治面貌
	IDCard        string         `gorm:"size:20" json:"id_card"`        // 身份证号
	Birthday      *time.Time     `json:"birthday"`                      // 出生日期
	Native        string         `gorm:"size:100" json:"native"`        // 籍贯
	Address       string         `gorm:"size:255" json:"address"`       // 现居地址
	Files         string         `gorm:"type:text" json:"files"`        // 档案附件，JSON数组
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// RewardPunishmentRecord 奖惩记录
type RewardPunishmentRecord struct {
	ID                 uint           `gorm:"primarykey" json:"id"`
	EmployeeID         uint           `gorm:"not null" json:"employee_id"`
	RewardPunishmentID uint           `gorm:"not null" json:"reward_punishment_id"`
	Amount             float64        `gorm:"type:decimal(10,2)" json:"amount"`
	Reason             string         `gorm:"size:500" json:"reason"`
	Date               *time.Time     `json:"date"`
	Remark             string         `gorm:"size:500" json:"remark"`
	Files              string         `gorm:"type:text" json:"files"` // 附件，JSON数组
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CareRecord 关怀记录
type CareRecord struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	EmployeeID    uint           `gorm:"not null" json:"employee_id"`
	CareProjectID uint           `gorm:"not null" json:"care_project_id"`
	Amount        float64        `gorm:"type:decimal(10,2)" json:"amount"`
	Date          *time.Time     `json:"date"`
	Remark        string         `gorm:"size:500" json:"remark"`
	Files         string         `gorm:"type:text" json:"files"` // 附件，JSON数组
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Transfer 人事调动
type Transfer struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	EmployeeID       uint           `gorm:"not null" json:"employee_id"`
	OldDepartmentID  uint           `gorm:"not null" json:"old_department_id"`
	NewDepartmentID  uint           `gorm:"not null" json:"new_department_id"`
	OldPositionID    uint           `gorm:"not null" json:"old_position_id"`
	NewPositionID    uint           `gorm:"not null" json:"new_position_id"`
	EffectiveDate    *time.Time     `json:"effective_date"`
	Reason           string         `gorm:"size:500" json:"reason"`
	Status           int            `gorm:"default:1" json:"status"` // 1:待审批 2:已通过 3:已驳回
	ApprovalRecordID *uint          `json:"approval_record_id"`      // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`  // 附件，JSON数组
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Resignation 离职档案
type Resignation struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	EmployeeID       uint           `gorm:"not null" json:"employee_id"`
	ResignationType  int            `gorm:"not null" json:"resignation_type"` // 1:主动离职 2:被动离职
	Reason           string         `gorm:"size:500" json:"reason"`
	LastWorkingDay   *time.Time     `json:"last_working_day"`
	HandoverTo       uint           `json:"handover_to"`                       // 工作交接人
	HandoverContent  string         `gorm:"type:text" json:"handover_content"` // 工作交接内容
	Status           int            `gorm:"default:1" json:"status"`           // 1:待审批 2:已通过 3:已驳回
	ApprovalRecordID *uint          `json:"approval_record_id"`                // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`            // 附件，JSON数组
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Contract 员工合同
type Contract struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	EmployeeID uint           `gorm:"not null" json:"employee_id"`
	ContractNo string         `gorm:"size:50;unique" json:"contract_no"` // 合同编号
	Type       int            `gorm:"not null" json:"type"`              // 1:固定期限 2:无固定期限 3:实习
	StartDate  *time.Time     `json:"start_date"`                        // 合同开始日期
	EndDate    *time.Time     `json:"end_date"`                          // 合同结束日期
	Status     int            `gorm:"default:1" json:"status"`           // 1:生效中 2:已终止 3:已到期
	Files      string         `gorm:"type:text" json:"files"`            // 附件，JSON数组
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Probation 转正
type Probation struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	EmployeeID       uint           `gorm:"not null" json:"employee_id"`
	ProbationEndDate *time.Time     `json:"probation_end_date"`          // 试用期结束日期
	Assessment       string         `gorm:"type:text" json:"assessment"` // 试用期评估
	Status           int            `gorm:"default:1" json:"status"`     // 1:待审批 2:已通过 3:已驳回
	ApprovalRecordID *uint          `json:"approval_record_id"`          // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`      // 附件，JSON数组
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Position) TableName() string {
	return "positions"
}

func (EmployeeArchive) TableName() string {
	return "employee_archives"
}

func (RewardPunishmentRecord) TableName() string {
	return "reward_punishment_records"
}

func (CareRecord) TableName() string {
	return "care_records"
}

func (Transfer) TableName() string {
	return "transfers"
}

func (Resignation) TableName() string {
	return "resignations"
}

func (Contract) TableName() string {
	return "contracts"
}

func (Probation) TableName() string {
	return "probations"
}
