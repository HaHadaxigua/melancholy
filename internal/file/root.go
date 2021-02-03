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
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"github.com/gin-gonic/gin"
)

var Module melancholy.IModule

type module struct {
	FolderService service.IFolderService
}

func new() *module {
	folderService := service.NewFolderService()
	return &module{
		FolderService: folderService,
	}
}

func (m module) StartService(router gin.IRouter) {
	Module = new()
	handler.SetupFileRouters(router)
}

func (m module) SetupStore(client *ent.Client, ctx context.Context) {
	store.FolderStore = store.NewFolderStore(client, ctx)
}
