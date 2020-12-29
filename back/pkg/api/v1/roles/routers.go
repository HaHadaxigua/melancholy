package roles

import (
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoleRouters(r *gin.RouterGroup) {
	secured := r.Group("/admin", middleware.JWT, middleware.Authorize)
	secured.POST("/roles", AddRole)
	secured.GET("/roles", GetAllRoles)
	secured.POST("/roles/addRole", AddUserRoles)

}
