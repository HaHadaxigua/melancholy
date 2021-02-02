package internal

import (
	"context"
	"fmt"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/basic"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	log "github.com/sirupsen/logrus"
	_ "gorm.io/driver/mysql"
	"net/url"
)

//var client *ent.Client
//var ctx context.Context

func SetupEnt() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&time_zone=%s",
		conf.C.Database.Username,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.Name,
		url.QueryEscape("'Asia/Shanghai'"))

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	ctx := context.Background()

	setupStore(client, ctx)
}

func setupStore(client *ent.Client, ctx context.Context) {
	basic.Module.SetupStore(client, ctx)

	//FolderStore = file.store.NewFolderStore(client, ctx)
	//UserStore = NewUserStore(client, ctx)
	//RoleStore = NewRoleStore(client, ctx)
	//ExitLogStore = NewExitLogStore(client, ctx)
}
