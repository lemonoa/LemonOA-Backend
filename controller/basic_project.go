package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type BasicProjectController struct {
	basicProjectService *service.BasicProjectService
}

func NewBasicProjectController(basicProjectService *service.BasicProjectService) *BasicProjectController {
	return &BasicProjectController{
		basicProjectService: basicProjectService,
	}
}

// RegisterRoutes 注册路由
func (c *BasicProjectController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/basic/project")
	{
		// 项目阶段
		api.GET("/project-stages", c.GetProjectStageList)
		api.GET("/project-stages/:id", c.GetProjectStageByID)
		api.POST("/project-stages", c.CreateProjectStage)
		api.PUT("/project-stages/:id", c.UpdateProjectStage)
		api.DELETE("/project-stages/:id", c.DeleteProjectStage)

		// 项目分类
		api.GET("/project-categories", c.GetProjectCategoryList)
		api.GET("/project-categories/:id", c.GetProjectCategoryByID)
		api.POST("/project-categories", c.CreateProjectCategory)
		api.PUT("/project-categories/:id", c.UpdateProjectCategory)
		api.DELETE("/project-categories/:id", c.DeleteProjectCategory)

		// 工作类型
		api.GET("/work-types", c.GetWorkTypeList)
		api.GET("/work-types/:id", c.GetWorkTypeByID)
		api.POST("/work-types", c.CreateWorkType)
		api.PUT("/work-types/:id", c.UpdateWorkType)
		api.DELETE("/work-types/:id", c.DeleteWorkType)
	}
}

// GetProjectStageList 获取项目阶段列表
func (c *BasicProjectController) GetProjectStageList(ctx *gin.Context) {
	stages, err := c.basicProjectService.GetProjectStageList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stages)
}

// GetProjectStageByID 根据ID获取项目阶段
func (c *BasicProjectController) GetProjectStageByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	stage, err := c.basicProjectService.GetProjectStageByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stage)
}

// CreateProjectStage 创建项目阶段
func (c *BasicProjectController) CreateProjectStage(ctx *gin.Context) {
	var stage model.ProjectStage
	if err := ctx.ShouldBindJSON(&stage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicProjectService.CreateProjectStage(&stage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, stage)
}

// UpdateProjectStage 更新项目阶段
func (c *BasicProjectController) UpdateProjectStage(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var stage model.ProjectStage
	if err := ctx.ShouldBindJSON(&stage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stage.ID = uint(id)
	if err := c.basicProjectService.UpdateProjectStage(&stage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stage)
}

// DeleteProjectStage 删除项目阶段
func (c *BasicProjectController) DeleteProjectStage(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicProjectService.DeleteProjectStage(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetProjectCategoryList 获取项目分类列表
func (c *BasicProjectController) GetProjectCategoryList(ctx *gin.Context) {
	categories, err := c.basicProjectService.GetProjectCategoryList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// GetProjectCategoryByID 根据ID获取项目分类
func (c *BasicProjectController) GetProjectCategoryByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	category, err := c.basicProjectService.GetProjectCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// CreateProjectCategory 创建项目分类
func (c *BasicProjectController) CreateProjectCategory(ctx *gin.Context) {
	var category model.ProjectCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicProjectService.CreateProjectCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

// UpdateProjectCategory 更新项目分类
func (c *BasicProjectController) UpdateProjectCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var category model.ProjectCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = uint(id)
	if err := c.basicProjectService.UpdateProjectCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// DeleteProjectCategory 删除项目分类
func (c *BasicProjectController) DeleteProjectCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicProjectService.DeleteProjectCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetWorkTypeList 获取工作类型列表
func (c *BasicProjectController) GetWorkTypeList(ctx *gin.Context) {
	types, err := c.basicProjectService.GetWorkTypeList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types)
}

// GetWorkTypeByID 根据ID获取工作类型
func (c *BasicProjectController) GetWorkTypeByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	workType, err := c.basicProjectService.GetWorkTypeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workType)
}

// CreateWorkType 创建工作类型
func (c *BasicProjectController) CreateWorkType(ctx *gin.Context) {
	var workType model.WorkType
	if err := ctx.ShouldBindJSON(&workType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicProjectService.CreateWorkType(&workType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, workType)
}

// UpdateWorkType 更新工作类型
func (c *BasicProjectController) UpdateWorkType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var workType model.WorkType
	if err := ctx.ShouldBindJSON(&workType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workType.ID = uint(id)
	if err := c.basicProjectService.UpdateWorkType(&workType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, workType)
}

// DeleteWorkType 删除工作类型
func (c *BasicProjectController) DeleteWorkType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicProjectService.DeleteWorkType(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
