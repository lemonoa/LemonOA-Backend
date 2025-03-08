package model

import (
	"time"

	"gorm.io/gorm"
)

// 系统管理权限
const (
	// 用户管理
	PermissionUserList     = "system:user:list"
	PermissionUserCreate   = "system:user:create"
	PermissionUserUpdate   = "system:user:update"
	PermissionUserDelete   = "system:user:delete"
	PermissionUserResetPwd = "system:user:reset-password"

	// 角色管理
	PermissionRoleList        = "system:role:list"
	PermissionRoleCreate      = "system:role:create"
	PermissionRoleUpdate      = "system:role:update"
	PermissionRoleDelete      = "system:role:delete"
	PermissionRoleGetPerms    = "system:role:get-permissions"
	PermissionRoleUpdatePerms = "system:role:update-permissions"

	// 权限管理
	PermissionPermList   = "system:permission:list"
	PermissionPermCreate = "system:permission:create"
	PermissionPermUpdate = "system:permission:update"
	PermissionPermDelete = "system:permission:delete"

	// 考勤管理
	PermissionAttendanceRuleList     = "attendance:rule:list"
	PermissionAttendanceRuleCreate   = "attendance:rule:create"
	PermissionAttendanceRuleUpdate   = "attendance:rule:update"
	PermissionAttendanceRuleDelete   = "attendance:rule:delete"
	PermissionAttendanceRecordList   = "attendance:record:list"
	PermissionAttendanceRecordCreate = "attendance:record:create"
	PermissionLeaveList              = "attendance:leave:list"
	PermissionLeaveCreate            = "attendance:leave:create"
	PermissionLeaveApprove           = "attendance:leave:approve"

	// 会议室管理
	PermissionMeetingRoomList       = "meeting:room:list"
	PermissionMeetingRoomCreate     = "meeting:room:create"
	PermissionMeetingRoomUpdate     = "meeting:room:update"
	PermissionMeetingRoomDelete     = "meeting:room:delete"
	PermissionMeetingReserveList    = "meeting:reserve:list"
	PermissionMeetingReserveCreate  = "meeting:reserve:create"
	PermissionMeetingReserveApprove = "meeting:reserve:approve"

	// 文档管理
	PermissionDocumentList    = "document:list"
	PermissionDocumentCreate  = "document:create"
	PermissionDocumentUpdate  = "document:update"
	PermissionDocumentDelete  = "document:delete"
	PermissionDocumentApprove = "document:approve"
	PermissionDocumentArchive = "document:archive"
)

// User 用户表
type User struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Username    string         `gorm:"size:50;not null;unique" json:"username"` // 用户名
	Password    string         `gorm:"size:100;not null" json:"-"`              // 密码
	Salt        string         `gorm:"size:32;not null" json:"-"`               // 密码盐值
	RealName    string         `gorm:"size:50" json:"real_name"`                // 真实姓名
	Avatar      string         `gorm:"size:255" json:"avatar"`                  // 头像
	Email       string         `gorm:"size:100" json:"email"`                   // 邮箱
	Mobile      string         `gorm:"size:20" json:"mobile"`                   // 手机号
	Status      int            `gorm:"default:1" json:"status"`                 // 1:正常 2:禁用
	LastLoginAt *time.Time     `json:"last_login_at"`                           // 最后登录时间
	LastLoginIP string         `gorm:"size:50" json:"last_login_ip"`            // 最后登录IP
	CreatedBy   uint           `gorm:"not null" json:"created_by"`              // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// UserRole 用户角色关联表
type UserRole struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`    // 用户ID
	RoleID    uint           `gorm:"not null" json:"role_id"`    // 角色ID
	CreatedBy uint           `gorm:"not null" json:"created_by"` // 创建人ID
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Permission 权限表
type Permission struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`        // 权限名称
	Code        string         `gorm:"size:50;not null;unique" json:"code"` // 权限编码
	Type        int            `gorm:"not null" json:"type"`                // 1:菜单 2:按钮 3:接口
	ParentID    *uint          `json:"parent_id"`                           // 父级ID
	Path        string         `gorm:"size:200" json:"path"`                // 路由路径
	Component   string         `gorm:"size:200" json:"component"`           // 前端组件
	Icon        string         `gorm:"size:50" json:"icon"`                 // 图标
	Sort        int            `gorm:"default:0" json:"sort"`               // 排序
	Status      int            `gorm:"default:1" json:"status"`             // 1:启用 2:禁用
	Description string         `gorm:"size:200" json:"description"`         // 描述
	CreatedBy   uint           `gorm:"not null" json:"created_by"`          // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// RolePermission 角色权限关联表
type RolePermission struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	RoleID       uint           `gorm:"not null" json:"role_id"`       // 角色ID
	PermissionID uint           `gorm:"not null" json:"permission_id"` // 权限ID
	CreatedBy    uint           `gorm:"not null" json:"created_by"`    // 创建人ID
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// LoginLog 登录日志
type LoginLog struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`    // 用户ID
	IP        string         `gorm:"size:50" json:"ip"`          // 登录IP
	UserAgent string         `gorm:"size:500" json:"user_agent"` // User-Agent
	Status    int            `gorm:"default:1" json:"status"`    // 1:成功 2:失败
	Message   string         `gorm:"size:200" json:"message"`    // 失败原因
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Role 角色表
type Role struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`        // 角色名称
	Code        string         `gorm:"size:50;not null;unique" json:"code"` // 角色编码
	Description string         `gorm:"size:255" json:"description"`         // 描述
	Status      int            `gorm:"default:1" json:"status"`             // 1:启用 2:禁用
	CreatedBy   uint           `gorm:"not null" json:"created_by"`          // 创建人ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

func (UserRole) TableName() string {
	return "user_roles"
}

func (Permission) TableName() string {
	return "permissions"
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

func (LoginLog) TableName() string {
	return "login_logs"
}
