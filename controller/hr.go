package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type HRController struct {
	hrService *service.HRService
}

func NewHRController(hrService *service.HRService) *HRController {
	return &HRController{
		hrService: hrService,
	}
}

// RegisterRoutes 注册路由
func (c *HRController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/hr")
	{
		// 岗位职称管理
		api.GET("/positions", c.GetPositionList)
		api.GET("/positions/:id", c.GetPositionByID)
		api.POST("/positions", c.CreatePosition)
		api.PUT("/positions/:id", c.UpdatePosition)
		api.DELETE("/positions/:id", c.DeletePosition)

		// 员工档案管理
		api.GET("/archives", c.GetEmployeeArchiveList)
		api.GET("/archives/employee/:employee_id", c.GetEmployeeArchiveByEmployeeID)
		api.POST("/archives", c.CreateEmployeeArchive)
		api.PUT("/archives/:id", c.UpdateEmployeeArchive)
		api.DELETE("/archives/:id", c.DeleteEmployeeArchive)

		// 奖惩记录管理
		api.GET("/reward-punishments", c.GetRewardPunishmentRecordList)
		api.GET("/reward-punishments/:id", c.GetRewardPunishmentRecordByID)
		api.POST("/reward-punishments", c.CreateRewardPunishmentRecord)
		api.PUT("/reward-punishments/:id", c.UpdateRewardPunishmentRecord)
		api.DELETE("/reward-punishments/:id", c.DeleteRewardPunishmentRecord)

		// 关怀记录管理
		api.GET("/care-records", c.GetCareRecordList)
		api.GET("/care-records/:id", c.GetCareRecordByID)
		api.POST("/care-records", c.CreateCareRecord)
		api.PUT("/care-records/:id", c.UpdateCareRecord)
		api.DELETE("/care-records/:id", c.DeleteCareRecord)

		// 人事调动管理
		api.GET("/transfers", c.GetTransferList)
		api.GET("/transfers/:id", c.GetTransferByID)
		api.POST("/transfers", c.CreateTransfer)
		api.PUT("/transfers/:id", c.UpdateTransfer)
		api.DELETE("/transfers/:id", c.DeleteTransfer)
		api.PUT("/transfers/:id/approve", c.ApproveTransfer)
		api.PUT("/transfers/:id/reject", c.RejectTransfer)

		// 离职档案管理
		api.GET("/resignations", c.GetResignationList)
		api.GET("/resignations/:id", c.GetResignationByID)
		api.POST("/resignations", c.CreateResignation)
		api.PUT("/resignations/:id", c.UpdateResignation)
		api.DELETE("/resignations/:id", c.DeleteResignation)
		api.PUT("/resignations/:id/approve", c.ApproveResignation)
		api.PUT("/resignations/:id/reject", c.RejectResignation)

		// 员工合同管理
		api.GET("/contracts", c.GetContractList)
		api.GET("/contracts/:id", c.GetContractByID)
		api.POST("/contracts", c.CreateContract)
		api.PUT("/contracts/:id", c.UpdateContract)
		api.DELETE("/contracts/:id", c.DeleteContract)
		api.PUT("/contracts/:id/terminate", c.TerminateContract)

		// 转正管理
		api.GET("/probations", c.GetProbationList)
		api.GET("/probations/:id", c.GetProbationByID)
		api.POST("/probations", c.CreateProbation)
		api.PUT("/probations/:id", c.UpdateProbation)
		api.DELETE("/probations/:id", c.DeleteProbation)
		api.PUT("/probations/:id/approve", c.ApproveProbation)
		api.PUT("/probations/:id/reject", c.RejectProbation)
	}
}

// GetPositionList 获取岗位职称列表
func (c *HRController) GetPositionList(ctx *gin.Context) {
	positions, err := c.hrService.GetPositionList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, positions)
}

// GetPositionByID 根据ID获取岗位职称
func (c *HRController) GetPositionByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	position, err := c.hrService.GetPositionByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, position)
}

// CreatePosition 创建岗位职称
func (c *HRController) CreatePosition(ctx *gin.Context) {
	var position model.Position
	if err := ctx.ShouldBindJSON(&position); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreatePosition(&position); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, position)
}

