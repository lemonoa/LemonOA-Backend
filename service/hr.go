package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type HRService struct {
	db *gorm.DB
}

func NewHRService(db *gorm.DB) *HRService {
	return &HRService{db: db}
}

// GetPositionList 获取岗位职称列表
func (s *HRService) GetPositionList() ([]model.Position, error) {
	var positions []model.Position
	err := s.db.Order("sort asc").Find(&positions).Error
	return positions, err
}

// GetPositionByID 根据ID获取岗位职称
func (s *HRService) GetPositionByID(id uint) (*model.Position, error) {
	var position model.Position
	err := s.db.First(&position, id).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}

// CreatePosition 创建岗位职称
func (s *HRService) CreatePosition(position *model.Position) error {
	return s.db.Create(position).Error
}

// UpdatePosition 更新岗位职称
func (s *HRService) UpdatePosition(position *model.Position) error {
	if position.ID == 0 {
		return errors.New("position id is required")
	}
	return s.db.Model(position).Updates(position).Error
}

// DeletePosition 删除岗位职称
func (s *HRService) DeletePosition(id uint) error {
	// 检查是否有员工使用该岗位
	var count int64
	if err := s.db.Model(&model.Employee{}).Where("position_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete position with associated employees")
	}

	return s.db.Delete(&model.Position{}, id).Error
}

// GetEmployeeArchiveByEmployeeID 根据员工ID获取档案
func (s *HRService) GetEmployeeArchiveByEmployeeID(employeeID uint) (*model.EmployeeArchive, error) {
	var archive model.EmployeeArchive
	err := s.db.Where("employee_id = ?", employeeID).First(&archive).Error
	if err != nil {
		return nil, err
	}
	return &archive, nil
}

// CreateEmployeeArchive 创建员工档案
func (s *HRService) CreateEmployeeArchive(archive *model.EmployeeArchive) error {
	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, archive.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查是否已存在档案
	var count int64
	if err := s.db.Model(&model.EmployeeArchive{}).Where("employee_id = ?", archive.EmployeeID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("employee archive already exists")
	}

	return s.db.Create(archive).Error
}

// UpdateEmployeeArchive 更新员工档案
func (s *HRService) UpdateEmployeeArchive(archive *model.EmployeeArchive) error {
	if archive.ID == 0 {
		return errors.New("archive id is required")
	}

	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, archive.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	return s.db.Model(archive).Updates(archive).Error
}

// DeleteEmployeeArchive 删除员工档案
func (s *HRService) DeleteEmployeeArchive(id uint) error {
	return s.db.Delete(&model.EmployeeArchive{}, id).Error
}

