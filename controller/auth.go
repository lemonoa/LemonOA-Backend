package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/lemonoa/LemonOA-Go/middleware"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// RegisterRoutes 注册路由
func (c *AuthController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/auth")
	{
		// 不需要认证的接口
		api.POST("/login", c.Login)

		// 需要认证的接口
		auth := api.Use(middleware.JWT())
		{
			auth.GET("/user-info", c.GetUserInfo)
			auth.GET("/permissions", c.GetUserPermissions)
			auth.POST("/change-password", c.ChangePassword)
		}
	}

	// 用户管理接口，需要认证和权限
	users := r.Group("/api/users").Use(middleware.JWT())
	{
		users.GET("", middleware.RequirePermission("system:user:list"), c.GetUserList)
		users.POST("", middleware.RequirePermission("system:user:create"), c.CreateUser)
		users.PUT("/:id", middleware.RequirePermission("system:user:update"), c.UpdateUser)
		users.DELETE("/:id", middleware.RequirePermission("system:user:delete"), c.DeleteUser)
		users.PUT("/:id/reset-password", middleware.RequirePermission("system:user:reset-password"), c.ResetPassword)
	}

	// 角色管理接口，需要认证和权限
	roles := r.Group("/api/roles").Use(middleware.JWT())
	{
		roles.GET("", middleware.RequirePermission("system:role:list"), c.GetRoleList)
		roles.POST("", middleware.RequirePermission("system:role:create"), c.CreateRole)
		roles.PUT("/:id", middleware.RequirePermission("system:role:update"), c.UpdateRole)
		roles.DELETE("/:id", middleware.RequirePermission("system:role:delete"), c.DeleteRole)
		roles.GET("/:id/permissions", middleware.RequirePermission("system:role:get-permissions"), c.GetRolePermissions)
		roles.PUT("/:id/permissions", middleware.RequirePermission("system:role:update-permissions"), c.UpdateRolePermissions)
	}

	// 权限管理接口，需要认证和权限
	permissions := r.Group("/api/permissions").Use(middleware.JWT())
	{
		permissions.GET("", middleware.RequirePermission("system:permission:list"), c.GetPermissionList)
		permissions.POST("", middleware.RequirePermission("system:permission:create"), c.CreatePermission)
		permissions.PUT("/:id", middleware.RequirePermission("system:permission:update"), c.UpdatePermission)
		permissions.DELETE("/:id", middleware.RequirePermission("system:permission:delete"), c.DeletePermission)
	}
}

// Login 用户登录
func (c *AuthController) Login(ctx *gin.Context) {
	var params struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.authService.Login(params.Username, params.Password, ctx.ClientIP(), ctx.Request.UserAgent())
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// GetUserInfo 获取用户信息
func (c *AuthController) GetUserInfo(ctx *gin.Context) {
	// 从JWT中获取userID
	userID := uint(1) // TODO: 从JWT中获取

	user, err := c.authService.GetUserInfo(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// GetUserPermissions 获取用户权限列表
func (c *AuthController) GetUserPermissions(ctx *gin.Context) {
	// 从JWT中获取userID
	userID := uint(1) // TODO: 从JWT中获取

	permissions, err := c.authService.GetUserPermissions(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, permissions)
}

// ChangePassword 修改密码
func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var params struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从JWT中获取userID
	userID := uint(1) // TODO: 从JWT中获取

	if err := c.authService.ChangePassword(userID, params.OldPassword, params.NewPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetUserList 获取用户列表
func (c *AuthController) GetUserList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	keyword := ctx.Query("keyword")
	status, _ := strconv.Atoi(ctx.Query("status"))

	users, total, err := c.authService.GetUserList(status, keyword, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  users,
		"total": total,
	})
}

// CreateUser 创建用户
func (c *AuthController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	user.CreatedBy = uint(1)

	if err := c.authService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// UpdateUser 更新用户
func (c *AuthController) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uint(id)
	if err := c.authService.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser 删除用户
func (c *AuthController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.authService.DeleteUser(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ResetPassword 重置用户密码
func (c *AuthController) ResetPassword(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.authService.ResetPassword(uint(id), req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetRoleList 获取角色列表
func (c *AuthController) GetRoleList(ctx *gin.Context) {
	roles, err := c.authService.GetRoleList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

// CreateRole 创建角色
func (c *AuthController) CreateRole(ctx *gin.Context) {
	var role model.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	role.CreatedBy = uint(1)

	if err := c.authService.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

// UpdateRole 更新角色
func (c *AuthController) UpdateRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var role model.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.ID = uint(id)
	if err := c.authService.UpdateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

// DeleteRole 删除角色
func (c *AuthController) DeleteRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.authService.DeleteRole(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetRolePermissions 获取角色权限
func (c *AuthController) GetRolePermissions(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	permissions, err := c.authService.GetRolePermissions(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, permissions)
}

// UpdateRolePermissions 更新角色权限
func (c *AuthController) UpdateRolePermissions(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req struct {
		PermissionIDs []uint `json:"permission_ids" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.authService.UpdateRolePermissions(uint(id), req.PermissionIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetPermissionList 获取权限列表
func (c *AuthController) GetPermissionList(ctx *gin.Context) {
	permissions, err := c.authService.GetPermissionList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, permissions)
}

// CreatePermission 创建权限
func (c *AuthController) CreatePermission(ctx *gin.Context) {
	var permission model.Permission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取当前用户ID
	permission.CreatedBy = uint(1)

	if err := c.authService.CreatePermission(&permission); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, permission)
}

// UpdatePermission 更新权限
func (c *AuthController) UpdatePermission(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var permission model.Permission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission.ID = uint(id)
	if err := c.authService.UpdatePermission(&permission); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, permission)
}

// DeletePermission 删除权限
func (c *AuthController) DeletePermission(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := c.authService.DeletePermission(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
