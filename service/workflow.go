package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type WorkflowService struct {
	db *gorm.DB
}

func NewWorkflowService(db *gorm.DB) *WorkflowService {
	return &WorkflowService{db: db}
}

// GetWorkflowTypeList 获取流程类型列表
func (s *WorkflowService) GetWorkflowTypeList(status int, keyword string, page, pageSize int) ([]model.WorkflowType, int64, error) {
	var types []model.WorkflowType
	var total int64

	query := s.db.Model(&model.WorkflowType{})
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

	err = query.Order("sort asc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&types).Error
	if err != nil {
		return nil, 0, err
	}

	return types, total, nil
}

// GetWorkflowTypeByID 根据ID获取流程类型
func (s *WorkflowService) GetWorkflowTypeByID(id uint) (*model.WorkflowType, error) {
	var workflowType model.WorkflowType
	err := s.db.First(&workflowType, id).Error
	if err != nil {
		return nil, err
	}
	return &workflowType, nil
}

// CreateWorkflowType 创建流程类型
func (s *WorkflowService) CreateWorkflowType(workflowType *model.WorkflowType) error {
	return s.db.Create(workflowType).Error
}

// UpdateWorkflowType 更新流程类型
func (s *WorkflowService) UpdateWorkflowType(workflowType *model.WorkflowType) error {
	if workflowType.ID == 0 {
		return errors.New("workflow type id is required")
	}
	return s.db.Model(workflowType).Updates(workflowType).Error
}

// DeleteWorkflowType 删除流程类型
func (s *WorkflowService) DeleteWorkflowType(id uint) error {
	// 检查是否有关联的流程定义
	var count int64
	if err := s.db.Model(&model.WorkflowDefinition{}).Where("type_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete workflow type with associated definitions")
	}

	return s.db.Delete(&model.WorkflowType{}, id).Error
}

// GetWorkflowDefinitionList 获取流程定义列表
func (s *WorkflowService) GetWorkflowDefinitionList(typeID uint, status int, keyword string, page, pageSize int) ([]model.WorkflowDefinition, int64, error) {
	var definitions []model.WorkflowDefinition
	var total int64

	query := s.db.Model(&model.WorkflowDefinition{})
	if typeID > 0 {
		query = query.Where("type_id = ?", typeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&definitions).Error
	if err != nil {
		return nil, 0, err
	}

	return definitions, total, nil
}

// GetWorkflowDefinitionByID 根据ID获取流程定义
func (s *WorkflowService) GetWorkflowDefinitionByID(id uint) (*model.WorkflowDefinition, error) {
	var definition model.WorkflowDefinition
	err := s.db.First(&definition, id).Error
	if err != nil {
		return nil, err
	}
	return &definition, nil
}

// CreateWorkflowDefinition 创建流程定义
func (s *WorkflowService) CreateWorkflowDefinition(definition *model.WorkflowDefinition) error {
	// 检查流程类型是否存在
	var workflowType model.WorkflowType
	if err := s.db.First(&workflowType, definition.TypeID).Error; err != nil {
		return errors.New("workflow type not found")
	}

	return s.db.Create(definition).Error
}

// UpdateWorkflowDefinition 更新流程定义
func (s *WorkflowService) UpdateWorkflowDefinition(definition *model.WorkflowDefinition) error {
	if definition.ID == 0 {
		return errors.New("workflow definition id is required")
	}

	// 检查流程类型是否存在
	var workflowType model.WorkflowType
	if err := s.db.First(&workflowType, definition.TypeID).Error; err != nil {
		return errors.New("workflow type not found")
	}

	return s.db.Model(definition).Updates(definition).Error
}

// DeleteWorkflowDefinition 删除流程定义
func (s *WorkflowService) DeleteWorkflowDefinition(id uint) error {
	// 检查是否有关联的流程实例
	var count int64
	if err := s.db.Model(&model.WorkflowInstance{}).Where("definition_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete workflow definition with associated instances")
	}

	return s.db.Delete(&model.WorkflowDefinition{}, id).Error
}

// PublishWorkflowDefinition 发布流程定义
func (s *WorkflowService) PublishWorkflowDefinition(id uint) error {
	return s.db.Model(&model.WorkflowDefinition{}).Where("id = ?", id).Update("status", 2).Error
}

