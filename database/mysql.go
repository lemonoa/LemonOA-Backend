package database

import (
	"fmt"

	"github.com/lemonoa/LemonOA-Go/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitMySQL 初始化MySQL连接
func InitMySQL() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.charset"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))

	// 自动迁移数据库表
	err = db.AutoMigrate(
		// 认证管理
		&model.User{},
		&model.UserRole{},
		&model.Permission{},
		&model.RolePermission{},
		&model.LoginLog{},

		// 系统管理
		&model.SystemConfig{},
		&model.Module{},
		&model.ModuleConfig{},
		&model.FunctionNode{},
		&model.Role{},
		&model.RoleFunction{},
		&model.OperationLog{},
		&model.Attachment{},
		&model.BackupRecord{},
		&model.ScheduledTask{},

		// 工作台
		&model.Department{},
		&model.Employee{},
		&model.Notification{},
		&model.ApprovalType{},
		&model.ApprovalFlow{},
		&model.ApprovalNode{},
		&model.ApprovalRecord{},
		&model.ApprovalNodeRecord{},
		&model.Todo{},

		// 基础数据-公共模块
		&model.Enterprise{},
		&model.Region{},
		&model.MessageTemplate{},

		// 基础数据-人事模块
		&model.RewardPunishment{},
		&model.CareProject{},
		&model.CommonData{},

		// 基础数据-行政模块
		&model.AssetCategory{},
		&model.AssetBrand{},
		&model.AssetUnit{},
		&model.SealType{},
		&model.VehicleExpense{},
		&model.NoticeType{},

		// 基础数据-财务模块
		&model.ExpenseType{},

		// 基础数据-客户模块
		&model.CustomerLevel{},
		&model.CustomerChannel{},
		&model.Industry{},
		&model.CustomerStatus{},
		&model.CustomerIntention{},
		&model.FollowUpMethod{},
		&model.SalesStage{},

		// 基础数据-合同模块
		&model.ContractCategory{},
		&model.ProductCategory{},
		&model.Product{},
		&model.ServiceContent{},
		&model.Supplier{},
		&model.PurchaseCategory{},
		&model.PurchaseItem{},

		// 基础数据-项目模块
		&model.ProjectStage{},
		&model.ProjectCategory{},
		&model.WorkType{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto migrate tables: %v", err)
	}

	DB = db
	return nil
}
