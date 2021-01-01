package store

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/user"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
)

// CreateUser 创建用户
func CreateUser(req *ent.User) (*ent.User, error) {
	client := GetClient()
	ctx := GetCtx()
	u, err := client.User.Create().
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetSalt(req.Salt).
		SetEmail(req.Email).
		SetState(req.State).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetUserById 根据用户id搜索用户
func GetUserById(id int) (*ent.User, error) {
	client := GetClient()
	ctx := GetCtx()
	u, err := client.User.Query().Where(user.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err){
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

// GetUserByName 根据用户名找到用户
func GetUserByName(name string) (*ent.User, error) {
	client := GetClient()
	ctx := GetCtx()
	u, err := client.User.Query().Where(user.UsernameEQ(name)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err){
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

// GetUserByEmail 根据邮箱找到用户
func GetUserByEmail(email string) (*ent.User, error) {
	client := GetClient()
	ctx := GetCtx()
	u, err := client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil{
		if ent.IsNotFound(err){
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

// GetAllUsers  找到所有的用户
func GetAllUsers() ([]*ent.User, error) {
	client := GetClient()
	ctx := GetCtx()
	users, err := client.User.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err){
			return []*ent.User{}, nil
		}
		return nil, err
	}
	return users, nil
}

// CheckUserExist判断用户是否存在, 存在则返回用户id, 不存在则返回-1
func CheckUserExist(email, password string) int {
	client := GetClient()
	ctx := GetCtx()
	u, err := client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
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
