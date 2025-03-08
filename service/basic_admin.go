package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

// BasicAdminService 基础数据-行政模块服务
type BasicAdminService struct {
	db *gorm.DB
}

func NewBasicAdminService(db *gorm.DB) *BasicAdminService {
	return &BasicAdminService{db: db}
}

// GetAssetCategoryList 获取资产分类列表
func (s *BasicAdminService) GetAssetCategoryList(parentID *uint) ([]model.AssetCategory, error) {
	var categories []model.AssetCategory
	query := s.db.Model(&model.AssetCategory{})
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	}
	err := query.Order("sort asc").Find(&categories).Error
	return categories, err
}

// GetAssetCategoryByID 根据ID获取资产分类
func (s *BasicAdminService) GetAssetCategoryByID(id uint) (*model.AssetCategory, error) {
	var category model.AssetCategory
	err := s.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateAssetCategory 创建资产分类
func (s *BasicAdminService) CreateAssetCategory(category *model.AssetCategory) error {
	return s.db.Create(category).Error
}

// UpdateAssetCategory 更新资产分类
func (s *BasicAdminService) UpdateAssetCategory(category *model.AssetCategory) error {
	if category.ID == 0 {
		return errors.New("asset category id is required")
	}
	return s.db.Model(category).Updates(category).Error
}

// DeleteAssetCategory 删除资产分类
func (s *BasicAdminService) DeleteAssetCategory(id uint) error {
	// 检查是否有子分类
	var count int64
	if err := s.db.Model(&model.AssetCategory{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete category with sub-categories")
	}

	return s.db.Delete(&model.AssetCategory{}, id).Error
}

// GetAssetBrandList 获取资产品牌列表
func (s *BasicAdminService) GetAssetBrandList() ([]model.AssetBrand, error) {
	var brands []model.AssetBrand
	err := s.db.Order("sort asc").Find(&brands).Error
	return brands, err
}

// GetAssetBrandByID 根据ID获取资产品牌
func (s *BasicAdminService) GetAssetBrandByID(id uint) (*model.AssetBrand, error) {
	var brand model.AssetBrand
	err := s.db.First(&brand, id).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

// CreateAssetBrand 创建资产品牌
func (s *BasicAdminService) CreateAssetBrand(brand *model.AssetBrand) error {
	return s.db.Create(brand).Error
}

// UpdateAssetBrand 更新资产品牌
func (s *BasicAdminService) UpdateAssetBrand(brand *model.AssetBrand) error {
	if brand.ID == 0 {
		return errors.New("asset brand id is required")
	}
	return s.db.Model(brand).Updates(brand).Error
}

// DeleteAssetBrand 删除资产品牌
func (s *BasicAdminService) DeleteAssetBrand(id uint) error {
	return s.db.Delete(&model.AssetBrand{}, id).Error
}

// GetAssetUnitList 获取资产单位列表
func (s *BasicAdminService) GetAssetUnitList() ([]model.AssetUnit, error) {
	var units []model.AssetUnit
	err := s.db.Order("sort asc").Find(&units).Error
	return units, err
}

// GetAssetUnitByID 根据ID获取资产单位
func (s *BasicAdminService) GetAssetUnitByID(id uint) (*model.AssetUnit, error) {
	var unit model.AssetUnit
	err := s.db.First(&unit, id).Error
	if err != nil {
		return nil, err
	}
	return &unit, nil
}

// CreateAssetUnit 创建资产单位
func (s *BasicAdminService) CreateAssetUnit(unit *model.AssetUnit) error {
	return s.db.Create(unit).Error
}

// UpdateAssetUnit 更新资产单位
func (s *BasicAdminService) UpdateAssetUnit(unit *model.AssetUnit) error {
	if unit.ID == 0 {
		return errors.New("asset unit id is required")
	}
	return s.db.Model(unit).Updates(unit).Error
}

// DeleteAssetUnit 删除资产单位
func (s *BasicAdminService) DeleteAssetUnit(id uint) error {
	return s.db.Delete(&model.AssetUnit{}, id).Error
}

// GetSealTypeList 获取印章类型列表
func (s *BasicAdminService) GetSealTypeList() ([]model.SealType, error) {
	var types []model.SealType
	err := s.db.Order("sort asc").Find(&types).Error
	return types, err
}

// GetSealTypeByID 根据ID获取印章类型
func (s *BasicAdminService) GetSealTypeByID(id uint) (*model.SealType, error) {
	var sealType model.SealType
	err := s.db.First(&sealType, id).Error
	if err != nil {
		return nil, err
	}
	return &sealType, nil
}

// CreateSealType 创建印章类型
func (s *BasicAdminService) CreateSealType(sealType *model.SealType) error {
	return s.db.Create(sealType).Error
}

// UpdateSealType 更新印章类型
func (s *BasicAdminService) UpdateSealType(sealType *model.SealType) error {
	if sealType.ID == 0 {
		return errors.New("seal type id is required")
	}
	return s.db.Model(sealType).Updates(sealType).Error
}

// DeleteSealType 删除印章类型
func (s *BasicAdminService) DeleteSealType(id uint) error {
	return s.db.Delete(&model.SealType{}, id).Error
}

// GetVehicleExpenseList 获取车辆费用列表
func (s *BasicAdminService) GetVehicleExpenseList() ([]model.VehicleExpense, error) {
	var expenses []model.VehicleExpense
	err := s.db.Order("sort asc").Find(&expenses).Error
	return expenses, err
}

// GetVehicleExpenseByID 根据ID获取车辆费用
func (s *BasicAdminService) GetVehicleExpenseByID(id uint) (*model.VehicleExpense, error) {
	var expense model.VehicleExpense
	err := s.db.First(&expense, id).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

// CreateVehicleExpense 创建车辆费用
func (s *BasicAdminService) CreateVehicleExpense(expense *model.VehicleExpense) error {
	return s.db.Create(expense).Error
}

// UpdateVehicleExpense 更新车辆费用
func (s *BasicAdminService) UpdateVehicleExpense(expense *model.VehicleExpense) error {
	if expense.ID == 0 {
		return errors.New("vehicle expense id is required")
	}
	return s.db.Model(expense).Updates(expense).Error
}

// DeleteVehicleExpense 删除车辆费用
func (s *BasicAdminService) DeleteVehicleExpense(id uint) error {
	return s.db.Delete(&model.VehicleExpense{}, id).Error
}

// GetNoticeTypeList 获取公告类型列表
func (s *BasicAdminService) GetNoticeTypeList() ([]model.NoticeType, error) {
	var types []model.NoticeType
	err := s.db.Order("sort asc").Find(&types).Error
	return types, err
}

// GetNoticeTypeByID 根据ID获取公告类型
func (s *BasicAdminService) GetNoticeTypeByID(id uint) (*model.NoticeType, error) {
	var noticeType model.NoticeType
	err := s.db.First(&noticeType, id).Error
	if err != nil {
		return nil, err
	}
	return &noticeType, nil
}

// CreateNoticeType 创建公告类型
func (s *BasicAdminService) CreateNoticeType(noticeType *model.NoticeType) error {
	return s.db.Create(noticeType).Error
}

// UpdateNoticeType 更新公告类型
func (s *BasicAdminService) UpdateNoticeType(noticeType *model.NoticeType) error {
	if noticeType.ID == 0 {
		return errors.New("notice type id is required")
	}
	return s.db.Model(noticeType).Updates(noticeType).Error
}

// DeleteNoticeType 删除公告类型
func (s *BasicAdminService) DeleteNoticeType(id uint) error {
	return s.db.Delete(&model.NoticeType{}, id).Error
}
