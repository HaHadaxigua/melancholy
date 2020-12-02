package api

import "github.com/gin-gonic/gin"

func SetupBasicRouters(e *gin.Engine) {
	e.GET("/auth", Login)
	e.POST("/register", Register) // 注册用户
	e.GET("/logout", Logout)
}
