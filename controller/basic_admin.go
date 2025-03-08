package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type BasicAdminController struct {
	basicAdminService *service.BasicAdminService
}

func NewBasicAdminController(basicAdminService *service.BasicAdminService) *BasicAdminController {
	return &BasicAdminController{
		basicAdminService: basicAdminService,
	}
}

// RegisterRoutes 注册路由
func (c *BasicAdminController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/basic/admin")
	{
		// 资产分类
		api.GET("/asset-categories", c.GetAssetCategoryList)
		api.GET("/asset-categories/:id", c.GetAssetCategoryByID)
		api.POST("/asset-categories", c.CreateAssetCategory)
		api.PUT("/asset-categories/:id", c.UpdateAssetCategory)
		api.DELETE("/asset-categories/:id", c.DeleteAssetCategory)

		// 资产品牌
		api.GET("/asset-brands", c.GetAssetBrandList)
		api.GET("/asset-brands/:id", c.GetAssetBrandByID)
		api.POST("/asset-brands", c.CreateAssetBrand)
		api.PUT("/asset-brands/:id", c.UpdateAssetBrand)
		api.DELETE("/asset-brands/:id", c.DeleteAssetBrand)

		// 资产单位
		api.GET("/asset-units", c.GetAssetUnitList)
		api.GET("/asset-units/:id", c.GetAssetUnitByID)
		api.POST("/asset-units", c.CreateAssetUnit)
		api.PUT("/asset-units/:id", c.UpdateAssetUnit)
		api.DELETE("/asset-units/:id", c.DeleteAssetUnit)

		// 印章类型
		api.GET("/seal-types", c.GetSealTypeList)
		api.GET("/seal-types/:id", c.GetSealTypeByID)
		api.POST("/seal-types", c.CreateSealType)
		api.PUT("/seal-types/:id", c.UpdateSealType)
		api.DELETE("/seal-types/:id", c.DeleteSealType)

		// 车辆费用
		api.GET("/vehicle-expenses", c.GetVehicleExpenseList)
		api.GET("/vehicle-expenses/:id", c.GetVehicleExpenseByID)
		api.POST("/vehicle-expenses", c.CreateVehicleExpense)
		api.PUT("/vehicle-expenses/:id", c.UpdateVehicleExpense)
		api.DELETE("/vehicle-expenses/:id", c.DeleteVehicleExpense)

		// 公告类型
		api.GET("/notice-types", c.GetNoticeTypeList)
		api.GET("/notice-types/:id", c.GetNoticeTypeByID)
		api.POST("/notice-types", c.CreateNoticeType)
		api.PUT("/notice-types/:id", c.UpdateNoticeType)
		api.DELETE("/notice-types/:id", c.DeleteNoticeType)
	}
}

// GetAssetCategoryList 获取资产分类列表
func (c *BasicAdminController) GetAssetCategoryList(ctx *gin.Context) {
	var parentID *uint
	if parentIDStr := ctx.Query("parent_id"); parentIDStr != "" {
		if id, err := strconv.ParseUint(parentIDStr, 10, 32); err == nil {
			pid := uint(id)
			parentID = &pid
		}
	}

	categories, err := c.basicAdminService.GetAssetCategoryList(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// GetAssetCategoryByID 根据ID获取资产分类
func (c *BasicAdminController) GetAssetCategoryByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	category, err := c.basicAdminService.GetAssetCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// CreateAssetCategory 创建资产分类
func (c *BasicAdminController) CreateAssetCategory(ctx *gin.Context) {
	var category model.AssetCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicAdminService.CreateAssetCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

// UpdateAssetCategory 更新资产分类
func (c *BasicAdminController) UpdateAssetCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var category model.AssetCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = uint(id)
	if err := c.basicAdminService.UpdateAssetCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// DeleteAssetCategory 删除资产分类
func (c *BasicAdminController) DeleteAssetCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicAdminService.DeleteAssetCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetAssetBrandList 获取资产品牌列表
func (c *BasicAdminController) GetAssetBrandList(ctx *gin.Context) {
	brands, err := c.basicAdminService.GetAssetBrandList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, brands)
}

// GetAssetBrandByID 根据ID获取资产品牌
func (c *BasicAdminController) GetAssetBrandByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	brand, err := c.basicAdminService.GetAssetBrandByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, brand)
}

// CreateAssetBrand 创建资产品牌
func (c *BasicAdminController) CreateAssetBrand(ctx *gin.Context) {
	var brand model.AssetBrand
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicAdminService.CreateAssetBrand(&brand); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, brand)
}

// UpdateAssetBrand 更新资产品牌
func (c *BasicAdminController) UpdateAssetBrand(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var brand model.AssetBrand
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand.ID = uint(id)
	if err := c.basicAdminService.UpdateAssetBrand(&brand); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, brand)
}

