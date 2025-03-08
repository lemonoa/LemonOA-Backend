package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type NoticeService struct {
	db *gorm.DB
}

func NewNoticeService(db *gorm.DB) *NoticeService {
	return &NoticeService{db: db}
}

// GetNoticeList 获取公告列表
func (s *NoticeService) GetNoticeList(typeID uint, status int, keyword string, page, pageSize int) ([]model.Notice, int64, error) {
	var notices []model.Notice
	var total int64

	query := s.db.Model(&model.Notice{})
	if typeID > 0 {
		query = query.Where("type_id = ?", typeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&notices).Error
	if err != nil {
		return nil, 0, err
	}

	return notices, total, nil
}

// GetNoticeByID 根据ID获取公告
func (s *NoticeService) GetNoticeByID(id uint) (*model.Notice, error) {
	var notice model.Notice
	err := s.db.First(&notice, id).Error
	if err != nil {
		return nil, err
	}
	return &notice, nil
}

// CreateNotice 创建公告
func (s *NoticeService) CreateNotice(notice *model.Notice) error {
	// 检查公告类型是否存在
	var noticeType model.NoticeType
	if err := s.db.First(&noticeType, notice.TypeID).Error; err != nil {
		return errors.New("notice type not found")
	}

	return s.db.Create(notice).Error
}

// UpdateNotice 更新公告
func (s *NoticeService) UpdateNotice(notice *model.Notice) error {
	if notice.ID == 0 {
		return errors.New("notice id is required")
	}

	// 检查公告类型是否存在
	var noticeType model.NoticeType
	if err := s.db.First(&noticeType, notice.TypeID).Error; err != nil {
		return errors.New("notice type not found")
	}

	return s.db.Model(notice).Updates(notice).Error
}

// DeleteNotice 删除公告
func (s *NoticeService) DeleteNotice(id uint) error {
	// 检查是否有阅读记录
	var count int64
	if err := s.db.Model(&model.NoticeRead{}).Where("notice_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete notice with read records")
	}

	return s.db.Delete(&model.Notice{}, id).Error
}

// PublishNotice 发布公告
func (s *NoticeService) PublishNotice(id uint) error {
	return s.db.Model(&model.Notice{}).Where("id = ?", id).Update("status", 2).Error
}

// RecallNotice 撤回公告
func (s *NoticeService) RecallNotice(id uint) error {
	return s.db.Model(&model.Notice{}).Where("id = ?", id).Update("status", 3).Error
}

// GetNoticeReadList 获取公告阅读记录列表
func (s *NoticeService) GetNoticeReadList(noticeID uint, page, pageSize int) ([]model.NoticeRead, int64, error) {
	var reads []model.NoticeRead
	var total int64

	query := s.db.Model(&model.NoticeRead{}).Where("notice_id = ?", noticeID)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&reads).Error
	if err != nil {
		return nil, 0, err
	}

	return reads, total, nil
}

// ReadNotice 阅读公告
func (s *NoticeService) ReadNotice(noticeID, userID uint) error {
	// 检查公告是否存在
	var notice model.Notice
	if err := s.db.First(&notice, noticeID).Error; err != nil {
		return errors.New("notice not found")
	}

	// 检查是否已阅读
	var count int64
	if err := s.db.Model(&model.NoticeRead{}).Where("notice_id = ? AND user_id = ?", noticeID, userID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	// 创建阅读记录
	now := time.Now()
	read := &model.NoticeRead{
		NoticeID: noticeID,
		UserID:   userID,
		ReadTime: &now,
	}
	return s.db.Create(read).Error
}

// GetUnreadNoticeCount 获取未读公告数量
func (s *NoticeService) GetUnreadNoticeCount(userID uint) (int64, error) {
	var count int64
	err := s.db.Raw(`
		SELECT COUNT(*) FROM notices n
		WHERE n.status = 2
		AND n.id NOT IN (
			SELECT notice_id FROM notice_reads
			WHERE user_id = ?
		)
	`, userID).Count(&count).Error
	return count, err
}
