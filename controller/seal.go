package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type SealController struct {
	sealService *service.SealService
}

func NewSealController(sealService *service.SealService) *SealController {
	return &SealController{
		sealService: sealService,
	}
}

// RegisterRoutes 注册路由
func (c *SealController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/seals")
	{
		// 印章管理
		api.GET("", c.GetSealList)
		api.GET("/:id", c.GetSealByID)
		api.POST("", c.CreateSeal)
		api.PUT("/:id", c.UpdateSeal)
		api.DELETE("/:id", c.DeleteSeal)

		// 用印申请管理
		api.GET("/applications", c.GetSealApplicationList)
		api.GET("/applications/:id", c.GetSealApplicationByID)
		api.POST("/applications", c.CreateSealApplication)
		api.PUT("/applications/:id", c.UpdateSealApplication)
		api.DELETE("/applications/:id", c.DeleteSealApplication)
		api.PUT("/applications/:id/approve", c.ApproveSealApplication)
		api.PUT("/applications/:id/reject", c.RejectSealApplication)
		api.PUT("/applications/:id/cancel", c.CancelSealApplication)

		// 用印记录管理
		api.GET("/records", c.GetSealRecordList)
		api.GET("/records/:id", c.GetSealRecordByID)
		api.POST("/records", c.CreateSealRecord)
		api.PUT("/records/:id", c.UpdateSealRecord)
		api.DELETE("/records/:id", c.DeleteSealRecord)
		api.PUT("/records/:id/return", c.ReturnSeal)
	}
}

// GetSealList 获取印章列表
func (c *SealController) GetSealList(ctx *gin.Context) {
	typeID, _ := strconv.ParseUint(ctx.Query("type_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	seals, total, err := c.sealService.GetSealList(uint(typeID), status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  seals,
		"total": total,
	})
}

// GetSealByID 根据ID获取印章
func (c *SealController) GetSealByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	seal, err := c.sealService.GetSealByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seal)
}

// CreateSeal 创建印章
func (c *SealController) CreateSeal(ctx *gin.Context) {
	var seal model.Seal
	if err := ctx.ShouldBindJSON(&seal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	seal.CreatedBy = uint(1)

	if err := c.sealService.CreateSeal(&seal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, seal)
}

// UpdateSeal 更新印章
func (c *SealController) UpdateSeal(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var seal model.Seal
	if err := ctx.ShouldBindJSON(&seal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seal.ID = uint(id)
	if err := c.sealService.UpdateSeal(&seal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seal)
}

// DeleteSeal 删除印章
func (c *SealController) DeleteSeal(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.sealService.DeleteSeal(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetSealApplicationList 获取用印申请列表
func (c *SealController) GetSealApplicationList(ctx *gin.Context) {
	sealID, _ := strconv.ParseUint(ctx.Query("seal_id"), 10, 32)
	userID, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	applications, total, err := c.sealService.GetSealApplicationList(uint(sealID), uint(userID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  applications,
		"total": total,
	})
}

// GetSealApplicationByID 根据ID获取用印申请
func (c *SealController) GetSealApplicationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	application, err := c.sealService.GetSealApplicationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// CreateSealApplication 创建用印申请
func (c *SealController) CreateSealApplication(ctx *gin.Context) {
	var application model.SealApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.sealService.CreateSealApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, application)
}

// UpdateSealApplication 更新用印申请
func (c *SealController) UpdateSealApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var application model.SealApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application.ID = uint(id)
	if err := c.sealService.UpdateSealApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// DeleteSealApplication 删除用印申请
func (c *SealController) DeleteSealApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.sealService.DeleteSealApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveSealApplication 审批通过用印申请
func (c *SealController) ApproveSealApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.sealService.ApproveSealApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectSealApplication 审批驳回用印申请
func (c *SealController) RejectSealApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.sealService.RejectSealApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CancelSealApplication 取消用印申请
func (c *SealController) CancelSealApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.sealService.CancelSealApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetSealRecordList 获取用印记录列表
func (c *SealController) GetSealRecordList(ctx *gin.Context) {
	applicationID, _ := strconv.ParseUint(ctx.Query("application_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	records, total, err := c.sealService.GetSealRecordList(uint(applicationID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
	})
}

// GetSealRecordByID 根据ID获取用印记录
func (c *SealController) GetSealRecordByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	record, err := c.sealService.GetSealRecordByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// CreateSealRecord 创建用印记录
func (c *SealController) CreateSealRecord(ctx *gin.Context) {
	var record model.SealRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	record.CreatedBy = uint(1)

	if err := c.sealService.CreateSealRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

// UpdateSealRecord 更新用印记录
func (c *SealController) UpdateSealRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var record model.SealRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = uint(id)
	if err := c.sealService.UpdateSealRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// DeleteSealRecord 删除用印记录
func (c *SealController) DeleteSealRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.sealService.DeleteSealRecord(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ReturnSeal 归还印章
func (c *SealController) ReturnSeal(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.sealService.ReturnSeal(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
