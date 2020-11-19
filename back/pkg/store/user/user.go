package user

import (
	model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

// CreateUser 创建用户
func CreateUser(u *model.User) error {
	db := store.GetConn()
	result := db.Create(u)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

//GetUserById 根据用户id搜索用户
func GetUserById(id int) (*model.User, error) {
	db := store.GetConn()
	var u *model.User
	if err := db.Model(&u).Where("id = ? AND deleted_at is null", id).Scan(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

//GetUserByName 根据用户名找到用户
func GetUserByName(name string) (*model.User, error) {
	db := store.GetConn()
	var u *model.User
	if err := db.Model(&u).Where("name = ? AND deleted_at is null", name).Scan(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

//GetUserByEmail 根据邮箱找到用户
func GetUserByEmail(email string) (*model.User, error) {
	db := store.GetConn()
	u := &model.User{}
	result := db.Model(&u).Where("email = ? AND deleted_at is null", email).Scan(u)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected >= 1 {
		return u, nil
	}else {
		return nil, nil
	}
}

//GetAllUsers  找到所有的用户
func GetAllUsers() ([]*model.User, error) {
	db := store.GetConn()
	var users []*model.User
	var u *model.User
	if err := db.Model(&u).Find(&u).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
