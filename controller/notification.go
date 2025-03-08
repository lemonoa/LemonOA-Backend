package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationService *service.NotificationService
}

func NewNotificationController(notificationService *service.NotificationService) *NotificationController {
	return &NotificationController{
		notificationService: notificationService,
	}
}

// RegisterRoutes 注册路由
func (c *NotificationController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/notifications")
	{
		api.GET("", c.GetNotificationList)
		api.GET("/unread-count", c.GetUnreadCount)
		api.POST("", c.CreateNotification)
		api.PUT("/:id/read", c.MarkAsRead)
		api.PUT("/read-all", c.MarkAllAsRead)
		api.DELETE("/:id", c.DeleteNotification)
	}
}

// GetNotificationList 获取消息列表
func (c *NotificationController) GetNotificationList(ctx *gin.Context) {
	// TODO: 从JWT中获取userID
	userID := uint(1)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	notifications, total, err := c.notificationService.GetNotificationList(userID, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  notifications,
		"total": total,
	})
}

// GetUnreadCount 获取未读消息数量
func (c *NotificationController) GetUnreadCount(ctx *gin.Context) {
	// TODO: 从JWT中获取userID
	userID := uint(1)

	count, err := c.notificationService.GetUnreadCount(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}

// CreateNotification 创建消息
func (c *NotificationController) CreateNotification(ctx *gin.Context) {
	var notification model.Notification
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.notificationService.CreateNotification(&notification); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, notification)
}

// MarkAsRead 标记消息为已读
func (c *NotificationController) MarkAsRead(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取userID
	userID := uint(1)

	if err := c.notificationService.MarkAsRead(uint(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// MarkAllAsRead 标记所有消息为已读
func (c *NotificationController) MarkAllAsRead(ctx *gin.Context) {
	// TODO: 从JWT中获取userID
	userID := uint(1)

	if err := c.notificationService.MarkAllAsRead(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DeleteNotification 删除消息
func (c *NotificationController) DeleteNotification(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取userID
	userID := uint(1)

	if err := c.notificationService.DeleteNotification(uint(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
