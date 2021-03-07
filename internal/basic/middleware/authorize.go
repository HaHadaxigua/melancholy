package middleware

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	gConst "github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/gin-gonic/gin"
)

// rbac权限认证
func Authorize(c *gin.Context) {
	userID := c.GetInt(gConst.UserID)
	if userID <= 0 {
		Forbidden()
	}

	user, err := service.User.GetUserByID(userID, true)
	if err != nil {
		Forbidden()
	}

	roles := service.FunctionalRoleFilter(user.Roles, func(r *model.Role) bool {
		if r.Name == consts.Admin {
			return true
		}
		return false
	})
	if len(roles) > 0 {
		c.Next()
		return
	}
	c.Abort()
	Forbidden()
}

// CheckRole 判断用户是否具有响应的角色
func CheckRole(roleID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		_user, ok := c.Get(gConst.User)
		if !ok {
			c.Abort()
			Forbidden()
			return
		}
		user := _user.(*model.User)
		roles := service.FunctionalRoleFilter(user.Roles, func(r *model.Role) bool {
			if r.ID == roleID {
				return true
			}
			return false
		})
		if len(roles) > 0 {
			c.Next()
			return
		}
		c.Abort()
		Forbidden()
		return
	}
}