// UpdatePosition 更新岗位职称
func (c *HRController) UpdatePosition(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var position model.Position
	if err := ctx.ShouldBindJSON(&position); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	position.ID = uint(id)
	if err := c.hrService.UpdatePosition(&position); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, position)
}

// DeletePosition 删除岗位职称
func (c *HRController) DeletePosition(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeletePosition(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetEmployeeArchiveList 获取员工档案列表
func (c *HRController) GetEmployeeArchiveList(ctx *gin.Context) {
	departmentID, _ := strconv.ParseUint(ctx.Query("department_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	archives, total, err := c.hrService.GetEmployeeArchiveList(uint(departmentID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  archives,
		"total": total,
	})
}

// GetEmployeeArchiveByEmployeeID 根据员工ID获取档案
func (c *HRController) GetEmployeeArchiveByEmployeeID(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Param("employee_id"), 10, 32)
	archive, err := c.hrService.GetEmployeeArchiveByEmployeeID(uint(employeeID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, archive)
}

// CreateEmployeeArchive 创建员工档案
func (c *HRController) CreateEmployeeArchive(ctx *gin.Context) {
	var archive model.EmployeeArchive
	if err := ctx.ShouldBindJSON(&archive); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreateEmployeeArchive(&archive); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, archive)
}

// UpdateEmployeeArchive 更新员工档案
func (c *HRController) UpdateEmployeeArchive(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var archive model.EmployeeArchive
	if err := ctx.ShouldBindJSON(&archive); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	archive.ID = uint(id)
	if err := c.hrService.UpdateEmployeeArchive(&archive); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, archive)
}

