package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type BasicCustomerController struct {
	basicCustomerService *service.BasicCustomerService
}

func NewBasicCustomerController(basicCustomerService *service.BasicCustomerService) *BasicCustomerController {
	return &BasicCustomerController{
		basicCustomerService: basicCustomerService,
	}
}

// RegisterRoutes 注册路由
func (c *BasicCustomerController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/basic/customer")
	{
		// 客户等级
		api.GET("/customer-levels", c.GetCustomerLevelList)
		api.GET("/customer-levels/:id", c.GetCustomerLevelByID)
		api.POST("/customer-levels", c.CreateCustomerLevel)
		api.PUT("/customer-levels/:id", c.UpdateCustomerLevel)
		api.DELETE("/customer-levels/:id", c.DeleteCustomerLevel)

		// 客户渠道
		api.GET("/customer-channels", c.GetCustomerChannelList)
		api.GET("/customer-channels/:id", c.GetCustomerChannelByID)
		api.POST("/customer-channels", c.CreateCustomerChannel)
		api.PUT("/customer-channels/:id", c.UpdateCustomerChannel)
		api.DELETE("/customer-channels/:id", c.DeleteCustomerChannel)

		// 行业类型
		api.GET("/industries", c.GetIndustryList)
		api.GET("/industries/:id", c.GetIndustryByID)
		api.POST("/industries", c.CreateIndustry)
		api.PUT("/industries/:id", c.UpdateIndustry)
		api.DELETE("/industries/:id", c.DeleteIndustry)

		// 客户状态
		api.GET("/customer-statuses", c.GetCustomerStatusList)
		api.GET("/customer-statuses/:id", c.GetCustomerStatusByID)
		api.POST("/customer-statuses", c.CreateCustomerStatus)
		api.PUT("/customer-statuses/:id", c.UpdateCustomerStatus)
		api.DELETE("/customer-statuses/:id", c.DeleteCustomerStatus)

		// 客户意向
		api.GET("/customer-intentions", c.GetCustomerIntentionList)
		api.GET("/customer-intentions/:id", c.GetCustomerIntentionByID)
		api.POST("/customer-intentions", c.CreateCustomerIntention)
		api.PUT("/customer-intentions/:id", c.UpdateCustomerIntention)
		api.DELETE("/customer-intentions/:id", c.DeleteCustomerIntention)

		// 跟进方式
		api.GET("/follow-up-methods", c.GetFollowUpMethodList)
		api.GET("/follow-up-methods/:id", c.GetFollowUpMethodByID)
		api.POST("/follow-up-methods", c.CreateFollowUpMethod)
		api.PUT("/follow-up-methods/:id", c.UpdateFollowUpMethod)
		api.DELETE("/follow-up-methods/:id", c.DeleteFollowUpMethod)

		// 销售阶段
		api.GET("/sales-stages", c.GetSalesStageList)
		api.GET("/sales-stages/:id", c.GetSalesStageByID)
		api.POST("/sales-stages", c.CreateSalesStage)
		api.PUT("/sales-stages/:id", c.UpdateSalesStage)
		api.DELETE("/sales-stages/:id", c.DeleteSalesStage)
	}
}

// GetCustomerLevelList 获取客户等级列表
func (c *BasicCustomerController) GetCustomerLevelList(ctx *gin.Context) {
	levels, err := c.basicCustomerService.GetCustomerLevelList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, levels)
}

// GetCustomerLevelByID 根据ID获取客户等级
func (c *BasicCustomerController) GetCustomerLevelByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	level, err := c.basicCustomerService.GetCustomerLevelByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, level)
}

// CreateCustomerLevel 创建客户等级
func (c *BasicCustomerController) CreateCustomerLevel(ctx *gin.Context) {
	var level model.CustomerLevel
	if err := ctx.ShouldBindJSON(&level); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCustomerService.CreateCustomerLevel(&level); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, level)
}

