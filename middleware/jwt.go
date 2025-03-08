package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/lemonoa/LemonOA-Go/database"

	"github.com/lemonoa/LemonOA-Go/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// JWT中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供token"})
			c.Abort()
			return
		}

		// 从Bearer token中提取JWT
		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token格式错误"})
			c.Abort()
			return
		}

		// 解析JWT
		claims, err := parseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// 将用户ID存储在上下文中
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}

// 解析JWT token
func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})

	if err != nil {
		return nil, errors.New("无效的token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("无效的token")
	}

	return claims, nil
}

// RequirePermission 权限验证中间件
func RequirePermission(permissionCode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
			c.Abort()
			return
		}

		// 查询用户角色
		var userRoles []model.UserRole
		if err := database.DB.Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户角色失败"})
			c.Abort()
			return
		}

		// 如果没有任何角色
		if len(userRoles) == 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "没有任何角色权限"})
			c.Abort()
			return
		}

		// 获取角色ID列表
		var roleIDs []uint
		for _, ur := range userRoles {
			roleIDs = append(roleIDs, ur.RoleID)
		}

		// 查询角色是否包含超级管理员
		var count int64
		if err := database.DB.Model(&model.Role{}).Where("id IN ? AND code = ?", roleIDs, "super_admin").Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "验证超级管理员失败"})
			c.Abort()
			return
		}

		// 如果是超级管理员，直接放行
		if count > 0 {
			c.Next()
			return
		}

		// 查询角色权限
		var hasPermission bool
		err := database.DB.Raw(`
			SELECT EXISTS (
				SELECT 1 FROM permissions p
				INNER JOIN role_permissions rp ON p.id = rp.permission_id
				WHERE rp.role_id IN ?
				AND p.code = ?
				AND p.status = 1
			)
		`, roleIDs, permissionCode).Scan(&hasPermission).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "验证权限失败"})
			c.Abort()
			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "没有操作权限"})
			c.Abort()
			return
		}

		c.Next()
	}
}
