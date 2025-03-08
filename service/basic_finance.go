package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

// BasicFinanceService 基础数据-财务模块服务
type BasicFinanceService struct {
	db *gorm.DB
}

func NewBasicFinanceService(db *gorm.DB) *BasicFinanceService {
	return &BasicFinanceService{db: db}
}

// GetExpenseTypeList 获取费用类型列表
func (s *BasicFinanceService) GetExpenseTypeList(parentID *uint) ([]model.ExpenseType, error) {
	var types []model.ExpenseType
	query := s.db.Model(&model.ExpenseType{})
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	}
	err := query.Order("sort asc").Find(&types).Error
	return types, err
}

// GetExpenseTypeByID 根据ID获取费用类型
func (s *BasicFinanceService) GetExpenseTypeByID(id uint) (*model.ExpenseType, error) {
	var expenseType model.ExpenseType
	err := s.db.First(&expenseType, id).Error
	if err != nil {
		return nil, err
	}
	return &expenseType, nil
}

// CreateExpenseType 创建费用类型
func (s *BasicFinanceService) CreateExpenseType(expenseType *model.ExpenseType) error {
	return s.db.Create(expenseType).Error
}

// UpdateExpenseType 更新费用类型
func (s *BasicFinanceService) UpdateExpenseType(expenseType *model.ExpenseType) error {
	if expenseType.ID == 0 {
		return errors.New("expense type id is required")
	}
	return s.db.Model(expenseType).Updates(expenseType).Error
}

// DeleteExpenseType 删除费用类型
func (s *BasicFinanceService) DeleteExpenseType(id uint) error {
	// 检查是否有子类型
	var count int64
	if err := s.db.Model(&model.ExpenseType{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete expense type with sub-types")
	}

	return s.db.Delete(&model.ExpenseType{}, id).Error
}
