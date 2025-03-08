package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type SealService struct {
	db *gorm.DB
}

func NewSealService(db *gorm.DB) *SealService {
	return &SealService{db: db}
}

// GetSealList 获取印章列表
func (s *SealService) GetSealList(typeID uint, status int, keyword string, page, pageSize int) ([]model.Seal, int64, error) {
	var seals []model.Seal
	var total int64

	query := s.db.Model(&model.Seal{})
	if typeID > 0 {
		query = query.Where("type_id = ?", typeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("name LIKE ? OR code LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&seals).Error
	if err != nil {
		return nil, 0, err
	}

	return seals, total, nil
}

// GetSealByID 根据ID获取印章
func (s *SealService) GetSealByID(id uint) (*model.Seal, error) {
	var seal model.Seal
	err := s.db.First(&seal, id).Error
	if err != nil {
		return nil, err
	}
	return &seal, nil
}

// CreateSeal 创建印章
func (s *SealService) CreateSeal(seal *model.Seal) error {
	// 检查印章类型是否存在
	var sealType model.SealType
	if err := s.db.First(&sealType, seal.TypeID).Error; err != nil {
		return errors.New("seal type not found")
	}

	// 检查保管人是否存在
	var keeper model.Employee
	if err := s.db.First(&keeper, seal.KeeperID).Error; err != nil {
		return errors.New("keeper not found")
	}

	return s.db.Create(seal).Error
}

// UpdateSeal 更新印章
func (s *SealService) UpdateSeal(seal *model.Seal) error {
	if seal.ID == 0 {
		return errors.New("seal id is required")
	}

	// 检查印章类型是否存在
	var sealType model.SealType
	if err := s.db.First(&sealType, seal.TypeID).Error; err != nil {
		return errors.New("seal type not found")
	}

	// 检查保管人是否存在
	var keeper model.Employee
	if err := s.db.First(&keeper, seal.KeeperID).Error; err != nil {
		return errors.New("keeper not found")
	}

	return s.db.Model(seal).Updates(seal).Error
}

// DeleteSeal 删除印章
func (s *SealService) DeleteSeal(id uint) error {
	// 检查是否有用印申请记录
	var count int64
	if err := s.db.Model(&model.SealApplication{}).Where("seal_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete seal with applications")
	}

	return s.db.Delete(&model.Seal{}, id).Error
}

// GetSealApplicationList 获取用印申请列表
func (s *SealService) GetSealApplicationList(sealID, userID uint, status int, page, pageSize int) ([]model.SealApplication, int64, error) {
	var applications []model.SealApplication
	var total int64

	query := s.db.Model(&model.SealApplication{})
	if sealID > 0 {
		query = query.Where("seal_id = ?", sealID)
	}
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applications).Error
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

// GetSealApplicationByID 根据ID获取用印申请
func (s *SealService) GetSealApplicationByID(id uint) (*model.SealApplication, error) {
	var application model.SealApplication
	err := s.db.First(&application, id).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

// CreateSealApplication 创建用印申请
func (s *SealService) CreateSealApplication(application *model.SealApplication) error {
	// 检查印章是否存在
	var seal model.Seal
	if err := s.db.First(&seal, application.SealID).Error; err != nil {
		return errors.New("seal not found")
	}

	// 检查印章是否可用
	if seal.Status != 1 {
		return errors.New("seal is not available")
	}

	// 检查申请人是否存在
	var user model.Employee
	if err := s.db.First(&user, application.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	// 检查申请部门是否存在
	var department model.Department
	if err := s.db.First(&department, application.DepartmentID).Error; err != nil {
		return errors.New("department not found")
	}

	// 检查时间段内是否有其他申请
	var count int64
	err := s.db.Model(&model.SealApplication{}).
		Where("seal_id = ? AND status IN (1,2) AND ((start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))",
			application.SealID,
			application.StartTime, application.EndTime,
			application.StartTime, application.EndTime).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("seal is already applied during this period")
	}

	return s.db.Create(application).Error
}

// UpdateSealApplication 更新用印申请
func (s *SealService) UpdateSealApplication(application *model.SealApplication) error {
	if application.ID == 0 {
		return errors.New("application id is required")
	}

	// 检查印章是否存在
	var seal model.Seal
	if err := s.db.First(&seal, application.SealID).Error; err != nil {
		return errors.New("seal not found")
	}

	// 检查申请人是否存在
	var user model.Employee
	if err := s.db.First(&user, application.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	// 检查申请部门是否存在
	var department model.Department
	if err := s.db.First(&department, application.DepartmentID).Error; err != nil {
		return errors.New("department not found")
	}

	// 检查时间段内是否有其他申请
	var count int64
	err := s.db.Model(&model.SealApplication{}).
		Where("id != ? AND seal_id = ? AND status IN (1,2) AND ((start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))",
			application.ID, application.SealID,
			application.StartTime, application.EndTime,
			application.StartTime, application.EndTime).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("seal is already applied during this period")
	}

	return s.db.Model(application).Updates(application).Error
}

// DeleteSealApplication 删除用印申请
func (s *SealService) DeleteSealApplication(id uint) error {
	// 只能删除待审批的申请
	result := s.db.Where("id = ? AND status = ?", id, 1).Delete(&model.SealApplication{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only delete pending applications")
	}
	return nil
}

// ApproveSealApplication 审批通过用印申请
func (s *SealService) ApproveSealApplication(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取申请记录
		var application model.SealApplication
		if err := tx.First(&application, id).Error; err != nil {
			return err
		}

		// 检查印章是否可用
		var seal model.Seal
		if err := tx.First(&seal, application.SealID).Error; err != nil {
			return err
		}
		if seal.Status != 1 {
			return errors.New("seal is not available")
		}

		// 更新申请状态为已通过
		if err := tx.Model(&application).Update("status", 2).Error; err != nil {
			return err
		}

		// 更新印章状态为借出
		if err := tx.Model(&seal).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// RejectSealApplication 审批驳回用印申请
func (s *SealService) RejectSealApplication(id uint) error {
	result := s.db.Model(&model.SealApplication{}).
		Where("id = ? AND status = ?", id, 1).
		Update("status", 3)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only reject pending applications")
	}
	return nil
}

// CancelSealApplication 取消用印申请
func (s *SealService) CancelSealApplication(id uint) error {
	result := s.db.Model(&model.SealApplication{}).
		Where("id = ? AND status IN (1,2)", id).
		Update("status", 4)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only cancel pending or approved applications")
	}
	return nil
}

// GetSealRecordList 获取用印记录列表
func (s *SealService) GetSealRecordList(applicationID uint, status int, page, pageSize int) ([]model.SealRecord, int64, error) {
	var records []model.SealRecord
	var total int64

	query := s.db.Model(&model.SealRecord{})
	if applicationID > 0 {
		query = query.Where("application_id = ?", applicationID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// GetSealRecordByID 根据ID获取用印记录
func (s *SealService) GetSealRecordByID(id uint) (*model.SealRecord, error) {
	var record model.SealRecord
	err := s.db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// CreateSealRecord 创建用印记录
func (s *SealService) CreateSealRecord(record *model.SealRecord) error {
	// 检查申请是否存在且已通过
	var application model.SealApplication
	if err := s.db.First(&application, record.ApplicationID).Error; err != nil {
		return errors.New("application not found")
	}
	if application.Status != 2 {
		return errors.New("application is not approved")
	}

	now := time.Now()
	record.BorrowTime = &now

	return s.db.Create(record).Error
}

// UpdateSealRecord 更新用印记录
func (s *SealService) UpdateSealRecord(record *model.SealRecord) error {
	if record.ID == 0 {
		return errors.New("record id is required")
	}

	// 检查申请是否存在
	var application model.SealApplication
	if err := s.db.First(&application, record.ApplicationID).Error; err != nil {
		return errors.New("application not found")
	}

	return s.db.Model(record).Updates(record).Error
}

// DeleteSealRecord 删除用印记录
func (s *SealService) DeleteSealRecord(id uint) error {
	return s.db.Delete(&model.SealRecord{}, id).Error
}

// ReturnSeal 归还印章
func (s *SealService) ReturnSeal(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取用印记录
		var record model.SealRecord
		if err := tx.First(&record, id).Error; err != nil {
			return err
		}

		// 获取申请记录
		var application model.SealApplication
		if err := tx.First(&application, record.ApplicationID).Error; err != nil {
			return err
		}

		now := time.Now()

		// 更新用印记录状态为已归还
		if err := tx.Model(&record).Updates(map[string]interface{}{
			"status":      2,
			"return_time": &now,
		}).Error; err != nil {
			return err
		}

		// 更新印章状态为在库
		if err := tx.Model(&model.Seal{}).Where("id = ?", application.SealID).Update("status", 1).Error; err != nil {
			return err
		}

		return nil
	})
}
