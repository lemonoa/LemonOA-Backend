package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type AssetController struct {
	assetService *service.AssetService
}

func NewAssetController(assetService *service.AssetService) *AssetController {
	return &AssetController{
		assetService: assetService,
	}
}

// RegisterRoutes 注册路由
func (c *AssetController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/assets")
	{
		// 资产管理
		api.GET("", c.GetAssetList)
		api.GET("/:id", c.GetAssetByID)
		api.POST("", c.CreateAsset)
		api.PUT("/:id", c.UpdateAsset)
		api.DELETE("/:id", c.DeleteAsset)

		// 维修记录管理
		api.GET("/repairs", c.GetAssetRepairList)
		api.GET("/repairs/:id", c.GetAssetRepairByID)
		api.POST("/repairs", c.CreateAssetRepair)
		api.PUT("/repairs/:id", c.UpdateAssetRepair)
		api.DELETE("/repairs/:id", c.DeleteAssetRepair)
		api.PUT("/repairs/:id/complete", c.CompleteAssetRepair)

		// 领用记录管理
		api.GET("/borrows", c.GetAssetBorrowList)
		api.GET("/borrows/:id", c.GetAssetBorrowByID)
		api.POST("/borrows", c.CreateAssetBorrow)
		api.PUT("/borrows/:id", c.UpdateAssetBorrow)
		api.DELETE("/borrows/:id", c.DeleteAssetBorrow)
		api.PUT("/borrows/:id/return", c.ReturnAsset)

		// 报废记录管理
		api.GET("/disposals", c.GetAssetDisposalList)
		api.GET("/disposals/:id", c.GetAssetDisposalByID)
		api.POST("/disposals", c.CreateAssetDisposal)
		api.PUT("/disposals/:id", c.UpdateAssetDisposal)
		api.DELETE("/disposals/:id", c.DeleteAssetDisposal)
		api.PUT("/disposals/:id/approve", c.ApproveAssetDisposal)
		api.PUT("/disposals/:id/reject", c.RejectAssetDisposal)
	}
}

// GetAssetList 获取资产列表
func (c *AssetController) GetAssetList(ctx *gin.Context) {
	categoryID, _ := strconv.ParseUint(ctx.Query("category_id"), 10, 32)
	brandID, _ := strconv.ParseUint(ctx.Query("brand_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	assets, total, err := c.assetService.GetAssetList(uint(categoryID), uint(brandID), status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  assets,
		"total": total,
	})
}

// GetAssetByID 根据ID获取资产
func (c *AssetController) GetAssetByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	asset, err := c.assetService.GetAssetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, asset)
}

// CreateAsset 创建资产
func (c *AssetController) CreateAsset(ctx *gin.Context) {
	var asset model.Asset
	if err := ctx.ShouldBindJSON(&asset); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	asset.CreatedBy = uint(1)

	if err := c.assetService.CreateAsset(&asset); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, asset)
}

// UpdateAsset 更新资产
func (c *AssetController) UpdateAsset(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var asset model.Asset
	if err := ctx.ShouldBindJSON(&asset); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asset.ID = uint(id)
	if err := c.assetService.UpdateAsset(&asset); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, asset)
}

// DeleteAsset 删除资产
func (c *AssetController) DeleteAsset(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.DeleteAsset(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetAssetRepairList 获取资产维修记录列表
func (c *AssetController) GetAssetRepairList(ctx *gin.Context) {
	assetID, _ := strconv.ParseUint(ctx.Query("asset_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	repairs, total, err := c.assetService.GetAssetRepairList(uint(assetID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  repairs,
		"total": total,
	})
}

// GetAssetRepairByID 根据ID获取资产维修记录
func (c *AssetController) GetAssetRepairByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	repair, err := c.assetService.GetAssetRepairByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, repair)
}

// CreateAssetRepair 创建资产维修记录
func (c *AssetController) CreateAssetRepair(ctx *gin.Context) {
	var repair model.AssetRepair
	if err := ctx.ShouldBindJSON(&repair); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	repair.CreatedBy = uint(1)

	if err := c.assetService.CreateAssetRepair(&repair); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, repair)
}

