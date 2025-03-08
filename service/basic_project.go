package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

// BasicProjectService 基础数据-项目模块服务
type BasicProjectService struct {
	db *gorm.DB
}

func NewBasicProjectService(db *gorm.DB) *BasicProjectService {
	return &BasicProjectService{db: db}
}

// GetProjectStageList 获取项目阶段列表
func (s *BasicProjectService) GetProjectStageList() ([]model.ProjectStage, error) {
	var stages []model.ProjectStage
	err := s.db.Order("sort asc").Find(&stages).Error
	return stages, err
}

// GetProjectStageByID 根据ID获取项目阶段
func (s *BasicProjectService) GetProjectStageByID(id uint) (*model.ProjectStage, error) {
	var stage model.ProjectStage
	err := s.db.First(&stage, id).Error
	if err != nil {
		return nil, err
	}
	return &stage, nil
}

// CreateProjectStage 创建项目阶段
func (s *BasicProjectService) CreateProjectStage(stage *model.ProjectStage) error {
	return s.db.Create(stage).Error
}

// UpdateProjectStage 更新项目阶段
func (s *BasicProjectService) UpdateProjectStage(stage *model.ProjectStage) error {
	if stage.ID == 0 {
		return errors.New("project stage id is required")
	}
	return s.db.Model(stage).Updates(stage).Error
}

// DeleteProjectStage 删除项目阶段
func (s *BasicProjectService) DeleteProjectStage(id uint) error {
	return s.db.Delete(&model.ProjectStage{}, id).Error
}

// GetProjectCategoryList 获取项目分类列表
func (s *BasicProjectService) GetProjectCategoryList() ([]model.ProjectCategory, error) {
	var categories []model.ProjectCategory
	err := s.db.Order("sort asc").Find(&categories).Error
	return categories, err
}

// GetProjectCategoryByID 根据ID获取项目分类
func (s *BasicProjectService) GetProjectCategoryByID(id uint) (*model.ProjectCategory, error) {
	var category model.ProjectCategory
	err := s.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateProjectCategory 创建项目分类
func (s *BasicProjectService) CreateProjectCategory(category *model.ProjectCategory) error {
	return s.db.Create(category).Error
}

// UpdateProjectCategory 更新项目分类
func (s *BasicProjectService) UpdateProjectCategory(category *model.ProjectCategory) error {
	if category.ID == 0 {
		return errors.New("project category id is required")
	}
	return s.db.Model(category).Updates(category).Error
}

// DeleteProjectCategory 删除项目分类
func (s *BasicProjectService) DeleteProjectCategory(id uint) error {
	return s.db.Delete(&model.ProjectCategory{}, id).Error
}

// GetWorkTypeList 获取工作类型列表
func (s *BasicProjectService) GetWorkTypeList() ([]model.WorkType, error) {
	var types []model.WorkType
	err := s.db.Order("sort asc").Find(&types).Error
	return types, err
}

// GetWorkTypeByID 根据ID获取工作类型
func (s *BasicProjectService) GetWorkTypeByID(id uint) (*model.WorkType, error) {
	var workType model.WorkType
	err := s.db.First(&workType, id).Error
	if err != nil {
		return nil, err
	}
	return &workType, nil
}

// CreateWorkType 创建工作类型
func (s *BasicProjectService) CreateWorkType(workType *model.WorkType) error {
	return s.db.Create(workType).Error
}

// UpdateWorkType 更新工作类型
func (s *BasicProjectService) UpdateWorkType(workType *model.WorkType) error {
	if workType.ID == 0 {
		return errors.New("work type id is required")
	}
	return s.db.Model(workType).Updates(workType).Error
}

// DeleteWorkType 删除工作类型
func (s *BasicProjectService) DeleteWorkType(id uint) error {
	return s.db.Delete(&model.WorkType{}, id).Error
}
