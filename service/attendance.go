package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type AttendanceService struct {
	db *gorm.DB
}

func NewAttendanceService(db *gorm.DB) *AttendanceService {
	return &AttendanceService{db: db}
}

// GetAttendanceRuleList 获取考勤规则列表
func (s *AttendanceService) GetAttendanceRuleList(status int, keyword string, page, pageSize int) ([]model.AttendanceRule, int64, error) {
	var rules []model.AttendanceRule
	var total int64

	query := s.db.Model(&model.AttendanceRule{})
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&rules).Error
	if err != nil {
		return nil, 0, err
	}

	return rules, total, nil
}

// GetAttendanceRuleByID 根据ID获取考勤规则
func (s *AttendanceService) GetAttendanceRuleByID(id uint) (*model.AttendanceRule, error) {
	var rule model.AttendanceRule
	err := s.db.First(&rule, id).Error
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

// CreateAttendanceRule 创建考勤规则
func (s *AttendanceService) CreateAttendanceRule(rule *model.AttendanceRule) error {
	return s.db.Create(rule).Error
}

// UpdateAttendanceRule 更新考勤规则
func (s *AttendanceService) UpdateAttendanceRule(rule *model.AttendanceRule) error {
	if rule.ID == 0 {
		return errors.New("attendance rule id is required")
	}
	return s.db.Model(rule).Updates(rule).Error
}

// DeleteAttendanceRule 删除考勤规则
func (s *AttendanceService) DeleteAttendanceRule(id uint) error {
	return s.db.Delete(&model.AttendanceRule{}, id).Error
}

// GetAttendanceRecordList 获取考勤记录列表
func (s *AttendanceService) GetAttendanceRecordList(employeeID uint, status int, startDate, endDate *time.Time, page, pageSize int) ([]model.AttendanceRecord, int64, error) {
	var records []model.AttendanceRecord
	var total int64

	query := s.db.Model(&model.AttendanceRecord{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if startDate != nil {
		query = query.Where("date >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("date <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("date desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// GetAttendanceRecordByID 根据ID获取考勤记录
func (s *AttendanceService) GetAttendanceRecordByID(id uint) (*model.AttendanceRecord, error) {
	var record model.AttendanceRecord
	err := s.db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// CreateAttendanceRecord 创建考勤记录
func (s *AttendanceService) CreateAttendanceRecord(record *model.AttendanceRecord) error {
	return s.db.Create(record).Error
}

// UpdateAttendanceRecord 更新考勤记录
func (s *AttendanceService) UpdateAttendanceRecord(record *model.AttendanceRecord) error {
	if record.ID == 0 {
		return errors.New("attendance record id is required")
	}
	return s.db.Model(record).Updates(record).Error
}

// DeleteAttendanceRecord 删除考勤记录
func (s *AttendanceService) DeleteAttendanceRecord(id uint) error {
	return s.db.Delete(&model.AttendanceRecord{}, id).Error
}

// GetLeaveApplicationList 获取请假申请列表
func (s *AttendanceService) GetLeaveApplicationList(employeeID uint, status int, startDate, endDate *time.Time, page, pageSize int) ([]model.LeaveApplication, int64, error) {
	var applications []model.LeaveApplication
	var total int64

	query := s.db.Model(&model.LeaveApplication{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if startDate != nil {
		query = query.Where("start_time >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("end_time <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applications).Error
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

// GetLeaveApplicationByID 根据ID获取请假申请
func (s *AttendanceService) GetLeaveApplicationByID(id uint) (*model.LeaveApplication, error) {
	var application model.LeaveApplication
	err := s.db.First(&application, id).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

// CreateLeaveApplication 创建请假申请
func (s *AttendanceService) CreateLeaveApplication(application *model.LeaveApplication) error {
	// 计算请假天数
	days := application.EndTime.Sub(*application.StartTime).Hours() / 24
	application.Days = float64(days)

	return s.db.Create(application).Error
}

// UpdateLeaveApplication 更新请假申请
func (s *AttendanceService) UpdateLeaveApplication(application *model.LeaveApplication) error {
	if application.ID == 0 {
		return errors.New("leave application id is required")
	}

	// 计算请假天数
	days := application.EndTime.Sub(*application.StartTime).Hours() / 24
	application.Days = float64(days)

	return s.db.Model(application).Updates(application).Error
}

// DeleteLeaveApplication 删除请假申请
func (s *AttendanceService) DeleteLeaveApplication(id uint) error {
	return s.db.Delete(&model.LeaveApplication{}, id).Error
}

// GetOvertimeApplicationList 获取加班申请列表
func (s *AttendanceService) GetOvertimeApplicationList(employeeID uint, status int, startDate, endDate *time.Time, page, pageSize int) ([]model.OvertimeApplication, int64, error) {
	var applications []model.OvertimeApplication
	var total int64

	query := s.db.Model(&model.OvertimeApplication{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if startDate != nil {
		query = query.Where("start_time >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("end_time <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applications).Error
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

// GetOvertimeApplicationByID 根据ID获取加班申请
func (s *AttendanceService) GetOvertimeApplicationByID(id uint) (*model.OvertimeApplication, error) {
	var application model.OvertimeApplication
	err := s.db.First(&application, id).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

// CreateOvertimeApplication 创建加班申请
func (s *AttendanceService) CreateOvertimeApplication(application *model.OvertimeApplication) error {
	// 计算加班小时数
	hours := application.EndTime.Sub(*application.StartTime).Hours()
	application.Hours = float64(hours)

	return s.db.Create(application).Error
}

// UpdateOvertimeApplication 更新加班申请
func (s *AttendanceService) UpdateOvertimeApplication(application *model.OvertimeApplication) error {
	if application.ID == 0 {
		return errors.New("overtime application id is required")
	}

	// 计算加班小时数
	hours := application.EndTime.Sub(*application.StartTime).Hours()
	application.Hours = float64(hours)

	return s.db.Model(application).Updates(application).Error
}

// DeleteOvertimeApplication 删除加班申请
func (s *AttendanceService) DeleteOvertimeApplication(id uint) error {
	return s.db.Delete(&model.OvertimeApplication{}, id).Error
}

// GetBusinessTripApplicationList 获取出差申请列表
func (s *AttendanceService) GetBusinessTripApplicationList(employeeID uint, status int, startDate, endDate *time.Time, page, pageSize int) ([]model.BusinessTripApplication, int64, error) {
	var applications []model.BusinessTripApplication
	var total int64

	query := s.db.Model(&model.BusinessTripApplication{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if startDate != nil {
		query = query.Where("start_time >= ?", startDate)
	}
	if endDate != nil {
		query = query.Where("end_time <= ?", endDate)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applications).Error
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

// GetBusinessTripApplicationByID 根据ID获取出差申请
func (s *AttendanceService) GetBusinessTripApplicationByID(id uint) (*model.BusinessTripApplication, error) {
	var application model.BusinessTripApplication
	err := s.db.First(&application, id).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

// CreateBusinessTripApplication 创建出差申请
func (s *AttendanceService) CreateBusinessTripApplication(application *model.BusinessTripApplication) error {
	// 计算出差天数
	days := application.EndTime.Sub(*application.StartTime).Hours() / 24
	application.Days = float64(days)

	return s.db.Create(application).Error
}

// UpdateBusinessTripApplication 更新出差申请
func (s *AttendanceService) UpdateBusinessTripApplication(application *model.BusinessTripApplication) error {
	if application.ID == 0 {
		return errors.New("business trip application id is required")
	}

	// 计算出差天数
	days := application.EndTime.Sub(*application.StartTime).Hours() / 24
	application.Days = float64(days)

	return s.db.Model(application).Updates(application).Error
}

// DeleteBusinessTripApplication 删除出差申请
func (s *AttendanceService) DeleteBusinessTripApplication(id uint) error {
	return s.db.Delete(&model.BusinessTripApplication{}, id).Error
}

// ApproveLeaveApplication 审批通过请假申请
func (s *AttendanceService) ApproveLeaveApplication(id, approverID uint) error {
	var leave model.LeaveApplication
	if err := s.db.First(&leave, id).Error; err != nil {
		return err
	}

	// 更新状态为已通过
	leave.Status = 2
	if err := s.db.Model(&leave).Updates(map[string]interface{}{
		"status": leave.Status,
	}).Error; err != nil {
		return err
	}

	return nil
}

// RejectLeaveApplication 驳回请假申请
func (s *AttendanceService) RejectLeaveApplication(id, approverID uint) error {
	var leave model.LeaveApplication
	if err := s.db.First(&leave, id).Error; err != nil {
		return err
	}

	// 更新状态为已驳回
	leave.Status = 3
	if err := s.db.Model(&leave).Updates(map[string]interface{}{
		"status": leave.Status,
	}).Error; err != nil {
		return err
	}

	return nil
}

// ApproveOvertimeApplication 审批通过加班申请
func (s *AttendanceService) ApproveOvertimeApplication(id, approverID uint) error {
	var overtime model.OvertimeApplication
	if err := s.db.First(&overtime, id).Error; err != nil {
		return err
	}

	// 更新状态为已通过
	overtime.Status = 2
	if err := s.db.Model(&overtime).Updates(map[string]interface{}{
		"status": overtime.Status,
	}).Error; err != nil {
		return err
	}

	return nil
}

// RejectOvertimeApplication 驳回加班申请
func (s *AttendanceService) RejectOvertimeApplication(id, approverID uint) error {
	var overtime model.OvertimeApplication
	if err := s.db.First(&overtime, id).Error; err != nil {
		return err
	}

	// 更新状态为已驳回
	overtime.Status = 3
	if err := s.db.Model(&overtime).Updates(map[string]interface{}{
		"status": overtime.Status,
	}).Error; err != nil {
		return err
	}

	return nil
}

// ApproveBusinessTripApplication 审批通过出差申请
func (s *AttendanceService) ApproveBusinessTripApplication(id, approverID uint) error {
	var trip model.BusinessTripApplication
	if err := s.db.First(&trip, id).Error; err != nil {
		return err
	}

	// 更新状态为已通过
	trip.Status = 2
	if err := s.db.Model(&trip).Updates(map[string]interface{}{
		"status": trip.Status,
	}).Error; err != nil {
		return err
	}

	return nil
}

// RejectBusinessTripApplication 驳回出差申请
func (s *AttendanceService) RejectBusinessTripApplication(id, approverID uint) error {
	var trip model.BusinessTripApplication
	if err := s.db.First(&trip, id).Error; err != nil {
		return err
	}

	// 更新状态为已驳回
	trip.Status = 3
	if err := s.db.Model(&trip).Updates(map[string]interface{}{
		"status": trip.Status,
	}).Error; err != nil {
		return err
	}

	return nil
}
