package store

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"log"
)


var enforcer *casbin.Enforcer

func CasbinSetup() {
	a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3307)/melancholy?charset=utf8", true)
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}
	e, err := casbin.NewEnforcer("etc/rbac_models.conf", a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}
	enforcer = e
}

func GetEnforcer() *casbin.Enforcer{
	return enforcer
}