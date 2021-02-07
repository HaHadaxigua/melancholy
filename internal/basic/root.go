/******
** @date : 2/2/2021 11:50 PM
** @author : zrx
** @description:
******/
package basic

import (
	"context"
	"github.com/HaHadaxigua/melancholy"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/basic/handler"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/gin-gonic/gin"
)

var Module melancholy.IModule

type module struct {
	RoleService service.IRoleService
	UserService service.IUserService
}

func New(client *ent.Client, ctx context.Context) *module {
	roleService := service.NewRoleService(client, ctx)
	userService := service.NewUserService(client, ctx)
	return &module{
		RoleService: roleService,
		UserService: userService,
	}
}

func (m module) InitService(router gin.IRouter) {
	handler.SetupAuthRouters(router)
	handler.SetupRoleRouters(router)
}
