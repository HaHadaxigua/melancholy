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
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/gin-gonic/gin"
)

var Module melancholy.IModule

type module struct {
	RoleService service.IRoleService
	UserService service.IUserService
}

func new() *module {
	roleService := service.NewRoleService()
	userService := service.NewUserService()
	return &module{
		RoleService: roleService,
		UserService: userService,
	}
}

func (m module) StartService(router gin.IRouter) {
	Module = new()
	handler.SetupAuthRouters(router)
	handler.SetupRoleRouters(router)
}

func (m module) SetupStore(client *ent.Client, ctx context.Context) {
	store.RoleStore = store.NewRoleStore(client, ctx)
	store.UserStore = store.NewUserStore(client, ctx)
}
