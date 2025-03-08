package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type BasicHRController struct {
	basicHRService *service.BasicHRService
}

func NewBasicHRController(basicHRService *service.BasicHRService) *BasicHRController {
	return &BasicHRController{
		basicHRService: basicHRService,
	}
}

// RegisterRoutes 注册路由
func (c *BasicHRController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/basic/hr")
	{
		// 奖惩项目
		api.GET("/reward-punishments", c.GetRewardPunishmentList)
		api.GET("/reward-punishments/:id", c.GetRewardPunishmentByID)
		api.POST("/reward-punishments", c.CreateRewardPunishment)
		api.PUT("/reward-punishments/:id", c.UpdateRewardPunishment)
		api.DELETE("/reward-punishments/:id", c.DeleteRewardPunishment)

		// 关怀项目
		api.GET("/care-projects", c.GetCareProjectList)
		api.GET("/care-projects/:id", c.GetCareProjectByID)
		api.POST("/care-projects", c.CreateCareProject)
		api.PUT("/care-projects/:id", c.UpdateCareProject)
		api.DELETE("/care-projects/:id", c.DeleteCareProject)

		// 常规数据
		api.GET("/common-data", c.GetCommonDataList)
		api.GET("/common-data/:id", c.GetCommonDataByID)
		api.GET("/common-data/code/:code", c.GetCommonDataByCode)
		api.POST("/common-data", c.CreateCommonData)
		api.PUT("/common-data/:id", c.UpdateCommonData)
		api.DELETE("/common-data/:id", c.DeleteCommonData)
	}
}

// GetRewardPunishmentList 获取奖惩项目列表
func (c *BasicHRController) GetRewardPunishmentList(ctx *gin.Context) {
	rewardType, _ := strconv.Atoi(ctx.Query("type"))
	items, err := c.basicHRService.GetRewardPunishmentList(rewardType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// GetRewardPunishmentByID 根据ID获取奖惩项目
func (c *BasicHRController) GetRewardPunishmentByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	item, err := c.basicHRService.GetRewardPunishmentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// CreateRewardPunishment 创建奖惩项目
func (c *BasicHRController) CreateRewardPunishment(ctx *gin.Context) {
	var item model.RewardPunishment
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicHRService.CreateRewardPunishment(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

// UpdateRewardPunishment 更新奖惩项目
func (c *BasicHRController) UpdateRewardPunishment(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var item model.RewardPunishment
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	if err := c.basicHRService.UpdateRewardPunishment(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// DeleteRewardPunishment 删除奖惩项目
func (c *BasicHRController) DeleteRewardPunishment(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicHRService.DeleteRewardPunishment(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetCareProjectList 获取关怀项目列表
func (c *BasicHRController) GetCareProjectList(ctx *gin.Context) {
	careType, _ := strconv.Atoi(ctx.Query("type"))
	items, err := c.basicHRService.GetCareProjectList(careType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// GetCareProjectByID 根据ID获取关怀项目
func (c *BasicHRController) GetCareProjectByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	item, err := c.basicHRService.GetCareProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// CreateCareProject 创建关怀项目
func (c *BasicHRController) CreateCareProject(ctx *gin.Context) {
	var item model.CareProject
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicHRService.CreateCareProject(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

// UpdateCareProject 更新关怀项目
func (c *BasicHRController) UpdateCareProject(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var item model.CareProject
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	if err := c.basicHRService.UpdateCareProject(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// DeleteCareProject 删除关怀项目
func (c *BasicHRController) DeleteCareProject(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicHRService.DeleteCareProject(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetCommonDataList 获取常规数据列表
func (c *BasicHRController) GetCommonDataList(ctx *gin.Context) {
	dataType := ctx.Query("type")
	items, err := c.basicHRService.GetCommonDataList(dataType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// GetCommonDataByID 根据ID获取常规数据
func (c *BasicHRController) GetCommonDataByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	item, err := c.basicHRService.GetCommonDataByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// GetCommonDataByCode 根据Code获取常规数据
func (c *BasicHRController) GetCommonDataByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	item, err := c.basicHRService.GetCommonDataByCode(code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// CreateCommonData 创建常规数据
func (c *BasicHRController) CreateCommonData(ctx *gin.Context) {
	var item model.CommonData
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicHRService.CreateCommonData(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

// UpdateCommonData 更新常规数据
func (c *BasicHRController) UpdateCommonData(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var item model.CommonData
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	if err := c.basicHRService.UpdateCommonData(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// DeleteCommonData 删除常规数据
func (c *BasicHRController) DeleteCommonData(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicHRService.DeleteCommonData(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
