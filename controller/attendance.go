package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/lemonoa/LemonOA-Go/middleware"

	"github.com/gin-gonic/gin"
)

type AttendanceController struct {
	attendanceService *service.AttendanceService
}

func NewAttendanceController(attendanceService *service.AttendanceService) *AttendanceController {
	return &AttendanceController{
		attendanceService: attendanceService,
	}
}

// RegisterRoutes 注册路由
func (c *AttendanceController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.JWT())

	// 考勤规则管理
	rules := api.Group("/attendance/rules")
	{
		rules.GET("", middleware.RequirePermission(model.PermissionAttendanceRuleList), c.GetAttendanceRuleList)
		rules.POST("", middleware.RequirePermission(model.PermissionAttendanceRuleCreate), c.CreateAttendanceRule)
		rules.PUT("/:id", middleware.RequirePermission(model.PermissionAttendanceRuleUpdate), c.UpdateAttendanceRule)
		rules.DELETE("/:id", middleware.RequirePermission(model.PermissionAttendanceRuleDelete), c.DeleteAttendanceRule)
	}

	// 考勤记录管理
	records := api.Group("/attendance/records")
	{
		records.GET("", middleware.RequirePermission(model.PermissionAttendanceRecordList), c.GetAttendanceRecordList)
		records.POST("", middleware.RequirePermission(model.PermissionAttendanceRecordCreate), c.CreateAttendanceRecord)
		records.PUT("/:id", middleware.RequirePermission(model.PermissionAttendanceRecordCreate), c.UpdateAttendanceRecord)
		records.DELETE("/:id", middleware.RequirePermission(model.PermissionAttendanceRecordCreate), c.DeleteAttendanceRecord)
	}

	// 请假管理
	leaves := api.Group("/attendance/leaves")
	{
		leaves.GET("", middleware.RequirePermission(model.PermissionLeaveList), c.GetLeaveApplicationList)
		leaves.POST("", middleware.RequirePermission(model.PermissionLeaveCreate), c.CreateLeaveApplication)
		leaves.PUT("/:id", middleware.RequirePermission(model.PermissionLeaveCreate), c.UpdateLeaveApplication)
		leaves.DELETE("/:id", middleware.RequirePermission(model.PermissionLeaveCreate), c.DeleteLeaveApplication)
		leaves.PUT("/:id/approve", middleware.RequirePermission(model.PermissionLeaveApprove), c.ApproveLeaveApplication)
		leaves.PUT("/:id/reject", middleware.RequirePermission(model.PermissionLeaveApprove), c.RejectLeaveApplication)
	}

	// 加班管理
	overtimes := api.Group("/attendance/overtimes")
	{
		overtimes.GET("", middleware.RequirePermission(model.PermissionLeaveList), c.GetOvertimeApplicationList)
		overtimes.POST("", middleware.RequirePermission(model.PermissionLeaveCreate), c.CreateOvertimeApplication)
		overtimes.PUT("/:id", middleware.RequirePermission(model.PermissionLeaveCreate), c.UpdateOvertimeApplication)
		overtimes.DELETE("/:id", middleware.RequirePermission(model.PermissionLeaveCreate), c.DeleteOvertimeApplication)
		overtimes.PUT("/:id/approve", middleware.RequirePermission(model.PermissionLeaveApprove), c.ApproveOvertimeApplication)
		overtimes.PUT("/:id/reject", middleware.RequirePermission(model.PermissionLeaveApprove), c.RejectOvertimeApplication)
	}

	// 出差管理
	trips := api.Group("/attendance/trips")
	{
		trips.GET("", middleware.RequirePermission(model.PermissionLeaveList), c.GetBusinessTripApplicationList)
		trips.POST("", middleware.RequirePermission(model.PermissionLeaveCreate), c.CreateBusinessTripApplication)
		trips.PUT("/:id", middleware.RequirePermission(model.PermissionLeaveCreate), c.UpdateBusinessTripApplication)
		trips.DELETE("/:id", middleware.RequirePermission(model.PermissionLeaveCreate), c.DeleteBusinessTripApplication)
		trips.PUT("/:id/approve", middleware.RequirePermission(model.PermissionLeaveApprove), c.ApproveBusinessTripApplication)
		trips.PUT("/:id/reject", middleware.RequirePermission(model.PermissionLeaveApprove), c.RejectBusinessTripApplication)
	}
}

