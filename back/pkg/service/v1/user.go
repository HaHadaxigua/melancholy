package v1

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/user"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
)


// NewAccount
func NewUser(username, password, email string) (*ent.User, error) {
	newSalt, err := tools.GenerateSalt()
	if err != nil {
		return nil, err
	}

	encodePwd, err := tools.EncryptPassword(password, newSalt)
	if err != nil {
		return nil, err
	}

	nu := &ent.User{
		Username: username,
		Password: encodePwd,
		Email:    email,
		State:    user.State0,
		Salt:     newSalt,
	}
	return nu, nil
}


// CreateUser
func CreateUser(r *msg.UserRequest) (*ent.User, error) {
	valid, err := CheckReq(r)
	if !valid && err != nil {
		return nil, err
	}

	user, err := store.GetUserByEmail(r.Email)
	if err != nil {
		e := msg.UserHasExistedErr
		e.Data = err.Error()
		return nil, e
	} else if user != nil {
		e := msg.UserHasExistedErr
		e.Data = "邮箱已被注册"
		return nil, e
	}

	newUser, err := NewUser(r.Username, r.Password, r.Email)
	u, err := store.CreateUser(newUser)
	if err != nil {
		e := msg.UserCreateErr
		e.Data = err.Error()
		return nil, e
	}
	return u, nil
}

// FindUserByUsername
func FindUserByUsername(r *msg.UserRequest) (*ent.User, error) {
	if !CheckUsername(r.Username) {
		return nil, msg.UserNameIllegalErr
	}
	tu, err := store.GetUserByName(r.Username)
	if err != nil {
		return nil, err
	}
	return tu, nil
}

// ListAllUser
func ListAllUser() ([]*ent.User, error) {
	users, err := store.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// VerifyReq
func CheckReq(r *msg.UserRequest) (bool, error) {
	if !CheckUsername(r.Username) {
		return false, msg.UserNameOrPwdIncorrectlyErr
	}
	if !CheckPassword(r.Password) {
		return false, msg.UserPwdIllegalErr
	}
	if !CheckEmail(r.Email) {
		return false, msg.UserEmailIllegalErr
	}
	return true, msg.OK
}

// todo 用户名合法性
func CheckUsername(username string) bool {
	return true
}

// todo 密码合法性
func CheckPassword(password string) bool {
	return true
}

// todo 邮箱合法性
func CheckEmail(email string) bool {
	return true
}