// UpdateCustomerLevel 更新客户等级
func (c *BasicCustomerController) UpdateCustomerLevel(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var level model.CustomerLevel
	if err := ctx.ShouldBindJSON(&level); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	level.ID = uint(id)
	if err := c.basicCustomerService.UpdateCustomerLevel(&level); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, level)
}

// DeleteCustomerLevel 删除客户等级
func (c *BasicCustomerController) DeleteCustomerLevel(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCustomerService.DeleteCustomerLevel(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetCustomerChannelList 获取客户渠道列表
func (c *BasicCustomerController) GetCustomerChannelList(ctx *gin.Context) {
	channels, err := c.basicCustomerService.GetCustomerChannelList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channels)
}

// GetCustomerChannelByID 根据ID获取客户渠道
func (c *BasicCustomerController) GetCustomerChannelByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	channel, err := c.basicCustomerService.GetCustomerChannelByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channel)
}

// CreateCustomerChannel 创建客户渠道
func (c *BasicCustomerController) CreateCustomerChannel(ctx *gin.Context) {
	var channel model.CustomerChannel
	if err := ctx.ShouldBindJSON(&channel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCustomerService.CreateCustomerChannel(&channel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, channel)
}

// UpdateCustomerChannel 更新客户渠道
func (c *BasicCustomerController) UpdateCustomerChannel(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var channel model.CustomerChannel
	if err := ctx.ShouldBindJSON(&channel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channel.ID = uint(id)
	if err := c.basicCustomerService.UpdateCustomerChannel(&channel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channel)
}

// DeleteCustomerChannel 删除客户渠道
func (c *BasicCustomerController) DeleteCustomerChannel(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCustomerService.DeleteCustomerChannel(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetIndustryList 获取行业类型列表
func (c *BasicCustomerController) GetIndustryList(ctx *gin.Context) {
	var parentID *uint
	if parentIDStr := ctx.Query("parent_id"); parentIDStr != "" {
		if id, err := strconv.ParseUint(parentIDStr, 10, 32); err == nil {
			pid := uint(id)
			parentID = &pid
		}
	}

	industries, err := c.basicCustomerService.GetIndustryList(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, industries)
}

// GetIndustryByID 根据ID获取行业类型
func (c *BasicCustomerController) GetIndustryByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	industry, err := c.basicCustomerService.GetIndustryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, industry)
}

// CreateIndustry 创建行业类型
func (c *BasicCustomerController) CreateIndustry(ctx *gin.Context) {
	var industry model.Industry
	if err := ctx.ShouldBindJSON(&industry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCustomerService.CreateIndustry(&industry); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, industry)
}

// UpdateIndustry 更新行业类型
func (c *BasicCustomerController) UpdateIndustry(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var industry model.Industry
	if err := ctx.ShouldBindJSON(&industry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	industry.ID = uint(id)
	if err := c.basicCustomerService.UpdateIndustry(&industry); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, industry)
}

// DeleteIndustry 删除行业类型
func (c *BasicCustomerController) DeleteIndustry(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCustomerService.DeleteIndustry(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetCustomerStatusList 获取客户状态列表
func (c *BasicCustomerController) GetCustomerStatusList(ctx *gin.Context) {
	statuses, err := c.basicCustomerService.GetCustomerStatusList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, statuses)
}

// GetCustomerStatusByID 根据ID获取客户状态
func (c *BasicCustomerController) GetCustomerStatusByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	status, err := c.basicCustomerService.GetCustomerStatusByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, status)
}

// CreateCustomerStatus 创建客户状态
func (c *BasicCustomerController) CreateCustomerStatus(ctx *gin.Context) {
	var status model.CustomerStatus
	if err := ctx.ShouldBindJSON(&status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCustomerService.CreateCustomerStatus(&status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, status)
}

// UpdateCustomerStatus 更新客户状态
func (c *BasicCustomerController) UpdateCustomerStatus(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var status model.CustomerStatus
	if err := ctx.ShouldBindJSON(&status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status.ID = uint(id)
	if err := c.basicCustomerService.UpdateCustomerStatus(&status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, status)
}

// DeleteCustomerStatus 删除客户状态
func (c *BasicCustomerController) DeleteCustomerStatus(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCustomerService.DeleteCustomerStatus(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetCustomerIntentionList 获取客户意向列表
func (c *BasicCustomerController) GetCustomerIntentionList(ctx *gin.Context) {
	intentions, err := c.basicCustomerService.GetCustomerIntentionList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, intentions)
}

// GetCustomerIntentionByID 根据ID获取客户意向
func (c *BasicCustomerController) GetCustomerIntentionByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	intention, err := c.basicCustomerService.GetCustomerIntentionByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, intention)
}

// CreateCustomerIntention 创建客户意向
func (c *BasicCustomerController) CreateCustomerIntention(ctx *gin.Context) {
	var intention model.CustomerIntention
	if err := ctx.ShouldBindJSON(&intention); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCustomerService.CreateCustomerIntention(&intention); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, intention)
}

// UpdateCustomerIntention 更新客户意向
func (c *BasicCustomerController) UpdateCustomerIntention(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var intention model.CustomerIntention
	if err := ctx.ShouldBindJSON(&intention); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intention.ID = uint(id)
	if err := c.basicCustomerService.UpdateCustomerIntention(&intention); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, intention)
}