// DisableWorkflowDefinition 停用流程定义
func (s *WorkflowService) DisableWorkflowDefinition(id uint) error {
	return s.db.Model(&model.WorkflowDefinition{}).Where("id = ?", id).Update("status", 3).Error
}

// GetWorkflowNodeList 获取流程节点列表
func (s *WorkflowService) GetWorkflowNodeList(definitionID uint) ([]model.WorkflowNode, error) {
	var nodes []model.WorkflowNode
	err := s.db.Where("definition_id = ?", definitionID).Order("sort asc").Find(&nodes).Error
	return nodes, err
}

// GetWorkflowNodeByID 根据ID获取流程节点
func (s *WorkflowService) GetWorkflowNodeByID(id uint) (*model.WorkflowNode, error) {
	var node model.WorkflowNode
	err := s.db.First(&node, id).Error
	if err != nil {
		return nil, err
	}
	return &node, nil
}

// CreateWorkflowNode 创建流程节点
func (s *WorkflowService) CreateWorkflowNode(node *model.WorkflowNode) error {
	// 检查流程定义是否存在
	var definition model.WorkflowDefinition
	if err := s.db.First(&definition, node.DefinitionID).Error; err != nil {
		return errors.New("workflow definition not found")
	}

	return s.db.Create(node).Error
}

// UpdateWorkflowNode 更新流程节点
func (s *WorkflowService) UpdateWorkflowNode(node *model.WorkflowNode) error {
	if node.ID == 0 {
		return errors.New("workflow node id is required")
	}

	// 检查流程定义是否存在
	var definition model.WorkflowDefinition
	if err := s.db.First(&definition, node.DefinitionID).Error; err != nil {
		return errors.New("workflow definition not found")
	}

	return s.db.Model(node).Updates(node).Error
}

// DeleteWorkflowNode 删除流程节点
func (s *WorkflowService) DeleteWorkflowNode(id uint) error {
	// 检查是否有关联的任务
	var count int64
	if err := s.db.Model(&model.WorkflowTask{}).Where("node_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete workflow node with associated tasks")
	}

	return s.db.Delete(&model.WorkflowNode{}, id).Error
}

