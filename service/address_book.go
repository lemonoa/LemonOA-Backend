package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type AddressBookService struct {
	db *gorm.DB
}

func NewAddressBookService(db *gorm.DB) *AddressBookService {
	return &AddressBookService{db: db}
}

// GetEmployeeList 获取员工列表
func (s *AddressBookService) GetEmployeeList(departmentID uint, page, pageSize int) ([]model.Employee, int64, error) {
	var employees []model.Employee
	var total int64

	query := s.db.Model(&model.Employee{})
	if departmentID > 0 {
		query = query.Where("department_id = ?", departmentID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&employees).Error
	if err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

// GetDepartmentList 获取部门列表
func (s *AddressBookService) GetDepartmentList() ([]model.Department, error) {
	var departments []model.Department
	err := s.db.Order("level asc, sort asc").Find(&departments).Error
	return departments, err
}

// CreateEmployee 创建员工
func (s *AddressBookService) CreateEmployee(employee *model.Employee) error {
	return s.db.Create(employee).Error
}

// UpdateEmployee 更新员工信息
func (s *AddressBookService) UpdateEmployee(employee *model.Employee) error {
	if employee.ID == 0 {
		return errors.New("employee id is required")
	}
	return s.db.Model(employee).Updates(employee).Error
}

// DeleteEmployee 删除员工
func (s *AddressBookService) DeleteEmployee(id uint) error {
	return s.db.Delete(&model.Employee{}, id).Error
}

// CreateDepartment 创建部门
func (s *AddressBookService) CreateDepartment(department *model.Department) error {
	if department.ParentID != nil {
		var parent model.Department
		if err := s.db.First(&parent, *department.ParentID).Error; err != nil {
			return errors.New("parent department not found")
		}
		department.Level = parent.Level + 1
	} else {
		department.Level = 1
	}
	return s.db.Create(department).Error
}

// UpdateDepartment 更新部门信息
func (s *AddressBookService) UpdateDepartment(department *model.Department) error {
	if department.ID == 0 {
		return errors.New("department id is required")
	}
	return s.db.Model(department).Updates(department).Error
}

// DeleteDepartment 删除部门
func (s *AddressBookService) DeleteDepartment(id uint) error {
	// 检查是否有子部门
	var count int64
	if err := s.db.Model(&model.Department{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete department with sub-departments")
	}

	// 检查是否有员工
	if err := s.db.Model(&model.Employee{}).Where("department_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete department with employees")
	}

	return s.db.Delete(&model.Department{}, id).Error
}
