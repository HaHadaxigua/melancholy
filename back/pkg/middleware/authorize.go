package middleware

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	//获取请求的URI
	obj := c.Request.URL.RequestURI()
	//获取请求方法
	act := c.Request.Method
	//获取用户的角色
	sub := strconv.Itoa(userID)
	//判断策略中是否存在
	if ok, err := e.Enforce(sub, obj, act); ok {
		log.Printf("userID:%d,authorize success", userID)
		c.Next()
	} else {
		e := msg.AuthorizeFailedErr
		log.Printf("userID:%d, %v, caused by:%v", userID, msg.AuthorizeFailedMsg, err)
		c.JSON(http.StatusUnauthorized, e)
		c.Abort()
	}
}