// GetWorkflowInstanceList 获取流程实例列表
func (s *WorkflowService) GetWorkflowInstanceList(definitionID uint, status int, keyword string, page, pageSize int) ([]model.WorkflowInstance, int64, error) {
	var instances []model.WorkflowInstance
	var total int64

	query := s.db.Model(&model.WorkflowInstance{})
	if definitionID > 0 {
		query = query.Where("definition_id = ?", definitionID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&instances).Error
	if err != nil {
		return nil, 0, err
	}

	return instances, total, nil
}

// GetWorkflowInstanceByID 根据ID获取流程实例
func (s *WorkflowService) GetWorkflowInstanceByID(id uint) (*model.WorkflowInstance, error) {
	var instance model.WorkflowInstance
	err := s.db.First(&instance, id).Error
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

// CreateWorkflowInstance 创建流程实例
func (s *WorkflowService) CreateWorkflowInstance(instance *model.WorkflowInstance) error {
	// 检查流程定义是否存在
	var definition model.WorkflowDefinition
	if err := s.db.First(&definition, instance.DefinitionID).Error; err != nil {
		return errors.New("workflow definition not found")
	}

	// 设置开始时间
	now := time.Now()
	instance.StartTime = &now

	return s.db.Create(instance).Error
}

// UpdateWorkflowInstance 更新流程实例
func (s *WorkflowService) UpdateWorkflowInstance(instance *model.WorkflowInstance) error {
	if instance.ID == 0 {
		return errors.New("workflow instance id is required")
	}

	// 检查流程定义是否存在
	var definition model.WorkflowDefinition
	if err := s.db.First(&definition, instance.DefinitionID).Error; err != nil {
		return errors.New("workflow definition not found")
	}

	return s.db.Model(instance).Updates(instance).Error
}

// DeleteWorkflowInstance 删除流程实例
func (s *WorkflowService) DeleteWorkflowInstance(id uint) error {
	// 检查是否有关联的任务
	var count int64
	if err := s.db.Model(&model.WorkflowTask{}).Where("instance_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete workflow instance with associated tasks")
	}

	return s.db.Delete(&model.WorkflowInstance{}, id).Error
}

// CancelWorkflowInstance 取消流程实例
func (s *WorkflowService) CancelWorkflowInstance(id uint) error {
	// 设置结束时间
	now := time.Now()
	return s.db.Model(&model.WorkflowInstance{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":   3,
		"end_time": &now,
	}).Error
}

// GetWorkflowTaskList 获取流程任务列表
func (s *WorkflowService) GetWorkflowTaskList(instanceID, assigneeID uint, status int, page, pageSize int) ([]model.WorkflowTask, int64, error) {
	var tasks []model.WorkflowTask
	var total int64

	query := s.db.Model(&model.WorkflowTask{})
	if instanceID > 0 {
		query = query.Where("instance_id = ?", instanceID)
	}
	if assigneeID > 0 {
		query = query.Where("assignee_id = ?", assigneeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

// GetWorkflowTaskByID 根据ID获取流程任务
func (s *WorkflowService) GetWorkflowTaskByID(id uint) (*model.WorkflowTask, error) {
	var task model.WorkflowTask
	err := s.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateWorkflowTask 创建流程任务
func (s *WorkflowService) CreateWorkflowTask(task *model.WorkflowTask) error {
	// 检查流程实例是否存在
	var instance model.WorkflowInstance
	if err := s.db.First(&instance, task.InstanceID).Error; err != nil {
		return errors.New("workflow instance not found")
	}

	// 检查流程节点是否存在
	var node model.WorkflowNode
	if err := s.db.First(&node, task.NodeID).Error; err != nil {
		return errors.New("workflow node not found")
	}

	return s.db.Create(task).Error
}

// UpdateWorkflowTask 更新流程任务
func (s *WorkflowService) UpdateWorkflowTask(task *model.WorkflowTask) error {
	if task.ID == 0 {
		return errors.New("workflow task id is required")
	}

	// 检查流程实例是否存在
	var instance model.WorkflowInstance
	if err := s.db.First(&instance, task.InstanceID).Error; err != nil {
		return errors.New("workflow instance not found")
	}

	// 检查流程节点是否存在
	var node model.WorkflowNode
	if err := s.db.First(&node, task.NodeID).Error; err != nil {
		return errors.New("workflow node not found")
	}

	return s.db.Model(task).Updates(task).Error
}

// DeleteWorkflowTask 删除流程任务
func (s *WorkflowService) DeleteWorkflowTask(id uint) error {
	return s.db.Delete(&model.WorkflowTask{}, id).Error
}

// HandleWorkflowTask 处理流程任务
func (s *WorkflowService) HandleWorkflowTask(id uint, action int, comment string) error {
	// 获取任务信息
	var task model.WorkflowTask
	if err := s.db.First(&task, id).Error; err != nil {
		return errors.New("workflow task not found")
	}

	// 设置处理时间
	now := time.Now()
	task.HandleTime = &now
	task.Action = action
	task.Comment = comment
	task.Status = 2

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 更新任务状态
		if err := tx.Model(&task).Updates(task).Error; err != nil {
			return err
		}

		// 如果是驳回，则结束流程实例
		if action == 2 {
			return tx.Model(&model.WorkflowInstance{}).Where("id = ?", task.InstanceID).Updates(map[string]interface{}{
				"status":   2,
				"end_time": &now,
			}).Error
		}

		// 如果是同意，则检查是否还有未处理的任务
		var count int64
		if err := tx.Model(&model.WorkflowTask{}).Where("instance_id = ? AND status = ?", task.InstanceID, 1).Count(&count).Error; err != nil {
			return err
		}

		// 如果没有未处理的任务，则结束流程实例
		if count == 0 {
			return tx.Model(&model.WorkflowInstance{}).Where("id = ?", task.InstanceID).Updates(map[string]interface{}{
				"status":   2,
				"end_time": &now,
			}).Error
		}

		return nil
	})
}

// TransferWorkflowTask 转办流程任务
func (s *WorkflowService) TransferWorkflowTask(id, assigneeID uint) error {
	// 获取任务信息
	var task model.WorkflowTask
	if err := s.db.First(&task, id).Error; err != nil {
		return errors.New("workflow task not found")
	}

	// 创建新任务
	newTask := &model.WorkflowTask{
		InstanceID: task.InstanceID,
		NodeID:     task.NodeID,
		AssigneeID: assigneeID,
		Status:     1,
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建新任务
		if err := tx.Create(newTask).Error; err != nil {
			return err
		}

		// 更新原任务状态
		return tx.Model(&task).Updates(map[string]interface{}{
			"status": 3,
		}).Error
	})
}