// UpdateAssetRepair 更新资产维修记录
func (c *AssetController) UpdateAssetRepair(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var repair model.AssetRepair
	if err := ctx.ShouldBindJSON(&repair); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repair.ID = uint(id)
	if err := c.assetService.UpdateAssetRepair(&repair); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, repair)
}

// DeleteAssetRepair 删除资产维修记录
func (c *AssetController) DeleteAssetRepair(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.DeleteAssetRepair(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CompleteAssetRepair 完成资产维修
func (c *AssetController) CompleteAssetRepair(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.CompleteAssetRepair(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetAssetBorrowList 获取资产领用记录列表
func (c *AssetController) GetAssetBorrowList(ctx *gin.Context) {
	assetID, _ := strconv.ParseUint(ctx.Query("asset_id"), 10, 32)
	borrowerID, _ := strconv.ParseUint(ctx.Query("borrower_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	borrows, total, err := c.assetService.GetAssetBorrowList(uint(assetID), uint(borrowerID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  borrows,
		"total": total,
	})
}

// GetAssetBorrowByID 根据ID获取资产领用记录
func (c *AssetController) GetAssetBorrowByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	borrow, err := c.assetService.GetAssetBorrowByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, borrow)
}

// CreateAssetBorrow 创建资产领用记录
func (c *AssetController) CreateAssetBorrow(ctx *gin.Context) {
	var borrow model.AssetBorrow
	if err := ctx.ShouldBindJSON(&borrow); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	borrow.CreatedBy = uint(1)

	if err := c.assetService.CreateAssetBorrow(&borrow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, borrow)
}

// UpdateAssetBorrow 更新资产领用记录
func (c *AssetController) UpdateAssetBorrow(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var borrow model.AssetBorrow
	if err := ctx.ShouldBindJSON(&borrow); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrow.ID = uint(id)
	if err := c.assetService.UpdateAssetBorrow(&borrow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, borrow)
}

// DeleteAssetBorrow 删除资产领用记录
func (c *AssetController) DeleteAssetBorrow(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.DeleteAssetBorrow(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ReturnAsset 归还资产
func (c *AssetController) ReturnAsset(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.ReturnAsset(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetAssetDisposalList 获取资产报废记录列表
func (c *AssetController) GetAssetDisposalList(ctx *gin.Context) {
	assetID, _ := strconv.ParseUint(ctx.Query("asset_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	disposals, total, err := c.assetService.GetAssetDisposalList(uint(assetID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  disposals,
		"total": total,
	})
}

// GetAssetDisposalByID 根据ID获取资产报废记录
func (c *AssetController) GetAssetDisposalByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	disposal, err := c.assetService.GetAssetDisposalByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, disposal)
}

// CreateAssetDisposal 创建资产报废记录
func (c *AssetController) CreateAssetDisposal(ctx *gin.Context) {
	var disposal model.AssetDisposal
	if err := ctx.ShouldBindJSON(&disposal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	disposal.CreatedBy = uint(1)

	if err := c.assetService.CreateAssetDisposal(&disposal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, disposal)
}

// UpdateAssetDisposal 更新资产报废记录
func (c *AssetController) UpdateAssetDisposal(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var disposal model.AssetDisposal
	if err := ctx.ShouldBindJSON(&disposal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	disposal.ID = uint(id)
	if err := c.assetService.UpdateAssetDisposal(&disposal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, disposal)
}

// DeleteAssetDisposal 删除资产报废记录
func (c *AssetController) DeleteAssetDisposal(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.DeleteAssetDisposal(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveAssetDisposal 审批通过资产报废
func (c *AssetController) ApproveAssetDisposal(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.ApproveAssetDisposal(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectAssetDisposal 驳回资产报废
func (c *AssetController) RejectAssetDisposal(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.assetService.RejectAssetDisposal(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
