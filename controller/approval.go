package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type ApprovalController struct {
	approvalService *service.ApprovalService
}

func NewApprovalController(approvalService *service.ApprovalService) *ApprovalController {
	return &ApprovalController{
		approvalService: approvalService,
	}
}

// RegisterRoutes 注册路由
func (c *ApprovalController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/approvals")
	{
		// 审批类型管理
		api.GET("/types", c.GetApprovalTypeList)
		api.POST("/types", c.CreateApprovalType)
		api.PUT("/types/:id", c.UpdateApprovalType)
		api.DELETE("/types/:id", c.DeleteApprovalType)

		// 审批流程管理
		api.GET("/flows", c.GetApprovalFlowList)
		api.POST("/flows", c.CreateApprovalFlow)
		api.PUT("/flows/:id", c.UpdateApprovalFlow)
		api.DELETE("/flows/:id", c.DeleteApprovalFlow)

		// 审批节点管理
		api.GET("/nodes", c.GetApprovalNodeList)
		api.POST("/nodes", c.CreateApprovalNode)
		api.PUT("/nodes/:id", c.UpdateApprovalNode)
		api.DELETE("/nodes/:id", c.DeleteApprovalNode)

		// 审批记录管理
		api.GET("/records", c.GetApprovalRecordList)
		api.POST("/records", c.CreateApprovalRecord)
		api.PUT("/records/:id/approve", c.ApproveRecord)
		api.PUT("/records/:id/reject", c.RejectRecord)

		// 待审批列表
		api.GET("/pending", c.GetPendingApprovalList)
	}
}

// GetApprovalTypeList 获取审批类型列表
func (c *ApprovalController) GetApprovalTypeList(ctx *gin.Context) {
	types, err := c.approvalService.GetApprovalTypeList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types)
}

// CreateApprovalType 创建审批类型
func (c *ApprovalController) CreateApprovalType(ctx *gin.Context) {
	var approvalType model.ApprovalType
	if err := ctx.ShouldBindJSON(&approvalType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.approvalService.CreateApprovalType(&approvalType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, approvalType)
}

// UpdateApprovalType 更新审批类型
func (c *ApprovalController) UpdateApprovalType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var approvalType model.ApprovalType
	if err := ctx.ShouldBindJSON(&approvalType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	approvalType.ID = uint(id)
	if err := c.approvalService.UpdateApprovalType(&approvalType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, approvalType)
}

// DeleteApprovalType 删除审批类型
func (c *ApprovalController) DeleteApprovalType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.approvalService.DeleteApprovalType(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetApprovalFlowList 获取审批流程列表
func (c *ApprovalController) GetApprovalFlowList(ctx *gin.Context) {
	typeID, _ := strconv.ParseUint(ctx.Query("type_id"), 10, 32)
	flows, err := c.approvalService.GetApprovalFlowList(uint(typeID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, flows)
}

// CreateApprovalFlow 创建审批流程
func (c *ApprovalController) CreateApprovalFlow(ctx *gin.Context) {
	var flow model.ApprovalFlow
	if err := ctx.ShouldBindJSON(&flow); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.approvalService.CreateApprovalFlow(&flow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, flow)
}

// UpdateApprovalFlow 更新审批流程
func (c *ApprovalController) UpdateApprovalFlow(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var flow model.ApprovalFlow
	if err := ctx.ShouldBindJSON(&flow); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flow.ID = uint(id)
	if err := c.approvalService.UpdateApprovalFlow(&flow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, flow)
}

// DeleteApprovalFlow 删除审批流程
func (c *ApprovalController) DeleteApprovalFlow(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.approvalService.DeleteApprovalFlow(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetApprovalNodeList 获取审批节点列表
func (c *ApprovalController) GetApprovalNodeList(ctx *gin.Context) {
	flowID, _ := strconv.ParseUint(ctx.Query("flow_id"), 10, 32)
	nodes, err := c.approvalService.GetApprovalNodeList(uint(flowID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nodes)
}

// CreateApprovalNode 创建审批节点
func (c *ApprovalController) CreateApprovalNode(ctx *gin.Context) {
	var node model.ApprovalNode
	if err := ctx.ShouldBindJSON(&node); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.approvalService.CreateApprovalNode(&node); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, node)
}

// UpdateApprovalNode 更新审批节点
func (c *ApprovalController) UpdateApprovalNode(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var node model.ApprovalNode
	if err := ctx.ShouldBindJSON(&node); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node.ID = uint(id)
	if err := c.approvalService.UpdateApprovalNode(&node); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, node)
}

// DeleteApprovalNode 删除审批节点
func (c *ApprovalController) DeleteApprovalNode(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.approvalService.DeleteApprovalNode(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetApprovalRecordList 获取审批记录列表
func (c *ApprovalController) GetApprovalRecordList(ctx *gin.Context) {
	// TODO: 从JWT中获取userID
	userID := uint(1)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	records, total, err := c.approvalService.GetApprovalRecordList(userID, status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
	})
}

// CreateApprovalRecord 创建审批记录
func (c *ApprovalController) CreateApprovalRecord(ctx *gin.Context) {
	var record model.ApprovalRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取userID
	record.ApplicantID = uint(1)

	if err := c.approvalService.CreateApprovalRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

// ApproveRecord 审批通过
func (c *ApprovalController) ApproveRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		NodeID  uint   `json:"node_id" binding:"required"`
		Comment string `json:"comment"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取userID
	approverID := uint(1)

	if err := c.approvalService.ApproveRecord(uint(id), req.NodeID, approverID, req.Comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectRecord 审批驳回
func (c *ApprovalController) RejectRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		NodeID  uint   `json:"node_id" binding:"required"`
		Comment string `json:"comment" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取userID
	approverID := uint(1)

	if err := c.approvalService.RejectRecord(uint(id), req.NodeID, approverID, req.Comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetPendingApprovalList 获取待审批列表
func (c *ApprovalController) GetPendingApprovalList(ctx *gin.Context) {
	// TODO: 从JWT中获取userID
	approverID := uint(1)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	records, total, err := c.approvalService.GetPendingApprovalList(approverID, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
	})
}
