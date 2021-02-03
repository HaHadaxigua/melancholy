package middleware

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// rbac权限认证
func Authorize(c *gin.Context) {
	e := store.GetEnforcer()
	//从DB加载策略
	err := e.LoadPolicy()
	if err != nil {
		c.Abort()
		return
	}
	userID := c.GetInt("user_id")
	if userID <= 0 {
		c.Abort()
		c.JSON(http.StatusInternalServerError, nil)
	}

	//获取请求的URI
	obj := strings.Split(c.Request.URL.RequestURI(), "/")[3]
	//获取请求方法
	act := c.Request.Method
	//获取用户的角色
	roles, err := service.RoleService.GetRolesByUserID(userID)

	for _, role := range roles {
		//判断策略中是否存在
		if ok, err := e.Enforce(role.Name, obj, act); ok {
			log.Printf("userID:%d,authorize success", userID)
			c.Next()
			return
		} else {
			err = response.AuthorizeFailedErr
			log.Printf("userID:%d, %v, caused by:%v", userID, response.AuthorizeFailedMsg, err)
		}
	}
	c.Abort()
	me := response.AuthorizeFailedErr
	me.Data = err
	c.JSON(http.StatusUnauthorized, me)
}
