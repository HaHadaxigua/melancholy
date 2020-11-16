package main

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

// this is the entrance of the backend
func main() {
	//_ = pkg.S.Run()
	pkg.StartServer()
}

func init() {
	fmt.Printf(">>>>>>>>>>>>>>>Hello %s<<<<<<<<<<<<<<<<<<<<<\n", conf.C.Application.Name)
	conf.Setup()
	store.Setup()
}