// DeleteAssetBrand 删除资产品牌
func (c *BasicAdminController) DeleteAssetBrand(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicAdminService.DeleteAssetBrand(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetAssetUnitList 获取资产单位列表
func (c *BasicAdminController) GetAssetUnitList(ctx *gin.Context) {
	units, err := c.basicAdminService.GetAssetUnitList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, units)
}

// GetAssetUnitByID 根据ID获取资产单位
func (c *BasicAdminController) GetAssetUnitByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	unit, err := c.basicAdminService.GetAssetUnitByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, unit)
}

// CreateAssetUnit 创建资产单位
func (c *BasicAdminController) CreateAssetUnit(ctx *gin.Context) {
	var unit model.AssetUnit
	if err := ctx.ShouldBindJSON(&unit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicAdminService.CreateAssetUnit(&unit); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, unit)
}

// UpdateAssetUnit 更新资产单位
func (c *BasicAdminController) UpdateAssetUnit(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var unit model.AssetUnit
	if err := ctx.ShouldBindJSON(&unit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unit.ID = uint(id)
	if err := c.basicAdminService.UpdateAssetUnit(&unit); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, unit)
}

// DeleteAssetUnit 删除资产单位
func (c *BasicAdminController) DeleteAssetUnit(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicAdminService.DeleteAssetUnit(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetSealTypeList 获取印章类型列表
func (c *BasicAdminController) GetSealTypeList(ctx *gin.Context) {
	types, err := c.basicAdminService.GetSealTypeList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types)
}

// GetSealTypeByID 根据ID获取印章类型
func (c *BasicAdminController) GetSealTypeByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	sealType, err := c.basicAdminService.GetSealTypeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sealType)
}

// CreateSealType 创建印章类型
func (c *BasicAdminController) CreateSealType(ctx *gin.Context) {
	var sealType model.SealType
	if err := ctx.ShouldBindJSON(&sealType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicAdminService.CreateSealType(&sealType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, sealType)
}

// UpdateSealType 更新印章类型
func (c *BasicAdminController) UpdateSealType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var sealType model.SealType
	if err := ctx.ShouldBindJSON(&sealType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sealType.ID = uint(id)
	if err := c.basicAdminService.UpdateSealType(&sealType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sealType)
}

// DeleteSealType 删除印章类型
func (c *BasicAdminController) DeleteSealType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicAdminService.DeleteSealType(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleExpenseList 获取车辆费用列表
func (c *BasicAdminController) GetVehicleExpenseList(ctx *gin.Context) {
	expenses, err := c.basicAdminService.GetVehicleExpenseList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expenses)
}

// GetVehicleExpenseByID 根据ID获取车辆费用
func (c *BasicAdminController) GetVehicleExpenseByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	expense, err := c.basicAdminService.GetVehicleExpenseByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

// CreateVehicleExpense 创建车辆费用
func (c *BasicAdminController) CreateVehicleExpense(ctx *gin.Context) {
	var expense model.VehicleExpense
	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicAdminService.CreateVehicleExpense(&expense); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, expense)
}

// UpdateVehicleExpense 更新车辆费用
func (c *BasicAdminController) UpdateVehicleExpense(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var expense model.VehicleExpense
	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense.ID = uint(id)
	if err := c.basicAdminService.UpdateVehicleExpense(&expense); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

// DeleteVehicleExpense 删除车辆费用
func (c *BasicAdminController) DeleteVehicleExpense(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicAdminService.DeleteVehicleExpense(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetNoticeTypeList 获取公告类型列表
func (c *BasicAdminController) GetNoticeTypeList(ctx *gin.Context) {
	types, err := c.basicAdminService.GetNoticeTypeList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types)
}

// GetNoticeTypeByID 根据ID获取公告类型
func (c *BasicAdminController) GetNoticeTypeByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	noticeType, err := c.basicAdminService.GetNoticeTypeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, noticeType)
}

// CreateNoticeType 创建公告类型
func (c *BasicAdminController) CreateNoticeType(ctx *gin.Context) {
	var noticeType model.NoticeType
	if err := ctx.ShouldBindJSON(&noticeType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicAdminService.CreateNoticeType(&noticeType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, noticeType)
}

// UpdateNoticeType 更新公告类型
func (c *BasicAdminController) UpdateNoticeType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var noticeType model.NoticeType
	if err := ctx.ShouldBindJSON(&noticeType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	noticeType.ID = uint(id)
	if err := c.basicAdminService.UpdateNoticeType(&noticeType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, noticeType)
}

// DeleteNoticeType 删除公告类型
func (c *BasicAdminController) DeleteNoticeType(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicAdminService.DeleteNoticeType(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
