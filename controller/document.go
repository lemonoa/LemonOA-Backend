package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/lemonoa/LemonOA-Go/middleware"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService *service.DocumentService
}

func NewDocumentController(documentService *service.DocumentService) *DocumentController {
	return &DocumentController{
		documentService: documentService,
	}
}

// RegisterRoutes 注册路由
func (c *DocumentController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.JWT())

	// 文档管理
	docs := api.Group("/documents")
	{
		docs.GET("", middleware.RequirePermission(model.PermissionDocumentList), c.GetDocumentList)
		docs.POST("", middleware.RequirePermission(model.PermissionDocumentCreate), c.CreateDocument)
		docs.PUT("/:id", middleware.RequirePermission(model.PermissionDocumentUpdate), c.UpdateDocument)
		docs.DELETE("/:id", middleware.RequirePermission(model.PermissionDocumentDelete), c.DeleteDocument)
		docs.POST("/:id/submit", middleware.RequirePermission(model.PermissionDocumentCreate), c.SubmitDocument)
		docs.PUT("/:id/approve", middleware.RequirePermission(model.PermissionDocumentApprove), c.ApproveDocument)
		docs.PUT("/:id/reject", middleware.RequirePermission(model.PermissionDocumentApprove), c.RejectDocument)
		docs.POST("/:id/distribute", middleware.RequirePermission(model.PermissionDocumentCreate), c.DistributeDocument)
		docs.PUT("/:id/read", middleware.RequirePermission(model.PermissionDocumentList), c.ReadDocument)
		docs.POST("/:id/archive", middleware.RequirePermission(model.PermissionDocumentArchive), c.ArchiveDocument)
		docs.POST("/:id/borrow", middleware.RequirePermission(model.PermissionDocumentList), c.BorrowDocument)
		docs.PUT("/:id/return", middleware.RequirePermission(model.PermissionDocumentList), c.ReturnDocument)
		docs.PUT("/:id/destroy", middleware.RequirePermission(model.PermissionDocumentArchive), c.DestroyDocument)
	}
}

// GetDocumentList 获取公文列表
func (c *DocumentController) GetDocumentList(ctx *gin.Context) {
	typeID, _ := strconv.ParseUint(ctx.Query("type_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	documents, total, err := c.documentService.GetDocumentList(uint(typeID), status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  documents,
		"total": total,
	})
}

// GetDocumentByID 根据ID获取公文
func (c *DocumentController) GetDocumentByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	document, err := c.documentService.GetDocumentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, document)
}

// CreateDocument 创建公文
func (c *DocumentController) CreateDocument(ctx *gin.Context) {
	var document model.Document
	if err := ctx.ShouldBindJSON(&document); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	document.CreatedBy = uint(1)

	if err := c.documentService.CreateDocument(&document); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, document)
}

// UpdateDocument 更新公文
func (c *DocumentController) UpdateDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var document model.Document
	if err := ctx.ShouldBindJSON(&document); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	document.ID = uint(id)
	if err := c.documentService.UpdateDocument(&document); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, document)
}

// DeleteDocument 删除公文
func (c *DocumentController) DeleteDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.documentService.DeleteDocument(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// SubmitDocument 提交公文审批
func (c *DocumentController) SubmitDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var approvers []model.DocumentApproval
	if err := ctx.ShouldBindJSON(&approvers); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.documentService.SubmitDocument(uint(id), approvers); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveDocument 审批通过公文
func (c *DocumentController) ApproveDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var data struct {
		Comment string `json:"comment"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.documentService.ApproveDocument(uint(id), approverID, data.Comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectDocument 审批驳回公文
func (c *DocumentController) RejectDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var data struct {
		Comment string `json:"comment"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.documentService.RejectDocument(uint(id), approverID, data.Comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DistributeDocument 分发公文
func (c *DocumentController) DistributeDocument(ctx *gin.Context) {
	var distributions []model.DocumentDistribution
	if err := ctx.ShouldBindJSON(&distributions); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	for i := range distributions {
		distributions[i].CreatedBy = uint(1)
	}

	if err := c.documentService.DistributeDocument(distributions); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ReadDocument 阅读公文
func (c *DocumentController) ReadDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	receiverID := uint(1)

	if err := c.documentService.ReadDocument(uint(id), receiverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ArchiveDocument 归档公文
func (c *DocumentController) ArchiveDocument(ctx *gin.Context) {
	var archive model.DocumentArchive
	if err := ctx.ShouldBindJSON(&archive); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	archive.CreatedBy = uint(1)

	if err := c.documentService.ArchiveDocument(&archive); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, archive)
}

// BorrowDocument 借阅公文
func (c *DocumentController) BorrowDocument(ctx *gin.Context) {
	var borrow model.DocumentBorrow
	if err := ctx.ShouldBindJSON(&borrow); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	borrow.CreatedBy = uint(1)

	if err := c.documentService.BorrowDocument(&borrow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, borrow)
}

// ReturnDocument 归还公文
func (c *DocumentController) ReturnDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.documentService.ReturnDocument(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DestroyDocument 销毁公文
func (c *DocumentController) DestroyDocument(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.documentService.DestroyDocument(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
