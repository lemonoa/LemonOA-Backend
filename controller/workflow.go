package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type WorkflowController struct {
	workflowService *service.WorkflowService
}

func NewWorkflowController(workflowService *service.WorkflowService) *WorkflowController {
	return &WorkflowController{
		workflowService: workflowService,
	}
}

// RegisterRoutes 注册路由
func (c *WorkflowController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/workflows")
	{
		// 流程类型管理
		api.GET("/types", c.GetWorkflowTypeList)
		api.GET("/types/:id", c.GetWorkflowTypeByID)
		api.POST("/types", c.CreateWorkflowType)
		api.PUT("/types/:id", c.UpdateWorkflowType)
		api.DELETE("/types/:id", c.DeleteWorkflowType)

		// 流程定义管理
		api.GET("/definitions", c.GetWorkflowDefinitionList)
		api.GET("/definitions/:id", c.GetWorkflowDefinitionByID)
		api.POST("/definitions", c.CreateWorkflowDefinition)
		api.PUT("/definitions/:id", c.UpdateWorkflowDefinition)
		api.DELETE("/definitions/:id", c.DeleteWorkflowDefinition)
		api.PUT("/definitions/:id/publish", c.PublishWorkflowDefinition)
		api.PUT("/definitions/:id/disable", c.DisableWorkflowDefinition)

		// 流程节点管理
		api.GET("/definitions/:id/nodes", c.GetWorkflowNodeList)
		api.GET("/nodes/:id", c.GetWorkflowNodeByID)
		api.POST("/nodes", c.CreateWorkflowNode)
		api.PUT("/nodes/:id", c.UpdateWorkflowNode)
		api.DELETE("/nodes/:id", c.DeleteWorkflowNode)

		// 流程实例管理
		api.GET("/instances", c.GetWorkflowInstanceList)
		api.GET("/instances/:id", c.GetWorkflowInstanceByID)
		api.POST("/instances", c.CreateWorkflowInstance)
		api.PUT("/instances/:id", c.UpdateWorkflowInstance)
		api.DELETE("/instances/:id", c.DeleteWorkflowInstance)
		api.PUT("/instances/:id/cancel", c.CancelWorkflowInstance)

		// 流程任务管理
		api.GET("/tasks", c.GetWorkflowTaskList)
		api.GET("/tasks/:id", c.GetWorkflowTaskByID)
		api.POST("/tasks", c.CreateWorkflowTask)
		api.PUT("/tasks/:id", c.UpdateWorkflowTask)
		api.DELETE("/tasks/:id", c.DeleteWorkflowTask)
		api.PUT("/tasks/:id/handle", c.HandleWorkflowTask)
		api.PUT("/tasks/:id/transfer", c.TransferWorkflowTask)
	}
}

// GetWorkflowTypeList 获取流程类型列表
func (c *WorkflowController) GetWorkflowTypeList(ctx *gin.Context) {
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	types, total, err := c.workflowService.GetWorkflowTypeList(status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  types,
		"total": total,
	})
}

// GetWorkflowTypeByID 根据ID获取流程类型
func (c *WorkflowController) GetWorkflowTypeByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	workflowType, err := c.workflowService.GetWorkflowTypeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workflowType)
}

// CreateWorkflowType 创建流程类型
func (c *WorkflowController) CreateWorkflowType(ctx *gin.Context) {
	var workflowType model.WorkflowType
	if err := ctx.ShouldBindJSON(&workflowType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	workflowType.CreatedBy = uint(1)

	if err := c.workflowService.CreateWorkflowType(&workflowType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, workflowType)
}

// UpdateWorkflowType 更新流程类型
func (c *WorkflowController) UpdateWorkflowType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var workflowType model.WorkflowType
	if err := ctx.ShouldBindJSON(&workflowType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workflowType.ID = uint(id)
	if err := c.workflowService.UpdateWorkflowType(&workflowType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workflowType)
}

// DeleteWorkflowType 删除流程类型
func (c *WorkflowController) DeleteWorkflowType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.DeleteWorkflowType(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetWorkflowDefinitionList 获取流程定义列表
func (c *WorkflowController) GetWorkflowDefinitionList(ctx *gin.Context) {
	typeID, _ := strconv.ParseUint(ctx.Query("type_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	definitions, total, err := c.workflowService.GetWorkflowDefinitionList(uint(typeID), status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  definitions,
		"total": total,
	})
}

// GetWorkflowDefinitionByID 根据ID获取流程定义
func (c *WorkflowController) GetWorkflowDefinitionByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	definition, err := c.workflowService.GetWorkflowDefinitionByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, definition)
}

// CreateWorkflowDefinition 创建流程定义
func (c *WorkflowController) CreateWorkflowDefinition(ctx *gin.Context) {
	var definition model.WorkflowDefinition
	if err := ctx.ShouldBindJSON(&definition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	definition.CreatedBy = uint(1)

	if err := c.workflowService.CreateWorkflowDefinition(&definition); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, definition)
}

