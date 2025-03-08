package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type MeetingService struct {
	db *gorm.DB
}

func NewMeetingService(db *gorm.DB) *MeetingService {
	return &MeetingService{db: db}
}

// GetMeetingRoomList 获取会议室列表
func (s *MeetingService) GetMeetingRoomList(status int, keyword string, page, pageSize int) ([]model.MeetingRoom, int64, error) {
	var rooms []model.MeetingRoom
	var total int64

	query := s.db.Model(&model.MeetingRoom{})
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("name LIKE ? OR location LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("sort asc, created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&rooms).Error
	if err != nil {
		return nil, 0, err
	}

	return rooms, total, nil
}

// GetMeetingRoomByID 根据ID获取会议室
func (s *MeetingService) GetMeetingRoomByID(id uint) (*model.MeetingRoom, error) {
	var room model.MeetingRoom
	err := s.db.First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

// CreateMeetingRoom 创建会议室
func (s *MeetingService) CreateMeetingRoom(room *model.MeetingRoom) error {
	return s.db.Create(room).Error
}

// UpdateMeetingRoom 更新会议室
func (s *MeetingService) UpdateMeetingRoom(room *model.MeetingRoom) error {
	if room.ID == 0 {
		return errors.New("meeting room id is required")
	}
	return s.db.Model(room).Updates(room).Error
}

// DeleteMeetingRoom 删除会议室
func (s *MeetingService) DeleteMeetingRoom(id uint) error {
	// 检查是否有预约记录
	var count int64
	if err := s.db.Model(&model.MeetingReservation{}).Where("room_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete meeting room with reservations")
	}

	return s.db.Delete(&model.MeetingRoom{}, id).Error
}

// GetMeetingReservationList 获取会议室预约列表
func (s *MeetingService) GetMeetingReservationList(roomID, userID uint, status int, page, pageSize int) ([]model.MeetingReservation, int64, error) {
	var reservations []model.MeetingReservation
	var total int64

	query := s.db.Model(&model.MeetingReservation{})
	if roomID > 0 {
		query = query.Where("room_id = ?", roomID)
	}
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&reservations).Error
	if err != nil {
		return nil, 0, err
	}

	return reservations, total, nil
}

