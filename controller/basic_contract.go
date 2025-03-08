package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type BasicContractController struct {
	basicContractService *service.BasicContractService
}

func NewBasicContractController(basicContractService *service.BasicContractService) *BasicContractController {
	return &BasicContractController{
		basicContractService: basicContractService,
	}
}

// RegisterRoutes 注册路由
func (c *BasicContractController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/basic/contract")
	{
		// 合同分类
		api.GET("/contract-categories", c.GetContractCategoryList)
		api.GET("/contract-categories/:id", c.GetContractCategoryByID)
		api.POST("/contract-categories", c.CreateContractCategory)
		api.PUT("/contract-categories/:id", c.UpdateContractCategory)
		api.DELETE("/contract-categories/:id", c.DeleteContractCategory)

		// 产品分类
		api.GET("/product-categories", c.GetProductCategoryList)
		api.GET("/product-categories/:id", c.GetProductCategoryByID)
		api.POST("/product-categories", c.CreateProductCategory)
		api.PUT("/product-categories/:id", c.UpdateProductCategory)
		api.DELETE("/product-categories/:id", c.DeleteProductCategory)

		// 产品列表
		api.GET("/products", c.GetProductList)
		api.GET("/products/:id", c.GetProductByID)
		api.POST("/products", c.CreateProduct)
		api.PUT("/products/:id", c.UpdateProduct)
		api.DELETE("/products/:id", c.DeleteProduct)

		// 服务内容
		api.GET("/service-contents", c.GetServiceContentList)
		api.GET("/service-contents/:id", c.GetServiceContentByID)
		api.POST("/service-contents", c.CreateServiceContent)
		api.PUT("/service-contents/:id", c.UpdateServiceContent)
		api.DELETE("/service-contents/:id", c.DeleteServiceContent)

		// 供应商列表
		api.GET("/suppliers", c.GetSupplierList)
		api.GET("/suppliers/:id", c.GetSupplierByID)
		api.POST("/suppliers", c.CreateSupplier)
		api.PUT("/suppliers/:id", c.UpdateSupplier)
		api.DELETE("/suppliers/:id", c.DeleteSupplier)

		// 采购品分类
		api.GET("/purchase-categories", c.GetPurchaseCategoryList)
		api.GET("/purchase-categories/:id", c.GetPurchaseCategoryByID)
		api.POST("/purchase-categories", c.CreatePurchaseCategory)
		api.PUT("/purchase-categories/:id", c.UpdatePurchaseCategory)
		api.DELETE("/purchase-categories/:id", c.DeletePurchaseCategory)

		// 采购品列表
		api.GET("/purchase-items", c.GetPurchaseItemList)
		api.GET("/purchase-items/:id", c.GetPurchaseItemByID)
		api.POST("/purchase-items", c.CreatePurchaseItem)
		api.PUT("/purchase-items/:id", c.UpdatePurchaseItem)
		api.DELETE("/purchase-items/:id", c.DeletePurchaseItem)
	}
}

