/******
** @date : 2/2/2021 11:50 PM
** @author : zrx
** @description:
******/
package file

import (
	"github.com/HaHadaxigua/melancholy"
	"github.com/HaHadaxigua/melancholy/internal/file/handler"
	"github.com/HaHadaxigua/melancholy/internal/file/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Module melancholy.IModule

type module struct {
	FileService service.FileService
}

func New(conn *gorm.DB) *module {
	fileService := service.NewFileService(conn)
	service.File = fileService
	return &module{
		FileService: fileService,
	}
}

func (m module) InitService(router gin.IRouter) {
	handler.SetupFileRouters(router)
}
