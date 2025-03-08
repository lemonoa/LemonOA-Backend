package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type BasicCommonController struct {
	basicCommonService *service.BasicCommonService
}

func NewBasicCommonController(basicCommonService *service.BasicCommonService) *BasicCommonController {
	return &BasicCommonController{
		basicCommonService: basicCommonService,
	}
}

// RegisterRoutes 注册路由
func (c *BasicCommonController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/basic/common")
	{
		// 企业主体
		api.GET("/enterprises", c.GetEnterpriseList)
		api.GET("/enterprises/:id", c.GetEnterpriseByID)
		api.POST("/enterprises", c.CreateEnterprise)
		api.PUT("/enterprises/:id", c.UpdateEnterprise)
		api.DELETE("/enterprises/:id", c.DeleteEnterprise)

		// 地区
		api.GET("/regions", c.GetRegionList)
		api.GET("/regions/:id", c.GetRegionByID)
		api.POST("/regions", c.CreateRegion)
		api.PUT("/regions/:id", c.UpdateRegion)
		api.DELETE("/regions/:id", c.DeleteRegion)

		// 消息模板
		api.GET("/message-templates", c.GetMessageTemplateList)
		api.GET("/message-templates/:id", c.GetMessageTemplateByID)
		api.GET("/message-templates/code/:code", c.GetMessageTemplateByCode)
		api.POST("/message-templates", c.CreateMessageTemplate)
		api.PUT("/message-templates/:id", c.UpdateMessageTemplate)
		api.DELETE("/message-templates/:id", c.DeleteMessageTemplate)
	}
}

// GetEnterpriseList 获取企业主体列表
func (c *BasicCommonController) GetEnterpriseList(ctx *gin.Context) {
	enterprises, err := c.basicCommonService.GetEnterpriseList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, enterprises)
}

// GetEnterpriseByID 根据ID获取企业主体
func (c *BasicCommonController) GetEnterpriseByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	enterprise, err := c.basicCommonService.GetEnterpriseByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, enterprise)
}

// CreateEnterprise 创建企业主体
func (c *BasicCommonController) CreateEnterprise(ctx *gin.Context) {
	var enterprise model.Enterprise
	if err := ctx.ShouldBindJSON(&enterprise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCommonService.CreateEnterprise(&enterprise); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, enterprise)
}

// UpdateEnterprise 更新企业主体
func (c *BasicCommonController) UpdateEnterprise(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var enterprise model.Enterprise
	if err := ctx.ShouldBindJSON(&enterprise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enterprise.ID = uint(id)
	if err := c.basicCommonService.UpdateEnterprise(&enterprise); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, enterprise)
}

// DeleteEnterprise 删除企业主体
func (c *BasicCommonController) DeleteEnterprise(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCommonService.DeleteEnterprise(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetRegionList 获取地区列表
func (c *BasicCommonController) GetRegionList(ctx *gin.Context) {
	var parentID *uint
	if parentIDStr := ctx.Query("parent_id"); parentIDStr != "" {
		if id, err := strconv.ParseUint(parentIDStr, 10, 32); err == nil {
			pid := uint(id)
			parentID = &pid
		}
	}

	regions, err := c.basicCommonService.GetRegionList(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, regions)
}

// GetRegionByID 根据ID获取地区
func (c *BasicCommonController) GetRegionByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	region, err := c.basicCommonService.GetRegionByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, region)
}

// CreateRegion 创建地区
func (c *BasicCommonController) CreateRegion(ctx *gin.Context) {
	var region model.Region
	if err := ctx.ShouldBindJSON(&region); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCommonService.CreateRegion(&region); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, region)
}

// UpdateRegion 更新地区
func (c *BasicCommonController) UpdateRegion(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var region model.Region
	if err := ctx.ShouldBindJSON(&region); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	region.ID = uint(id)
	if err := c.basicCommonService.UpdateRegion(&region); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, region)
}

// DeleteRegion 删除地区
func (c *BasicCommonController) DeleteRegion(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCommonService.DeleteRegion(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetMessageTemplateList 获取消息模板列表
func (c *BasicCommonController) GetMessageTemplateList(ctx *gin.Context) {
	templateType, _ := strconv.Atoi(ctx.Query("type"))
	templates, err := c.basicCommonService.GetMessageTemplateList(templateType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, templates)
}

// GetMessageTemplateByID 根据ID获取消息模板
func (c *BasicCommonController) GetMessageTemplateByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	template, err := c.basicCommonService.GetMessageTemplateByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, template)
}

// GetMessageTemplateByCode 根据Code获取消息模板
func (c *BasicCommonController) GetMessageTemplateByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	template, err := c.basicCommonService.GetMessageTemplateByCode(code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, template)
}

// CreateMessageTemplate 创建消息模板
func (c *BasicCommonController) CreateMessageTemplate(ctx *gin.Context) {
	var template model.MessageTemplate
	if err := ctx.ShouldBindJSON(&template); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicCommonService.CreateMessageTemplate(&template); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, template)
}

// UpdateMessageTemplate 更新消息模板
func (c *BasicCommonController) UpdateMessageTemplate(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var template model.MessageTemplate
	if err := ctx.ShouldBindJSON(&template); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	template.ID = uint(id)
	if err := c.basicCommonService.UpdateMessageTemplate(&template); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, template)
}

// DeleteMessageTemplate 删除消息模板
func (c *BasicCommonController) DeleteMessageTemplate(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicCommonService.DeleteMessageTemplate(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