// GetEmployeeArchiveList 获取员工档案列表
func (s *HRService) GetEmployeeArchiveList(departmentID uint, page, pageSize int) ([]model.EmployeeArchive, int64, error) {
	var archives []model.EmployeeArchive
	var total int64

	query := s.db.Model(&model.EmployeeArchive{}).
		Joins("LEFT JOIN employees ON employee_archives.employee_id = employees.id")

	if departmentID > 0 {
		query = query.Where("employees.department_id = ?", departmentID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&archives).Error
	if err != nil {
		return nil, 0, err
	}

	return archives, total, nil
}

// GetRewardPunishmentRecordList 获取奖惩记录列表
func (s *HRService) GetRewardPunishmentRecordList(employeeID uint, page, pageSize int) ([]model.RewardPunishmentRecord, int64, error) {
	var records []model.RewardPunishmentRecord
	var total int64

	query := s.db.Model(&model.RewardPunishmentRecord{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
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

// GetRewardPunishmentRecordByID 根据ID获取奖惩记录
func (s *HRService) GetRewardPunishmentRecordByID(id uint) (*model.RewardPunishmentRecord, error) {
	var record model.RewardPunishmentRecord
	err := s.db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// CreateRewardPunishmentRecord 创建奖惩记录
func (s *HRService) CreateRewardPunishmentRecord(record *model.RewardPunishmentRecord) error {
	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, record.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查奖惩项目是否存在
	var rewardPunishment model.RewardPunishment
	if err := s.db.First(&rewardPunishment, record.RewardPunishmentID).Error; err != nil {
		return errors.New("reward punishment not found")
	}

	return s.db.Create(record).Error
}

// UpdateRewardPunishmentRecord 更新奖惩记录
func (s *HRService) UpdateRewardPunishmentRecord(record *model.RewardPunishmentRecord) error {
	if record.ID == 0 {
		return errors.New("record id is required")
	}

	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, record.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查奖惩项目是否存在
	var rewardPunishment model.RewardPunishment
	if err := s.db.First(&rewardPunishment, record.RewardPunishmentID).Error; err != nil {
		return errors.New("reward punishment not found")
	}

	return s.db.Model(record).Updates(record).Error
}

// DeleteRewardPunishmentRecord 删除奖惩记录
func (s *HRService) DeleteRewardPunishmentRecord(id uint) error {
	return s.db.Delete(&model.RewardPunishmentRecord{}, id).Error
}

// GetCareRecordList 获取关怀记录列表
func (s *HRService) GetCareRecordList(employeeID uint, page, pageSize int) ([]model.CareRecord, int64, error) {
	var records []model.CareRecord
	var total int64

	query := s.db.Model(&model.CareRecord{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
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

// GetCareRecordByID 根据ID获取关怀记录
func (s *HRService) GetCareRecordByID(id uint) (*model.CareRecord, error) {
	var record model.CareRecord
	err := s.db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// CreateCareRecord 创建关怀记录
func (s *HRService) CreateCareRecord(record *model.CareRecord) error {
	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, record.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查关怀项目是否存在
	var careProject model.CareProject
	if err := s.db.First(&careProject, record.CareProjectID).Error; err != nil {
		return errors.New("care project not found")
	}

	return s.db.Create(record).Error
}

// UpdateCareRecord 更新关怀记录
func (s *HRService) UpdateCareRecord(record *model.CareRecord) error {
	if record.ID == 0 {
		return errors.New("record id is required")
	}

	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, record.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查关怀项目是否存在
	var careProject model.CareProject
	if err := s.db.First(&careProject, record.CareProjectID).Error; err != nil {
		return errors.New("care project not found")
	}

	return s.db.Model(record).Updates(record).Error
}

// DeleteCareRecord 删除关怀记录
func (s *HRService) DeleteCareRecord(id uint) error {
	return s.db.Delete(&model.CareRecord{}, id).Error
}

// GetTransferList 获取人事调动列表
func (s *HRService) GetTransferList(employeeID uint, status int, page, pageSize int) ([]model.Transfer, int64, error) {
	var transfers []model.Transfer
	var total int64

	query := s.db.Model(&model.Transfer{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&transfers).Error
	if err != nil {
		return nil, 0, err
	}

	return transfers, total, nil
}

// GetTransferByID 根据ID获取人事调动
func (s *HRService) GetTransferByID(id uint) (*model.Transfer, error) {
	var transfer model.Transfer
	err := s.db.First(&transfer, id).Error
	if err != nil {
		return nil, err
	}
	return &transfer, nil
}

// CreateTransfer 创建人事调动
func (s *HRService) CreateTransfer(transfer *model.Transfer) error {
	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, transfer.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查原部门是否存在
	var oldDepartment model.Department
	if err := s.db.First(&oldDepartment, transfer.OldDepartmentID).Error; err != nil {
		return errors.New("old department not found")
	}

	// 检查新部门是否存在
	var newDepartment model.Department
	if err := s.db.First(&newDepartment, transfer.NewDepartmentID).Error; err != nil {
		return errors.New("new department not found")
	}

	// 检查原岗位是否存在
	var oldPosition model.Position
	if err := s.db.First(&oldPosition, transfer.OldPositionID).Error; err != nil {
		return errors.New("old position not found")
	}

	// 检查新岗位是否存在
	var newPosition model.Position
	if err := s.db.First(&newPosition, transfer.NewPositionID).Error; err != nil {
		return errors.New("new position not found")
	}

	return s.db.Create(transfer).Error
}

// UpdateTransfer 更新人事调动
func (s *HRService) UpdateTransfer(transfer *model.Transfer) error {
	if transfer.ID == 0 {
		return errors.New("transfer id is required")
	}

	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, transfer.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查原部门是否存在
	var oldDepartment model.Department
	if err := s.db.First(&oldDepartment, transfer.OldDepartmentID).Error; err != nil {
		return errors.New("old department not found")
	}

	// 检查新部门是否存在
	var newDepartment model.Department
	if err := s.db.First(&newDepartment, transfer.NewDepartmentID).Error; err != nil {
		return errors.New("new department not found")
	}

	// 检查原岗位是否存在
	var oldPosition model.Position
	if err := s.db.First(&oldPosition, transfer.OldPositionID).Error; err != nil {
		return errors.New("old position not found")
	}

	// 检查新岗位是否存在
	var newPosition model.Position
	if err := s.db.First(&newPosition, transfer.NewPositionID).Error; err != nil {
		return errors.New("new position not found")
	}

	return s.db.Model(transfer).Updates(transfer).Error
}

// DeleteTransfer 删除人事调动
func (s *HRService) DeleteTransfer(id uint) error {
	return s.db.Delete(&model.Transfer{}, id).Error
}

// ApproveTransfer 审批通过人事调动
func (s *HRService) ApproveTransfer(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var transfer model.Transfer
		if err := tx.First(&transfer, id).Error; err != nil {
			return err
		}

		// 更新调动状态
		if err := tx.Model(&transfer).Update("status", 2).Error; err != nil {
			return err
		}

		// 更新员工信息
		if err := tx.Model(&model.Employee{}).Where("id = ?", transfer.EmployeeID).Updates(map[string]interface{}{
			"department_id": transfer.NewDepartmentID,
			"position_id":   transfer.NewPositionID,
		}).Error; err != nil {
			return err
		}

		return nil
	})
}

// RejectTransfer 审批驳回人事调动
func (s *HRService) RejectTransfer(id uint) error {
	return s.db.Model(&model.Transfer{}).Where("id = ?", id).Update("status", 3).Error
}

// GetResignationList 获取离职档案列表
func (s *HRService) GetResignationList(employeeID uint, status int, page, pageSize int) ([]model.Resignation, int64, error) {
	var resignations []model.Resignation
	var total int64

	query := s.db.Model(&model.Resignation{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&resignations).Error
	if err != nil {
		return nil, 0, err
	}

	return resignations, total, nil
}

// GetResignationByID 根据ID获取离职档案
func (s *HRService) GetResignationByID(id uint) (*model.Resignation, error) {
	var resignation model.Resignation
	err := s.db.First(&resignation, id).Error
	if err != nil {
		return nil, err
	}
	return &resignation, nil
}

// CreateResignation 创建离职档案
func (s *HRService) CreateResignation(resignation *model.Resignation) error {
	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, resignation.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查工作交接人是否存在
	if resignation.HandoverTo > 0 {
		var handoverEmployee model.Employee
		if err := s.db.First(&handoverEmployee, resignation.HandoverTo).Error; err != nil {
			return errors.New("handover employee not found")
		}
	}

	return s.db.Create(resignation).Error
}

// UpdateResignation 更新离职档案
func (s *HRService) UpdateResignation(resignation *model.Resignation) error {
	if resignation.ID == 0 {
		return errors.New("resignation id is required")
	}

	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, resignation.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查工作交接人是否存在
	if resignation.HandoverTo > 0 {
		var handoverEmployee model.Employee
		if err := s.db.First(&handoverEmployee, resignation.HandoverTo).Error; err != nil {
			return errors.New("handover employee not found")
		}
	}

	return s.db.Model(resignation).Updates(resignation).Error
}

// DeleteResignation 删除离职档案
func (s *HRService) DeleteResignation(id uint) error {
	return s.db.Delete(&model.Resignation{}, id).Error
}

// ApproveResignation 审批通过离职档案
func (s *HRService) ApproveResignation(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var resignation model.Resignation
		if err := tx.First(&resignation, id).Error; err != nil {
			return err
		}

		// 更新离职状态
		if err := tx.Model(&resignation).Update("status", 2).Error; err != nil {
			return err
		}

		// 更新员工状态为离职
		if err := tx.Model(&model.Employee{}).Where("id = ?", resignation.EmployeeID).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// RejectResignation 审批驳回离职档案
func (s *HRService) RejectResignation(id uint) error {
	return s.db.Model(&model.Resignation{}).Where("id = ?", id).Update("status", 3).Error
}

// GetContractList 获取员工合同列表
func (s *HRService) GetContractList(employeeID uint, status int, page, pageSize int) ([]model.Contract, int64, error) {
	var contracts []model.Contract
	var total int64

	query := s.db.Model(&model.Contract{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&contracts).Error
	if err != nil {
		return nil, 0, err
	}

	return contracts, total, nil
}

// GetContractByID 根据ID获取员工合同
func (s *HRService) GetContractByID(id uint) (*model.Contract, error) {
	var contract model.Contract
	err := s.db.First(&contract, id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

// CreateContract 创建员工合同
func (s *HRService) CreateContract(contract *model.Contract) error {
	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, contract.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查合同编号是否重复
	var count int64
	if err := s.db.Model(&model.Contract{}).Where("contract_no = ?", contract.ContractNo).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("contract no already exists")
	}

	return s.db.Create(contract).Error
}

// UpdateContract 更新员工合同
func (s *HRService) UpdateContract(contract *model.Contract) error {
	if contract.ID == 0 {
		return errors.New("contract id is required")
	}

	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, contract.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查合同编号是否重复
	var count int64
	if err := s.db.Model(&model.Contract{}).Where("contract_no = ? AND id != ?", contract.ContractNo, contract.ID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("contract no already exists")
	}

	return s.db.Model(contract).Updates(contract).Error
}

// DeleteContract 删除员工合同
func (s *HRService) DeleteContract(id uint) error {
	return s.db.Delete(&model.Contract{}, id).Error
}

// TerminateContract 终止员工合同
func (s *HRService) TerminateContract(id uint) error {
	return s.db.Model(&model.Contract{}).Where("id = ?", id).Update("status", 2).Error
}

// GetProbationList 获取转正列表
func (s *HRService) GetProbationList(employeeID uint, status int, page, pageSize int) ([]model.Probation, int64, error) {
	var probations []model.Probation
	var total int64

	query := s.db.Model(&model.Probation{})
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&probations).Error
	if err != nil {
		return nil, 0, err
	}

	return probations, total, nil
}

// GetProbationByID 根据ID获取转正
func (s *HRService) GetProbationByID(id uint) (*model.Probation, error) {
	var probation model.Probation
	err := s.db.First(&probation, id).Error
	if err != nil {
		return nil, err
	}
	return &probation, nil
}

// CreateProbation 创建转正
func (s *HRService) CreateProbation(probation *model.Probation) error {
	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, probation.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// 检查是否已存在转正记录
	var count int64
	if err := s.db.Model(&model.Probation{}).Where("employee_id = ? AND status != ?", probation.EmployeeID, 3).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("probation record already exists")
	}

	return s.db.Create(probation).Error
}

// UpdateProbation 更新转正
func (s *HRService) UpdateProbation(probation *model.Probation) error {
	if probation.ID == 0 {
		return errors.New("probation id is required")
	}

	// 检查员工是否存在
	var employee model.Employee
	if err := s.db.First(&employee, probation.EmployeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	return s.db.Model(probation).Updates(probation).Error
}

// DeleteProbation 删除转正
func (s *HRService) DeleteProbation(id uint) error {
	return s.db.Delete(&model.Probation{}, id).Error
}

// ApproveProbation 审批通过转正
func (s *HRService) ApproveProbation(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var probation model.Probation
		if err := tx.First(&probation, id).Error; err != nil {
			return err
		}

		// 更新转正状态
		if err := tx.Model(&probation).Update("status", 2).Error; err != nil {
			return err
		}

		// 更新员工试用期状态
		if err := tx.Model(&model.Employee{}).Where("id = ?", probation.EmployeeID).Update("probation_status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// RejectProbation 审批驳回转正
func (s *HRService) RejectProbation(id uint) error {
	return s.db.Model(&model.Probation{}).Where("id = ?", id).Update("status", 3).Error
}
