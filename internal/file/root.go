/******
** @date : 2/2/2021 11:50 PM
** @author : zrx
** @description:
******/
package file

import (
	"context"
	"github.com/HaHadaxigua/melancholy"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/file/handler"
	"github.com/HaHadaxigua/melancholy/internal/file/service"
	"github.com/gin-gonic/gin"
)

var Module melancholy.IModule

type module struct {
	FolderService service.IFolderService
}

func New(client *ent.Client, ctx context.Context) *module {
	folderService := service.NewFolderService(client, ctx)
	return &module{
		FolderService: folderService,
	}
}

func (m module) InitService(router gin.IRouter) {
	handler.SetupFileRouters(router)
}