// GetAttendanceRuleList 获取考勤规则列表
func (c *AttendanceController) GetAttendanceRuleList(ctx *gin.Context) {
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	rules, total, err := c.attendanceService.GetAttendanceRuleList(status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  rules,
		"total": total,
	})
}

// GetAttendanceRuleByID 根据ID获取考勤规则
func (c *AttendanceController) GetAttendanceRuleByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	rule, err := c.attendanceService.GetAttendanceRuleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rule)
}

// CreateAttendanceRule 创建考勤规则
func (c *AttendanceController) CreateAttendanceRule(ctx *gin.Context) {
	var rule model.AttendanceRule
	if err := ctx.ShouldBindJSON(&rule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	rule.CreatedBy = uint(1)

	if err := c.attendanceService.CreateAttendanceRule(&rule); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, rule)
}

// UpdateAttendanceRule 更新考勤规则
func (c *AttendanceController) UpdateAttendanceRule(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var rule model.AttendanceRule
	if err := ctx.ShouldBindJSON(&rule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rule.ID = uint(id)
	if err := c.attendanceService.UpdateAttendanceRule(&rule); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rule)
}

// DeleteAttendanceRule 删除考勤规则
func (c *AttendanceController) DeleteAttendanceRule(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.attendanceService.DeleteAttendanceRule(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetAttendanceRecordList 获取考勤记录列表
func (c *AttendanceController) GetAttendanceRecordList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	startDate, _ := time.Parse("2006-01-02", ctx.Query("start_date"))
	endDate, _ := time.Parse("2006-01-02", ctx.Query("end_date"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	var startPtr, endPtr *time.Time
	if !startDate.IsZero() {
		startPtr = &startDate
	}
	if !endDate.IsZero() {
		endPtr = &endDate
	}

	records, total, err := c.attendanceService.GetAttendanceRecordList(uint(employeeID), status, startPtr, endPtr, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
	})
}

// GetAttendanceRecordByID 根据ID获取考勤记录
func (c *AttendanceController) GetAttendanceRecordByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	record, err := c.attendanceService.GetAttendanceRecordByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// CreateAttendanceRecord 创建考勤记录
func (c *AttendanceController) CreateAttendanceRecord(ctx *gin.Context) {
	var record model.AttendanceRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.attendanceService.CreateAttendanceRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

// UpdateAttendanceRecord 更新考勤记录
func (c *AttendanceController) UpdateAttendanceRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var record model.AttendanceRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = uint(id)
	if err := c.attendanceService.UpdateAttendanceRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// DeleteAttendanceRecord 删除考勤记录
func (c *AttendanceController) DeleteAttendanceRecord(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.attendanceService.DeleteAttendanceRecord(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetLeaveApplicationList 获取请假申请列表
func (c *AttendanceController) GetLeaveApplicationList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	startDate, _ := time.Parse("2006-01-02", ctx.Query("start_date"))
	endDate, _ := time.Parse("2006-01-02", ctx.Query("end_date"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	var startPtr, endPtr *time.Time
	if !startDate.IsZero() {
		startPtr = &startDate
	}
	if !endDate.IsZero() {
		endPtr = &endDate
	}

	applications, total, err := c.attendanceService.GetLeaveApplicationList(uint(employeeID), status, startPtr, endPtr, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  applications,
		"total": total,
	})
}

// GetLeaveApplicationByID 根据ID获取请假申请
func (c *AttendanceController) GetLeaveApplicationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	application, err := c.attendanceService.GetLeaveApplicationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// CreateLeaveApplication 创建请假申请
func (c *AttendanceController) CreateLeaveApplication(ctx *gin.Context) {
	var application model.LeaveApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.attendanceService.CreateLeaveApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, application)
}

// UpdateLeaveApplication 更新请假申请
func (c *AttendanceController) UpdateLeaveApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var application model.LeaveApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application.ID = uint(id)
	if err := c.attendanceService.UpdateLeaveApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// DeleteLeaveApplication 删除请假申请
func (c *AttendanceController) DeleteLeaveApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.attendanceService.DeleteLeaveApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetOvertimeApplicationList 获取加班申请列表
func (c *AttendanceController) GetOvertimeApplicationList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	startDate, _ := time.Parse("2006-01-02", ctx.Query("start_date"))
	endDate, _ := time.Parse("2006-01-02", ctx.Query("end_date"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	var startPtr, endPtr *time.Time
	if !startDate.IsZero() {
		startPtr = &startDate
	}
	if !endDate.IsZero() {
		endPtr = &endDate
	}

	applications, total, err := c.attendanceService.GetOvertimeApplicationList(uint(employeeID), status, startPtr, endPtr, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  applications,
		"total": total,
	})
}

// GetOvertimeApplicationByID 根据ID获取加班申请
func (c *AttendanceController) GetOvertimeApplicationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	application, err := c.attendanceService.GetOvertimeApplicationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// CreateOvertimeApplication 创建加班申请
func (c *AttendanceController) CreateOvertimeApplication(ctx *gin.Context) {
	var application model.OvertimeApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.attendanceService.CreateOvertimeApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, application)
}

// UpdateOvertimeApplication 更新加班申请
func (c *AttendanceController) UpdateOvertimeApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var application model.OvertimeApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application.ID = uint(id)
	if err := c.attendanceService.UpdateOvertimeApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// DeleteOvertimeApplication 删除加班申请
func (c *AttendanceController) DeleteOvertimeApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.attendanceService.DeleteOvertimeApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetBusinessTripApplicationList 获取出差申请列表
func (c *AttendanceController) GetBusinessTripApplicationList(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Query("employee_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	startDate, _ := time.Parse("2006-01-02", ctx.Query("start_date"))
	endDate, _ := time.Parse("2006-01-02", ctx.Query("end_date"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	var startPtr, endPtr *time.Time
	if !startDate.IsZero() {
		startPtr = &startDate
	}
	if !endDate.IsZero() {
		endPtr = &endDate
	}

	applications, total, err := c.attendanceService.GetBusinessTripApplicationList(uint(employeeID), status, startPtr, endPtr, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  applications,
		"total": total,
	})
}

// GetBusinessTripApplicationByID 根据ID获取出差申请
func (c *AttendanceController) GetBusinessTripApplicationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	application, err := c.attendanceService.GetBusinessTripApplicationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// CreateBusinessTripApplication 创建出差申请
func (c *AttendanceController) CreateBusinessTripApplication(ctx *gin.Context) {
	var application model.BusinessTripApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.attendanceService.CreateBusinessTripApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, application)
}

// UpdateBusinessTripApplication 更新出差申请
func (c *AttendanceController) UpdateBusinessTripApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var application model.BusinessTripApplication
	if err := ctx.ShouldBindJSON(&application); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application.ID = uint(id)
	if err := c.attendanceService.UpdateBusinessTripApplication(&application); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, application)
}

// DeleteBusinessTripApplication 删除出差申请
func (c *AttendanceController) DeleteBusinessTripApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.attendanceService.DeleteBusinessTripApplication(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveLeaveApplication 审批请假申请
func (c *AttendanceController) ApproveLeaveApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.attendanceService.ApproveLeaveApplication(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectLeaveApplication 驳回请假申请
func (c *AttendanceController) RejectLeaveApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.attendanceService.RejectLeaveApplication(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveOvertimeApplication 审批加班申请
func (c *AttendanceController) ApproveOvertimeApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.attendanceService.ApproveOvertimeApplication(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectOvertimeApplication 驳回加班申请
func (c *AttendanceController) RejectOvertimeApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.attendanceService.RejectOvertimeApplication(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveBusinessTripApplication 审批出差申请
func (c *AttendanceController) ApproveBusinessTripApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.attendanceService.ApproveBusinessTripApplication(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectBusinessTripApplication 驳回出差申请
func (c *AttendanceController) RejectBusinessTripApplication(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.attendanceService.RejectBusinessTripApplication(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
