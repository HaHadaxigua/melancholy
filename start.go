/******
** @date : 2/3/2021 12:19 AM
** @author : zrx
** @description:
******/
package melancholy

import (
	"context"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/gin-gonic/gin"
)

type IModule interface {
	StartService(router gin.IRouter)
	SetupStore(client *ent.Client, ctx context.Context)
}