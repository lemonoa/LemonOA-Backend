package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type AddressBookController struct {
	addressBookService *service.AddressBookService
}

func NewAddressBookController(addressBookService *service.AddressBookService) *AddressBookController {
	return &AddressBookController{
		addressBookService: addressBookService,
	}
}

// RegisterRoutes 注册路由
func (c *AddressBookController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/address-book")
	{
		api.GET("/employees", c.GetEmployeeList)
		api.POST("/employees", c.CreateEmployee)
		api.PUT("/employees/:id", c.UpdateEmployee)
		api.DELETE("/employees/:id", c.DeleteEmployee)

		api.GET("/departments", c.GetDepartmentList)
		api.POST("/departments", c.CreateDepartment)
		api.PUT("/departments/:id", c.UpdateDepartment)
		api.DELETE("/departments/:id", c.DeleteDepartment)
	}
}

// GetEmployeeList 获取员工列表
func (c *AddressBookController) GetEmployeeList(ctx *gin.Context) {
	departmentID, _ := strconv.ParseUint(ctx.Query("department_id"), 10, 32)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	employees, total, err := c.addressBookService.GetEmployeeList(uint(departmentID), page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  employees,
		"total": total,
	})
}

// CreateEmployee 创建员工
func (c *AddressBookController) CreateEmployee(ctx *gin.Context) {
	var employee model.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.addressBookService.CreateEmployee(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, employee)
}

// UpdateEmployee 更新员工信息
func (c *AddressBookController) UpdateEmployee(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var employee model.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.ID = uint(id)
	if err := c.addressBookService.UpdateEmployee(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// DeleteEmployee 删除员工
func (c *AddressBookController) DeleteEmployee(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.addressBookService.DeleteEmployee(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetDepartmentList 获取部门列表
func (c *AddressBookController) GetDepartmentList(ctx *gin.Context) {
	departments, err := c.addressBookService.GetDepartmentList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, departments)
}

// CreateDepartment 创建部门
func (c *AddressBookController) CreateDepartment(ctx *gin.Context) {
	var department model.Department
	if err := ctx.ShouldBindJSON(&department); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.addressBookService.CreateDepartment(&department); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, department)
}

// UpdateDepartment 更新部门信息
func (c *AddressBookController) UpdateDepartment(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var department model.Department
	if err := ctx.ShouldBindJSON(&department); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department.ID = uint(id)
	if err := c.addressBookService.UpdateDepartment(&department); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, department)
}

// DeleteDepartment 删除部门
func (c *AddressBookController) DeleteDepartment(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.addressBookService.DeleteDepartment(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
