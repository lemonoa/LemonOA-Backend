package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type NoticeController struct {
	noticeService *service.NoticeService
}

func NewNoticeController(noticeService *service.NoticeService) *NoticeController {
	return &NoticeController{
		noticeService: noticeService,
	}
}

// RegisterRoutes 注册路由
func (c *NoticeController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/notices")
	{
		// 公告管理
		api.GET("", c.GetNoticeList)
		api.GET("/:id", c.GetNoticeByID)
		api.POST("", c.CreateNotice)
		api.PUT("/:id", c.UpdateNotice)
		api.DELETE("/:id", c.DeleteNotice)
		api.PUT("/:id/publish", c.PublishNotice)
		api.PUT("/:id/recall", c.RecallNotice)

		// 阅读记录管理
		api.GET("/:id/reads", c.GetNoticeReadList)
		api.PUT("/:id/read", c.ReadNotice)
		api.GET("/unread-count", c.GetUnreadNoticeCount)
	}
}

// GetNoticeList 获取公告列表
func (c *NoticeController) GetNoticeList(ctx *gin.Context) {
	typeID, _ := strconv.ParseUint(ctx.Query("type_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	notices, total, err := c.noticeService.GetNoticeList(uint(typeID), status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  notices,
		"total": total,
	})
}

// GetNoticeByID 根据ID获取公告
func (c *NoticeController) GetNoticeByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	notice, err := c.noticeService.GetNoticeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notice)
}

// CreateNotice 创建公告
func (c *NoticeController) CreateNotice(ctx *gin.Context) {
	var notice model.Notice
	if err := ctx.ShouldBindJSON(&notice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	notice.CreatedBy = uint(1)

	if err := c.noticeService.CreateNotice(&notice); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, notice)
}

// UpdateNotice 更新公告
func (c *NoticeController) UpdateNotice(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var notice model.Notice
	if err := ctx.ShouldBindJSON(&notice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notice.ID = uint(id)
	if err := c.noticeService.UpdateNotice(&notice); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notice)
}

// DeleteNotice 删除公告
func (c *NoticeController) DeleteNotice(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.noticeService.DeleteNotice(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// PublishNotice 发布公告
func (c *NoticeController) PublishNotice(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.noticeService.PublishNotice(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RecallNotice 撤回公告
func (c *NoticeController) RecallNotice(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.noticeService.RecallNotice(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetNoticeReadList 获取公告阅读记录列表
func (c *NoticeController) GetNoticeReadList(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	reads, total, err := c.noticeService.GetNoticeReadList(uint(id), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  reads,
		"total": total,
	})
}

// ReadNotice 阅读公告
func (c *NoticeController) ReadNotice(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	userID := uint(1)

	if err := c.noticeService.ReadNotice(uint(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetUnreadNoticeCount 获取未读公告数量
func (c *NoticeController) GetUnreadNoticeCount(ctx *gin.Context) {
	// TODO: 从JWT中获取当前用户ID
	userID := uint(1)

	count, err := c.noticeService.GetUnreadNoticeCount(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}
