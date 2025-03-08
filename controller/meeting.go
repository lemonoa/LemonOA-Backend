package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/lemonoa/LemonOA-Go/middleware"

	"github.com/gin-gonic/gin"
)

type MeetingController struct {
	meetingService *service.MeetingService
}

func NewMeetingController(meetingService *service.MeetingService) *MeetingController {
	return &MeetingController{
		meetingService: meetingService,
	}
}

// RegisterRoutes 注册路由
func (c *MeetingController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.JWT())

	// 会议室管理
	rooms := api.Group("/meeting/rooms")
	{
		rooms.GET("", middleware.RequirePermission(model.PermissionMeetingRoomList), c.GetMeetingRoomList)
		rooms.POST("", middleware.RequirePermission(model.PermissionMeetingRoomCreate), c.CreateMeetingRoom)
		rooms.PUT("/:id", middleware.RequirePermission(model.PermissionMeetingRoomUpdate), c.UpdateMeetingRoom)
		rooms.DELETE("/:id", middleware.RequirePermission(model.PermissionMeetingRoomDelete), c.DeleteMeetingRoom)
	}

	// 会议预约管理
	reservations := api.Group("/meeting/reservations")
	{
		reservations.GET("", middleware.RequirePermission(model.PermissionMeetingReserveList), c.GetMeetingReservationList)
		reservations.POST("", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.CreateMeetingReservation)
		reservations.PUT("/:id", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.UpdateMeetingReservation)
		reservations.DELETE("/:id", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.DeleteMeetingReservation)
		reservations.PUT("/:id/approve", middleware.RequirePermission(model.PermissionMeetingReserveApprove), c.ApproveMeetingReservation)
		reservations.PUT("/:id/reject", middleware.RequirePermission(model.PermissionMeetingReserveApprove), c.RejectMeetingReservation)
		reservations.PUT("/:id/cancel", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.CancelMeetingReservation)
		reservations.PUT("/:id/check-in", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.CheckInMeeting)
		reservations.PUT("/:id/check-out", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.CheckOutMeeting)
	}

	// 会议纪要管理
	minutes := api.Group("/meeting/minutes")
	{
		minutes.GET("", middleware.RequirePermission(model.PermissionMeetingReserveList), c.GetMeetingMinutesList)
		minutes.POST("", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.CreateMeetingMinutes)
		minutes.PUT("/:id", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.UpdateMeetingMinutes)
		minutes.DELETE("/:id", middleware.RequirePermission(model.PermissionMeetingReserveCreate), c.DeleteMeetingMinutes)
	}

	// 会议室维护管理
	maintenance := api.Group("/meeting/maintenance")
	{
		maintenance.GET("", middleware.RequirePermission(model.PermissionMeetingRoomList), c.GetMeetingRoomMaintenanceList)
		maintenance.POST("", middleware.RequirePermission(model.PermissionMeetingRoomCreate), c.CreateMeetingRoomMaintenance)
		maintenance.PUT("/:id", middleware.RequirePermission(model.PermissionMeetingRoomUpdate), c.UpdateMeetingRoomMaintenance)
		maintenance.DELETE("/:id", middleware.RequirePermission(model.PermissionMeetingRoomDelete), c.DeleteMeetingRoomMaintenance)
		maintenance.PUT("/:id/complete", middleware.RequirePermission(model.PermissionMeetingRoomUpdate), c.CompleteMeetingRoomMaintenance)
	}
}

// GetMeetingRoomList 获取会议室列表
func (c *MeetingController) GetMeetingRoomList(ctx *gin.Context) {
	status, _ := strconv.Atoi(ctx.Query("status"))
	keyword := ctx.Query("keyword")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	rooms, total, err := c.meetingService.GetMeetingRoomList(status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  rooms,
		"total": total,
	})
}

// GetMeetingRoomByID 根据ID获取会议室
func (c *MeetingController) GetMeetingRoomByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	room, err := c.meetingService.GetMeetingRoomByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, room)
}

// CreateMeetingRoom 创建会议室
func (c *MeetingController) CreateMeetingRoom(ctx *gin.Context) {
	var room model.MeetingRoom
	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	room.CreatedBy = uint(1)

	if err := c.meetingService.CreateMeetingRoom(&room); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, room)
}

// UpdateMeetingRoom 更新会议室
func (c *MeetingController) UpdateMeetingRoom(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var room model.MeetingRoom
	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room.ID = uint(id)
	if err := c.meetingService.UpdateMeetingRoom(&room); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, room)
}

