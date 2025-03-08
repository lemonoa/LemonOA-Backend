package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type SystemService struct {
	db *gorm.DB
}

func NewSystemService(db *gorm.DB) *SystemService {
	return &SystemService{db: db}
}

// GetSystemConfigList 获取系统配置列表
func (s *SystemService) GetSystemConfigList() ([]model.SystemConfig, error) {
	var configs []model.SystemConfig
	err := s.db.Find(&configs).Error
	return configs, err
}

// GetSystemConfigByKey 根据Key获取系统配置
func (s *SystemService) GetSystemConfigByKey(key string) (*model.SystemConfig, error) {
	var config model.SystemConfig
	err := s.db.Where("key = ?", key).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// UpdateSystemConfig 更新系统配置
func (s *SystemService) UpdateSystemConfig(config *model.SystemConfig) error {
	if config.ID == 0 {
		return errors.New("system config id is required")
	}
	return s.db.Model(config).Updates(config).Error
}

// GetModuleList 获取功能模块列表
func (s *SystemService) GetModuleList() ([]model.Module, error) {
	var modules []model.Module
	err := s.db.Order("sort asc").Find(&modules).Error
	return modules, err
}

// CreateModule 创建功能模块
func (s *SystemService) CreateModule(module *model.Module) error {
	return s.db.Create(module).Error
}

// UpdateModule 更新功能模块
func (s *SystemService) UpdateModule(module *model.Module) error {
	if module.ID == 0 {
		return errors.New("module id is required")
	}
	return s.db.Model(module).Updates(module).Error
}

// DeleteModule 删除功能模块
func (s *SystemService) DeleteModule(id uint) error {
	// 检查是否有子模块
	var count int64
	if err := s.db.Model(&model.Module{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete module with sub-modules")
	}

	// 检查是否有关联的功能节点
	if err := s.db.Model(&model.FunctionNode{}).Where("module_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete module with associated function nodes")
	}

	return s.db.Delete(&model.Module{}, id).Error
}

// GetModuleConfigList 获取模块配置列表
func (s *SystemService) GetModuleConfigList(moduleID uint) ([]model.ModuleConfig, error) {
	var configs []model.ModuleConfig
	query := s.db.Model(&model.ModuleConfig{})
	if moduleID > 0 {
		query = query.Where("module_id = ?", moduleID)
	}
	err := query.Find(&configs).Error
	return configs, err
}

// UpdateModuleConfig 更新模块配置
func (s *SystemService) UpdateModuleConfig(config *model.ModuleConfig) error {
	if config.ID == 0 {
		return errors.New("module config id is required")
	}
	return s.db.Model(config).Updates(config).Error
}

// GetFunctionNodeList 获取功能节点列表
func (s *SystemService) GetFunctionNodeList(moduleID uint) ([]model.FunctionNode, error) {
	var nodes []model.FunctionNode
	query := s.db.Model(&model.FunctionNode{})
	if moduleID > 0 {
		query = query.Where("module_id = ?", moduleID)
	}
	err := query.Order("sort asc").Find(&nodes).Error
	return nodes, err
}

// CreateFunctionNode 创建功能节点
func (s *SystemService) CreateFunctionNode(node *model.FunctionNode) error {
	return s.db.Create(node).Error
}

// UpdateFunctionNode 更新功能节点
func (s *SystemService) UpdateFunctionNode(node *model.FunctionNode) error {
	if node.ID == 0 {
		return errors.New("function node id is required")
	}
	return s.db.Model(node).Updates(node).Error
}

// DeleteFunctionNode 删除功能节点
func (s *SystemService) DeleteFunctionNode(id uint) error {
	// 检查是否有角色关联
	var count int64
	if err := s.db.Model(&model.RoleFunction{}).Where("function_node_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete function node with associated roles")
	}

	return s.db.Delete(&model.FunctionNode{}, id).Error
}

// GetRoleList 获取角色列表
func (s *SystemService) GetRoleList() ([]model.Role, error) {
	var roles []model.Role
	err := s.db.Find(&roles).Error
	return roles, err
}

// CreateRole 创建角色
func (s *SystemService) CreateRole(role *model.Role) error {
	return s.db.Create(role).Error
}

// UpdateRole 更新角色
func (s *SystemService) UpdateRole(role *model.Role) error {
	if role.ID == 0 {
		return errors.New("role id is required")
	}
	return s.db.Model(role).Updates(role).Error
}

// DeleteRole 删除角色
func (s *SystemService) DeleteRole(id uint) error {
	// 删除角色的同时删除角色的功能权限
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.RoleFunction{}, "role_id = ?", id).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Role{}, id).Error
	})
}

