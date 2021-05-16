/******
** @date : 2/2/2021 11:50 PM
** @author : zrx
** @description:
******/
package user

import (
	"github.com/HaHadaxigua/melancholy"
	"github.com/HaHadaxigua/melancholy/internal/user/handler"
	"github.com/HaHadaxigua/melancholy/internal/user/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Module melancholy.IModule

type module struct {
	UserService service.UserService
}

func New(conn *gorm.DB) *module {
	userService := service.NewUserService(conn)
	service.UserSvc = userService
	return &module{
		UserService: userService,
	}
}

func (m module) InitService(router gin.IRouter) {
	handler.SetupUserRouters(router)
}
