package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type ApprovalService struct {
	db *gorm.DB
}

func NewApprovalService(db *gorm.DB) *ApprovalService {
	return &ApprovalService{db: db}
}

// GetApprovalTypeList 获取审批类型列表
func (s *ApprovalService) GetApprovalTypeList() ([]model.ApprovalType, error) {
	var types []model.ApprovalType
	err := s.db.Order("sort asc").Find(&types).Error
	return types, err
}

// CreateApprovalType 创建审批类型
func (s *ApprovalService) CreateApprovalType(approvalType *model.ApprovalType) error {
	return s.db.Create(approvalType).Error
}

// UpdateApprovalType 更新审批类型
func (s *ApprovalService) UpdateApprovalType(approvalType *model.ApprovalType) error {
	if approvalType.ID == 0 {
		return errors.New("approval type id is required")
	}
	return s.db.Model(approvalType).Updates(approvalType).Error
}

// DeleteApprovalType 删除审批类型
func (s *ApprovalService) DeleteApprovalType(id uint) error {
	// 检查是否有关联的审批流程
	var count int64
	if err := s.db.Model(&model.ApprovalFlow{}).Where("approval_type_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete approval type with associated flows")
	}

	return s.db.Delete(&model.ApprovalType{}, id).Error
}

// GetApprovalFlowList 获取审批流程列表
func (s *ApprovalService) GetApprovalFlowList(typeID uint) ([]model.ApprovalFlow, error) {
	var flows []model.ApprovalFlow
	query := s.db.Model(&model.ApprovalFlow{})
	if typeID > 0 {
		query = query.Where("approval_type_id = ?", typeID)
	}
	err := query.Find(&flows).Error
	return flows, err
}

// CreateApprovalFlow 创建审批流程
func (s *ApprovalService) CreateApprovalFlow(flow *model.ApprovalFlow) error {
	return s.db.Create(flow).Error
}

// UpdateApprovalFlow 更新审批流程
func (s *ApprovalService) UpdateApprovalFlow(flow *model.ApprovalFlow) error {
	if flow.ID == 0 {
		return errors.New("approval flow id is required")
	}
	return s.db.Model(flow).Updates(flow).Error
}

// DeleteApprovalFlow 删除审批流程
func (s *ApprovalService) DeleteApprovalFlow(id uint) error {
	// 检查是否有关联的审批节点
	var count int64
	if err := s.db.Model(&model.ApprovalNode{}).Where("approval_flow_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete approval flow with associated nodes")
	}

	return s.db.Delete(&model.ApprovalFlow{}, id).Error
}

// GetApprovalNodeList 获取审批节点列表
func (s *ApprovalService) GetApprovalNodeList(flowID uint) ([]model.ApprovalNode, error) {
	var nodes []model.ApprovalNode
	err := s.db.Where("approval_flow_id = ?", flowID).Order("sort asc").Find(&nodes).Error
	return nodes, err
}

// CreateApprovalNode 创建审批节点
func (s *ApprovalService) CreateApprovalNode(node *model.ApprovalNode) error {
	return s.db.Create(node).Error
}

// UpdateApprovalNode 更新审批节点
func (s *ApprovalService) UpdateApprovalNode(node *model.ApprovalNode) error {
	if node.ID == 0 {
		return errors.New("approval node id is required")
	}
	return s.db.Model(node).Updates(node).Error
}

// DeleteApprovalNode 删除审批节点
func (s *ApprovalService) DeleteApprovalNode(id uint) error {
	return s.db.Delete(&model.ApprovalNode{}, id).Error
}

