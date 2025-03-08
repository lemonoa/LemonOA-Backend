package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

type VehicleService struct {
	db *gorm.DB
}

func NewVehicleService(db *gorm.DB) *VehicleService {
	return &VehicleService{db: db}
}

// GetVehicleList 获取车辆列表
func (s *VehicleService) GetVehicleList(status int, keyword string, page, pageSize int) ([]model.Vehicle, int64, error) {
	var vehicles []model.Vehicle
	var total int64

	query := s.db.Model(&model.Vehicle{})
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("plate_number LIKE ? OR brand LIKE ? OR model LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&vehicles).Error
	if err != nil {
		return nil, 0, err
	}

	return vehicles, total, nil
}

// GetVehicleByID 根据ID获取车辆
func (s *VehicleService) GetVehicleByID(id uint) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	err := s.db.First(&vehicle, id).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// CreateVehicle 创建车辆
func (s *VehicleService) CreateVehicle(vehicle *model.Vehicle) error {
	// 检查车牌号是否重复
	var count int64
	if err := s.db.Model(&model.Vehicle{}).Where("plate_number = ?", vehicle.PlateNumber).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("plate number already exists")
	}

	// 检查使用人是否存在
	if vehicle.UserID != nil {
		var user model.Employee
		if err := s.db.First(&user, vehicle.UserID).Error; err != nil {
			return errors.New("user not found")
		}
	}

	// 检查使用部门是否存在
	if vehicle.DepartmentID != nil {
		var department model.Department
		if err := s.db.First(&department, vehicle.DepartmentID).Error; err != nil {
			return errors.New("department not found")
		}
	}

	return s.db.Create(vehicle).Error
}

// UpdateVehicle 更新车辆
func (s *VehicleService) UpdateVehicle(vehicle *model.Vehicle) error {
	if vehicle.ID == 0 {
		return errors.New("vehicle id is required")
	}

	// 检查车牌号是否重复
	var count int64
	if err := s.db.Model(&model.Vehicle{}).Where("plate_number = ? AND id != ?", vehicle.PlateNumber, vehicle.ID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("plate number already exists")
	}

	// 检查使用人是否存在
	if vehicle.UserID != nil {
		var user model.Employee
		if err := s.db.First(&user, vehicle.UserID).Error; err != nil {
			return errors.New("user not found")
		}
	}

	// 检查使用部门是否存在
	if vehicle.DepartmentID != nil {
		var department model.Department
		if err := s.db.First(&department, vehicle.DepartmentID).Error; err != nil {
			return errors.New("department not found")
		}
	}

	return s.db.Model(vehicle).Updates(vehicle).Error
}

// DeleteVehicle 删除车辆
func (s *VehicleService) DeleteVehicle(id uint) error {
	// 检查是否有维修记录
	var count int64
	if err := s.db.Model(&model.VehicleRepair{}).Where("vehicle_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete vehicle with repair records")
	}

	// 检查是否有保养记录
	if err := s.db.Model(&model.VehicleMaintenance{}).Where("vehicle_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete vehicle with maintenance records")
	}

	// 检查是否有里程记录
	if err := s.db.Model(&model.VehicleMileage{}).Where("vehicle_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete vehicle with mileage records")
	}

	// 检查是否有费用记录
	if err := s.db.Model(&model.VehicleExpenseRecord{}).Where("vehicle_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete vehicle with expense records")
	}

	// 检查是否有违章记录
	if err := s.db.Model(&model.VehicleViolation{}).Where("vehicle_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete vehicle with violation records")
	}

	// 检查是否有事故记录
	if err := s.db.Model(&model.VehicleAccident{}).Where("vehicle_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete vehicle with accident records")
	}

	// 检查是否有用车申请
	if err := s.db.Model(&model.VehicleApplication{}).Where("vehicle_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete vehicle with applications")
	}

	return s.db.Delete(&model.Vehicle{}, id).Error
}

// GetVehicleRepairList 获取车辆维修记录列表
func (s *VehicleService) GetVehicleRepairList(vehicleID uint, status int, page, pageSize int) ([]model.VehicleRepair, int64, error) {
	var repairs []model.VehicleRepair
	var total int64

	query := s.db.Model(&model.VehicleRepair{})
	if vehicleID > 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&repairs).Error
	if err != nil {
		return nil, 0, err
	}

	return repairs, total, nil
}

