package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	vehicleService *service.VehicleService
}

func NewVehicleController(vehicleService *service.VehicleService) *VehicleController {
	return &VehicleController{
		vehicleService: vehicleService,
	}
}

// RegisterRoutes 注册路由
func (c *VehicleController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/vehicles")
	{
		// 车辆管理
		api.GET("", c.GetVehicleList)
		api.GET("/:id", c.GetVehicleByID)
		api.POST("", c.CreateVehicle)
		api.PUT("/:id", c.UpdateVehicle)
		api.DELETE("/:id", c.DeleteVehicle)

		// 维修记录管理
		api.GET("/repairs", c.GetVehicleRepairList)
		api.GET("/repairs/:id", c.GetVehicleRepairByID)
		api.POST("/repairs", c.CreateVehicleRepair)
		api.PUT("/repairs/:id", c.UpdateVehicleRepair)
		api.DELETE("/repairs/:id", c.DeleteVehicleRepair)
		api.PUT("/repairs/:id/complete", c.CompleteVehicleRepair)

		// 保养记录管理
		api.GET("/maintenances", c.GetVehicleMaintenanceList)
		api.GET("/maintenances/:id", c.GetVehicleMaintenanceByID)
		api.POST("/maintenances", c.CreateVehicleMaintenance)
		api.PUT("/maintenances/:id", c.UpdateVehicleMaintenance)
		api.DELETE("/maintenances/:id", c.DeleteVehicleMaintenance)
		api.PUT("/maintenances/:id/complete", c.CompleteVehicleMaintenance)

		// 里程记录管理
		api.GET("/mileages", c.GetVehicleMileageList)
		api.GET("/mileages/:id", c.GetVehicleMileageByID)
		api.POST("/mileages", c.CreateVehicleMileage)
		api.PUT("/mileages/:id", c.UpdateVehicleMileage)
		api.DELETE("/mileages/:id", c.DeleteVehicleMileage)

		// 费用记录管理
		api.GET("/expenses", c.GetVehicleExpenseList)
		api.GET("/expenses/:id", c.GetVehicleExpenseByID)
		api.POST("/expenses", c.CreateVehicleExpense)
		api.PUT("/expenses/:id", c.UpdateVehicleExpense)
		api.DELETE("/expenses/:id", c.DeleteVehicleExpense)

		// 违章记录管理
		api.GET("/violations", c.GetVehicleViolationList)
		api.GET("/violations/:id", c.GetVehicleViolationByID)
		api.POST("/violations", c.CreateVehicleViolation)
		api.PUT("/violations/:id", c.UpdateVehicleViolation)
		api.DELETE("/violations/:id", c.DeleteVehicleViolation)
		api.PUT("/violations/:id/handle", c.HandleVehicleViolation)

		// 事故记录管理
		api.GET("/accidents", c.GetVehicleAccidentList)
		api.GET("/accidents/:id", c.GetVehicleAccidentByID)
		api.POST("/accidents", c.CreateVehicleAccident)
		api.PUT("/accidents/:id", c.UpdateVehicleAccident)
		api.DELETE("/accidents/:id", c.DeleteVehicleAccident)
		api.PUT("/accidents/:id/handle", c.HandleVehicleAccident)

		// 用车申请管理
		api.GET("/applications", c.GetVehicleApplicationList)
		api.GET("/applications/:id", c.GetVehicleApplicationByID)
		api.POST("/applications", c.CreateVehicleApplication)
		api.PUT("/applications/:id", c.UpdateVehicleApplication)
		api.DELETE("/applications/:id", c.DeleteVehicleApplication)
		api.PUT("/applications/:id/approve", c.ApproveVehicleApplication)
		api.PUT("/applications/:id/reject", c.RejectVehicleApplication)

		// 车辆归还管理
		api.GET("/returns", c.GetVehicleReturnList)
		api.GET("/returns/:id", c.GetVehicleReturnByID)
		api.POST("/returns", c.CreateVehicleReturn)
		api.PUT("/returns/:id", c.UpdateVehicleReturn)
		api.DELETE("/returns/:id", c.DeleteVehicleReturn)
	}
}

// GetVehicleList 获取车辆列表
func (c *VehicleController) GetVehicleList(ctx *gin.Context) {
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	vehicles, total, err := c.vehicleService.GetVehicleList(status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  vehicles,
		"total": total,
	})
}

// GetVehicleByID 根据ID获取车辆
func (c *VehicleController) GetVehicleByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	vehicle, err := c.vehicleService.GetVehicleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

// CreateVehicle 创建车辆
func (c *VehicleController) CreateVehicle(ctx *gin.Context) {
	var vehicle model.Vehicle
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	vehicle.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicle(&vehicle); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, vehicle)
}

// UpdateVehicle 更新车辆
func (c *VehicleController) UpdateVehicle(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var vehicle model.Vehicle
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle.ID = uint(id)
	if err := c.vehicleService.UpdateVehicle(&vehicle); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

// DeleteVehicle 删除车辆
func (c *VehicleController) DeleteVehicle(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicle(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleRepairList 获取车辆维修记录列表
func (c *VehicleController) GetVehicleRepairList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	repairs, total, err := c.vehicleService.GetVehicleRepairList(uint(vehicleID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  repairs,
		"total": total,
	})
}

// GetVehicleRepairByID 根据ID获取车辆维修记录
func (c *VehicleController) GetVehicleRepairByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	repair, err := c.vehicleService.GetVehicleRepairByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, repair)
}

