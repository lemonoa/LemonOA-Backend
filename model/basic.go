package model

import (
	"time"

	"gorm.io/gorm"
)

// Enterprise 企业主体
type Enterprise struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	Address     string         `gorm:"size:255" json:"address"`
	Phone       string         `gorm:"size:20" json:"phone"`
	Email       string         `gorm:"size:100" json:"email"`
	Website     string         `gorm:"size:255" json:"website"`
	Logo        string         `gorm:"size:255" json:"logo"`
	Description string         `gorm:"size:500" json:"description"`
	Status      int            `gorm:"default:1" json:"status"` // 1:正常 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Region 地区
type Region struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Code      string         `gorm:"size:20;unique" json:"code"`
	ParentID  *uint          `json:"parent_id"`
	Level     int            `gorm:"default:1" json:"level"` // 1:省 2:市 3:区
	Sort      int            `gorm:"default:0" json:"sort"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// MessageTemplate 消息模板
type MessageTemplate struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Code      string         `gorm:"size:50;unique" json:"code"`
	Type      int            `gorm:"not null" json:"type"` // 1:邮件 2:短信
	Content   string         `gorm:"type:text" json:"content"`
	Params    string         `gorm:"type:text" json:"params"` // JSON格式的参数列表
	Status    int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// RewardPunishment 奖惩项目
type RewardPunishment struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Type        int            `gorm:"not null" json:"type"` // 1:奖励 2:惩罚
	Amount      float64        `gorm:"type:decimal(10,2)" json:"amount"`
	Description string         `gorm:"size:255" json:"description"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CareProject 关怀项目
type CareProject struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Type        int            `gorm:"not null" json:"type"` // 1:生日 2:节日 3:其他
	Amount      float64        `gorm:"type:decimal(10,2)" json:"amount"`
	Description string         `gorm:"size:255" json:"description"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CommonData 常规数据
type CommonData struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Code      string         `gorm:"size:50;not null" json:"code"`
	Type      string         `gorm:"size:50;not null" json:"type"` // education:学历 marriage:婚姻状况 等
	Value     string         `gorm:"size:50" json:"value"`
	Sort      int            `gorm:"default:0" json:"sort"`
	Status    int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// AssetCategory 资产分类
type AssetCategory struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	ParentID    *uint          `json:"parent_id"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// AssetBrand 资产品牌
type AssetBrand struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Logo        string         `gorm:"size:255" json:"logo"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// AssetUnit 资产单位
type AssetUnit struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Code      string         `gorm:"size:20;unique" json:"code"`
	Sort      int            `gorm:"default:0" json:"sort"`
	Status    int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// SealType 印章类型
type SealType struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// VehicleExpense 车辆费用
type VehicleExpense struct {
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

// NoticeType 公告类型
type NoticeType struct {
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

// ExpenseType 费用类型
type ExpenseType struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	ParentID    *uint          `json:"parent_id"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CustomerLevel 客户等级
type CustomerLevel struct {
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

// CustomerChannel 客户渠道
type CustomerChannel struct {
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

// Industry 行业类型
type Industry struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	ParentID    *uint          `json:"parent_id"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CustomerStatus 客户状态
type CustomerStatus struct {
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

// CustomerIntention 客户意向
type CustomerIntention struct {
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

// FollowUpMethod 跟进方式
type FollowUpMethod struct {
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

// SalesStage 销售阶段
type SalesStage struct {
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

// ContractCategory 合同分类
type ContractCategory struct {
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

// ProductCategory 产品分类
type ProductCategory struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	ParentID    *uint          `json:"parent_id"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Product 产品
type Product struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Price       float64        `gorm:"type:decimal(10,2)" json:"price"`
	Unit        string         `gorm:"size:20" json:"unit"`
	Description string         `gorm:"size:500" json:"description"`
	Status      int            `gorm:"default:1" json:"status"` // 1:上架 2:下架
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ServiceContent 服务内容
type ServiceContent struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	Description string         `gorm:"size:500" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Supplier 供应商
type Supplier struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	Contact     string         `gorm:"size:50" json:"contact"`
	Phone       string         `gorm:"size:20" json:"phone"`
	Email       string         `gorm:"size:100" json:"email"`
	Address     string         `gorm:"size:255" json:"address"`
	Description string         `gorm:"size:500" json:"description"`
	Status      int            `gorm:"default:1" json:"status"` // 1:正常 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// PurchaseCategory 采购品分类
type PurchaseCategory struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	ParentID    *uint          `json:"parent_id"`
	Description string         `gorm:"size:255" json:"description"`
	Sort        int            `gorm:"default:0" json:"sort"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// PurchaseItem 采购品
type PurchaseItem struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	Code        string         `gorm:"size:50;unique" json:"code"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Price       float64        `gorm:"type:decimal(10,2)" json:"price"`
	Unit        string         `gorm:"size:20" json:"unit"`
	Description string         `gorm:"size:500" json:"description"`
	Status      int            `gorm:"default:1" json:"status"` // 1:启用 2:禁用
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ProjectStage 项目阶段
type ProjectStage struct {
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

// ProjectCategory 项目分类
type ProjectCategory struct {
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

// WorkType 工作类型
type WorkType struct {
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

// TableName 指定表名
func (Enterprise) TableName() string {
	return "enterprises"
}

func (Region) TableName() string {
	return "regions"
}

func (MessageTemplate) TableName() string {
	return "message_templates"
}

func (RewardPunishment) TableName() string {
	return "reward_punishments"
}

func (CareProject) TableName() string {
	return "care_projects"
}

func (CommonData) TableName() string {
	return "common_data"
}

func (AssetCategory) TableName() string {
	return "asset_categories"
}

func (AssetBrand) TableName() string {
	return "asset_brands"
}

func (AssetUnit) TableName() string {
	return "asset_units"
}

func (SealType) TableName() string {
	return "seal_types"
}

func (VehicleExpense) TableName() string {
	return "vehicle_expenses"
}

func (NoticeType) TableName() string {
	return "notice_types"
}

func (ExpenseType) TableName() string {
	return "expense_types"
}

func (CustomerLevel) TableName() string {
	return "customer_levels"
}

func (CustomerChannel) TableName() string {
	return "customer_channels"
}

func (Industry) TableName() string {
	return "industries"
}

func (CustomerStatus) TableName() string {
	return "customer_statuses"
}

func (CustomerIntention) TableName() string {
	return "customer_intentions"
}

func (FollowUpMethod) TableName() string {
	return "follow_up_methods"
}

func (SalesStage) TableName() string {
	return "sales_stages"
}

func (ContractCategory) TableName() string {
	return "contract_categories"
}

func (ProductCategory) TableName() string {
	return "product_categories"
}

func (Product) TableName() string {
	return "products"
}

func (ServiceContent) TableName() string {
	return "service_contents"
}

func (Supplier) TableName() string {
	return "suppliers"
}

func (PurchaseCategory) TableName() string {
	return "purchase_categories"
}

func (PurchaseItem) TableName() string {
	return "purchase_items"
}

func (ProjectStage) TableName() string {
	return "project_stages"
}

func (ProjectCategory) TableName() string {
	return "project_categories"
}

func (WorkType) TableName() string {
	return "work_types"
}
