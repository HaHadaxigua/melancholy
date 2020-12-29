package auth

import (
	"github.com/gin-gonic/gin"
)

func SetupAuthRouters(r *gin.RouterGroup) {
	// open
	r.GET("/login", Login)
	r.POST("/register", Register) // 注册用户
	r.GET("/logout", Logout)
}
