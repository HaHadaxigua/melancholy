package main

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/conf"
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
// @BasePath /handler/v1

// this is the entrance of the backend
func main() {
	internal.StartServer()
}

func init() {
	fmt.Printf(">>>>>>>>>>>>>>>Hello %s<<<<<<<<<<<<<<<<<<<<<\n", conf.C.Application.Name)
	conf.Setup()
	store.CasbinSetup()
	internal.SetupEnt()
}
