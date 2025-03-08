package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
	systemService *service.SystemService
}

func NewSystemController(systemService *service.SystemService) *SystemController {
	return &SystemController{
		systemService: systemService,
	}
}

// RegisterRoutes 注册路由
func (c *SystemController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/system")
	{
		// 系统配置
		api.GET("/configs", c.GetSystemConfigList)
		api.GET("/configs/:key", c.GetSystemConfigByKey)
		api.PUT("/configs/:id", c.UpdateSystemConfig)

		// 功能模块
		api.GET("/modules", c.GetModuleList)
		api.POST("/modules", c.CreateModule)
		api.PUT("/modules/:id", c.UpdateModule)
		api.DELETE("/modules/:id", c.DeleteModule)

		// 模块配置
		api.GET("/module-configs", c.GetModuleConfigList)
		api.PUT("/module-configs/:id", c.UpdateModuleConfig)

		// 功能节点
		api.GET("/function-nodes", c.GetFunctionNodeList)
		api.POST("/function-nodes", c.CreateFunctionNode)
		api.PUT("/function-nodes/:id", c.UpdateFunctionNode)
		api.DELETE("/function-nodes/:id", c.DeleteFunctionNode)

		// 角色管理
		api.GET("/roles", c.GetRoleList)
		api.POST("/roles", c.CreateRole)
		api.PUT("/roles/:id", c.UpdateRole)
		api.DELETE("/roles/:id", c.DeleteRole)
		api.GET("/roles/:id/functions", c.GetRoleFunctions)
		api.PUT("/roles/:id/functions", c.UpdateRoleFunctions)

		// 操作日志
		api.GET("/operation-logs", c.GetOperationLogList)

		// 附件管理
		api.GET("/attachments", c.GetAttachmentList)
		api.POST("/attachments", c.CreateAttachment)
		api.DELETE("/attachments/:id", c.DeleteAttachment)

		// 备份管理
		api.GET("/backup-records", c.GetBackupRecordList)
		api.POST("/backup-records", c.CreateBackupRecord)
		api.PUT("/backup-records/:id", c.UpdateBackupRecord)
		api.DELETE("/backup-records/:id", c.DeleteBackupRecord)

		// 定时任务
		api.GET("/scheduled-tasks", c.GetScheduledTaskList)
		api.POST("/scheduled-tasks", c.CreateScheduledTask)
		api.PUT("/scheduled-tasks/:id", c.UpdateScheduledTask)
		api.DELETE("/scheduled-tasks/:id", c.DeleteScheduledTask)
		api.PUT("/scheduled-tasks/:id/status", c.UpdateTaskStatus)
	}
}

// GetSystemConfigList 获取系统配置列表
func (c *SystemController) GetSystemConfigList(ctx *gin.Context) {
	configs, err := c.systemService.GetSystemConfigList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, configs)
}

// GetSystemConfigByKey 根据Key获取系统配置
func (c *SystemController) GetSystemConfigByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	config, err := c.systemService.GetSystemConfigByKey(key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, config)
}

// UpdateSystemConfig 更新系统配置
func (c *SystemController) UpdateSystemConfig(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var config model.SystemConfig
	if err := ctx.ShouldBindJSON(&config); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.ID = uint(id)
	if err := c.systemService.UpdateSystemConfig(&config); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, config)
}

// GetModuleList 获取功能模块列表
func (c *SystemController) GetModuleList(ctx *gin.Context) {
	modules, err := c.systemService.GetModuleList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, modules)
}

// CreateModule 创建功能模块
func (c *SystemController) CreateModule(ctx *gin.Context) {
	var module model.Module
	if err := ctx.ShouldBindJSON(&module); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.systemService.CreateModule(&module); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, module)
}

// UpdateModule 更新功能模块
func (c *SystemController) UpdateModule(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var module model.Module
	if err := ctx.ShouldBindJSON(&module); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module.ID = uint(id)
	if err := c.systemService.UpdateModule(&module); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, module)
}

