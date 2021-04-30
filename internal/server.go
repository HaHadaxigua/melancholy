/******
** @date : 2/3/2021 12:04 AM
** @author : zrx
** @description:
******/
package internal

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic"
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/common/oss"
	"github.com/HaHadaxigua/melancholy/internal/common/oss/aliyun"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/HaHadaxigua/melancholy/internal/file"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

type SE struct {
	*gin.Engine
}

var Se SE

func StartServer() {
	gin.SetMode(conf.C.Mode)

	Se.Engine = gin.Default()

	initOssConfig()
	startService(Se.Engine)

	hs := &http.Server{
		Addr:           conf.C.Application.Domain,
		Handler:        Se.Engine,
		ReadTimeout:    time.Duration(conf.C.Application.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.C.Application.WriterTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := hs.ListenAndServe()
	if err != nil {
		logrus.Panicf("Start server failed [%v]", err.Error())
	}
}

// 初始化对象存储配置
func initOssConfig() {
	fmt.Println(conf.C.Oss.EndPoint, conf.C.Oss.AccessKeyID, conf.C.Oss.AccessKeySecret)
	oss.AliyunOss, _ = aliyun.NewAliyunOss(conf.C.Oss.EndPoint, conf.C.Oss.AccessKeyID, conf.C.Oss.AccessKeySecret)
	if oss.AliyunOss == nil {
		panic("init oss failed")
	}
}

func startService(e *gin.Engine) {
	// support cors
	e.Use(middleware.Cors)
	// swagger-path
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := e.Group(consts.ApiV1)

	conn := GetConn()
	basic.Module = basic.New(conn)
	basic.Module.InitService(router)

	file.Module = file.New(conn)
	file.Module.InitService(router)
}
