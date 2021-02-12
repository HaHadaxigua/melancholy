package store

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(req *model.User) error
	FindUserById(id int, withRole bool) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	GetUserByName(name string) ([]*model.User, error)
	ListUsers(req *msg.ReqUserFilter) ([]*model.User, int, error)
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
	var user = &model.User{}
	query := s.db.Model(&model.User{ID: id})
	if withRole {
		query.Preload("Roles").Preload("Roles.Permissions")
	}
	if err := query.Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByName 根据用户名找到用户
func (s *userStore) GetUserByName(name string) ([]*model.User, error) {
	var users []*model.User
	if err := s.db.Model(&model.User{}).Where("username like %?%", name).Find(&users).Error; err != nil {
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

func (s *userStore) ListUsers(req *msg.ReqUserFilter) ([]*model.User, int, error) {
	var users []*model.User
	query := s.db.Model(&model.User{})

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Offset(req.Offset).Limit(req.Limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, int(total), nil
}
