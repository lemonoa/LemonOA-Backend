package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

// BasicCommonService 基础数据-公共模块服务
type BasicCommonService struct {
	db *gorm.DB
}

func NewBasicCommonService(db *gorm.DB) *BasicCommonService {
	return &BasicCommonService{db: db}
}

// GetEnterpriseList 获取企业主体列表
func (s *BasicCommonService) GetEnterpriseList() ([]model.Enterprise, error) {
	var enterprises []model.Enterprise
	err := s.db.Find(&enterprises).Error
	return enterprises, err
}

// GetEnterpriseByID 根据ID获取企业主体
func (s *BasicCommonService) GetEnterpriseByID(id uint) (*model.Enterprise, error) {
	var enterprise model.Enterprise
	err := s.db.First(&enterprise, id).Error
	if err != nil {
		return nil, err
	}
	return &enterprise, nil
}

// CreateEnterprise 创建企业主体
func (s *BasicCommonService) CreateEnterprise(enterprise *model.Enterprise) error {
	return s.db.Create(enterprise).Error
}

// UpdateEnterprise 更新企业主体
func (s *BasicCommonService) UpdateEnterprise(enterprise *model.Enterprise) error {
	if enterprise.ID == 0 {
		return errors.New("enterprise id is required")
	}
	return s.db.Model(enterprise).Updates(enterprise).Error
}

// DeleteEnterprise 删除企业主体
func (s *BasicCommonService) DeleteEnterprise(id uint) error {
	return s.db.Delete(&model.Enterprise{}, id).Error
}

// GetRegionList 获取地区列表
func (s *BasicCommonService) GetRegionList(parentID *uint) ([]model.Region, error) {
	var regions []model.Region
	query := s.db.Model(&model.Region{})
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	}
	err := query.Order("sort asc").Find(&regions).Error
	return regions, err
}

// GetRegionByID 根据ID获取地区
func (s *BasicCommonService) GetRegionByID(id uint) (*model.Region, error) {
	var region model.Region
	err := s.db.First(&region, id).Error
	if err != nil {
		return nil, err
	}
	return &region, nil
}

// CreateRegion 创建地区
func (s *BasicCommonService) CreateRegion(region *model.Region) error {
	if region.ParentID != nil {
		var parent model.Region
		if err := s.db.First(&parent, *region.ParentID).Error; err != nil {
			return errors.New("parent region not found")
		}
		region.Level = parent.Level + 1
	} else {
		region.Level = 1
	}
	return s.db.Create(region).Error
}

// UpdateRegion 更新地区
func (s *BasicCommonService) UpdateRegion(region *model.Region) error {
	if region.ID == 0 {
		return errors.New("region id is required")
	}
	return s.db.Model(region).Updates(region).Error
}

// DeleteRegion 删除地区
func (s *BasicCommonService) DeleteRegion(id uint) error {
	// 检查是否有子地区
	var count int64
	if err := s.db.Model(&model.Region{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete region with sub-regions")
	}

	return s.db.Delete(&model.Region{}, id).Error
}

// GetMessageTemplateList 获取消息模板列表
func (s *BasicCommonService) GetMessageTemplateList(templateType int) ([]model.MessageTemplate, error) {
	var templates []model.MessageTemplate
	query := s.db.Model(&model.MessageTemplate{})
	if templateType > 0 {
		query = query.Where("type = ?", templateType)
	}
	err := query.Find(&templates).Error
	return templates, err
}

// GetMessageTemplateByID 根据ID获取消息模板
func (s *BasicCommonService) GetMessageTemplateByID(id uint) (*model.MessageTemplate, error) {
	var template model.MessageTemplate
	err := s.db.First(&template, id).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// GetMessageTemplateByCode 根据Code获取消息模板
func (s *BasicCommonService) GetMessageTemplateByCode(code string) (*model.MessageTemplate, error) {
	var template model.MessageTemplate
	err := s.db.Where("code = ?", code).First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// CreateMessageTemplate 创建消息模板
func (s *BasicCommonService) CreateMessageTemplate(template *model.MessageTemplate) error {
	return s.db.Create(template).Error
}

// UpdateMessageTemplate 更新消息模板
func (s *BasicCommonService) UpdateMessageTemplate(template *model.MessageTemplate) error {
	if template.ID == 0 {
		return errors.New("message template id is required")
	}
	return s.db.Model(template).Updates(template).Error
}

// DeleteMessageTemplate 删除消息模板
func (s *BasicCommonService) DeleteMessageTemplate(id uint) error {
	return s.db.Delete(&model.MessageTemplate{}, id).Error
}