// GetVehicleRepairByID 根据ID获取车辆维修记录
func (s *VehicleService) GetVehicleRepairByID(id uint) (*model.VehicleRepair, error) {
	var repair model.VehicleRepair
	err := s.db.First(&repair, id).Error
	if err != nil {
		return nil, err
	}
	return &repair, nil
}

// CreateVehicleRepair 创建车辆维修记录
func (s *VehicleService) CreateVehicleRepair(repair *model.VehicleRepair) error {
	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, repair.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建维修记录
		if err := tx.Create(repair).Error; err != nil {
			return err
		}

		// 更新车辆状态为维修中
		if err := tx.Model(&vehicle).Update("status", 3).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateVehicleRepair 更新车辆维修记录
func (s *VehicleService) UpdateVehicleRepair(repair *model.VehicleRepair) error {
	if repair.ID == 0 {
		return errors.New("repair id is required")
	}

	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, repair.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Model(repair).Updates(repair).Error
}

// DeleteVehicleRepair 删除车辆维修记录
func (s *VehicleService) DeleteVehicleRepair(id uint) error {
	return s.db.Delete(&model.VehicleRepair{}, id).Error
}

// CompleteVehicleRepair 完成车辆维修
func (s *VehicleService) CompleteVehicleRepair(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取维修记录
		var repair model.VehicleRepair
		if err := tx.First(&repair, id).Error; err != nil {
			return err
		}

		// 更新维修记录状态为已完成
		if err := tx.Model(&repair).Update("status", 3).Error; err != nil {
			return err
		}

		// 更新车辆状态为在用
		if err := tx.Model(&model.Vehicle{}).Where("id = ?", repair.VehicleID).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetVehicleMaintenanceList 获取车辆保养记录列表
func (s *VehicleService) GetVehicleMaintenanceList(vehicleID uint, status int, page, pageSize int) ([]model.VehicleMaintenance, int64, error) {
	var maintenances []model.VehicleMaintenance
	var total int64

	query := s.db.Model(&model.VehicleMaintenance{})
	if vehicleID > 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
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

// GetVehicleMaintenanceByID 根据ID获取车辆保养记录
func (s *VehicleService) GetVehicleMaintenanceByID(id uint) (*model.VehicleMaintenance, error) {
	var maintenance model.VehicleMaintenance
	err := s.db.First(&maintenance, id).Error
	if err != nil {
		return nil, err
	}
	return &maintenance, nil
}

// CreateVehicleMaintenance 创建车辆保养记录
func (s *VehicleService) CreateVehicleMaintenance(maintenance *model.VehicleMaintenance) error {
	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, maintenance.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建保养记录
		if err := tx.Create(maintenance).Error; err != nil {
			return err
		}

		// 更新车辆状态为保养中
		if err := tx.Model(&vehicle).Update("status", 3).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateVehicleMaintenance 更新车辆保养记录
func (s *VehicleService) UpdateVehicleMaintenance(maintenance *model.VehicleMaintenance) error {
	if maintenance.ID == 0 {
		return errors.New("maintenance id is required")
	}

	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, maintenance.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Model(maintenance).Updates(maintenance).Error
}

// DeleteVehicleMaintenance 删除车辆保养记录
func (s *VehicleService) DeleteVehicleMaintenance(id uint) error {
	return s.db.Delete(&model.VehicleMaintenance{}, id).Error
}

// CompleteVehicleMaintenance 完成车辆保养
func (s *VehicleService) CompleteVehicleMaintenance(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取保养记录
		var maintenance model.VehicleMaintenance
		if err := tx.First(&maintenance, id).Error; err != nil {
			return err
		}

		// 更新保养记录状态为已完成
		if err := tx.Model(&maintenance).Update("status", 3).Error; err != nil {
			return err
		}

		// 更新车辆状态为在用
		if err := tx.Model(&model.Vehicle{}).Where("id = ?", maintenance.VehicleID).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetVehicleMileageList 获取车辆里程记录列表
func (s *VehicleService) GetVehicleMileageList(vehicleID uint, page, pageSize int) ([]model.VehicleMileage, int64, error) {
	var mileages []model.VehicleMileage
	var total int64

	query := s.db.Model(&model.VehicleMileage{})
	if vehicleID > 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("date desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&mileages).Error
	if err != nil {
		return nil, 0, err
	}

	return mileages, total, nil
}

// GetVehicleMileageByID 根据ID获取车辆里程记录
func (s *VehicleService) GetVehicleMileageByID(id uint) (*model.VehicleMileage, error) {
	var mileage model.VehicleMileage
	err := s.db.First(&mileage, id).Error
	if err != nil {
		return nil, err
	}
	return &mileage, nil
}

// CreateVehicleMileage 创建车辆里程记录
func (s *VehicleService) CreateVehicleMileage(mileage *model.VehicleMileage) error {
	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, mileage.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	// 计算行驶里程
	mileage.Distance = mileage.EndMileage - mileage.StartMileage

	return s.db.Create(mileage).Error
}

// UpdateVehicleMileage 更新车辆里程记录
func (s *VehicleService) UpdateVehicleMileage(mileage *model.VehicleMileage) error {
	if mileage.ID == 0 {
		return errors.New("mileage id is required")
	}

	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, mileage.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	// 计算行驶里程
	mileage.Distance = mileage.EndMileage - mileage.StartMileage

	return s.db.Model(mileage).Updates(mileage).Error
}

// DeleteVehicleMileage 删除车辆里程记录
func (s *VehicleService) DeleteVehicleMileage(id uint) error {
	return s.db.Delete(&model.VehicleMileage{}, id).Error
}

// GetVehicleExpenseList 获取车辆费用记录列表
func (s *VehicleService) GetVehicleExpenseList(vehicleID uint, expenseID uint, page, pageSize int) ([]model.VehicleExpenseRecord, int64, error) {
	var expenses []model.VehicleExpenseRecord
	var total int64

	query := s.db.Model(&model.VehicleExpenseRecord{})
	if vehicleID > 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
	}
	if expenseID > 0 {
		query = query.Where("expense_id = ?", expenseID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("date desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&expenses).Error
	if err != nil {
		return nil, 0, err
	}

	return expenses, total, nil
}

// GetVehicleExpenseByID 根据ID获取车辆费用记录
func (s *VehicleService) GetVehicleExpenseByID(id uint) (*model.VehicleExpenseRecord, error) {
	var expense model.VehicleExpenseRecord
	err := s.db.First(&expense, id).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

// CreateVehicleExpense 创建车辆费用记录
func (s *VehicleService) CreateVehicleExpense(expense *model.VehicleExpenseRecord) error {
	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, expense.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	// 检查费用类型是否存在
	var expenseType model.VehicleExpense
	if err := s.db.First(&expenseType, expense.ExpenseID).Error; err != nil {
		return errors.New("expense type not found")
	}

	return s.db.Create(expense).Error
}

// UpdateVehicleExpense 更新车辆费用记录
func (s *VehicleService) UpdateVehicleExpense(expense *model.VehicleExpenseRecord) error {
	if expense.ID == 0 {
		return errors.New("expense id is required")
	}

	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, expense.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	// 检查费用类型是否存在
	var expenseType model.VehicleExpense
	if err := s.db.First(&expenseType, expense.ExpenseID).Error; err != nil {
		return errors.New("expense type not found")
	}

	return s.db.Model(expense).Updates(expense).Error
}

// DeleteVehicleExpense 删除车辆费用记录
func (s *VehicleService) DeleteVehicleExpense(id uint) error {
	return s.db.Delete(&model.VehicleExpenseRecord{}, id).Error
}

// GetVehicleViolationList 获取车辆违章记录列表
func (s *VehicleService) GetVehicleViolationList(vehicleID uint, status int, page, pageSize int) ([]model.VehicleViolation, int64, error) {
	var violations []model.VehicleViolation
	var total int64

	query := s.db.Model(&model.VehicleViolation{})
	if vehicleID > 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("date desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&violations).Error
	if err != nil {
		return nil, 0, err
	}

	return violations, total, nil
}

// GetVehicleViolationByID 根据ID获取车辆违章记录
func (s *VehicleService) GetVehicleViolationByID(id uint) (*model.VehicleViolation, error) {
	var violation model.VehicleViolation
	err := s.db.First(&violation, id).Error
	if err != nil {
		return nil, err
	}
	return &violation, nil
}

// CreateVehicleViolation 创建车辆违章记录
func (s *VehicleService) CreateVehicleViolation(violation *model.VehicleViolation) error {
	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, violation.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Create(violation).Error
}

// UpdateVehicleViolation 更新车辆违章记录
func (s *VehicleService) UpdateVehicleViolation(violation *model.VehicleViolation) error {
	if violation.ID == 0 {
		return errors.New("violation id is required")
	}

	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, violation.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Model(violation).Updates(violation).Error
}

// DeleteVehicleViolation 删除车辆违章记录
func (s *VehicleService) DeleteVehicleViolation(id uint) error {
	return s.db.Delete(&model.VehicleViolation{}, id).Error
}

// HandleVehicleViolation 处理车辆违章
func (s *VehicleService) HandleVehicleViolation(id uint, status int) error {
	if status != 2 && status != 3 {
		return errors.New("invalid status")
	}

	return s.db.Model(&model.VehicleViolation{}).Where("id = ?", id).Update("status", status).Error
}

// GetVehicleAccidentList 获取车辆事故记录列表
func (s *VehicleService) GetVehicleAccidentList(vehicleID uint, status int, page, pageSize int) ([]model.VehicleAccident, int64, error) {
	var accidents []model.VehicleAccident
	var total int64

	query := s.db.Model(&model.VehicleAccident{})
	if vehicleID > 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("date desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&accidents).Error
	if err != nil {
		return nil, 0, err
	}

	return accidents, total, nil
}

// GetVehicleAccidentByID 根据ID获取车辆事故记录
func (s *VehicleService) GetVehicleAccidentByID(id uint) (*model.VehicleAccident, error) {
	var accident model.VehicleAccident
	err := s.db.First(&accident, id).Error
	if err != nil {
		return nil, err
	}
	return &accident, nil
}

// CreateVehicleAccident 创建车辆事故记录
func (s *VehicleService) CreateVehicleAccident(accident *model.VehicleAccident) error {
	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, accident.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Create(accident).Error
}

// UpdateVehicleAccident 更新车辆事故记录
func (s *VehicleService) UpdateVehicleAccident(accident *model.VehicleAccident) error {
	if accident.ID == 0 {
		return errors.New("accident id is required")
	}

	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, accident.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	return s.db.Model(accident).Updates(accident).Error
}

// DeleteVehicleAccident 删除车辆事故记录
func (s *VehicleService) DeleteVehicleAccident(id uint) error {
	return s.db.Delete(&model.VehicleAccident{}, id).Error
}

// HandleVehicleAccident 处理车辆事故
func (s *VehicleService) HandleVehicleAccident(id uint, status int) error {
	if status != 2 && status != 3 {
		return errors.New("invalid status")
	}

	return s.db.Model(&model.VehicleAccident{}).Where("id = ?", id).Update("status", status).Error
}

// GetVehicleApplicationList 获取用车申请列表
func (s *VehicleService) GetVehicleApplicationList(vehicleID, userID uint, status int, page, pageSize int) ([]model.VehicleApplication, int64, error) {
	var applications []model.VehicleApplication
	var total int64

	query := s.db.Model(&model.VehicleApplication{})
	if vehicleID > 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
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

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applications).Error
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

// GetVehicleApplicationByID 根据ID获取用车申请
func (s *VehicleService) GetVehicleApplicationByID(id uint) (*model.VehicleApplication, error) {
	var application model.VehicleApplication
	err := s.db.First(&application, id).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

// CreateVehicleApplication 创建用车申请
func (s *VehicleService) CreateVehicleApplication(application *model.VehicleApplication) error {
	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, application.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	// 检查申请人是否存在
	var user model.Employee
	if err := s.db.First(&user, application.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	// 检查申请部门是否存在
	var department model.Department
	if err := s.db.First(&department, application.DepartmentID).Error; err != nil {
		return errors.New("department not found")
	}

	// 检查车辆是否可用
	if vehicle.Status != 2 {
		return errors.New("vehicle is not available")
	}

	// 检查时间段内是否有其他申请
	var count int64
	err := s.db.Model(&model.VehicleApplication{}).
		Where("vehicle_id = ? AND status IN (1,2) AND ((start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?))",
			application.VehicleID,
			application.StartDate, application.EndDate,
			application.StartDate, application.EndDate).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("vehicle is already booked during this period")
	}

	return s.db.Create(application).Error
}

// UpdateVehicleApplication 更新用车申请
func (s *VehicleService) UpdateVehicleApplication(application *model.VehicleApplication) error {
	if application.ID == 0 {
		return errors.New("application id is required")
	}

	// 检查车辆是否存在
	var vehicle model.Vehicle
	if err := s.db.First(&vehicle, application.VehicleID).Error; err != nil {
		return errors.New("vehicle not found")
	}

	// 检查申请人是否存在
	var user model.Employee
	if err := s.db.First(&user, application.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	// 检查申请部门是否存在
	var department model.Department
	if err := s.db.First(&department, application.DepartmentID).Error; err != nil {
		return errors.New("department not found")
	}

	// 检查时间段内是否有其他申请
	var count int64
	err := s.db.Model(&model.VehicleApplication{}).
		Where("id != ? AND vehicle_id = ? AND status IN (1,2) AND ((start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?))",
			application.ID, application.VehicleID,
			application.StartDate, application.EndDate,
			application.StartDate, application.EndDate).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("vehicle is already booked during this period")
	}

	return s.db.Model(application).Updates(application).Error
}

// DeleteVehicleApplication 删除用车申请
func (s *VehicleService) DeleteVehicleApplication(id uint) error {
	// 只能删除待审批的申请
	result := s.db.Where("id = ? AND status = ?", id, 1).Delete(&model.VehicleApplication{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only delete pending applications")
	}
	return nil
}

// ApproveVehicleApplication 审批通过用车申请
func (s *VehicleService) ApproveVehicleApplication(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取申请记录
		var application model.VehicleApplication
		if err := tx.First(&application, id).Error; err != nil {
			return err
		}

		// 检查状态是否为待审批
		if application.Status != 1 {
			return errors.New("can only approve pending applications")
		}

		// 更新申请状态为已通过
		if err := tx.Model(&application).Update("status", 2).Error; err != nil {
			return err
		}

		// 更新车辆状态为使用中
		if err := tx.Model(&model.Vehicle{}).Where("id = ?", application.VehicleID).Update("status", 2).Error; err != nil {
			return err
		}

		return nil
	})
}

// RejectVehicleApplication 审批驳回用车申请
func (s *VehicleService) RejectVehicleApplication(id uint) error {
	// 只能驳回待审批的申请
	result := s.db.Model(&model.VehicleApplication{}).Where("id = ? AND status = ?", id, 1).Update("status", 3)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("can only reject pending applications")
	}
	return nil
}

// GetVehicleReturnList 获取车辆归还记录列表
func (s *VehicleService) GetVehicleReturnList(applicationID uint, status int, page, pageSize int) ([]model.VehicleReturn, int64, error) {
	var returns []model.VehicleReturn
	var total int64

	query := s.db.Model(&model.VehicleReturn{})
	if applicationID > 0 {
		query = query.Where("application_id = ?", applicationID)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&returns).Error
	if err != nil {
		return nil, 0, err
	}

	return returns, total, nil
}

// GetVehicleReturnByID 根据ID获取车辆归还记录
func (s *VehicleService) GetVehicleReturnByID(id uint) (*model.VehicleReturn, error) {
	var vehicleReturn model.VehicleReturn
	err := s.db.First(&vehicleReturn, id).Error
	if err != nil {
		return nil, err
	}
	return &vehicleReturn, nil
}

// CreateVehicleReturn 创建车辆归还记录
func (s *VehicleService) CreateVehicleReturn(vehicleReturn *model.VehicleReturn) error {
	// 检查用车申请是否存在
	var application model.VehicleApplication
	if err := s.db.First(&application, vehicleReturn.ApplicationID).Error; err != nil {
		return errors.New("application not found")
	}

	// 检查申请状态是否为已通过
	if application.Status != 2 {
		return errors.New("can only return vehicles from approved applications")
	}

	// 计算行驶里程
	vehicleReturn.Distance = vehicleReturn.EndMileage - vehicleReturn.StartMileage

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建归还记录
		if err := tx.Create(vehicleReturn).Error; err != nil {
			return err
		}

		// 更新车辆状态为闲置
		if err := tx.Model(&model.Vehicle{}).Where("id = ?", application.VehicleID).Update("status", 1).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateVehicleReturn 更新车辆归还记录
func (s *VehicleService) UpdateVehicleReturn(vehicleReturn *model.VehicleReturn) error {
	if vehicleReturn.ID == 0 {
		return errors.New("return id is required")
	}

	// 检查用车申请是否存在
	var application model.VehicleApplication
	if err := s.db.First(&application, vehicleReturn.ApplicationID).Error; err != nil {
		return errors.New("application not found")
	}

	// 计算行驶里程
	vehicleReturn.Distance = vehicleReturn.EndMileage - vehicleReturn.StartMileage

	return s.db.Model(vehicleReturn).Updates(vehicleReturn).Error
}

// DeleteVehicleReturn 删除车辆归还记录
func (s *VehicleService) DeleteVehicleReturn(id uint) error {
	return s.db.Delete(&model.VehicleReturn{}, id).Error
}
