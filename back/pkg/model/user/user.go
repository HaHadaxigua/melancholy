package user

import (
	"github.com/HaHadaxigua/melancholy/pkg/consts"
	"github.com/HaHadaxigua/melancholy/pkg/model"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
)

type User struct {
	model.Model
	Username    string `json:"username" gorm:"username"` // 用户昵称， 可以更改
	Password    string `json:"password"`
	PhoneNumber int    `json:"phoneNumber"`
	Email       string `json:"email"`
	State       int    `json:"state"` // 帐号状态 -30: 逻辑删除		-20: 封禁， -10: 未激活(需要邮箱激活)， 0：正常
	Salt        string `json:"salt"`  // 随机加入的盐
}

func (a *User) TableName() string {
	return `user`
}

//NewAccount 创建新的帐号
func NewUser(username, password, email string) (*User, error) {
	newSalt, err := tools.GenerateSalt()
	if err != nil {
		return nil, err
	}

	encodePwd, err := tools.EncryptPassword(password, newSalt)
	if err != nil {
		return nil, err
	}

	nu := &User{
		Username: username,
		Password: encodePwd,
		Email:    email,
		State:    consts.InActivated,
		Salt:     newSalt,
	}
	return nu, nil
}


