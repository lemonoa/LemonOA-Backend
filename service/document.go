package service

import (
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type DocumentService struct {
	db *gorm.DB
}

func NewDocumentService(db *gorm.DB) *DocumentService {
	return &DocumentService{db: db}
}

// GetDocumentList 获取公文列表
func (s *DocumentService) GetDocumentList(typeID uint, status int, keyword string, page, pageSize int) ([]model.Document, int64, error) {
	var documents []model.Document
	var total int64

	query := s.db.Model(&model.Document{})
	if typeID > 0 {
		query = query.Where("type_id = ?", typeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("title LIKE ? OR code LIKE ? OR keywords LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&documents).Error
	if err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}

// GetDocumentByID 根据ID获取公文
func (s *DocumentService) GetDocumentByID(id uint) (*model.Document, error) {
	var document model.Document
	err := s.db.First(&document, id).Error
	if err != nil {
		return nil, err
	}
	return &document, nil
}

// CreateDocument 创建公文
func (s *DocumentService) CreateDocument(document *model.Document) error {
	// 检查公文类型是否存在
	var documentType model.DocumentType
	if err := s.db.First(&documentType, document.TypeID).Error; err != nil {
		return errors.New("document type not found")
	}

	// 检查拟稿人是否存在
	var user model.Employee
	if err := s.db.First(&user, document.DraftUserID).Error; err != nil {
		return errors.New("draft user not found")
	}

	// 检查拟稿部门是否存在
	var department model.Department
	if err := s.db.First(&department, document.DraftDeptID).Error; err != nil {
		return errors.New("draft department not found")
	}

	now := time.Now()
	document.DraftDate = &now

	return s.db.Create(document).Error
}

// UpdateDocument 更新公文
func (s *DocumentService) UpdateDocument(document *model.Document) error {
	if document.ID == 0 {
		return errors.New("document id is required")
	}

	// 检查公文类型是否存在
	var documentType model.DocumentType
	if err := s.db.First(&documentType, document.TypeID).Error; err != nil {
		return errors.New("document type not found")
	}

	// 检查拟稿人是否存在
	var user model.Employee
	if err := s.db.First(&user, document.DraftUserID).Error; err != nil {
		return errors.New("draft user not found")
	}

	// 检查拟稿部门是否存在
	var department model.Department
	if err := s.db.First(&department, document.DraftDeptID).Error; err != nil {
		return errors.New("draft department not found")
	}

	return s.db.Model(document).Updates(document).Error
}

// DeleteDocument 删除公文
func (s *DocumentService) DeleteDocument(id uint) error {
	// 检查公文状态
	var document model.Document
	if err := s.db.First(&document, id).Error; err != nil {
		return err
	}
	if document.Status != 1 {
		return errors.New("can only delete draft documents")
	}

	return s.db.Delete(&model.Document{}, id).Error
}

// SubmitDocument 提交公文审批
func (s *DocumentService) SubmitDocument(id uint, approvers []model.DocumentApproval) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取公文
		var document model.Document
		if err := tx.First(&document, id).Error; err != nil {
			return err
		}

		// 检查公文状态
		if document.Status != 1 {
			return errors.New("can only submit draft documents")
		}

		// 更新公文状态为审批中
		if err := tx.Model(&document).Update("status", 2).Error; err != nil {
			return err
		}

		// 创建审批流程
		for _, approver := range approvers {
			approver.DocumentID = id
			if err := tx.Create(&approver).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// ApproveDocument 审批通过公文
func (s *DocumentService) ApproveDocument(id, approverID uint, comment string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取公文
		var document model.Document
		if err := tx.First(&document, id).Error; err != nil {
			return err
		}

		// 检查公文状态
		if document.Status != 2 {
			return errors.New("document is not in approval process")
		}

		// 获取当前审批节点
		var approval model.DocumentApproval
		err := tx.Where("document_id = ? AND approver_id = ? AND status = ?", id, approverID, 1).First(&approval).Error
		if err != nil {
			return errors.New("approval node not found")
		}

		now := time.Now()

		// 更新当前节点状态
		if err := tx.Model(&approval).Updates(map[string]interface{}{
			"status":        2,
			"comment":       comment,
			"approval_time": &now,
		}).Error; err != nil {
			return err
		}

		// 检查是否还有待审批的节点
		var count int64
		if err := tx.Model(&model.DocumentApproval{}).Where("document_id = ? AND status = ?", id, 1).Count(&count).Error; err != nil {
			return err
		}

		// 如果没有待审批的节点，更新公文状态为已签发
		if count == 0 {
			if err := tx.Model(&document).Updates(map[string]interface{}{
				"status":    3,
				"sign_date": &now,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// RejectDocument 审批驳回公文
func (s *DocumentService) RejectDocument(id, approverID uint, comment string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取公文
		var document model.Document
		if err := tx.First(&document, id).Error; err != nil {
			return err
		}

		// 检查公文状态
		if document.Status != 2 {
			return errors.New("document is not in approval process")
		}

		// 获取当前审批节点
		var approval model.DocumentApproval
		err := tx.Where("document_id = ? AND approver_id = ? AND status = ?", id, approverID, 1).First(&approval).Error
		if err != nil {
			return errors.New("approval node not found")
		}

		now := time.Now()

		// 更新当前节点状态
		if err := tx.Model(&approval).Updates(map[string]interface{}{
			"status":        3,
			"comment":       comment,
			"approval_time": &now,
		}).Error; err != nil {
			return err
		}

		// 更新公文状态为草稿
		if err := tx.Model(&document).Update("status", 1).Error; err != nil {
			return err
		}

		// 删除所有审批节点
		if err := tx.Where("document_id = ?", id).Delete(&model.DocumentApproval{}).Error; err != nil {
			return err
		}

		return nil
	})
}

// DistributeDocument 分发公文
func (s *DocumentService) DistributeDocument(distributions []model.DocumentDistribution) error {
	if len(distributions) == 0 {
		return errors.New("distributions is required")
	}

	// 检查公文是否存在且已签发
	var document model.Document
	if err := s.db.First(&document, distributions[0].DocumentID).Error; err != nil {
		return errors.New("document not found")
	}
	if document.Status != 3 {
		return errors.New("can only distribute signed documents")
	}

	return s.db.Create(&distributions).Error
}

// ReadDocument 阅读公文
func (s *DocumentService) ReadDocument(id, receiverID uint) error {
	now := time.Now()
	result := s.db.Model(&model.DocumentDistribution{}).
		Where("document_id = ? AND receiver_id = ? AND status = ?", id, receiverID, 1).
		Updates(map[string]interface{}{
			"status":    2,
			"read_time": &now,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("distribution record not found")
	}
	return nil
}

// ArchiveDocument 归档公文
func (s *DocumentService) ArchiveDocument(archive *model.DocumentArchive) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 检查公文是否存在且已签发
		var document model.Document
		if err := tx.First(&document, archive.DocumentID).Error; err != nil {
			return errors.New("document not found")
		}
		if document.Status != 3 {
			return errors.New("can only archive signed documents")
		}

		// 创建归档记录
		now := time.Now()
		archive.ArchiveDate = &now
		if err := tx.Create(archive).Error; err != nil {
			return err
		}

		// 更新公文状态为已归档
		if err := tx.Model(&document).Update("status", 4).Error; err != nil {
			return err
		}

		return nil
	})
}

// BorrowDocument 借阅公文
func (s *DocumentService) BorrowDocument(borrow *model.DocumentBorrow) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 检查公文是否存在且已归档
		var document model.Document
		if err := tx.First(&document, borrow.DocumentID).Error; err != nil {
			return errors.New("document not found")
		}
		if document.Status != 4 {
			return errors.New("can only borrow archived documents")
		}

		// 检查是否已被借阅
		var count int64
		if err := tx.Model(&model.DocumentBorrow{}).Where("document_id = ? AND status = ?", borrow.DocumentID, 1).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("document is already borrowed")
		}

		// 创建借阅记录
		now := time.Now()
		borrow.BorrowDate = &now
		if err := tx.Create(borrow).Error; err != nil {
			return err
		}

		// 更新归档记录状态为已借阅
		if err := tx.Model(&model.DocumentArchive{}).Where("document_id = ?", borrow.DocumentID).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// ReturnDocument 归还公文
func (s *DocumentService) ReturnDocument(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取借阅记录
		var borrow model.DocumentBorrow
		if err := tx.First(&borrow, id).Error; err != nil {
			return err
		}

		now := time.Now()

		// 更新借阅记录状态为已归还
		if err := tx.Model(&borrow).Updates(map[string]interface{}{
			"status":      2,
			"return_date": &now,
		}).Error; err != nil {
			return err
		}

		// 更新归档记录状态为已归档
		if err := tx.Model(&model.DocumentArchive{}).Where("document_id = ?", borrow.DocumentID).Update("status", 1).Error; err != nil {
			return err
		}

		return nil
	})
}

// DestroyDocument 销毁公文
func (s *DocumentService) DestroyDocument(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 检查公文是否存在且已归档
		var document model.Document
		if err := tx.First(&document, id).Error; err != nil {
			return errors.New("document not found")
		}
		if document.Status != 4 {
			return errors.New("can only destroy archived documents")
		}

		// 检查是否已被借阅
		var count int64
		if err := tx.Model(&model.DocumentBorrow{}).Where("document_id = ? AND status = ?", id, 1).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("document is currently borrowed")
		}

		now := time.Now()

		// 更新归档记录状态为已销毁
		if err := tx.Model(&model.DocumentArchive{}).Where("document_id = ?", id).Updates(map[string]interface{}{
			"status":       3,
			"destroy_date": &now,
		}).Error; err != nil {
			return err
		}

		// 更新公文状态为已作废
		if err := tx.Model(&document).Update("status", 5).Error; err != nil {
			return err
		}

		return nil
	})
}
