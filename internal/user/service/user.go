package service

import (
	basicStore "github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/user/store"
	"gorm.io/gorm"
)

var UserSvc UserService

type UserService interface {

}

type userService struct {
	basicStore basicStore.UserStore
	userStore  store.UserStore
}

func NewUserService(conn *gorm.DB) *userService {
	return &userService{
		basicStore: basicStore.NewUserStore(conn),
		userStore:  store.NewUserStore(conn),
	}
}

