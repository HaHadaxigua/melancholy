package roles

import (
	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.RouterGroup) {
	routers := r.Group("/admin")
	{
		routers.POST("/roles", AddRole)
		routers.GET("/roles", GetAllRoles)
		routers.POST("/roles/addRole", AddUserRoles)
	}
}