// GetApprovalRecordList 获取审批记录列表
func (s *ApprovalService) GetApprovalRecordList(userID uint, status int, page, pageSize int) ([]model.ApprovalRecord, int64, error) {
	var records []model.ApprovalRecord
	var total int64

	query := s.db.Model(&model.ApprovalRecord{})
	if userID > 0 {
		query = query.Where("applicant_id = ?", userID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// CreateApprovalRecord 创建审批记录
func (s *ApprovalService) CreateApprovalRecord(record *model.ApprovalRecord) error {
	// 获取审批流程的第一个节点
	var firstNode model.ApprovalNode
	err := s.db.Where("approval_flow_id = ?", record.ApprovalFlowID).Order("sort asc").First(&firstNode).Error
	if err != nil {
		return err
	}

	record.CurrentNodeID = firstNode.ID
	record.Status = 2 // 设置为审批中状态

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建审批记录
		if err := tx.Create(record).Error; err != nil {
			return err
		}

		// 创建节点记录
		nodeRecord := &model.ApprovalNodeRecord{
			ApprovalRecordID: record.ID,
			ApprovalNodeID:   firstNode.ID,
			ApproverID:       *firstNode.ApproverID, // 假设是指定人员审批
			Status:           1,                     // 待审批状态
		}
		return tx.Create(nodeRecord).Error
	})
}

// ApproveRecord 审批通过
func (s *ApprovalService) ApproveRecord(recordID, nodeID, approverID uint, comment string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 更新当前节点记录
		nodeRecord := &model.ApprovalNodeRecord{}
		err := tx.Where("approval_record_id = ? AND approval_node_id = ? AND approver_id = ?", recordID, nodeID, approverID).First(nodeRecord).Error
		if err != nil {
			return err
		}

		nodeRecord.Status = 2 // 已通过
		nodeRecord.Comment = comment
		if err := tx.Save(nodeRecord).Error; err != nil {
			return err
		}

		// 获取下一个节点
		var nextNode model.ApprovalNode
		var currentNode model.ApprovalNode
		if err := tx.First(&currentNode, nodeRecord.ApprovalNodeID).Error; err != nil {
			return err
		}
		err = tx.Where("approval_flow_id = ? AND sort > ?", currentNode.ApprovalFlowID, currentNode.Sort).Order("sort asc").First(&nextNode).Error
		if err == gorm.ErrRecordNotFound {
			// 没有下一个节点，审批流程结束
			return tx.Model(&model.ApprovalRecord{}).Where("id = ?", recordID).Update("status", 3).Error
		}
		if err != nil {
			return err
		}

		// 创建下一个节点记录
		newNodeRecord := &model.ApprovalNodeRecord{
			ApprovalRecordID: recordID,
			ApprovalNodeID:   nextNode.ID,
			ApproverID:       *nextNode.ApproverID, // 假设是指定人员审批
			Status:           1,                    // 待审批状态
		}
		if err := tx.Create(newNodeRecord).Error; err != nil {
			return err
		}

		// 更新审批记录的当前节点
		return tx.Model(&model.ApprovalRecord{}).Where("id = ?", recordID).Update("current_node_id", nextNode.ID).Error
	})
}

// RejectRecord 审批驳回
func (s *ApprovalService) RejectRecord(recordID, nodeID, approverID uint, comment string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 更新当前节点记录
		nodeRecord := &model.ApprovalNodeRecord{}
		err := tx.Where("approval_record_id = ? AND approval_node_id = ? AND approver_id = ?", recordID, nodeID, approverID).First(nodeRecord).Error
		if err != nil {
			return err
		}

		nodeRecord.Status = 3 // 已驳回
		nodeRecord.Comment = comment
		if err := tx.Save(nodeRecord).Error; err != nil {
			return err
		}

		// 更新审批记录状态为已驳回
		return tx.Model(&model.ApprovalRecord{}).Where("id = ?", recordID).Update("status", 4).Error
	})
}

// GetPendingApprovalList 获取待审批列表
func (s *ApprovalService) GetPendingApprovalList(approverID uint, page, pageSize int) ([]model.ApprovalRecord, int64, error) {
	var records []model.ApprovalRecord
	var total int64

	subQuery := s.db.Model(&model.ApprovalNodeRecord{}).
		Select("approval_record_id").
		Where("approver_id = ? AND status = ?", approverID, 1)

	query := s.db.Model(&model.ApprovalRecord{}).Where("id IN (?)", subQuery)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}
