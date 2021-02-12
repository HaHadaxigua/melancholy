package service

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"gorm.io/gorm"
)

var User UserService

type UserService interface {
	CreateUser(r *msg.ReqRegister) (*model.User, error)
	FindUserByUsername(username string) ([]*model.User, error)
	GetUserByID(userID int, withRole bool) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	ListUsers(req *msg.ReqUserFilter) (*msg.RspUserList, error)
}

type userService struct {
	store store.UserStore
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{
		store: store.NewUserStore(db),
	}
}

// NewAccount
func newUser(username, password, email string) (*model.User, error) {
	newSalt, err := tools.GenerateSalt()
	if err != nil {
		return nil, err
	}

	encodePwd, err := tools.EncryptPassword(password, newSalt)
	if err != nil {
		return nil, err
	}

	nu := &model.User{
		Username: username,
		Password: encodePwd,
		Email:    email,
		Salt:     newSalt,
	}
	return nu, nil
}

// CreateUser
func (s *userService) CreateUser(r *msg.ReqRegister) (*model.User, error) {
	valid, err := checkCreateUserReq(r)
	if !valid && err != nil {
		return nil, err
	}

	newUser, err := newUser(r.Username, r.Password, r.Email)
	if err = s.store.Create(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

// FindUserByUsername
func (s *userService) FindUserByUsername(username string) ([]*model.User, error) {
	if !tools.CheckUsername(username) {
		return nil, msg.ErrUserNameIllegal
	}
	users, err := s.store.GetUserByName(username)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// ListAllUser
func (s *userService) ListUsers(req *msg.ReqUserFilter) (*msg.RspUserList, error) {
	rsp := &msg.RspUserList{}
	users, total, err := s.store.ListUsers(req)
	if err != nil {
		return nil, err
	}

	rsp.Total = total
	rsp.List = (FunctionalUserMap(users, buildUserRsp)).([]*msg.RspUserListItem)

	return rsp, nil
}

// 找出一个用户所拥有的所有角色
func (s *userService) GetUserByID(userID int, withRole bool) (*model.User, error) {
	return s.store.FindUserById(userID, withRole)
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.store.FindUserByEmail(email)
}

func checkCreateUserReq(r *msg.ReqRegister) (bool, error) {
	if !tools.CheckUsername(r.Username) {
		return false, msg.ErrUserNameIllegal
	}
	if !tools.CheckPassword(r.Password) {
		return false, msg.ErrUserNameOrPwdIncorrectly
	}
	if !tools.CheckEmail(r.Email) {
		return false, msg.ErrUserNameIllegal
	}
	return true, nil
}