// GetRoleFunctions 获取角色的功能权限
func (s *SystemService) GetRoleFunctions(roleID uint) ([]model.FunctionNode, error) {
	var nodes []model.FunctionNode
	err := s.db.Joins("JOIN role_functions ON function_nodes.id = role_functions.function_node_id").
		Where("role_functions.role_id = ?", roleID).
		Find(&nodes).Error
	return nodes, err
}

// UpdateRoleFunctions 更新角色的功能权限
func (s *SystemService) UpdateRoleFunctions(roleID uint, functionNodeIDs []uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 删除原有的权限
		if err := tx.Delete(&model.RoleFunction{}, "role_id = ?", roleID).Error; err != nil {
			return err
		}

		// 添加新的权限
		for _, nodeID := range functionNodeIDs {
			roleFunction := &model.RoleFunction{
				RoleID:         roleID,
				FunctionNodeID: nodeID,
			}
			if err := tx.Create(roleFunction).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// GetOperationLogList 获取操作日志列表
func (s *SystemService) GetOperationLogList(userID uint, module string, page, pageSize int) ([]model.OperationLog, int64, error) {
	var logs []model.OperationLog
	var total int64

	query := s.db.Model(&model.OperationLog{})
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}
	if module != "" {
		query = query.Where("module = ?", module)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// CreateOperationLog 创建操作日志
func (s *SystemService) CreateOperationLog(log *model.OperationLog) error {
	return s.db.Create(log).Error
}

// GetAttachmentList 获取附件列表
func (s *SystemService) GetAttachmentList(module string, relatedID uint, page, pageSize int) ([]model.Attachment, int64, error) {
	var attachments []model.Attachment
	var total int64

	query := s.db.Model(&model.Attachment{})
	if module != "" {
		query = query.Where("module = ?", module)
	}
	if relatedID > 0 {
		query = query.Where("related_id = ?", relatedID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&attachments).Error
	if err != nil {
		return nil, 0, err
	}

	return attachments, total, nil
}

// CreateAttachment 创建附件
func (s *SystemService) CreateAttachment(attachment *model.Attachment) error {
	return s.db.Create(attachment).Error
}

// DeleteAttachment 删除附件
func (s *SystemService) DeleteAttachment(id uint) error {
	return s.db.Delete(&model.Attachment{}, id).Error
}

// GetBackupRecordList 获取备份记录列表
func (s *SystemService) GetBackupRecordList(page, pageSize int) ([]model.BackupRecord, int64, error) {
	var records []model.BackupRecord
	var total int64

	err := s.db.Model(&model.BackupRecord{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = s.db.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// CreateBackupRecord 创建备份记录
func (s *SystemService) CreateBackupRecord(record *model.BackupRecord) error {
	return s.db.Create(record).Error
}

// UpdateBackupRecord 更新备份记录
func (s *SystemService) UpdateBackupRecord(record *model.BackupRecord) error {
	if record.ID == 0 {
		return errors.New("backup record id is required")
	}
	return s.db.Model(record).Updates(record).Error
}

// DeleteBackupRecord 删除备份记录
func (s *SystemService) DeleteBackupRecord(id uint) error {
	return s.db.Delete(&model.BackupRecord{}, id).Error
}

// GetScheduledTaskList 获取定时任务列表
func (s *SystemService) GetScheduledTaskList() ([]model.ScheduledTask, error) {
	var tasks []model.ScheduledTask
	err := s.db.Find(&tasks).Error
	return tasks, err
}

// CreateScheduledTask 创建定时任务
func (s *SystemService) CreateScheduledTask(task *model.ScheduledTask) error {
	return s.db.Create(task).Error
}

// UpdateScheduledTask 更新定时任务
func (s *SystemService) UpdateScheduledTask(task *model.ScheduledTask) error {
	if task.ID == 0 {
		return errors.New("scheduled task id is required")
	}
	return s.db.Model(task).Updates(task).Error
}

// DeleteScheduledTask 删除定时任务
func (s *SystemService) DeleteScheduledTask(id uint) error {
	return s.db.Delete(&model.ScheduledTask{}, id).Error
}

// UpdateTaskStatus 更新任务状态
func (s *SystemService) UpdateTaskStatus(id uint, status int) error {
	return s.db.Model(&model.ScheduledTask{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateTaskLastRunAt 更新任务最后运行时间
func (s *SystemService) UpdateTaskLastRunAt(id uint, lastRunAt time.Time) error {
	return s.db.Model(&model.ScheduledTask{}).Where("id = ?", id).Update("last_run_at", lastRunAt).Error
}
