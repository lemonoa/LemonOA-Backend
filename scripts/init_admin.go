package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/lemonoa/LemonOA-Go/database"

	"github.com/lemonoa/LemonOA-Go/model"

	"github.com/spf13/viper"
)

func init() {
	// 加载配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 初始化数据库连接
	if err := database.InitMySQL(); err != nil {
		log.Fatalf("初始化数据库连接失败: %v", err)
	}
}

func main() {
	db := database.DB

	// 1. 创建超级管理员角色
	superAdminRole := &model.Role{
		Name:        "超级管理员",
		Code:        "super_admin",
		Description: "系统超级管理员,拥有所有权限",
		Status:      1,
		CreatedBy:   1,
	}
	if err := db.Create(superAdminRole).Error; err != nil {
		log.Fatalf("创建超级管理员角色失败: %v", err)
	}

	// 2. 创建系统管理员角色
	adminRole := &model.Role{
		Name:        "系统管理员",
		Code:        "admin",
		Description: "系统管理员,拥有大部分系统管理权限",
		Status:      1,
		CreatedBy:   1,
	}
	if err := db.Create(adminRole).Error; err != nil {
		log.Fatalf("创建系统管理员角色失败: %v", err)
	}

	// 3. 创建权限数据
	permissions := []model.Permission{
		// 系统管理权限
		{Name: "用户列表", Code: model.PermissionUserList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建用户", Code: model.PermissionUserCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "更新用户", Code: model.PermissionUserUpdate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "删除用户", Code: model.PermissionUserDelete, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "重置密码", Code: model.PermissionUserResetPwd, Type: 3, Status: 1, CreatedBy: 1},

		{Name: "角色列表", Code: model.PermissionRoleList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建角色", Code: model.PermissionRoleCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "更新角色", Code: model.PermissionRoleUpdate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "删除角色", Code: model.PermissionRoleDelete, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "获取角色权限", Code: model.PermissionRoleGetPerms, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "更新角色权限", Code: model.PermissionRoleUpdatePerms, Type: 3, Status: 1, CreatedBy: 1},

		{Name: "权限列表", Code: model.PermissionPermList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建权限", Code: model.PermissionPermCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "更新权限", Code: model.PermissionPermUpdate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "删除权限", Code: model.PermissionPermDelete, Type: 3, Status: 1, CreatedBy: 1},

		// 考勤管理权限
		{Name: "考勤规则列表", Code: model.PermissionAttendanceRuleList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建考勤规则", Code: model.PermissionAttendanceRuleCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "更新考勤规则", Code: model.PermissionAttendanceRuleUpdate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "删除考勤规则", Code: model.PermissionAttendanceRuleDelete, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "考勤记录列表", Code: model.PermissionAttendanceRecordList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建考勤记录", Code: model.PermissionAttendanceRecordCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "请假列表", Code: model.PermissionLeaveList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建请假", Code: model.PermissionLeaveCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "审批请假", Code: model.PermissionLeaveApprove, Type: 3, Status: 1, CreatedBy: 1},

		// 会议室管理权限
		{Name: "会议室列表", Code: model.PermissionMeetingRoomList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建会议室", Code: model.PermissionMeetingRoomCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "更新会议室", Code: model.PermissionMeetingRoomUpdate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "删除会议室", Code: model.PermissionMeetingRoomDelete, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "会议预约列表", Code: model.PermissionMeetingReserveList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建会议预约", Code: model.PermissionMeetingReserveCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "审批会议预约", Code: model.PermissionMeetingReserveApprove, Type: 3, Status: 1, CreatedBy: 1},

		// 文档管理权限
		{Name: "文档列表", Code: model.PermissionDocumentList, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "创建文档", Code: model.PermissionDocumentCreate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "更新文档", Code: model.PermissionDocumentUpdate, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "删除文档", Code: model.PermissionDocumentDelete, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "审批文档", Code: model.PermissionDocumentApprove, Type: 3, Status: 1, CreatedBy: 1},
		{Name: "归档文档", Code: model.PermissionDocumentArchive, Type: 3, Status: 1, CreatedBy: 1},
	}

	if err := db.Create(&permissions).Error; err != nil {
		log.Fatalf("创建权限数据失败: %v", err)
	}

	// 4. 为系统管理员角色分配权限
	var adminPermissions []model.RolePermission
	for _, p := range permissions {
		adminPermissions = append(adminPermissions, model.RolePermission{
			RoleID:       adminRole.ID,
			PermissionID: p.ID,
			CreatedBy:    1,
		})
	}
	if err := db.Create(&adminPermissions).Error; err != nil {
		log.Fatalf("分配系统管理员权限失败: %v", err)
	}

	// 5. 创建超级管理员用户
	salt := generateSalt()
	adminUser := &model.User{
		Username:  "admin",
		Password:  encryptPassword("admin123", salt),
		Salt:      salt,
		RealName:  "超级管理员",
		Status:    1,
		CreatedBy: 1,
	}
	if err := db.Create(adminUser).Error; err != nil {
		log.Fatalf("创建超级管理员用户失败: %v", err)
	}

	// 6. 为超级管理员用户分配角色
	userRole := &model.UserRole{
		UserID:    adminUser.ID,
		RoleID:    superAdminRole.ID,
		CreatedBy: 1,
	}
	if err := db.Create(userRole).Error; err != nil {
		log.Fatalf("分配超级管理员角色失败: %v", err)
	}

	fmt.Println("初始化完成!")
	fmt.Println("超级管理员账号: admin")
	fmt.Println("初始密码: admin123")
}

// 生成32位随机盐值
func generateSalt() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return time.Now().Format("20060102150405")
	}
	return hex.EncodeToString(bytes)
}

// 加密密码
func encryptPassword(password, salt string) string {
	hash := sha512.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}
