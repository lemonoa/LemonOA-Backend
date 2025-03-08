package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

// BasicCustomerService 基础数据-客户模块服务
type BasicCustomerService struct {
	db *gorm.DB
}

func NewBasicCustomerService(db *gorm.DB) *BasicCustomerService {
	return &BasicCustomerService{db: db}
}

// GetCustomerLevelList 获取客户等级列表
func (s *BasicCustomerService) GetCustomerLevelList() ([]model.CustomerLevel, error) {
	var levels []model.CustomerLevel
	err := s.db.Order("sort asc").Find(&levels).Error
	return levels, err
}

// GetCustomerLevelByID 根据ID获取客户等级
func (s *BasicCustomerService) GetCustomerLevelByID(id uint) (*model.CustomerLevel, error) {
	var level model.CustomerLevel
	err := s.db.First(&level, id).Error
	if err != nil {
		return nil, err
	}
	return &level, nil
}

// CreateCustomerLevel 创建客户等级
func (s *BasicCustomerService) CreateCustomerLevel(level *model.CustomerLevel) error {
	return s.db.Create(level).Error
}

// UpdateCustomerLevel 更新客户等级
func (s *BasicCustomerService) UpdateCustomerLevel(level *model.CustomerLevel) error {
	if level.ID == 0 {
		return errors.New("customer level id is required")
	}
	return s.db.Model(level).Updates(level).Error
}

// DeleteCustomerLevel 删除客户等级
func (s *BasicCustomerService) DeleteCustomerLevel(id uint) error {
	return s.db.Delete(&model.CustomerLevel{}, id).Error
}

// GetCustomerChannelList 获取客户渠道列表
func (s *BasicCustomerService) GetCustomerChannelList() ([]model.CustomerChannel, error) {
	var channels []model.CustomerChannel
	err := s.db.Order("sort asc").Find(&channels).Error
	return channels, err
}

// GetCustomerChannelByID 根据ID获取客户渠道
func (s *BasicCustomerService) GetCustomerChannelByID(id uint) (*model.CustomerChannel, error) {
	var channel model.CustomerChannel
	err := s.db.First(&channel, id).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

// CreateCustomerChannel 创建客户渠道
func (s *BasicCustomerService) CreateCustomerChannel(channel *model.CustomerChannel) error {
	return s.db.Create(channel).Error
}

// UpdateCustomerChannel 更新客户渠道
func (s *BasicCustomerService) UpdateCustomerChannel(channel *model.CustomerChannel) error {
	if channel.ID == 0 {
		return errors.New("customer channel id is required")
	}
	return s.db.Model(channel).Updates(channel).Error
}

// DeleteCustomerChannel 删除客户渠道
func (s *BasicCustomerService) DeleteCustomerChannel(id uint) error {
	return s.db.Delete(&model.CustomerChannel{}, id).Error
}

// GetIndustryList 获取行业类型列表
func (s *BasicCustomerService) GetIndustryList(parentID *uint) ([]model.Industry, error) {
	var industries []model.Industry
	query := s.db.Model(&model.Industry{})
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	}
	err := query.Order("sort asc").Find(&industries).Error
	return industries, err
}

// GetIndustryByID 根据ID获取行业类型
func (s *BasicCustomerService) GetIndustryByID(id uint) (*model.Industry, error) {
	var industry model.Industry
	err := s.db.First(&industry, id).Error
	if err != nil {
		return nil, err
	}
	return &industry, nil
}

// CreateIndustry 创建行业类型
func (s *BasicCustomerService) CreateIndustry(industry *model.Industry) error {
	return s.db.Create(industry).Error
}

// UpdateIndustry 更新行业类型
func (s *BasicCustomerService) UpdateIndustry(industry *model.Industry) error {
	if industry.ID == 0 {
		return errors.New("industry id is required")
	}
	return s.db.Model(industry).Updates(industry).Error
}