// GetContractCategoryList 获取合同分类列表
func (c *BasicContractController) GetContractCategoryList(ctx *gin.Context) {
	categories, err := c.basicContractService.GetContractCategoryList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// GetContractCategoryByID 根据ID获取合同分类
func (c *BasicContractController) GetContractCategoryByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	category, err := c.basicContractService.GetContractCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// CreateContractCategory 创建合同分类
func (c *BasicContractController) CreateContractCategory(ctx *gin.Context) {
	var category model.ContractCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicContractService.CreateContractCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

// UpdateContractCategory 更新合同分类
func (c *BasicContractController) UpdateContractCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var category model.ContractCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = uint(id)
	if err := c.basicContractService.UpdateContractCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// DeleteContractCategory 删除合同分类
func (c *BasicContractController) DeleteContractCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicContractService.DeleteContractCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetProductCategoryList 获取产品分类列表
func (c *BasicContractController) GetProductCategoryList(ctx *gin.Context) {
	var parentID *uint
	if parentIDStr := ctx.Query("parent_id"); parentIDStr != "" {
		if id, err := strconv.ParseUint(parentIDStr, 10, 32); err == nil {
			pid := uint(id)
			parentID = &pid
		}
	}

	categories, err := c.basicContractService.GetProductCategoryList(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// GetProductCategoryByID 根据ID获取产品分类
func (c *BasicContractController) GetProductCategoryByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	category, err := c.basicContractService.GetProductCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// CreateProductCategory 创建产品分类
func (c *BasicContractController) CreateProductCategory(ctx *gin.Context) {
	var category model.ProductCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicContractService.CreateProductCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

// UpdateProductCategory 更新产品分类
func (c *BasicContractController) UpdateProductCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var category model.ProductCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = uint(id)
	if err := c.basicContractService.UpdateProductCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// DeleteProductCategory 删除产品分类
func (c *BasicContractController) DeleteProductCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicContractService.DeleteProductCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetProductList 获取产品列表
func (c *BasicContractController) GetProductList(ctx *gin.Context) {
	categoryID, _ := strconv.ParseUint(ctx.Query("category_id"), 10, 32)
	products, err := c.basicContractService.GetProductList(uint(categoryID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetProductByID 根据ID获取产品
func (c *BasicContractController) GetProductByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	product, err := c.basicContractService.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// CreateProduct 创建产品
func (c *BasicContractController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicContractService.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// UpdateProduct 更新产品
func (c *BasicContractController) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	if err := c.basicContractService.UpdateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct 删除产品
func (c *BasicContractController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicContractService.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetServiceContentList 获取服务内容列表
func (c *BasicContractController) GetServiceContentList(ctx *gin.Context) {
	contents, err := c.basicContractService.GetServiceContentList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contents)
}

// GetServiceContentByID 根据ID获取服务内容
func (c *BasicContractController) GetServiceContentByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	content, err := c.basicContractService.GetServiceContentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, content)
}

// CreateServiceContent 创建服务内容
func (c *BasicContractController) CreateServiceContent(ctx *gin.Context) {
	var content model.ServiceContent
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicContractService.CreateServiceContent(&content); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, content)
}

// UpdateServiceContent 更新服务内容
func (c *BasicContractController) UpdateServiceContent(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var content model.ServiceContent
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content.ID = uint(id)
	if err := c.basicContractService.UpdateServiceContent(&content); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, content)
}

// DeleteServiceContent 删除服务内容
func (c *BasicContractController) DeleteServiceContent(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicContractService.DeleteServiceContent(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetSupplierList 获取供应商列表
func (c *BasicContractController) GetSupplierList(ctx *gin.Context) {
	suppliers, err := c.basicContractService.GetSupplierList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, suppliers)
}

// GetSupplierByID 根据ID获取供应商
func (c *BasicContractController) GetSupplierByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	supplier, err := c.basicContractService.GetSupplierByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, supplier)
}

// CreateSupplier 创建供应商
func (c *BasicContractController) CreateSupplier(ctx *gin.Context) {
	var supplier model.Supplier
	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicContractService.CreateSupplier(&supplier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, supplier)
}

// UpdateSupplier 更新供应商
func (c *BasicContractController) UpdateSupplier(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var supplier model.Supplier
	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier.ID = uint(id)
	if err := c.basicContractService.UpdateSupplier(&supplier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, supplier)
}

// DeleteSupplier 删除供应商
func (c *BasicContractController) DeleteSupplier(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicContractService.DeleteSupplier(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetPurchaseCategoryList 获取采购品分类列表
func (c *BasicContractController) GetPurchaseCategoryList(ctx *gin.Context) {
	var parentID *uint
	if parentIDStr := ctx.Query("parent_id"); parentIDStr != "" {
		if id, err := strconv.ParseUint(parentIDStr, 10, 32); err == nil {
			pid := uint(id)
			parentID = &pid
		}
	}

	categories, err := c.basicContractService.GetPurchaseCategoryList(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// GetPurchaseCategoryByID 根据ID获取采购品分类
func (c *BasicContractController) GetPurchaseCategoryByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	category, err := c.basicContractService.GetPurchaseCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// CreatePurchaseCategory 创建采购品分类
func (c *BasicContractController) CreatePurchaseCategory(ctx *gin.Context) {
	var category model.PurchaseCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicContractService.CreatePurchaseCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

// UpdatePurchaseCategory 更新采购品分类
func (c *BasicContractController) UpdatePurchaseCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var category model.PurchaseCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = uint(id)
	if err := c.basicContractService.UpdatePurchaseCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// DeletePurchaseCategory 删除采购品分类
func (c *BasicContractController) DeletePurchaseCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicContractService.DeletePurchaseCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetPurchaseItemList 获取采购品列表
func (c *BasicContractController) GetPurchaseItemList(ctx *gin.Context) {
	categoryID, _ := strconv.ParseUint(ctx.Query("category_id"), 10, 32)
	items, err := c.basicContractService.GetPurchaseItemList(uint(categoryID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// GetPurchaseItemByID 根据ID获取采购品
func (c *BasicContractController) GetPurchaseItemByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	item, err := c.basicContractService.GetPurchaseItemByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// CreatePurchaseItem 创建采购品
func (c *BasicContractController) CreatePurchaseItem(ctx *gin.Context) {
	var item model.PurchaseItem
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.basicContractService.CreatePurchaseItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

// UpdatePurchaseItem 更新采购品
func (c *BasicContractController) UpdatePurchaseItem(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var item model.PurchaseItem
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	if err := c.basicContractService.UpdatePurchaseItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// DeletePurchaseItem 删除采购品
func (c *BasicContractController) DeletePurchaseItem(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.basicContractService.DeletePurchaseItem(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