// CreateVehicleRepair 创建车辆维修记录
func (c *VehicleController) CreateVehicleRepair(ctx *gin.Context) {
	var repair model.VehicleRepair
	if err := ctx.ShouldBindJSON(&repair); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	repair.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleRepair(&repair); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, repair)
}

// UpdateVehicleRepair 更新车辆维修记录
func (c *VehicleController) UpdateVehicleRepair(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var repair model.VehicleRepair
	if err := ctx.ShouldBindJSON(&repair); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repair.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleRepair(&repair); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, repair)
}

// DeleteVehicleRepair 删除车辆维修记录
func (c *VehicleController) DeleteVehicleRepair(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleRepair(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CompleteVehicleRepair 完成车辆维修
func (c *VehicleController) CompleteVehicleRepair(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.CompleteVehicleRepair(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleMaintenanceList 获取车辆保养记录列表
func (c *VehicleController) GetVehicleMaintenanceList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	maintenances, total, err := c.vehicleService.GetVehicleMaintenanceList(uint(vehicleID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  maintenances,
		"total": total,
	})
}

// GetVehicleMaintenanceByID 根据ID获取车辆保养记录
func (c *VehicleController) GetVehicleMaintenanceByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	maintenance, err := c.vehicleService.GetVehicleMaintenanceByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, maintenance)
}

// CreateVehicleMaintenance 创建车辆保养记录
func (c *VehicleController) CreateVehicleMaintenance(ctx *gin.Context) {
	var maintenance model.VehicleMaintenance
	if err := ctx.ShouldBindJSON(&maintenance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	maintenance.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleMaintenance(&maintenance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, maintenance)
}

// UpdateVehicleMaintenance 更新车辆保养记录
func (c *VehicleController) UpdateVehicleMaintenance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var maintenance model.VehicleMaintenance
	if err := ctx.ShouldBindJSON(&maintenance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maintenance.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleMaintenance(&maintenance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, maintenance)
}

// DeleteVehicleMaintenance 删除车辆保养记录
func (c *VehicleController) DeleteVehicleMaintenance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleMaintenance(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CompleteVehicleMaintenance 完成车辆保养
func (c *VehicleController) CompleteVehicleMaintenance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.CompleteVehicleMaintenance(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleMileageList 获取车辆里程记录列表
func (c *VehicleController) GetVehicleMileageList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	mileages, total, err := c.vehicleService.GetVehicleMileageList(uint(vehicleID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  mileages,
		"total": total,
	})
}

// GetVehicleMileageByID 根据ID获取车辆里程记录
func (c *VehicleController) GetVehicleMileageByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	mileage, err := c.vehicleService.GetVehicleMileageByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mileage)
}

// CreateVehicleMileage 创建车辆里程记录
func (c *VehicleController) CreateVehicleMileage(ctx *gin.Context) {
	var mileage model.VehicleMileage
	if err := ctx.ShouldBindJSON(&mileage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	mileage.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleMileage(&mileage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, mileage)
}

// UpdateVehicleMileage 更新车辆里程记录
func (c *VehicleController) UpdateVehicleMileage(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var mileage model.VehicleMileage
	if err := ctx.ShouldBindJSON(&mileage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mileage.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleMileage(&mileage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mileage)
}

// DeleteVehicleMileage 删除车辆里程记录
func (c *VehicleController) DeleteVehicleMileage(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleMileage(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleExpenseList 获取车辆费用记录列表
func (c *VehicleController) GetVehicleExpenseList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	expenseID, _ := strconv.ParseUint(ctx.Query("expense_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	expenses, total, err := c.vehicleService.GetVehicleExpenseList(uint(vehicleID), uint(expenseID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  expenses,
		"total": total,
	})
}

// GetVehicleExpenseByID 根据ID获取车辆费用记录
func (c *VehicleController) GetVehicleExpenseByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	expense, err := c.vehicleService.GetVehicleExpenseByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

// CreateVehicleExpense 创建车辆费用记录
func (c *VehicleController) CreateVehicleExpense(ctx *gin.Context) {
	var expense model.VehicleExpenseRecord
	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	expense.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleExpense(&expense); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, expense)
}

// UpdateVehicleExpense 更新车辆费用记录
func (c *VehicleController) UpdateVehicleExpense(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var expense model.VehicleExpenseRecord
	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleExpense(&expense); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

// DeleteVehicleExpense 删除车辆费用记录
func (c *VehicleController) DeleteVehicleExpense(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleExpense(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleViolationList 获取车辆违章记录列表
func (c *VehicleController) GetVehicleViolationList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	violations, total, err := c.vehicleService.GetVehicleViolationList(uint(vehicleID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  violations,
		"total": total,
	})
}

// GetVehicleViolationByID 根据ID获取车辆违章记录
func (c *VehicleController) GetVehicleViolationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	violation, err := c.vehicleService.GetVehicleViolationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, violation)
}

// CreateVehicleViolation 创建车辆违章记录
func (c *VehicleController) CreateVehicleViolation(ctx *gin.Context) {
	var violation model.VehicleViolation
	if err := ctx.ShouldBindJSON(&violation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	violation.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleViolation(&violation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, violation)
}

// UpdateVehicleViolation 更新车辆违章记录
func (c *VehicleController) UpdateVehicleViolation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var violation model.VehicleViolation
	if err := ctx.ShouldBindJSON(&violation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	violation.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleViolation(&violation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, violation)
}

// DeleteVehicleViolation 删除车辆违章记录
func (c *VehicleController) DeleteVehicleViolation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleViolation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// HandleVehicleViolation 处理车辆违章
func (c *VehicleController) HandleVehicleViolation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		Status int `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.vehicleService.HandleVehicleViolation(uint(id), req.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleAccidentList 获取车辆事故记录列表
func (c *VehicleController) GetVehicleAccidentList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	accidents, total, err := c.vehicleService.GetVehicleAccidentList(uint(vehicleID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  accidents,
		"total": total,
	})
}

// GetVehicleAccidentByID 根据ID获取车辆事故记录
func (c *VehicleController) GetVehicleAccidentByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	accident, err := c.vehicleService.GetVehicleAccidentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accident)
}

// CreateVehicleAccident 创建车辆事故记录
func (c *VehicleController) CreateVehicleAccident(ctx *gin.Context) {
	var accident model.VehicleAccident
	if err := ctx.ShouldBindJSON(&accident); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	accident.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleAccident(&accident); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, accident)
}

// UpdateVehicleAccident 更新车辆事故记录
func (c *VehicleController) UpdateVehicleAccident(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var accident model.VehicleAccident
	if err := ctx.ShouldBindJSON(&accident); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accident.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleAccident(&accident); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accident)
}

// DeleteVehicleAccident 删除车辆事故记录
func (c *VehicleController) DeleteVehicleAccident(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleAccident(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// HandleVehicleAccident 处理车辆事故
func (c *VehicleController) HandleVehicleAccident(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		Status int `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.vehicleService.HandleVehicleAccident(uint(id), req.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleApplicationList 获取车辆用车申请列表
func (c *VehicleController) GetVehicleApplicationList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	status, _ := strconv.Atoi(ctx.Query("status"))
	applications, total, err := c.vehicleService.GetVehicleApplicationList(uint(vehicleID), uint(status), page, pageSize, 0) // Add the missing argument
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  applications,
		"total": total,
	})
}

// GetVehicleApplicationByID 根据ID获取车辆用车申请
func (c *VehicleController) GetVehicleApplicationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	application, err := c.vehicleService.GetVehicleApplicationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// CreateVehicleApplication 创建车辆用车申请
func (c *VehicleController) CreateVehicleApplication(ctx *gin.Context) {
	var application model.VehicleApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	application.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, application)
}

// UpdateVehicleApplication 更新车辆用车申请
func (c *VehicleController) UpdateVehicleApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var application model.VehicleApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// DeleteVehicleApplication 删除车辆用车申请
func (c *VehicleController) DeleteVehicleApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveVehicleApplication 批准车辆用车申请
func (c *VehicleController) ApproveVehicleApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.ApproveVehicleApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectVehicleApplication 拒绝车辆用车申请
func (c *VehicleController) RejectVehicleApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.RejectVehicleApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetVehicleReturnList 获取车辆归还记录列表
func (c *VehicleController) GetVehicleReturnList(ctx *gin.Context) {
	vehicleID, _ := strconv.ParseUint(ctx.Query("vehicle_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	returns, total, err := c.vehicleService.GetVehicleReturnList(uint(vehicleID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  returns,
		"total": total,
	})
}

// GetVehicleReturnByID 根据ID获取车辆归还记录
func (c *VehicleController) GetVehicleReturnByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	vehicleReturn, err := c.vehicleService.GetVehicleReturnByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicleReturn)
}

// CreateVehicleReturn 创建车辆归还记录
func (c *VehicleController) CreateVehicleReturn(ctx *gin.Context) {
	var vehicleReturn model.VehicleReturn
	if err := ctx.ShouldBindJSON(&vehicleReturn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	vehicleReturn.CreatedBy = uint(1)

	if err := c.vehicleService.CreateVehicleReturn(&vehicleReturn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, vehicleReturn)
}

// UpdateVehicleReturn 更新车辆归还记录
func (c *VehicleController) UpdateVehicleReturn(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var vehicleReturn model.VehicleReturn
	if err := ctx.ShouldBindJSON(&vehicleReturn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicleReturn.ID = uint(id)
	if err := c.vehicleService.UpdateVehicleReturn(&vehicleReturn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicleReturn)
}

// DeleteVehicleReturn 删除车辆归还记录
func (c *VehicleController) DeleteVehicleReturn(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.vehicleService.DeleteVehicleReturn(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
