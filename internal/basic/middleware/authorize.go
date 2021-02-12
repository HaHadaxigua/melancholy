package middleware

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// rbac权限认证
func Authorize(c *gin.Context) {
	userID := c.GetInt("user_id")
	if userID <= 0 {
		c.Abort()
		c.JSON(http.StatusInternalServerError, nil)
	}

	user, err := service.User.GetUserByID(userID, true)
	if err != nil {
		c.Abort()
	}

	roles := service.FunctionalRoleFilter(user.Roles, func(r *model.Role) bool {
		if r.Name == consts.Admin {
			return true
		}
		return false
	})
	if len(roles) > 0 {
		c.Next()
	} else {
		c.Abort()
	}
}