// UpdateWorkflowDefinition 更新流程定义
func (c *WorkflowController) UpdateWorkflowDefinition(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var definition model.WorkflowDefinition
	if err := ctx.ShouldBindJSON(&definition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	definition.ID = uint(id)
	if err := c.workflowService.UpdateWorkflowDefinition(&definition); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, definition)
}

// DeleteWorkflowDefinition 删除流程定义
func (c *WorkflowController) DeleteWorkflowDefinition(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.DeleteWorkflowDefinition(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// PublishWorkflowDefinition 发布流程定义
func (c *WorkflowController) PublishWorkflowDefinition(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.PublishWorkflowDefinition(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DisableWorkflowDefinition 停用流程定义
func (c *WorkflowController) DisableWorkflowDefinition(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.DisableWorkflowDefinition(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetWorkflowNodeList 获取流程节点列表
func (c *WorkflowController) GetWorkflowNodeList(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	nodes, err := c.workflowService.GetWorkflowNodeList(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nodes)
}

// GetWorkflowNodeByID 根据ID获取流程节点
func (c *WorkflowController) GetWorkflowNodeByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	node, err := c.workflowService.GetWorkflowNodeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, node)
}

// CreateWorkflowNode 创建流程节点
func (c *WorkflowController) CreateWorkflowNode(ctx *gin.Context) {
	var node model.WorkflowNode
	if err := ctx.ShouldBindJSON(&node); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	node.CreatedBy = uint(1)

	if err := c.workflowService.CreateWorkflowNode(&node); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, node)
}

// UpdateWorkflowNode 更新流程节点
func (c *WorkflowController) UpdateWorkflowNode(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var node model.WorkflowNode
	if err := ctx.ShouldBindJSON(&node); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node.ID = uint(id)
	if err := c.workflowService.UpdateWorkflowNode(&node); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, node)
}

// DeleteWorkflowNode 删除流程节点
func (c *WorkflowController) DeleteWorkflowNode(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.DeleteWorkflowNode(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetWorkflowInstanceList 获取流程实例列表
func (c *WorkflowController) GetWorkflowInstanceList(ctx *gin.Context) {
	definitionID, _ := strconv.ParseUint(ctx.Query("definition_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	instances, total, err := c.workflowService.GetWorkflowInstanceList(uint(definitionID), status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  instances,
		"total": total,
	})
}

// GetWorkflowInstanceByID 根据ID获取流程实例
func (c *WorkflowController) GetWorkflowInstanceByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	instance, err := c.workflowService.GetWorkflowInstanceByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, instance)
}

// CreateWorkflowInstance 创建流程实例
func (c *WorkflowController) CreateWorkflowInstance(ctx *gin.Context) {
	var instance model.WorkflowInstance
	if err := ctx.ShouldBindJSON(&instance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	instance.CreatedBy = uint(1)

	if err := c.workflowService.CreateWorkflowInstance(&instance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, instance)
}

// UpdateWorkflowInstance 更新流程实例
func (c *WorkflowController) UpdateWorkflowInstance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var instance model.WorkflowInstance
	if err := ctx.ShouldBindJSON(&instance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instance.ID = uint(id)
	if err := c.workflowService.UpdateWorkflowInstance(&instance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, instance)
}

// DeleteWorkflowInstance 删除流程实例
func (c *WorkflowController) DeleteWorkflowInstance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.DeleteWorkflowInstance(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CancelWorkflowInstance 取消流程实例
func (c *WorkflowController) CancelWorkflowInstance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.CancelWorkflowInstance(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetWorkflowTaskList 获取流程任务列表
func (c *WorkflowController) GetWorkflowTaskList(ctx *gin.Context) {
	instanceID, _ := strconv.ParseUint(ctx.Query("instance_id"), 10, 32)
	assigneeID, _ := strconv.ParseUint(ctx.Query("assignee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	tasks, total, err := c.workflowService.GetWorkflowTaskList(uint(instanceID), uint(assigneeID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  tasks,
		"total": total,
	})
}

// GetWorkflowTaskByID 根据ID获取流程任务
func (c *WorkflowController) GetWorkflowTaskByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	task, err := c.workflowService.GetWorkflowTaskByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// CreateWorkflowTask 创建流程任务
func (c *WorkflowController) CreateWorkflowTask(ctx *gin.Context) {
	var task model.WorkflowTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.workflowService.CreateWorkflowTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

// UpdateWorkflowTask 更新流程任务
func (c *WorkflowController) UpdateWorkflowTask(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var task model.WorkflowTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uint(id)
	if err := c.workflowService.UpdateWorkflowTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// DeleteWorkflowTask 删除流程任务
func (c *WorkflowController) DeleteWorkflowTask(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.workflowService.DeleteWorkflowTask(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// HandleWorkflowTask 处理流程任务
func (c *WorkflowController) HandleWorkflowTask(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var data struct {
		Action  int    `json:"action" binding:"required"`
		Comment string `json:"comment"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.workflowService.HandleWorkflowTask(uint(id), data.Action, data.Comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// TransferWorkflowTask 转办流程任务
func (c *WorkflowController) TransferWorkflowTask(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var data struct {
		AssigneeID uint `json:"assignee_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.workflowService.TransferWorkflowTask(uint(id), data.AssigneeID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
