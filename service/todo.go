package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

// GetTodoList 获取待办事项列表
func (s *TodoService) GetTodoList(userID uint, status int, page, pageSize int) ([]model.Todo, int64, error) {
	var todos []model.Todo
	var total int64

	query := s.db.Model(&model.Todo{}).Where("user_id = ?", userID)
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&todos).Error
	if err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}

// CreateTodo 创建待办事项
func (s *TodoService) CreateTodo(todo *model.Todo) error {
	return s.db.Create(todo).Error
}

// UpdateTodo 更新待办事项
func (s *TodoService) UpdateTodo(todo *model.Todo) error {
	if todo.ID == 0 {
		return errors.New("todo id is required")
	}
	return s.db.Model(todo).Updates(todo).Error
}

// DeleteTodo 删除待办事项
func (s *TodoService) DeleteTodo(id, userID uint) error {
	return s.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Todo{}).Error
}

// MarkAsCompleted 标记待办事项为已完成
func (s *TodoService) MarkAsCompleted(id, userID uint) error {
	now := time.Now()
	return s.db.Model(&model.Todo{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"status":       2,
			"completed_at": &now,
		}).Error
}

// MarkAsUncompleted 标记待办事项为未完成
func (s *TodoService) MarkAsUncompleted(id, userID uint) error {
	return s.db.Model(&model.Todo{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"status":       1,
			"completed_at": nil,
		}).Error
}

// CreateApprovalTodo 创建审批任务待办事项
func (s *TodoService) CreateApprovalTodo(approvalRecord *model.ApprovalRecord, approverID uint) error {
	todo := &model.Todo{
		Title:   approvalRecord.Title,
		Content: approvalRecord.Content,
		Type:    1, // 审批任务
		UserID:  approverID,
	}
	return s.db.Create(todo).Error
}

// DeleteApprovalTodo 删除审批任务待办事项
func (s *TodoService) DeleteApprovalTodo(approvalRecordID, approverID uint) error {
	return s.db.Where("type = ? AND user_id = ? AND content LIKE ?", 1, approverID, fmt.Sprintf("%%审批记录ID:%d%%", approvalRecordID)).
		Delete(&model.Todo{}).Error
}
