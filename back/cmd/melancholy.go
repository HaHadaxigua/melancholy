package main

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://adenzrx.xyz

// @contact.name Aden
// @contact.url https://adenzrx.xyz
// @contact.email 1213604254@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8990
// @BasePath /api/v1

// this is the entrance of the backend
func main() {
	//_ = pkg.S.Run()
	pkg.StartServer()
}

func init() {
	fmt.Printf(">>>>>>>>>>>>>>>Hello %s<<<<<<<<<<<<<<<<<<<<<\n", conf.C.Application.Name)
	conf.Setup()
	//store.Setup()
	store.CasbinSetup()
	store.SetupEnt()
}
