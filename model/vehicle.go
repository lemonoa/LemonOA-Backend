package model

import (
	"time"

	"gorm.io/gorm"
)

// Vehicle 车辆信息
type Vehicle struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	PlateNumber  string         `gorm:"size:20;not null;unique" json:"plate_number"` // 车牌号
	Brand        string         `gorm:"size:50;not null" json:"brand"`               // 品牌
	Model        string         `gorm:"size:100;not null" json:"model"`              // 型号
	Color        string         `gorm:"size:20" json:"color"`                        // 颜色
	PurchaseDate *time.Time     `json:"purchase_date"`                               // 购买日期
	Price        float64        `gorm:"type:decimal(10,2)" json:"price"`             // 购买价格
	EngineNumber string         `gorm:"size:50" json:"engine_number"`                // 发动机号
	VIN          string         `gorm:"size:50" json:"vin"`                          // 车架号
	Status       int            `gorm:"default:1" json:"status"`                     // 1:闲置 2:使用中 3:维修中 4:报废
	UserID       *uint          `json:"user_id"`                                     // 使用人ID
	DepartmentID *uint          `json:"department_id"`                               // 使用部门ID
	Files        string         `gorm:"type:text" json:"files"`                      // 附件，JSON数组
	Remark       string         `gorm:"size:500" json:"remark"`                      // 备注
	CreatedBy    uint           `gorm:"not null" json:"created_by"`                  // 创建人ID
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleRepair 车辆维修记录
type VehicleRepair struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	VehicleID   uint           `gorm:"not null" json:"vehicle_id"`     // 车辆ID
	Type        int            `gorm:"not null" json:"type"`           // 1:内部维修 2:外部维修
	Reason      string         `gorm:"size:500" json:"reason"`         // 维修原因
	Description string         `gorm:"size:500" json:"description"`    // 维修说明
	StartDate   *time.Time     `json:"start_date"`                     // 维修开始日期
	EndDate     *time.Time     `json:"end_date"`                       // 维修结束日期
	Cost        float64        `gorm:"type:decimal(10,2)" json:"cost"` // 维修费用
	RepairBy    string         `gorm:"size:100" json:"repair_by"`      // 维修人/维修单位
	Status      int            `gorm:"default:1" json:"status"`        // 1:待维修 2:维修中 3:已完成
	Result      string         `gorm:"size:500" json:"result"`         // 维修结果
	Files       string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	Remark      string         `gorm:"size:500" json:"remark"`         // 备注
	CreatedBy   uint           `gorm:"not null" json:"created_by"`     // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleMaintenance 车辆保养记录
type VehicleMaintenance struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	VehicleID     uint           `gorm:"not null" json:"vehicle_id"`     // 车辆ID
	Type          int            `gorm:"not null" json:"type"`           // 1:定期保养 2:临时保养
	Description   string         `gorm:"size:500" json:"description"`    // 保养说明
	StartDate     *time.Time     `json:"start_date"`                     // 保养开始日期
	EndDate       *time.Time     `json:"end_date"`                       // 保养结束日期
	Cost          float64        `gorm:"type:decimal(10,2)" json:"cost"` // 保养费用
	MaintenanceBy string         `gorm:"size:100" json:"maintenance_by"` // 保养人/保养单位
	Status        int            `gorm:"default:1" json:"status"`        // 1:待保养 2:保养中 3:已完成
	Result        string         `gorm:"size:500" json:"result"`         // 保养结果
	Files         string         `gorm:"type:text" json:"files"`         // 附件，JSON数组
	Remark        string         `gorm:"size:500" json:"remark"`         // 备注
	CreatedBy     uint           `gorm:"not null" json:"created_by"`     // 创建人ID
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleMileage 车辆里程记录
type VehicleMileage struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	VehicleID    uint           `gorm:"not null" json:"vehicle_id"`              // 车辆ID
	Date         *time.Time     `json:"date"`                                    // 记录日期
	StartMileage float64        `gorm:"type:decimal(10,2)" json:"start_mileage"` // 起始里程
	EndMileage   float64        `gorm:"type:decimal(10,2)" json:"end_mileage"`   // 结束里程
	Distance     float64        `gorm:"type:decimal(10,2)" json:"distance"`      // 行驶里程
	Files        string         `gorm:"type:text" json:"files"`                  // 附件，JSON数组
	Remark       string         `gorm:"size:500" json:"remark"`                  // 备注
	CreatedBy    uint           `gorm:"not null" json:"created_by"`              // 创建人ID
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleExpenseRecord 车辆费用记录
type VehicleExpenseRecord struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	VehicleID uint           `gorm:"not null" json:"vehicle_id"`       // 车辆ID
	ExpenseID uint           `gorm:"not null" json:"expense_id"`       // 费用类型ID
	Date      *time.Time     `json:"date"`                             // 费用日期
	Amount    float64        `gorm:"type:decimal(10,2)" json:"amount"` // 费用金额
	Files     string         `gorm:"type:text" json:"files"`           // 附件，JSON数组
	Remark    string         `gorm:"size:500" json:"remark"`           // 备注
	CreatedBy uint           `gorm:"not null" json:"created_by"`       // 创建人ID
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleViolation 车辆违章记录
type VehicleViolation struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	VehicleID   uint           `gorm:"not null" json:"vehicle_id"`       // 车辆ID
	Date        *time.Time     `json:"date"`                             // 违章日期
	Location    string         `gorm:"size:255" json:"location"`         // 违章地点
	Description string         `gorm:"size:500" json:"description"`      // 违章说明
	Points      int            `json:"points"`                           // 扣分
	Amount      float64        `gorm:"type:decimal(10,2)" json:"amount"` // 罚款金额
	Status      int            `gorm:"default:1" json:"status"`          // 1:未处理 2:处理中 3:已处理
	Files       string         `gorm:"type:text" json:"files"`           // 附件，JSON数组
	Remark      string         `gorm:"size:500" json:"remark"`           // 备注
	CreatedBy   uint           `gorm:"not null" json:"created_by"`       // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleAccident 车辆事故记录
