package main

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	s *http.Server
)

// this is the entrance of the backend
func main() {

}

func init() {
	startLog()
	startServer()
}

//startLog 启动日志
func startLog() {
	fmt.Printf(">>>>>>>>>>>>>>>Hello %s<<<<<<<<<<<<<<<<<<<<<", conf.Conf.Application.Name)
}

//startServer 启动服务器
func startServer() {
	engine := gin.Default()

	s = &http.Server{
		Addr:           conf.Conf.Application.Domain,
		Handler:        engine,
		ReadTimeout:    time.Duration(conf.Conf.Application.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.Conf.Application.WriterTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Panicf("failed to start server [%s]", err.Error())
	}

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("hello")
	})
}
