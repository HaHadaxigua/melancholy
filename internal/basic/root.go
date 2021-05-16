/******
** @date : 2/2/2021 11:50 PM
** @author : zrx
** @description:
******/
package basic

import (
	"github.com/HaHadaxigua/melancholy"
	"github.com/HaHadaxigua/melancholy/internal/basic/handler"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Module melancholy.IModule

type module struct {
	UserService       service.UserService
	RoleService       service.RoleService
	PermissionService service.PermissionService
}

func New(conn *gorm.DB) *module {
	userService := service.NewUserService(conn)
	roleService := service.NewRoleService(conn)
	permissionService := service.NewPermissionService(conn)

	service.User = userService
	service.Role = roleService
	service.Permission = permissionService

	return &module{
		RoleService:       roleService,
		UserService:       userService,
		PermissionService: permissionService,
	}
}

func (m module) InitService(router gin.IRouter) {
	handler.SetupRouters(router)
	handler.SetupBasicRouters(router)
	handler.SetupUserRouters(router)
}