type VehicleAccident struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	VehicleID      uint           `gorm:"not null" json:"vehicle_id"`       // 车辆ID
	Date           *time.Time     `json:"date"`                             // 事故日期
	Location       string         `gorm:"size:255" json:"location"`         // 事故地点
	Description    string         `gorm:"size:500" json:"description"`      // 事故说明
	Type           int            `gorm:"not null" json:"type"`             // 1:轻微事故 2:一般事故 3:重大事故
	Responsibility int            `gorm:"default:1" json:"responsibility"`  // 1:全责 2:主责 3:次责 4:无责
	Amount         float64        `gorm:"type:decimal(10,2)" json:"amount"` // 损失金额
	Status         int            `gorm:"default:1" json:"status"`          // 1:未处理 2:处理中 3:已处理
	Files          string         `gorm:"type:text" json:"files"`           // 附件，JSON数组
	Remark         string         `gorm:"size:500" json:"remark"`           // 备注
	CreatedBy      uint           `gorm:"not null" json:"created_by"`       // 创建人ID
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleApplication 用车申请
type VehicleApplication struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	VehicleID        uint           `gorm:"not null" json:"vehicle_id"`    // 车辆ID
	UserID           uint           `gorm:"not null" json:"user_id"`       // 申请人ID
	DepartmentID     uint           `gorm:"not null" json:"department_id"` // 申请部门ID
	StartDate        *time.Time     `json:"start_date"`                    // 用车开始日期
	EndDate          *time.Time     `json:"end_date"`                      // 用车结束日期
	Destination      string         `gorm:"size:255" json:"destination"`   // 目的地
	Purpose          string         `gorm:"size:500" json:"purpose"`       // 用车事由
	Passengers       string         `gorm:"size:500" json:"passengers"`    // 随行人员
	Status           int            `gorm:"default:1" json:"status"`       // 1:待审批 2:已通过 3:已驳回
	ApprovalRecordID *uint          `json:"approval_record_id"`            // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`        // 附件，JSON数组
	Remark           string         `gorm:"size:500" json:"remark"`        // 备注
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy        uint           `json:"created_by"`
}

// VehicleReturn 车辆归还记录
type VehicleReturn struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	ApplicationID uint           `gorm:"not null" json:"application_id"`          // 用车申请ID
	ReturnDate    *time.Time     `json:"return_date"`                             // 归还日期
	StartMileage  float64        `gorm:"type:decimal(10,2)" json:"start_mileage"` // 起始里程
	EndMileage    float64        `gorm:"type:decimal(10,2)" json:"end_mileage"`   // 结束里程
	Distance      float64        `gorm:"type:decimal(10,2)" json:"distance"`      // 行驶里程
	Status        int            `gorm:"default:1" json:"status"`                 // 1:正常 2:异常
	Problem       string         `gorm:"size:500" json:"problem"`                 // 问题说明
	Files         string         `gorm:"type:text" json:"files"`                  // 附件，JSON数组
	Remark        string         `gorm:"size:500" json:"remark"`                  // 备注
	CreatedBy     uint           `gorm:"not null" json:"created_by"`              // 创建人ID
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Vehicle) TableName() string {
	return "vehicles"
}

func (VehicleRepair) TableName() string {
	return "vehicle_repairs"
}

func (VehicleMaintenance) TableName() string {
	return "vehicle_maintenances"
}

func (VehicleMileage) TableName() string {
	return "vehicle_mileages"
}

func (VehicleExpenseRecord) TableName() string {
	return "vehicle_expense_records"
}

func (VehicleViolation) TableName() string {
	return "vehicle_violations"
}

func (VehicleAccident) TableName() string {
	return "vehicle_accidents"
}

func (VehicleApplication) TableName() string {
	return "vehicle_applications"
}

func (VehicleReturn) TableName() string {
	return "vehicle_returns"
}