// DeleteEmployeeArchive 删除员工档案
func (c *HRController) DeleteEmployeeArchive(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeleteEmployeeArchive(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetRewardPunishmentRecordList 获取奖惩记录列表
func (c *HRController) GetRewardPunishmentRecordList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	records, total, err := c.hrService.GetRewardPunishmentRecordList(uint(employeeID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
	})
}

// GetRewardPunishmentRecordByID 根据ID获取奖惩记录
func (c *HRController) GetRewardPunishmentRecordByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	record, err := c.hrService.GetRewardPunishmentRecordByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// CreateRewardPunishmentRecord 创建奖惩记录
func (c *HRController) CreateRewardPunishmentRecord(ctx *gin.Context) {
	var record model.RewardPunishmentRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreateRewardPunishmentRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

// UpdateRewardPunishmentRecord 更新奖惩记录
func (c *HRController) UpdateRewardPunishmentRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var record model.RewardPunishmentRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = uint(id)
	if err := c.hrService.UpdateRewardPunishmentRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// DeleteRewardPunishmentRecord 删除奖惩记录
func (c *HRController) DeleteRewardPunishmentRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeleteRewardPunishmentRecord(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetCareRecordList 获取关怀记录列表
func (c *HRController) GetCareRecordList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	records, total, err := c.hrService.GetCareRecordList(uint(employeeID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
	})
}

// GetCareRecordByID 根据ID获取关怀记录
func (c *HRController) GetCareRecordByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	record, err := c.hrService.GetCareRecordByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// CreateCareRecord 创建关怀记录
func (c *HRController) CreateCareRecord(ctx *gin.Context) {
	var record model.CareRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreateCareRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

// UpdateCareRecord 更新关怀记录
func (c *HRController) UpdateCareRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var record model.CareRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = uint(id)
	if err := c.hrService.UpdateCareRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// DeleteCareRecord 删除关怀记录
func (c *HRController) DeleteCareRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeleteCareRecord(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetTransferList 获取人事调动列表
func (c *HRController) GetTransferList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	transfers, total, err := c.hrService.GetTransferList(uint(employeeID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  transfers,
		"total": total,
	})
}

// GetTransferByID 根据ID获取人事调动
func (c *HRController) GetTransferByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	transfer, err := c.hrService.GetTransferByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}

// CreateTransfer 创建人事调动
func (c *HRController) CreateTransfer(ctx *gin.Context) {
	var transfer model.Transfer
	if err := ctx.ShouldBindJSON(&transfer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreateTransfer(&transfer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transfer)
}

// UpdateTransfer 更新人事调动
func (c *HRController) UpdateTransfer(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var transfer model.Transfer
	if err := ctx.ShouldBindJSON(&transfer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfer.ID = uint(id)
	if err := c.hrService.UpdateTransfer(&transfer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}

// DeleteTransfer 删除人事调动
func (c *HRController) DeleteTransfer(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeleteTransfer(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveTransfer 审批通过人事调动
func (c *HRController) ApproveTransfer(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.ApproveTransfer(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectTransfer 审批驳回人事调动
func (c *HRController) RejectTransfer(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.RejectTransfer(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetResignationList 获取离职档案列表
func (c *HRController) GetResignationList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	resignations, total, err := c.hrService.GetResignationList(uint(employeeID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  resignations,
		"total": total,
	})
}

// GetResignationByID 根据ID获取离职档案
func (c *HRController) GetResignationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	resignation, err := c.hrService.GetResignationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resignation)
}

// CreateResignation 创建离职档案
func (c *HRController) CreateResignation(ctx *gin.Context) {
	var resignation model.Resignation
	if err := ctx.ShouldBindJSON(&resignation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreateResignation(&resignation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resignation)
}

// UpdateResignation 更新离职档案
func (c *HRController) UpdateResignation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var resignation model.Resignation
	if err := ctx.ShouldBindJSON(&resignation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resignation.ID = uint(id)
	if err := c.hrService.UpdateResignation(&resignation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resignation)
}

// DeleteResignation 删除离职档案
func (c *HRController) DeleteResignation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeleteResignation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveResignation 审批通过离职档案
func (c *HRController) ApproveResignation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.ApproveResignation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectResignation 审批驳回离职档案
func (c *HRController) RejectResignation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.RejectResignation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetContractList 获取员工合同列表
func (c *HRController) GetContractList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	contracts, total, err := c.hrService.GetContractList(uint(employeeID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  contracts,
		"total": total,
	})
}

// GetContractByID 根据ID获取员工合同
func (c *HRController) GetContractByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	contract, err := c.hrService.GetContractByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contract)
}

// CreateContract 创建员工合同
func (c *HRController) CreateContract(ctx *gin.Context) {
	var contract model.Contract
	if err := ctx.ShouldBindJSON(&contract); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreateContract(&contract); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, contract)
}

// UpdateContract 更新员工合同
func (c *HRController) UpdateContract(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var contract model.Contract
	if err := ctx.ShouldBindJSON(&contract); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contract.ID = uint(id)
	if err := c.hrService.UpdateContract(&contract); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contract)
}

// DeleteContract 删除员工合同
func (c *HRController) DeleteContract(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeleteContract(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// TerminateContract 终止员工合同
func (c *HRController) TerminateContract(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.TerminateContract(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetProbationList 获取转正列表
func (c *HRController) GetProbationList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	probations, total, err := c.hrService.GetProbationList(uint(employeeID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  probations,
		"total": total,
	})
}

// GetProbationByID 根据ID获取转正
func (c *HRController) GetProbationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	probation, err := c.hrService.GetProbationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, probation)
}

// CreateProbation 创建转正
func (c *HRController) CreateProbation(ctx *gin.Context) {
	var probation model.Probation
	if err := ctx.ShouldBindJSON(&probation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.hrService.CreateProbation(&probation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, probation)
}

// UpdateProbation 更新转正
func (c *HRController) UpdateProbation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var probation model.Probation
	if err := ctx.ShouldBindJSON(&probation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	probation.ID = uint(id)
	if err := c.hrService.UpdateProbation(&probation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, probation)
}

// DeleteProbation 删除转正
func (c *HRController) DeleteProbation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.DeleteProbation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveProbation 审批通过转正
func (c *HRController) ApproveProbation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.ApproveProbation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectProbation 审批驳回转正
func (c *HRController) RejectProbation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.hrService.RejectProbation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
