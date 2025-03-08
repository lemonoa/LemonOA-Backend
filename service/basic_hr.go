package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

// BasicHRService 基础数据-人事模块服务
type BasicHRService struct {
	db *gorm.DB
}

func NewBasicHRService(db *gorm.DB) *BasicHRService {
	return &BasicHRService{db: db}
}

// GetRewardPunishmentList 获取奖惩项目列表
func (s *BasicHRService) GetRewardPunishmentList(rewardType int) ([]model.RewardPunishment, error) {
	var items []model.RewardPunishment
	query := s.db.Model(&model.RewardPunishment{})
	if rewardType > 0 {
		query = query.Where("type = ?", rewardType)
	}
	err := query.Find(&items).Error
	return items, err
}

// GetRewardPunishmentByID 根据ID获取奖惩项目
func (s *BasicHRService) GetRewardPunishmentByID(id uint) (*model.RewardPunishment, error) {
	var item model.RewardPunishment
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// CreateRewardPunishment 创建奖惩项目
func (s *BasicHRService) CreateRewardPunishment(item *model.RewardPunishment) error {
	return s.db.Create(item).Error
}

// UpdateRewardPunishment 更新奖惩项目
func (s *BasicHRService) UpdateRewardPunishment(item *model.RewardPunishment) error {
	if item.ID == 0 {
		return errors.New("reward punishment id is required")
	}
	return s.db.Model(item).Updates(item).Error
}

// DeleteRewardPunishment 删除奖惩项目
func (s *BasicHRService) DeleteRewardPunishment(id uint) error {
	return s.db.Delete(&model.RewardPunishment{}, id).Error
}

// GetCareProjectList 获取关怀项目列表
func (s *BasicHRService) GetCareProjectList(careType int) ([]model.CareProject, error) {
	var items []model.CareProject
	query := s.db.Model(&model.CareProject{})
	if careType > 0 {
		query = query.Where("type = ?", careType)
	}
	err := query.Find(&items).Error
	return items, err
}

// GetCareProjectByID 根据ID获取关怀项目
func (s *BasicHRService) GetCareProjectByID(id uint) (*model.CareProject, error) {
	var item model.CareProject
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// CreateCareProject 创建关怀项目
func (s *BasicHRService) CreateCareProject(item *model.CareProject) error {
	return s.db.Create(item).Error
}

// UpdateCareProject 更新关怀项目
func (s *BasicHRService) UpdateCareProject(item *model.CareProject) error {
	if item.ID == 0 {
		return errors.New("care project id is required")
	}
	return s.db.Model(item).Updates(item).Error
}

// DeleteCareProject 删除关怀项目
func (s *BasicHRService) DeleteCareProject(id uint) error {
	return s.db.Delete(&model.CareProject{}, id).Error
}

// GetCommonDataList 获取常规数据列表
func (s *BasicHRService) GetCommonDataList(dataType string) ([]model.CommonData, error) {
	var items []model.CommonData
	query := s.db.Model(&model.CommonData{})
	if dataType != "" {
		query = query.Where("type = ?", dataType)
	}
	err := query.Order("sort asc").Find(&items).Error
	return items, err
}

// GetCommonDataByID 根据ID获取常规数据
func (s *BasicHRService) GetCommonDataByID(id uint) (*model.CommonData, error) {
	var item model.CommonData
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetCommonDataByCode 根据Code获取常规数据
func (s *BasicHRService) GetCommonDataByCode(code string) (*model.CommonData, error) {
	var item model.CommonData
	err := s.db.Where("code = ?", code).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// CreateCommonData 创建常规数据
func (s *BasicHRService) CreateCommonData(item *model.CommonData) error {
	return s.db.Create(item).Error
}

// UpdateCommonData 更新常规数据
func (s *BasicHRService) UpdateCommonData(item *model.CommonData) error {
	if item.ID == 0 {
		return errors.New("common data id is required")
	}
	return s.db.Model(item).Updates(item).Error
}

// DeleteCommonData 删除常规数据
func (s *BasicHRService) DeleteCommonData(id uint) error {
	return s.db.Delete(&model.CommonData{}, id).Error
}
