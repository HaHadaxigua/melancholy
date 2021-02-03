package service

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/user"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
)

var UserService IUserService

type IUserService interface {
	CreateUser(r *response.UserRequest) (*ent.User, error)
	FindUserByUsername(r *response.UserRequest) (*ent.User, error)
	ListAllUser() ([]*ent.User, error)
	GetStore() store.IUserStore
	CheckUserExist(email, password string) int
}

type userService struct {
	userStore store.IUserStore
}

func NewUserService() *userService {
	return &userService{
		userStore: store.UserStore,
	}
}

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
func (us *userService) CreateUser(r *response.UserRequest) (*ent.User, error) {
	valid, err := CheckCreateUserReq(r)
	if !valid && err != nil {
		return nil, err
	}

	user, err := us.userStore.GetUserByEmail(r.Email)
	if err != nil {
		e := response.UserHasExistedErr
		e.Data = err.Error()
		return nil, e
	} else if user != nil {
		// fixme
		e := response.UserHasExistedErr
		e.Data = "邮箱已被注册"
		return nil, e
	}

	newUser, err := NewUser(r.Username, r.Password, r.Email)
	u, err := us.userStore.CreateUser(newUser)
	if err != nil {
		e := response.UserCreateErr
		e.Data = err.Error()
		return nil, e
	}
	return u, nil
}

// FindUserByUsername
func (us *userService) FindUserByUsername(r *response.UserRequest) (*ent.User, error) {
	if !tools.CheckUsername(r.Username) {
		return nil, response.UserNameIllegalErr
	}
	tu, err := us.userStore.GetUserByName(r.Username)
	if err != nil {
		return nil, err
	}
	return tu, nil
}

// ListAllUser
func (us *userService) ListAllUser() ([]*ent.User, error) {
	users, err := us.userStore.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userService) GetStore() store.IUserStore {
	return us.userStore
}

func (us *userService) CheckUserExist(email, password string) int {
	return us.userStore.CheckUserExist(email, password)
}

func CheckCreateUserReq(r *response.UserRequest) (bool, error) {
	if !tools.CheckUsername(r.Username) {
		return false, response.UserNameOrPwdIncorrectlyErr
	}
	if !tools.CheckPassword(r.Password) {
		return false, response.UserPwdIllegalErr
	}
	if !tools.CheckEmail(r.Email) {
		return false, response.UserEmailIllegalErr
	}
	return true, nil
}
