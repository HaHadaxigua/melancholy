package service

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	basicStore "github.com/HaHadaxigua/melancholy/internal/basic/store"
	fModel "github.com/HaHadaxigua/melancholy/internal/user/model"
	"github.com/HaHadaxigua/melancholy/internal/user/msg"
	"github.com/HaHadaxigua/melancholy/internal/user/store"
	"gorm.io/gorm"
)

var UserSvc UserService

type UserService interface {
	GetFriendList(req *msg.ReqFriendList) (*msg.RspFriendList, error)
	AddFriend(req *msg.ReqAddFriend) error
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

// GetFriendList 获取用户的好友列表
func (service userService) GetFriendList(req *msg.ReqFriendList) (*msg.RspFriendList, error) {
	friends, total, err := service.userStore.FindOnesFriend(req)
	if err != nil {
		return nil, err
	}

	users, err := service.basicStore.FindUsersByID(fModel.Friends(friends).GetFriendUserIDList(), false)
	if err != nil {
		return nil, err
	}
	usersIDMap := model.Users(users).ToIDMap()

	rsp := fModel.Friends(friends).ToRsp(usersIDMap, total)
	return rsp, nil
}

//  AddFriend 添加好友请求
func (service userService) AddFriend(req *msg.ReqAddFriend) error {
	// 1. 判断用户是否存在
	targetUser, err := service.basicStore.FindUserById(req.TargetUserID, false)
	if err != nil {
		return err
	}
	// 2. todo 用户存在， 判断是否存在好友关系
	fmt.Println(targetUser)
	return nil
}