// DeleteCustomerIntention 删除客户意向
func (c *BasicCustomerController) DeleteCustomerIntention(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCustomerService.DeleteCustomerIntention(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetFollowUpMethodList 获取跟进方式列表
func (c *BasicCustomerController) GetFollowUpMethodList(ctx *gin.Context) {
	methods, err := c.basicCustomerService.GetFollowUpMethodList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, methods)
}

// GetFollowUpMethodByID 根据ID获取跟进方式
func (c *BasicCustomerController) GetFollowUpMethodByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	method, err := c.basicCustomerService.GetFollowUpMethodByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, method)
}

// CreateFollowUpMethod 创建跟进方式
func (c *BasicCustomerController) CreateFollowUpMethod(ctx *gin.Context) {
	var method model.FollowUpMethod
	if err := ctx.ShouldBindJSON(&method); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCustomerService.CreateFollowUpMethod(&method); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, method)
}

// UpdateFollowUpMethod 更新跟进方式
func (c *BasicCustomerController) UpdateFollowUpMethod(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var method model.FollowUpMethod
	if err := ctx.ShouldBindJSON(&method); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	method.ID = uint(id)
	if err := c.basicCustomerService.UpdateFollowUpMethod(&method); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, method)
}

// DeleteFollowUpMethod 删除跟进方式
func (c *BasicCustomerController) DeleteFollowUpMethod(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCustomerService.DeleteFollowUpMethod(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetSalesStageList 获取销售阶段列表
func (c *BasicCustomerController) GetSalesStageList(ctx *gin.Context) {
	stages, err := c.basicCustomerService.GetSalesStageList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stages)
}

// GetSalesStageByID 根据ID获取销售阶段
func (c *BasicCustomerController) GetSalesStageByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	stage, err := c.basicCustomerService.GetSalesStageByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stage)
}

// CreateSalesStage 创建销售阶段
func (c *BasicCustomerController) CreateSalesStage(ctx *gin.Context) {
	var stage model.SalesStage
	if err := ctx.ShouldBindJSON(&stage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCustomerService.CreateSalesStage(&stage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, stage)
}

// UpdateSalesStage 更新销售阶段
func (c *BasicCustomerController) UpdateSalesStage(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var stage model.SalesStage
	if err := ctx.ShouldBindJSON(&stage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stage.ID = uint(id)
	if err := c.basicCustomerService.UpdateSalesStage(&stage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stage)
}

// DeleteSalesStage 删除销售阶段
func (c *BasicCustomerController) DeleteSalesStage(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCustomerService.DeleteSalesStage(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
