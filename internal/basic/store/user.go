package store

import (
	"context"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/user"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
)

type IUserStore interface {
	CreateUser(req *ent.User) (*ent.User, error)
	GetUserById(id int) (*ent.User, error)
	GetUserByName(name string) (*ent.User, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetAllUsers() ([]*ent.User, error)
	CheckUserExist(email, password string) int
}

type userStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewUserStore(client *ent.Client, ctx context.Context) *userStore {
	return &userStore{
		client: client,
		ctx:    ctx,
	}
}

// CreateUser 创建用户
func (us *userStore) CreateUser(req *ent.User) (*ent.User, error) {
	u, err := us.client.User.Create().
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetSalt(req.Salt).
		SetEmail(req.Email).
		SetState(req.State).
		Save(us.ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetUserById 根据用户id搜索用户
func (us *userStore) GetUserById(id int) (*ent.User, error) {
	u, err := us.client.User.Query().Where(user.IDEQ(id)).Only(us.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

// GetUserByName 根据用户名找到用户
func (us *userStore) GetUserByName(name string) (*ent.User, error) {
	u, err := us.client.User.Query().Where(user.UsernameEQ(name)).Only(us.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

// GetUserByEmail 根据邮箱找到用户
func (us *userStore) GetUserByEmail(email string) (*ent.User, error) {
	u, err := us.client.User.Query().Where(user.EmailEQ(email)).Only(us.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

// GetAllUsers  找到所有的用户
func (us *userStore) GetAllUsers() ([]*ent.User, error) {
	users, err := us.client.User.Query().All(us.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return []*ent.User{}, nil
		}
		return nil, err
	}
	return users, nil
}

// CheckUserExist判断用户是否存在, 存在则返回用户id, 不存在则返回-1
func (us *userStore) CheckUserExist(email, password string) int {
	u, err := us.client.User.Query().Where(user.EmailEQ(email)).Only(us.ctx)
	if err != nil {
		return -1
	}
	if u != nil {
		flag := tools.VerifyPassword(u.Password, password+u.Salt)
		if flag {
			return u.ID
		}
		return -1
	}
	return -1
}
