package service

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"time"

	"github.com/lemonoa/LemonOA-Go/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// Login 用户登录
func (s *AuthService) Login(username, password string, ip, userAgent string) (string, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("用户不存在")
		}
		return "", err
	}

	if user.Status != 1 {
		return "", errors.New("用户已被禁用")
	}

	// 验证密码
	if s.encryptPassword(password, user.Salt) != user.Password {
		// 记录登录失败日志
		s.createLoginLog(user.ID, ip, userAgent, 2, "密码错误")
		return "", errors.New("密码错误")
	}

	// 更新最后登录信息
	now := time.Now()
	s.db.Model(&user).Updates(map[string]interface{}{
		"last_login_at": &now,
		"last_login_ip": ip,
	})

	// 记录登录成功日志
	s.createLoginLog(user.ID, ip, userAgent, 1, "")

	// 生成JWT token
	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserInfo 获取用户信息
func (s *AuthService) GetUserInfo(userID uint) (*model.User, error) {
	var user model.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserPermissions 获取用户权限列表
func (s *AuthService) GetUserPermissions(userID uint) ([]model.Permission, error) {
	var permissions []model.Permission
	err := s.db.Raw(`
		SELECT DISTINCT p.* FROM permissions p
		INNER JOIN role_permissions rp ON p.id = rp.permission_id
		INNER JOIN user_roles ur ON rp.role_id = ur.role_id
		WHERE ur.user_id = ? AND p.status = 1
		ORDER BY p.sort ASC
	`, userID).Find(&permissions).Error
	return permissions, err
}

// ChangePassword 修改密码
func (s *AuthService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	var user model.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return err
	}

	if s.encryptPassword(oldPassword, user.Salt) != user.Password {
		return errors.New("原密码错误")
	}

	// 生成新的盐值和密码
	salt := s.generateSalt()
	password := s.encryptPassword(newPassword, salt)

	return s.db.Model(&user).Updates(map[string]interface{}{
		"salt":     salt,
		"password": password,
	}).Error
}

// 生成JWT token
func (s *AuthService) generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(viper.GetInt("jwt.expire")) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("jwt.secret")))
}

// 生成32位随机盐值
func (s *AuthService) generateSalt() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return time.Now().Format("20060102150405")
	}
	return hex.EncodeToString(bytes)
}

// 加密密码
func (s *AuthService) encryptPassword(password, salt string) string {
	hash := sha512.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// 创建登录日志
func (s *AuthService) createLoginLog(userID uint, ip, userAgent string, status int, message string) {
	log := &model.LoginLog{
		UserID:    userID,
		IP:        ip,
		UserAgent: userAgent,
		Status:    status,
		Message:   message,
	}
	s.db.Create(log)
}

// GetUserList 获取用户列表
func (s *AuthService) GetUserList(status int, keyword string, page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := s.db.Model(&model.User{})
	if status > 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("username LIKE ? OR real_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateUser 创建用户
func (s *AuthService) CreateUser(user *model.User) error {
	// 生成密码盐值
	user.Salt = s.generateSalt()
	// 加密密码
	user.Password = s.encryptPassword(user.Password, user.Salt)

	return s.db.Create(user).Error
}

// UpdateUser 更新用户
func (s *AuthService) UpdateUser(user *model.User) error {
	if user.ID == 0 {
		return errors.New("user id is required")
	}

	// 如果密码不为空,则需要重新加密
	if user.Password != "" {
		user.Salt = s.generateSalt()
		user.Password = s.encryptPassword(user.Password, user.Salt)
	}

	return s.db.Model(user).Updates(user).Error
}

// DeleteUser 删除用户
func (s *AuthService) DeleteUser(id uint) error {
	return s.db.Delete(&model.User{}, id).Error
}

// ResetPassword 重置用户密码
func (s *AuthService) ResetPassword(id uint, password string) error {
	user := &model.User{ID: id}
	if err := s.db.First(user).Error; err != nil {
		return err
	}

	// 生成新的盐值和密码
	user.Salt = s.generateSalt()
	user.Password = s.encryptPassword(password, user.Salt)

	return s.db.Model(user).Updates(map[string]interface{}{
		"salt":     user.Salt,
		"password": user.Password,
	}).Error
}

// GetRoleList 获取角色列表
func (s *AuthService) GetRoleList() ([]model.Role, error) {
	var roles []model.Role
	err := s.db.Find(&roles).Error
	return roles, err
}

// CreateRole 创建角色
func (s *AuthService) CreateRole(role *model.Role) error {
	return s.db.Create(role).Error
}

// UpdateRole 更新角色
func (s *AuthService) UpdateRole(role *model.Role) error {
	if role.ID == 0 {
		return errors.New("role id is required")
	}
	return s.db.Model(role).Updates(role).Error
}

// DeleteRole 删除角色
func (s *AuthService) DeleteRole(id uint) error {
	// 检查是否有用户关联此角色
	var count int64
	if err := s.db.Model(&model.UserRole{}).Where("role_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete role with associated users")
	}

	return s.db.Delete(&model.Role{}, id).Error
}

// GetRolePermissions 获取角色权限
func (s *AuthService) GetRolePermissions(roleID uint) ([]model.Permission, error) {
	var permissions []model.Permission
	err := s.db.Model(&model.Permission{}).
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ?", roleID).
		Find(&permissions).Error
	return permissions, err
}

// UpdateRolePermissions 更新角色权限
func (s *AuthService) UpdateRolePermissions(roleID uint, permissionIDs []uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 删除原有权限
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}

		// 添加新权限
		for _, pid := range permissionIDs {
			rp := &model.RolePermission{
				RoleID:       roleID,
				PermissionID: pid,
			}
			if err := tx.Create(rp).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// GetPermissionList 获取权限列表
func (s *AuthService) GetPermissionList() ([]model.Permission, error) {
	var permissions []model.Permission
	err := s.db.Find(&permissions).Error
	return permissions, err
}

// CreatePermission 创建权限
func (s *AuthService) CreatePermission(permission *model.Permission) error {
	return s.db.Create(permission).Error
}

// UpdatePermission 更新权限
func (s *AuthService) UpdatePermission(permission *model.Permission) error {
	if permission.ID == 0 {
		return errors.New("permission id is required")
	}
	return s.db.Model(permission).Updates(permission).Error
}

// DeletePermission 删除权限
func (s *AuthService) DeletePermission(id uint) error {
	// 检查是否有角色关联此权限
	var count int64
	if err := s.db.Model(&model.RolePermission{}).Where("permission_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete permission with associated roles")
	}

	return s.db.Delete(&model.Permission{}, id).Error
}
