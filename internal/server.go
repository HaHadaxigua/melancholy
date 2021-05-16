/******
** @date : 2/3/2021 12:04 AM
** @author : zrx
** @description:
******/
package internal

import (
	"github.com/HaHadaxigua/melancholy/internal/basic"
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/common/cloud"
	aliyunCloud "github.com/HaHadaxigua/melancholy/internal/common/cloud/aliyun"
	"github.com/HaHadaxigua/melancholy/internal/common/oss"
	aliyunOSS "github.com/HaHadaxigua/melancholy/internal/common/oss/aliyun"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/HaHadaxigua/melancholy/internal/file"
	"github.com/HaHadaxigua/melancholy/internal/user"
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
	initCloudConfig()
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

// initOssConfig 初始化对象存储配置
func initOssConfig() {
	oss.AliyunOss, _ = aliyunOSS.NewAliyunOss(conf.C.Oss.EndPoint, conf.C.Oss.AccessKeyID, conf.C.Oss.AccessKeySecret)
	if oss.AliyunOss == nil {
		panic("init oss failed")
	}
	logrus.Info("init ali oss config success!")
}

// initCloudConfig 初始化视频点播配置
func initCloudConfig() {
	cloud.AliyunCloud = aliyunCloud.NewAliyunCloud(conf.C.Cloud.AccessKeyID, conf.C.Cloud.AccessKeySecret)
	cloud.AliyunCloud.InitCloudClient(consts.RegionID)
	cloud.AliyunCloud.InitVodClient()
	logrus.Info("init ali cloud config success!")
}

func startService(e *gin.Engine) {
	e.MaxMultipartMemory = 5 << 30 // 5GB
	// support cors
	e.Use(middleware.Cors)
	// swagger-path
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := e.Group(consts.ApiV1)

	conn := GetConn()
	// 权限管理
	basic.Module = basic.New(conn)
	basic.Module.InitService(router)

	// 文件模块
	file.Module = file.New(conn)
	file.Module.InitService(router)

	// 初始化用户业务
	user.Module = user.New(conn)
	user.Module.InitService(router)
}
