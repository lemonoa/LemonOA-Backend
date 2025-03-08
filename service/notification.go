package service

import (
	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type NotificationService struct {
	db *gorm.DB
}

func NewNotificationService(db *gorm.DB) *NotificationService {
	return &NotificationService{db: db}
}

// GetNotificationList 获取消息列表
func (s *NotificationService) GetNotificationList(userID uint, page, pageSize int) ([]model.Notification, int64, error) {
	var notifications []model.Notification
	var total int64

	query := s.db.Model(&model.Notification{}).Where("user_id = ?", userID)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&notifications).Error
	if err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}

// GetUnreadCount 获取未读消息数量
func (s *NotificationService) GetUnreadCount(userID uint) (int64, error) {
	var count int64
	err := s.db.Model(&model.Notification{}).Where("user_id = ? AND status = ?", userID, 1).Count(&count).Error
	return count, err
}

// CreateNotification 创建消息
func (s *NotificationService) CreateNotification(notification *model.Notification) error {
	return s.db.Create(notification).Error
}

// MarkAsRead 标记消息为已读
func (s *NotificationService) MarkAsRead(id, userID uint) error {
	return s.db.Model(&model.Notification{}).Where("id = ? AND user_id = ?", id, userID).Update("status", 2).Error
}

// MarkAllAsRead 标记所有消息为已读
func (s *NotificationService) MarkAllAsRead(userID uint) error {
	return s.db.Model(&model.Notification{}).Where("user_id = ? AND status = ?", userID, 1).Update("status", 2).Error
}

// DeleteNotification 删除消息
func (s *NotificationService) DeleteNotification(id, userID uint) error {
	return s.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Notification{}).Error
}
