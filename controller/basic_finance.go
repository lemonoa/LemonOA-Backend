package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type BasicFinanceController struct {
	basicFinanceService *service.BasicFinanceService
}

func NewBasicFinanceController(basicFinanceService *service.BasicFinanceService) *BasicFinanceController {
	return &BasicFinanceController{
		basicFinanceService: basicFinanceService,
	}
}

// RegisterRoutes 注册路由
func (c *BasicFinanceController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/basic/finance")
	{
		// 费用类型
		api.GET("/expense-types", c.GetExpenseTypeList)
		api.GET("/expense-types/:id", c.GetExpenseTypeByID)
		api.POST("/expense-types", c.CreateExpenseType)
		api.PUT("/expense-types/:id", c.UpdateExpenseType)
		api.DELETE("/expense-types/:id", c.DeleteExpenseType)
	}
}

// GetExpenseTypeList 获取费用类型列表
func (c *BasicFinanceController) GetExpenseTypeList(ctx *gin.Context) {
	var parentID *uint
	if parentIDStr := ctx.Query("parent_id"); parentIDStr != "" {
		if id, err := strconv.ParseUint(parentIDStr, 10, 32); err == nil {
			pid := uint(id)
			parentID = &pid
		}
	}

	types, err := c.basicFinanceService.GetExpenseTypeList(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types)
}

// GetExpenseTypeByID 根据ID获取费用类型
func (c *BasicFinanceController) GetExpenseTypeByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	expenseType, err := c.basicFinanceService.GetExpenseTypeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expenseType)
}

// CreateExpenseType 创建费用类型
func (c *BasicFinanceController) CreateExpenseType(ctx *gin.Context) {
	var expenseType model.ExpenseType
	if err := ctx.ShouldBindJSON(&expenseType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicFinanceService.CreateExpenseType(&expenseType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, expenseType)
}

// UpdateExpenseType 更新费用类型
func (c *BasicFinanceController) UpdateExpenseType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var expenseType model.ExpenseType
	if err := ctx.ShouldBindJSON(&expenseType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenseType.ID = uint(id)
	if err := c.basicFinanceService.UpdateExpenseType(&expenseType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expenseType)
}

// DeleteExpenseType 删除费用类型
func (c *BasicFinanceController) DeleteExpenseType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicFinanceService.DeleteExpenseType(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
