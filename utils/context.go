package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GetCurrentUserID 从Gin Context中获取当前用户ID
func GetCurrentUserID(c *gin.Context) (uint, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("用户未登录")
	}

	// JWT claims中的user_id可能是float64类型
	switch v := userID.(type) {
	case uint:
		return v, nil
	case float64:
		return uint(v), nil
	case int:
		return uint(v), nil
	case int64:
		return uint(v), nil
	default:
		return 0, errors.New("无效的用户ID类型")
	}
}