// DeleteIndustry 删除行业类型
func (s *BasicCustomerService) DeleteIndustry(id uint) error {
	// 检查是否有子行业
	var count int64
	if err := s.db.Model(&model.Industry{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete industry with sub-industries")
	}

	return s.db.Delete(&model.Industry{}, id).Error
}

// GetCustomerStatusList 获取客户状态列表
func (s *BasicCustomerService) GetCustomerStatusList() ([]model.CustomerStatus, error) {
	var statuses []model.CustomerStatus
	err := s.db.Order("sort asc").Find(&statuses).Error
	return statuses, err
}

// GetCustomerStatusByID 根据ID获取客户状态
func (s *BasicCustomerService) GetCustomerStatusByID(id uint) (*model.CustomerStatus, error) {
	var status model.CustomerStatus
	err := s.db.First(&status, id).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

// CreateCustomerStatus 创建客户状态
func (s *BasicCustomerService) CreateCustomerStatus(status *model.CustomerStatus) error {
	return s.db.Create(status).Error
}

// UpdateCustomerStatus 更新客户状态
func (s *BasicCustomerService) UpdateCustomerStatus(status *model.CustomerStatus) error {
	if status.ID == 0 {
		return errors.New("customer status id is required")
	}
	return s.db.Model(status).Updates(status).Error
}

// DeleteCustomerStatus 删除客户状态
func (s *BasicCustomerService) DeleteCustomerStatus(id uint) error {
	return s.db.Delete(&model.CustomerStatus{}, id).Error
}

// GetCustomerIntentionList 获取客户意向列表
func (s *BasicCustomerService) GetCustomerIntentionList() ([]model.CustomerIntention, error) {
	var intentions []model.CustomerIntention
	err := s.db.Order("sort asc").Find(&intentions).Error
	return intentions, err
}

// GetCustomerIntentionByID 根据ID获取客户意向
func (s *BasicCustomerService) GetCustomerIntentionByID(id uint) (*model.CustomerIntention, error) {
	var intention model.CustomerIntention
	err := s.db.First(&intention, id).Error
	if err != nil {
		return nil, err
	}
	return &intention, nil
}

// CreateCustomerIntention 创建客户意向
func (s *BasicCustomerService) CreateCustomerIntention(intention *model.CustomerIntention) error {
	return s.db.Create(intention).Error
}

// UpdateCustomerIntention 更新客户意向
func (s *BasicCustomerService) UpdateCustomerIntention(intention *model.CustomerIntention) error {
	if intention.ID == 0 {
		return errors.New("customer intention id is required")
	}
	return s.db.Model(intention).Updates(intention).Error
}

// DeleteCustomerIntention 删除客户意向
func (s *BasicCustomerService) DeleteCustomerIntention(id uint) error {
	return s.db.Delete(&model.CustomerIntention{}, id).Error
}

// GetFollowUpMethodList 获取跟进方式列表
func (s *BasicCustomerService) GetFollowUpMethodList() ([]model.FollowUpMethod, error) {
	var methods []model.FollowUpMethod
	err := s.db.Order("sort asc").Find(&methods).Error
	return methods, err
}

// GetFollowUpMethodByID 根据ID获取跟进方式
func (s *BasicCustomerService) GetFollowUpMethodByID(id uint) (*model.FollowUpMethod, error) {
	var method model.FollowUpMethod
	err := s.db.First(&method, id).Error
	if err != nil {
		return nil, err
	}
	return &method, nil
}

// CreateFollowUpMethod 创建跟进方式
func (s *BasicCustomerService) CreateFollowUpMethod(method *model.FollowUpMethod) error {
	return s.db.Create(method).Error
}

// UpdateFollowUpMethod 更新跟进方式
func (s *BasicCustomerService) UpdateFollowUpMethod(method *model.FollowUpMethod) error {
	if method.ID == 0 {
		return errors.New("follow up method id is required")
	}
	return s.db.Model(method).Updates(method).Error
}

// DeleteFollowUpMethod 删除跟进方式
func (s *BasicCustomerService) DeleteFollowUpMethod(id uint) error {
	return s.db.Delete(&model.FollowUpMethod{}, id).Error
}

// GetSalesStageList 获取销售阶段列表
func (s *BasicCustomerService) GetSalesStageList() ([]model.SalesStage, error) {
	var stages []model.SalesStage
	err := s.db.Order("sort asc").Find(&stages).Error
	return stages, err
}

// GetSalesStageByID 根据ID获取销售阶段
func (s *BasicCustomerService) GetSalesStageByID(id uint) (*model.SalesStage, error) {
	var stage model.SalesStage
	err := s.db.First(&stage, id).Error
	if err != nil {
		return nil, err
	}
	return &stage, nil
}

// CreateSalesStage 创建销售阶段
func (s *BasicCustomerService) CreateSalesStage(stage *model.SalesStage) error {
	return s.db.Create(stage).Error
}

// UpdateSalesStage 更新销售阶段
func (s *BasicCustomerService) UpdateSalesStage(stage *model.SalesStage) error {
	if stage.ID == 0 {
		return errors.New("sales stage id is required")
	}
	return s.db.Model(stage).Updates(stage).Error
}

// DeleteSalesStage 删除销售阶段
func (s *BasicCustomerService) DeleteSalesStage(id uint) error {
	return s.db.Delete(&model.SalesStage{}, id).Error
}
