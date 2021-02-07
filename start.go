/******
** @date : 2/3/2021 12:19 AM
** @author : zrx
** @description:
******/
package melancholy

import (
	"github.com/gin-gonic/gin"
)

type IModule interface {
	InitService(router gin.IRouter)
}