// DeleteModule 删除功能模块
func (c *SystemController) DeleteModule(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.systemService.DeleteModule(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetModuleConfigList 获取模块配置列表
func (c *SystemController) GetModuleConfigList(ctx *gin.Context) {
	moduleID, _ := strconv.ParseUint(ctx.Query("module_id"), 10, 32)
	configs, err := c.systemService.GetModuleConfigList(uint(moduleID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, configs)
}

// UpdateModuleConfig 更新模块配置
func (c *SystemController) UpdateModuleConfig(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var config model.ModuleConfig
	if err := ctx.ShouldBindJSON(&config); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.ID = uint(id)
	if err := c.systemService.UpdateModuleConfig(&config); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, config)
}

// GetFunctionNodeList 获取功能节点列表
func (c *SystemController) GetFunctionNodeList(ctx *gin.Context) {
	moduleID, _ := strconv.ParseUint(ctx.Query("module_id"), 10, 32)
	nodes, err := c.systemService.GetFunctionNodeList(uint(moduleID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nodes)
}

// CreateFunctionNode 创建功能节点
func (c *SystemController) CreateFunctionNode(ctx *gin.Context) {
	var node model.FunctionNode
	if err := ctx.ShouldBindJSON(&node); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.systemService.CreateFunctionNode(&node); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, node)
}

// UpdateFunctionNode 更新功能节点
func (c *SystemController) UpdateFunctionNode(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var node model.FunctionNode
	if err := ctx.ShouldBindJSON(&node); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node.ID = uint(id)
	if err := c.systemService.UpdateFunctionNode(&node); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, node)
}

// DeleteFunctionNode 删除功能节点
func (c *SystemController) DeleteFunctionNode(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.systemService.DeleteFunctionNode(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetRoleList 获取角色列表
func (c *SystemController) GetRoleList(ctx *gin.Context) {
	roles, err := c.systemService.GetRoleList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

// CreateRole 创建角色
func (c *SystemController) CreateRole(ctx *gin.Context) {
	var role model.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.systemService.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

// UpdateRole 更新角色
func (c *SystemController) UpdateRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var role model.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.ID = uint(id)
	if err := c.systemService.UpdateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

// DeleteRole 删除角色
func (c *SystemController) DeleteRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.systemService.DeleteRole(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetRoleFunctions 获取角色的功能权限
func (c *SystemController) GetRoleFunctions(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	nodes, err := c.systemService.GetRoleFunctions(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nodes)
}

// UpdateRoleFunctions 更新角色的功能权限
func (c *SystemController) UpdateRoleFunctions(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		FunctionNodeIDs []uint `json:"function_node_ids" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.systemService.UpdateRoleFunctions(uint(id), req.FunctionNodeIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetOperationLogList 获取操作日志列表
func (c *SystemController) GetOperationLogList(ctx *gin.Context) {
	userID, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
	module := ctx.Query("module")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	logs, total, err := c.systemService.GetOperationLogList(uint(userID), module, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  logs,
		"total": total,
	})
}

// GetAttachmentList 获取附件列表
func (c *SystemController) GetAttachmentList(ctx *gin.Context) {
	module := ctx.Query("module")
	relatedID, _ := strconv.ParseUint(ctx.Query("related_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	attachments, total, err := c.systemService.GetAttachmentList(module, uint(relatedID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  attachments,
		"total": total,
	})
}

// CreateAttachment 创建附件
func (c *SystemController) CreateAttachment(ctx *gin.Context) {
	var attachment model.Attachment
	if err := ctx.ShouldBindJSON(&attachment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取userID
	attachment.UploadedBy = uint(1)

	if err := c.systemService.CreateAttachment(&attachment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, attachment)
}

// DeleteAttachment 删除附件
func (c *SystemController) DeleteAttachment(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.systemService.DeleteAttachment(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetBackupRecordList 获取备份记录列表
func (c *SystemController) GetBackupRecordList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	records, total, err := c.systemService.GetBackupRecordList(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
	})
}

// CreateBackupRecord 创建备份记录
func (c *SystemController) CreateBackupRecord(ctx *gin.Context) {
	var record model.BackupRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.systemService.CreateBackupRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

// UpdateBackupRecord 更新备份记录
func (c *SystemController) UpdateBackupRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var record model.BackupRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = uint(id)
	if err := c.systemService.UpdateBackupRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// DeleteBackupRecord 删除备份记录
func (c *SystemController) DeleteBackupRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.systemService.DeleteBackupRecord(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetScheduledTaskList 获取定时任务列表
func (c *SystemController) GetScheduledTaskList(ctx *gin.Context) {
	tasks, err := c.systemService.GetScheduledTaskList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// CreateScheduledTask 创建定时任务
func (c *SystemController) CreateScheduledTask(ctx *gin.Context) {
	var task model.ScheduledTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.systemService.CreateScheduledTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

// UpdateScheduledTask 更新定时任务
func (c *SystemController) UpdateScheduledTask(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var task model.ScheduledTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uint(id)
	if err := c.systemService.UpdateScheduledTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// DeleteScheduledTask 删除定时任务
func (c *SystemController) DeleteScheduledTask(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.systemService.DeleteScheduledTask(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// UpdateTaskStatus 更新任务状态
func (c *SystemController) UpdateTaskStatus(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		Status int `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.systemService.UpdateTaskStatus(uint(id), req.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
