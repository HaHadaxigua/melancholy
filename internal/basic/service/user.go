package service

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/common/oss"
	"github.com/HaHadaxigua/melancholy/internal/common/oss/aliyun"
	"github.com/HaHadaxigua/melancholy/utils"
	"gorm.io/gorm"
)

var User UserService

type UserService interface {
	CreateUser(r *msg.ReqRegister) (*model.User, error)
	FindUserByUsername(username string) ([]*model.User, error)
	GetUserByID(userID int, withRole bool) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	ListUsers(req *msg.ReqUserFilter, withRoles bool) (*msg.RspUserList, error)
	RoleManager(uid, rid, operation int) error

	SetUserInfo(req *msg.ReqSetUserInfo) error
	UpdateAvatar(req *msg.ReqUpdateAvatar) error
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

	newUser, err := newUser(utils.GenUUID(), r.Password, r.Email)
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
func (s *userService) ListUsers(req *msg.ReqUserFilter, withRoles bool) (*msg.RspUserList, error) {
	rsp := &msg.RspUserList{}
	users, total, err := s.store.ListUsers(req, withRoles)
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
	//if !tools.CheckUsername(r.Username) {
	//	return false, msg.ErrUserNameIllegal
	//}
	if !tools.CheckPassword(r.Password) {
		return false, msg.ErrUserNameOrPwdWrong
	}
	if !tools.CheckEmail(r.Email) {
		return false, msg.ErrUserNameIllegal
	}
	return true, nil
}

func (s *userService) RoleManager(uid, rid, operation int) error {
	user, err := s.GetUserByID(uid, true)
	if err != nil {
		return err
	}
	_, err = Role.GetRoleByID(rid, false)
	if err != nil {
		return err
	}

	roles := FunctionalRoleFilter(user.Roles, func(r *model.Role) bool {
		if r.ID == rid {
			return true
		}
		return false
	})

	switch operation {
	case consts.AppendRole:
		if len(roles) > 0 {
			return nil
		}
	case consts.RemoveRole:
		if len(roles) < 1 {
			return nil
		}
	}
	return s.store.RoleManager(uid, rid, operation)
}

// SetUserInfo 设置用户信息
func (s userService) SetUserInfo(req *msg.ReqSetUserInfo) error {
	mem, err := utils.StructToMap(req, "gorm")
	if err != nil {
		return err
	}
	if len(mem) < 1 {
		return fmt.Errorf("请求解析错误")
	}
	return s.store.UpdateUserInfo(mem, req.UserID)
}

func (s userService) UpdateAvatar(req *msg.ReqUpdateAvatar) error {
	user, err := s.GetUserByID(req.UserID, false)
	if err != nil {
		return err
	}

	if user.OssEndPoint == "" || user.OssAccessKey == "" || user.CloudAccessSecret == "" {
		return err
	}

	bucketName, ossAddress := oss.BuildBucketNameAndAddress(req.UserID, req.FileHeader.Filename)

	aliOss, err := aliyun.NewAliyunOss(user.OssEndPoint, user.OssAccessKey, user.OssAccessSecret)
	if err = aliOss.UploadBytes(bucketName, req.FileHeader.Filename, req.Data); err != nil {
		return err
	}

	if err = s.store.UpdateOneColumn("avatar", ossAddress, req.UserID); err != nil {
		return err
	}
	return nil
}
