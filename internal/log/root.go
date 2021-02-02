/******
** @date : 2/2/2021 11:50 PM
** @author : zrx
** @description:
******/
package log

import (
	"github.com/HaHadaxigua/melancholy/internal/log/service"
	"github.com/gin-gonic/gin"
)

var ModuleLog moduleLog

type moduleLog struct {
	LogService service.ILogService
}

func new() *moduleLog {
	return &moduleLog{
		LogService: service.NewLogService(),
	}
}

func (m moduleLog) StartService(router gin.IRouter) {
	new()
}
