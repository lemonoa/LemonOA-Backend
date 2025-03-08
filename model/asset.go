package model

import (
	"time"

	"gorm.io/gorm"
)

// Asset 固定资产
type Asset struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	Name         string         `gorm:"size:100;not null" json:"name"`   // 资产名称
	Code         string         `gorm:"size:50;unique" json:"code"`      // 资产编号
	CategoryID   uint           `gorm:"not null" json:"category_id"`     // 资产分类ID
	BrandID      uint           `gorm:"not null" json:"brand_id"`        // 品牌ID
	Model        string         `gorm:"size:100" json:"model"`           // 规格型号
	UnitID       uint           `gorm:"not null" json:"unit_id"`         // 单位ID
	Price        float64        `gorm:"type:decimal(10,2)" json:"price"` // 采购价格
	PurchaseDate *time.Time     `json:"purchase_date"`                   // 购买日期
	WarrantyDate *time.Time     `json:"warranty_date"`                   // 保修期限
	Status       int            `gorm:"default:1" json:"status"`         // 1:闲置 2:在用 3:维修中 4:报废
	UserID       *uint          `json:"user_id"`                         // 使用人ID
	DepartmentID *uint          `json:"department_id"`                   // 使用部门ID
	Location     string         `gorm:"size:255" json:"location"`        // 存放位置
	Description  string         `gorm:"size:500" json:"description"`     // 资产描述
	Files        string         `gorm:"type:text" json:"files"`          // 附件，JSON数组
	Remark       string         `gorm:"size:500" json:"remark"`          // 备注
	CreatedBy    uint           `gorm:"not null" json:"created_by"`      // 创建人ID
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// AssetRepair 资产维修记录
type AssetRepair struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	AssetID     uint           `gorm:"not null" json:"asset_id"`       // 资产ID
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

// AssetBorrow 资产领用记录
type AssetBorrow struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	AssetID      uint           `gorm:"not null" json:"asset_id"`      // 资产ID
	BorrowerID   uint           `gorm:"not null" json:"borrower_id"`   // 借用人ID
	DepartmentID uint           `gorm:"not null" json:"department_id"` // 借用部门ID
	Purpose      string         `gorm:"size:500" json:"purpose"`       // 借用用途
	BorrowDate   *time.Time     `json:"borrow_date"`                   // 借用日期
	ReturnDate   *time.Time     `json:"return_date"`                   // 归还日期
	Status       int            `gorm:"default:1" json:"status"`       // 1:已借出 2:已归还 3:已逾期
	Files        string         `gorm:"type:text" json:"files"`        // 附件，JSON数组
	Remark       string         `gorm:"size:500" json:"remark"`        // 备注
	CreatedBy    uint           `gorm:"not null" json:"created_by"`    // 创建人ID
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// AssetDisposal 资产报废记录
type AssetDisposal struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	AssetID          uint           `gorm:"not null" json:"asset_id"`         // 资产ID
	Reason           string         `gorm:"size:500" json:"reason"`           // 报废原因
	Method           int            `gorm:"not null" json:"method"`           // 1:销毁 2:捐赠 3:转卖
	Amount           float64        `gorm:"type:decimal(10,2)" json:"amount"` // 处置金额
	DisposalDate     *time.Time     `json:"disposal_date"`                    // 报废日期
	Status           int            `gorm:"default:1" json:"status"`          // 1:待审批 2:已通过 3:已驳回 4:已取消
	ApprovalRecordID *uint          `json:"approval_record_id"`               // 关联的审批记录ID
	Files            string         `gorm:"type:text" json:"files"`           // 附件，JSON数组
	Remark           string         `gorm:"size:500" json:"remark"`           // 备注
	CreatedBy        uint           `gorm:"not null" json:"created_by"`       // 创建人ID
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (Asset) TableName() string {
	return "assets"
}

func (AssetRepair) TableName() string {
	return "asset_repairs"
}

func (AssetBorrow) TableName() string {
	return "asset_borrows"
}

func (AssetDisposal) TableName() string {
	return "asset_disposals"
}