// GetMeetingReservationByID 根据ID获取会议室预约
func (s *MeetingService) GetMeetingReservationByID(id uint) (*model.MeetingReservation, error) {
	var reservation model.MeetingReservation
	err := s.db.First(&reservation, id).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

// CreateMeetingReservation 创建会议室预约
func (s *MeetingService) CreateMeetingReservation(reservation *model.MeetingReservation) error {
	// 检查会议室是否存在
	var room model.MeetingRoom
	if err := s.db.First(&room, reservation.RoomID).Error; err != nil {
		return errors.New("meeting room not found")
	}

	// 检查会议室是否可用
	if room.Status != 1 {
		return errors.New("meeting room is not available")
	}

	// 检查预约人是否存在
	var user model.Employee
	if err := s.db.First(&user, reservation.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	// 检查预约部门是否存在
	var department model.Department
	if err := s.db.First(&department, reservation.DepartmentID).Error; err != nil {
		return errors.New("department not found")
	}

	// 检查时间段内是否有其他预约
	var count int64
	err := s.db.Model(&model.MeetingReservation{}).
		Where("room_id = ? AND status IN (1,2) AND ((start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))",
			reservation.RoomID,
			reservation.StartTime, reservation.EndTime,
			reservation.StartTime, reservation.EndTime).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("meeting room is already booked during this period")
	}

	return s.db.Create(reservation).Error
}

// UpdateMeetingReservation 更新会议室预约
func (s *MeetingService) UpdateMeetingReservation(reservation *model.MeetingReservation) error {
	if reservation.ID == 0 {
		return errors.New("reservation id is required")
	}

	// 检查会议室是否存在
	var room model.MeetingRoom
	if err := s.db.First(&room, reservation.RoomID).Error; err != nil {
		return errors.New("meeting room not found")
	}

	// 检查预约人是否存在
	var user model.Employee
	if err := s.db.First(&user, reservation.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	// 检查预约部门是否存在
	var department model.Department
	if err := s.db.First(&department, reservation.DepartmentID).Error; err != nil {
		return errors.New("department not found")
	}

	// 检查时间段内是否有其他预约
	var count int64
	err := s.db.Model(&model.MeetingReservation{}).
		Where("id != ? AND room_id = ? AND status IN (1,2) AND ((start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))",
			reservation.ID, reservation.RoomID,
			reservation.StartTime, reservation.EndTime,
			reservation.StartTime, reservation.EndTime).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("meeting room is already booked during this period")
	}

	return s.db.Model(reservation).Updates(reservation).Error
}

// DeleteMeetingReservation 删除会议室预约
func (s *MeetingService) DeleteMeetingReservation(id uint) error {
	// 只能删除待审批的预约
	result := s.db.Where("id = ? AND status = ?", id, 1).Delete(&model.MeetingReservation{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only delete pending reservations")
	}
	return nil
}

// ApproveMeetingReservation 审批通过会议室预约
func (s *MeetingService) ApproveMeetingReservation(id, approverID uint) error {
	now := time.Now()
	result := s.db.Model(&model.MeetingReservation{}).
		Where("id = ? AND status = ?", id, 1).
		Updates(map[string]interface{}{
			"status":        2,
			"approver_id":   approverID,
			"approval_time": &now,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only approve pending reservations")
	}
	return nil
}

// RejectMeetingReservation 审批驳回会议室预约
func (s *MeetingService) RejectMeetingReservation(id, approverID uint) error {
	now := time.Now()
	result := s.db.Model(&model.MeetingReservation{}).
		Where("id = ? AND status = ?", id, 1).
		Updates(map[string]interface{}{
			"status":        3,
			"approver_id":   approverID,
			"approval_time": &now,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only reject pending reservations")
	}
	return nil
}

// CancelMeetingReservation 取消会议室预约
func (s *MeetingService) CancelMeetingReservation(id uint, reason string) error {
	now := time.Now()
	result := s.db.Model(&model.MeetingReservation{}).
		Where("id = ? AND status IN (1,2)", id).
		Updates(map[string]interface{}{
			"status":        4,
			"cancel_reason": reason,
			"cancel_time":   &now,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only cancel pending or approved reservations")
	}
	return nil
}

// CheckInMeeting 会议签到
func (s *MeetingService) CheckInMeeting(id uint) error {
	now := time.Now()
	result := s.db.Model(&model.MeetingReservation{}).
		Where("id = ? AND status = ? AND check_in_time IS NULL", id, 2).
		Update("check_in_time", &now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only check in approved meetings that haven't been checked in")
	}
	return nil
}

// CheckOutMeeting 会议签退
func (s *MeetingService) CheckOutMeeting(id uint) error {
	now := time.Now()
	result := s.db.Model(&model.MeetingReservation{}).
		Where("id = ? AND status = ? AND check_in_time IS NOT NULL AND check_out_time IS NULL", id, 2).
		Update("check_out_time", &now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only check out approved meetings that have been checked in but not checked out")
	}
	return nil
}

// GetMeetingMinutesList 获取会议纪要列表
func (s *MeetingService) GetMeetingMinutesList(reservationID uint, page, pageSize int) ([]model.MeetingMinutes, int64, error) {
	var minutes []model.MeetingMinutes
	var total int64

	query := s.db.Model(&model.MeetingMinutes{})
	if reservationID > 0 {
		query = query.Where("reservation_id = ?", reservationID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&minutes).Error
	if err != nil {
		return nil, 0, err
	}

	return minutes, total, nil
}

// GetMeetingMinutesByID 根据ID获取会议纪要
func (s *MeetingService) GetMeetingMinutesByID(id uint) (*model.MeetingMinutes, error) {
	var minutes model.MeetingMinutes
	err := s.db.First(&minutes, id).Error
	if err != nil {
		return nil, err
	}
	return &minutes, nil
}

// CreateMeetingMinutes 创建会议纪要
func (s *MeetingService) CreateMeetingMinutes(minutes *model.MeetingMinutes) error {
	// 检查会议预约是否存在
	var reservation model.MeetingReservation
	if err := s.db.First(&reservation, minutes.ReservationID).Error; err != nil {
		return errors.New("meeting reservation not found")
	}

	// 检查会议是否已签到
	if reservation.CheckInTime == nil {
		return errors.New("meeting has not been checked in")
	}

	return s.db.Create(minutes).Error
}

// UpdateMeetingMinutes 更新会议纪要
func (s *MeetingService) UpdateMeetingMinutes(minutes *model.MeetingMinutes) error {
	if minutes.ID == 0 {
		return errors.New("minutes id is required")
	}

	// 检查会议预约是否存在
	var reservation model.MeetingReservation
	if err := s.db.First(&reservation, minutes.ReservationID).Error; err != nil {
		return errors.New("meeting reservation not found")
	}

	return s.db.Model(minutes).Updates(minutes).Error
}

// DeleteMeetingMinutes 删除会议纪要
func (s *MeetingService) DeleteMeetingMinutes(id uint) error {
	return s.db.Delete(&model.MeetingMinutes{}, id).Error
}

// GetMeetingRoomMaintenanceList 获取会议室维护记录列表
func (s *MeetingService) GetMeetingRoomMaintenanceList(roomID uint, status int, page, pageSize int) ([]model.MeetingRoomMaintenance, int64, error) {
	var maintenances []model.MeetingRoomMaintenance
	var total int64

	query := s.db.Model(&model.MeetingRoomMaintenance{})
	if roomID > 0 {
		query = query.Where("room_id = ?", roomID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&maintenances).Error
	if err != nil {
		return nil, 0, err
	}

	return maintenances, total, nil
}

// GetMeetingRoomMaintenanceByID 根据ID获取会议室维护记录
func (s *MeetingService) GetMeetingRoomMaintenanceByID(id uint) (*model.MeetingRoomMaintenance, error) {
	var maintenance model.MeetingRoomMaintenance
	err := s.db.First(&maintenance, id).Error
	if err != nil {
		return nil, err
	}
	return &maintenance, nil
}

// CreateMeetingRoomMaintenance 创建会议室维护记录
func (s *MeetingService) CreateMeetingRoomMaintenance(maintenance *model.MeetingRoomMaintenance) error {
	// 检查会议室是否存在
	var room model.MeetingRoom
	if err := s.db.First(&room, maintenance.RoomID).Error; err != nil {
		return errors.New("meeting room not found")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建维护记录
		if err := tx.Create(maintenance).Error; err != nil {
			return err
		}

		// 更新会议室状态为维护中
		if err := tx.Model(&room).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateMeetingRoomMaintenance 更新会议室维护记录
func (s *MeetingService) UpdateMeetingRoomMaintenance(maintenance *model.MeetingRoomMaintenance) error {
	if maintenance.ID == 0 {
		return errors.New("maintenance id is required")
	}

	// 检查会议室是否存在
	var room model.MeetingRoom
	if err := s.db.First(&room, maintenance.RoomID).Error; err != nil {
		return errors.New("meeting room not found")
	}

	return s.db.Model(maintenance).Updates(maintenance).Error
}

// DeleteMeetingRoomMaintenance 删除会议室维护记录
func (s *MeetingService) DeleteMeetingRoomMaintenance(id uint) error {
	return s.db.Delete(&model.MeetingRoomMaintenance{}, id).Error
}

// CompleteMeetingRoomMaintenance 完成会议室维护
func (s *MeetingService) CompleteMeetingRoomMaintenance(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取维护记录
		var maintenance model.MeetingRoomMaintenance
		if err := tx.First(&maintenance, id).Error; err != nil {
			return err
		}

		// 更新维护记录状态为已完成
		if err := tx.Model(&maintenance).Update("status", 3).Error; err != nil {
			return err
		}

		// 更新会议室状态为可用
		if err := tx.Model(&model.MeetingRoom{}).Where("id = ?", maintenance.RoomID).Update("status", 1).Error; err != nil {
			return err
		}

		return nil
	})
}
