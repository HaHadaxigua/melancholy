package middleware

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/service/v1"
	"github.com/HaHadaxigua/melancholy/pkg/store"
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
	roles, err := v1.RoleService.GetRolesByUserID(userID)

	for _, role := range roles {
		//判断策略中是否存在
		if ok, err := e.Enforce(role.Name, obj, act); ok {
			log.Printf("userID:%d,authorize success", userID)
			c.Next()
			return
		} else {
			err = msg.AuthorizeFailedErr
			log.Printf("userID:%d, %v, caused by:%v", userID, msg.AuthorizeFailedMsg, err)
		}
	}
	c.Abort()
	me := msg.AuthorizeFailedErr
	me.Data = err
	c.JSON(http.StatusUnauthorized, me)
}
