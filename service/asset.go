package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type AssetService struct {
	db *gorm.DB
}

func NewAssetService(db *gorm.DB) *AssetService {
	return &AssetService{db: db}
}

// GetAssetList 获取资产列表
func (s *AssetService) GetAssetList(categoryID, brandID uint, status int, keyword string, page, pageSize int) ([]model.Asset, int64, error) {
	var assets []model.Asset
	var total int64

	query := s.db.Model(&model.Asset{})
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	if brandID > 0 {
		query = query.Where("brand_id = ?", brandID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("name LIKE ? OR code LIKE ? OR model LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&assets).Error
	if err != nil {
		return nil, 0, err
	}

	return assets, total, nil
}

// GetAssetByID 根据ID获取资产
func (s *AssetService) GetAssetByID(id uint) (*model.Asset, error) {
	var asset model.Asset
	err := s.db.First(&asset, id).Error
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

// CreateAsset 创建资产
func (s *AssetService) CreateAsset(asset *model.Asset) error {
	// 检查资产分类是否存在
	var category model.AssetCategory
	if err := s.db.First(&category, asset.CategoryID).Error; err != nil {
		return errors.New("asset category not found")
	}

	// 检查品牌是否存在
	var brand model.AssetBrand
	if err := s.db.First(&brand, asset.BrandID).Error; err != nil {
		return errors.New("asset brand not found")
	}

	// 检查单位是否存在
	var unit model.AssetUnit
	if err := s.db.First(&unit, asset.UnitID).Error; err != nil {
		return errors.New("asset unit not found")
	}

	// 检查使用人是否存在
	if asset.UserID != nil {
		var user model.Employee
		if err := s.db.First(&user, asset.UserID).Error; err != nil {
			return errors.New("user not found")
		}
	}

	// 检查使用部门是否存在
	if asset.DepartmentID != nil {
		var department model.Department
		if err := s.db.First(&department, asset.DepartmentID).Error; err != nil {
			return errors.New("department not found")
		}
	}

	return s.db.Create(asset).Error
}

// UpdateAsset 更新资产
func (s *AssetService) UpdateAsset(asset *model.Asset) error {
	if asset.ID == 0 {
		return errors.New("asset id is required")
	}

	// 检查资产分类是否存在
	var category model.AssetCategory
	if err := s.db.First(&category, asset.CategoryID).Error; err != nil {
		return errors.New("asset category not found")
	}

	// 检查品牌是否存在
	var brand model.AssetBrand
	if err := s.db.First(&brand, asset.BrandID).Error; err != nil {
		return errors.New("asset brand not found")
	}

	// 检查单位是否存在
	var unit model.AssetUnit
	if err := s.db.First(&unit, asset.UnitID).Error; err != nil {
		return errors.New("asset unit not found")
	}

	// 检查使用人是否存在
	if asset.UserID != nil {
		var user model.Employee
		if err := s.db.First(&user, asset.UserID).Error; err != nil {
			return errors.New("user not found")
		}
	}

	// 检查使用部门是否存在
	if asset.DepartmentID != nil {
		var department model.Department
		if err := s.db.First(&department, asset.DepartmentID).Error; err != nil {
			return errors.New("department not found")
		}
	}

	return s.db.Model(asset).Updates(asset).Error
}

// DeleteAsset 删除资产
func (s *AssetService) DeleteAsset(id uint) error {
	// 检查是否有维修记录
	var count int64
	if err := s.db.Model(&model.AssetRepair{}).Where("asset_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete asset with repair records")
	}

	return s.db.Delete(&model.Asset{}, id).Error
}

// GetAssetRepairList 获取资产维修记录列表
func (s *AssetService) GetAssetRepairList(assetID uint, status int, page, pageSize int) ([]model.AssetRepair, int64, error) {
	var repairs []model.AssetRepair
	var total int64

	query := s.db.Model(&model.AssetRepair{})
	if assetID > 0 {
		query = query.Where("asset_id = ?", assetID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&repairs).Error
	if err != nil {
		return nil, 0, err
	}

	return repairs, total, nil
}

// GetAssetRepairByID 根据ID获取资产维修记录
func (s *AssetService) GetAssetRepairByID(id uint) (*model.AssetRepair, error) {
	var repair model.AssetRepair
	err := s.db.First(&repair, id).Error
	if err != nil {
		return nil, err
	}
	return &repair, nil
}

// CreateAssetRepair 创建资产维修记录
func (s *AssetService) CreateAssetRepair(repair *model.AssetRepair) error {
	// 检查资产是否存在
	var asset model.Asset
	if err := s.db.First(&asset, repair.AssetID).Error; err != nil {
		return errors.New("asset not found")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建维修记录
		if err := tx.Create(repair).Error; err != nil {
			return err
		}

		// 更新资产状态为维修中
		if err := tx.Model(&asset).Update("status", 3).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateAssetRepair 更新资产维修记录
func (s *AssetService) UpdateAssetRepair(repair *model.AssetRepair) error {
	if repair.ID == 0 {
		return errors.New("repair id is required")
	}

	// 检查资产是否存在
	var asset model.Asset
	if err := s.db.First(&asset, repair.AssetID).Error; err != nil {
		return errors.New("asset not found")
	}

	return s.db.Model(repair).Updates(repair).Error
}

// DeleteAssetRepair 删除资产维修记录
func (s *AssetService) DeleteAssetRepair(id uint) error {
	return s.db.Delete(&model.AssetRepair{}, id).Error
}

// CompleteAssetRepair 完成资产维修
func (s *AssetService) CompleteAssetRepair(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取维修记录
		var repair model.AssetRepair
		if err := tx.First(&repair, id).Error; err != nil {
			return err
		}

		// 更新维修记录状态为已完成
		if err := tx.Model(&repair).Update("status", 3).Error; err != nil {
			return err
		}

		// 更新资产状态为在用
		if err := tx.Model(&model.Asset{}).Where("id = ?", repair.AssetID).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetAssetBorrowList 获取资产领用记录列表
func (s *AssetService) GetAssetBorrowList(assetID, borrowerID uint, status int, page, pageSize int) ([]model.AssetBorrow, int64, error) {
	var borrows []model.AssetBorrow
	var total int64

	query := s.db.Model(&model.AssetBorrow{})
	if assetID > 0 {
		query = query.Where("asset_id = ?", assetID)
	}
	if borrowerID > 0 {
		query = query.Where("borrower_id = ?", borrowerID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&borrows).Error
	if err != nil {
		return nil, 0, err
	}

	return borrows, total, nil
}

// GetAssetBorrowByID 根据ID获取资产领用记录
func (s *AssetService) GetAssetBorrowByID(id uint) (*model.AssetBorrow, error) {
	var borrow model.AssetBorrow
	err := s.db.First(&borrow, id).Error
	if err != nil {
		return nil, err
	}
	return &borrow, nil
}

// CreateAssetBorrow 创建资产领用记录
func (s *AssetService) CreateAssetBorrow(borrow *model.AssetBorrow) error {
	// 检查资产是否存在
	var asset model.Asset
	if err := s.db.First(&asset, borrow.AssetID).Error; err != nil {
		return errors.New("asset not found")
	}

	// 检查资产状态是否为闲置
	if asset.Status != 1 {
		return errors.New("asset is not available")
	}

	// 开启事务
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建领用记录
		if err := tx.Create(borrow).Error; err != nil {
			return err
		}

		// 更新资产状态为在用
		if err := tx.Model(&asset).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateAssetBorrow 更新资产领用记录
func (s *AssetService) UpdateAssetBorrow(borrow *model.AssetBorrow) error {
	if borrow.ID == 0 {
		return errors.New("asset borrow id is required")
	}
	return s.db.Model(borrow).Updates(borrow).Error
}

// DeleteAssetBorrow 删除资产领用记录
func (s *AssetService) DeleteAssetBorrow(id uint) error {
	return s.db.Delete(&model.AssetBorrow{}, id).Error
}

// ReturnAsset 归还资产
func (s *AssetService) ReturnAsset(id uint) error {
	// 获取领用记录
	var borrow model.AssetBorrow
	if err := s.db.First(&borrow, id).Error; err != nil {
		return err
	}

	// 检查资产是否已归还
	if borrow.Status == 2 {
		return errors.New("asset has already been returned")
	}

	// 开启事务
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 更新领用记录状态为已归还
		now := time.Now()
		if err := tx.Model(&borrow).Updates(map[string]interface{}{
			"status":      2,
			"return_date": &now,
		}).Error; err != nil {
			return err
		}

		// 更新资产状态为闲置
		if err := tx.Model(&model.Asset{}).Where("id = ?", borrow.AssetID).Update("status", 1).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetAssetDisposalList 获取资产报废记录列表
func (s *AssetService) GetAssetDisposalList(assetID uint, status int, page, pageSize int) ([]model.AssetDisposal, int64, error) {
	var disposals []model.AssetDisposal
	var total int64

	query := s.db.Model(&model.AssetDisposal{})
	if assetID > 0 {
		query = query.Where("asset_id = ?", assetID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&disposals).Error
	if err != nil {
		return nil, 0, err
	}

	return disposals, total, nil
}

// GetAssetDisposalByID 根据ID获取资产报废记录
func (s *AssetService) GetAssetDisposalByID(id uint) (*model.AssetDisposal, error) {
	var disposal model.AssetDisposal
	err := s.db.First(&disposal, id).Error
	if err != nil {
		return nil, err
	}
	return &disposal, nil
}

// CreateAssetDisposal 创建资产报废记录
func (s *AssetService) CreateAssetDisposal(disposal *model.AssetDisposal) error {
	// 检查资产是否存在
	var asset model.Asset
	if err := s.db.First(&asset, disposal.AssetID).Error; err != nil {
		return errors.New("asset not found")
	}

	// 检查资产状态是否为闲置
	if asset.Status != 1 {
		return errors.New("asset is not available")
	}

	return s.db.Create(disposal).Error
}

// UpdateAssetDisposal 更新资产报废记录
func (s *AssetService) UpdateAssetDisposal(disposal *model.AssetDisposal) error {
	if disposal.ID == 0 {
		return errors.New("asset disposal id is required")
	}
	return s.db.Model(disposal).Updates(disposal).Error
}

// DeleteAssetDisposal 删除资产报废记录
func (s *AssetService) DeleteAssetDisposal(id uint) error {
	return s.db.Delete(&model.AssetDisposal{}, id).Error
}

// ApproveAssetDisposal 审批通过资产报废
func (s *AssetService) ApproveAssetDisposal(id uint) error {
	// 获取报废记录
	var disposal model.AssetDisposal
	if err := s.db.First(&disposal, id).Error; err != nil {
		return err
	}

	// 检查状态是否为待审批
	if disposal.Status != 1 {
		return errors.New("disposal record is not pending approval")
	}

	// 开启事务
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 更新报废记录状态为已通过
		if err := tx.Model(&disposal).Update("status", 2).Error; err != nil {
			return err
		}

		// 更新资产状态为报废
		if err := tx.Model(&model.Asset{}).Where("id = ?", disposal.AssetID).Update("status", 4).Error; err != nil {
			return err
		}

		return nil
	})
}

// RejectAssetDisposal 驳回资产报废
func (s *AssetService) RejectAssetDisposal(id uint) error {
	// 获取报废记录
	var disposal model.AssetDisposal
	if err := s.db.First(&disposal, id).Error; err != nil {
		return err
	}

	// 检查状态是否为待审批
	if disposal.Status != 1 {
		return errors.New("disposal record is not pending approval")
	}

	return s.db.Model(&disposal).Update("status", 3).Error
}
