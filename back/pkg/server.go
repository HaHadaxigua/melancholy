package pkg

import (
	"github.com/HaHadaxigua/melancholy/pkg/api"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type SE struct {
	*gin.Engine
}

var Se SE

func StartServer() {
	//S = &Server{}
	gin.SetMode(conf.C.Mode)

	Se.Engine = gin.Default()

	api.SetupRouters(Se.Engine)

	hs := &http.Server{
		Addr:           conf.C.Application.Domain,
		Handler:        Se.Engine,
		ReadTimeout:    time.Duration(conf.C.Application.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.C.Application.WriterTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := hs.ListenAndServe()
	if err != nil {
		log.Panicf("Start server failed [%v]", err.Error())
	}
}
