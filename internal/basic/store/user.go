package store

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(req *model.User) error
	FindUserById(id int, withRole bool) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUsersByID(id []int, withRole bool) ([]*model.User, error)
	GetUserByName(name string) ([]*model.User, error)
	ListUsers(req *msg.ReqUserFilter, withRoles bool) ([]*model.User, int, error)
	RoleManager(uid, rid, operation int) error

	UpdateUserInfo(mem map[string]interface{}, userID int) error
	UpdateOneColumn(fieldName string, value interface{}, userID int) error // 给出列名和值进行更新
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{
		db: db,
	}
}

func (s *userStore) Create(user *model.User) error {
	return s.db.Create(user).Error
}

// GetUserById 根据用户id搜索用户
func (s *userStore) FindUserById(id int, withRole bool) (*model.User, error) {
	var user model.User
	query := s.db.Model(&model.User{}).Where("id = ?", id)
	if withRole {
		query.Preload("Roles").Preload("Roles.Permissions")
	}
	if err := query.Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUsersByID 根据id批量获取用户
func (s *userStore) FindUsersByID(id []int, withRole bool) ([]*model.User, error) {
	var users []*model.User
	query := s.db.Model(&model.User{}).Where("id in ?", id)
	if withRole {
		query.Preload("Roles").Preload("Roles.Permissions")
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByName 根据用户名找到用户
func (s *userStore) GetUserByName(name string) ([]*model.User, error) {
	var users []*model.User
	if err := s.db.Model(&model.User{}).Where("username like ?", "%"+name+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByEmail 根据邮箱找到用户
func (s *userStore) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := s.db.Model(&model.User{}).Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userStore) ListUsers(req *msg.ReqUserFilter, withRoles bool) ([]*model.User, int, error) {
	var users []*model.User
	query := s.db.Model(&model.User{})

	if withRoles {
		query = query.Preload("Roles")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Offset(req.Offset).Limit(req.Limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, int(total), nil
}

func (s *userStore) RoleManager(uid, rid, operation int) error {
	query := s.db.Model(&model.User{ID: uid}).Association("Roles")
	switch operation {
	case consts.AppendRole:
		return query.Append(&model.Role{ID: rid})
	case consts.RemoveRole:
		return query.Delete(&model.Role{ID: rid})
	default:
		return nil
	}
}

// UpdateUserInfo 更新多列用户信息
func (s userStore) UpdateUserInfo(mem map[string]interface{}, userID int) error {
	query := s.db.Model(&model.User{ID: userID})
	if err := query.Updates(mem).Error; err != nil {
		return err
	}
	return nil
}

// UpdateOneColumn 根据给定列进行更新
func (s userStore) UpdateOneColumn(fieldName string, value interface{}, userID int) error {
	query := s.db.Model(&model.User{ID: userID})
	if err := query.Where("id = ?", userID).Update(fieldName, value).Error; err != nil {
		return err
	}
	return nil
}