// DeleteMeetingRoom 删除会议室
func (c *MeetingController) DeleteMeetingRoom(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.meetingService.DeleteMeetingRoom(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetMeetingReservationList 获取会议室预约列表
func (c *MeetingController) GetMeetingReservationList(ctx *gin.Context) {
	roomID, _ := strconv.ParseUint(ctx.Query("room_id"), 10, 32)
	userID, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	reservations, total, err := c.meetingService.GetMeetingReservationList(uint(roomID), uint(userID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  reservations,
		"total": total,
	})
}

// GetMeetingReservationByID 根据ID获取会议室预约
func (c *MeetingController) GetMeetingReservationByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	reservation, err := c.meetingService.GetMeetingReservationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

// CreateMeetingReservation 创建会议室预约
func (c *MeetingController) CreateMeetingReservation(ctx *gin.Context) {
	var reservation model.MeetingReservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.meetingService.CreateMeetingReservation(&reservation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, reservation)
}

// UpdateMeetingReservation 更新会议室预约
func (c *MeetingController) UpdateMeetingReservation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var reservation model.MeetingReservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation.ID = uint(id)
	if err := c.meetingService.UpdateMeetingReservation(&reservation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

// DeleteMeetingReservation 删除会议室预约
func (c *MeetingController) DeleteMeetingReservation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.meetingService.DeleteMeetingReservation(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ApproveMeetingReservation 审批通过会议室预约
func (c *MeetingController) ApproveMeetingReservation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.meetingService.ApproveMeetingReservation(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RejectMeetingReservation 审批驳回会议室预约
func (c *MeetingController) RejectMeetingReservation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取当前用户ID
	approverID := uint(1)

	if err := c.meetingService.RejectMeetingReservation(uint(id), approverID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CancelMeetingReservation 取消会议室预约
func (c *MeetingController) CancelMeetingReservation(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var data struct {
		Reason string `json:"reason"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.meetingService.CancelMeetingReservation(uint(id), data.Reason); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CheckInMeeting 会议签到
func (c *MeetingController) CheckInMeeting(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.meetingService.CheckInMeeting(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CheckOutMeeting 会议签退
func (c *MeetingController) CheckOutMeeting(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.meetingService.CheckOutMeeting(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetMeetingMinutesList 获取会议纪要列表
func (c *MeetingController) GetMeetingMinutesList(ctx *gin.Context) {
	reservationID, _ := strconv.ParseUint(ctx.Query("reservation_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	minutes, total, err := c.meetingService.GetMeetingMinutesList(uint(reservationID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  minutes,
		"total": total,
	})
}

// GetMeetingMinutesByID 根据ID获取会议纪要
func (c *MeetingController) GetMeetingMinutesByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	minutes, err := c.meetingService.GetMeetingMinutesByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, minutes)
}

// CreateMeetingMinutes 创建会议纪要
func (c *MeetingController) CreateMeetingMinutes(ctx *gin.Context) {
	var minutes model.MeetingMinutes
	if err := ctx.ShouldBindJSON(&minutes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	minutes.CreatedBy = uint(1)

	if err := c.meetingService.CreateMeetingMinutes(&minutes); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, minutes)
}

// UpdateMeetingMinutes 更新会议纪要
func (c *MeetingController) UpdateMeetingMinutes(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var minutes model.MeetingMinutes
	if err := ctx.ShouldBindJSON(&minutes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	minutes.ID = uint(id)
	if err := c.meetingService.UpdateMeetingMinutes(&minutes); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, minutes)
}

// DeleteMeetingMinutes 删除会议纪要
func (c *MeetingController) DeleteMeetingMinutes(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.meetingService.DeleteMeetingMinutes(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetMeetingRoomMaintenanceList 获取会议室维护记录列表
func (c *MeetingController) GetMeetingRoomMaintenanceList(ctx *gin.Context) {
	roomID, _ := strconv.ParseUint(ctx.Query("room_id"), 10, 32)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	maintenances, total, err := c.meetingService.GetMeetingRoomMaintenanceList(uint(roomID), status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  maintenances,
		"total": total,
	})
}

// GetMeetingRoomMaintenanceByID 根据ID获取会议室维护记录
func (c *MeetingController) GetMeetingRoomMaintenanceByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	maintenance, err := c.meetingService.GetMeetingRoomMaintenanceByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, maintenance)
}

// CreateMeetingRoomMaintenance 创建会议室维护记录
func (c *MeetingController) CreateMeetingRoomMaintenance(ctx *gin.Context) {
	var maintenance model.MeetingRoomMaintenance
	if err := ctx.ShouldBindJSON(&maintenance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	maintenance.CreatedBy = uint(1)

	if err := c.meetingService.CreateMeetingRoomMaintenance(&maintenance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, maintenance)
}

// UpdateMeetingRoomMaintenance 更新会议室维护记录
func (c *MeetingController) UpdateMeetingRoomMaintenance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var maintenance model.MeetingRoomMaintenance
	if err := ctx.ShouldBindJSON(&maintenance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maintenance.ID = uint(id)
	if err := c.meetingService.UpdateMeetingRoomMaintenance(&maintenance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, maintenance)
}

// DeleteMeetingRoomMaintenance 删除会议室维护记录
func (c *MeetingController) DeleteMeetingRoomMaintenance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.meetingService.DeleteMeetingRoomMaintenance(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CompleteMeetingRoomMaintenance 完成会议室维护
func (c *MeetingController) CompleteMeetingRoomMaintenance(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.meetingService.CompleteMeetingRoomMaintenance(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
