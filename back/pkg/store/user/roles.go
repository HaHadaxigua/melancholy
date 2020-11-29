package user

import (
	"errors"
	model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	"gorm.io/gorm"
)

// CreateRoles 创建角色
func CreateRoles(name string) error {
	db := store.GetConn()
	r := &model.Role{Name: name}
	if err := db.Create(r).Error; err != nil {
		return err
	}
	return nil
}

// GetAllRoles 获取所有的角色
func GetAllRoles() ([]*model.Role, error) {
	db := store.GetConn()
	var roles []*model.Role
	if err := db.Model(&model.Role{}).Scan(&roles).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return []*model.Role{}, nil
		default:
			return nil, err
		}
	}
	return roles, nil
}

// AddUserRoles 添加角色给用户  todo：校验角色是否真实存在
func AddUserRoles(roleID, userID int) error {
	db := store.GetConn()
	ur := &model.XUserRoles{RoleID: roleID, UserID: userID}

	res := db.Where(&model.XUserRoles{RoleID: roleID, UserID: userID}).Find(&ur)
	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
	}

	if res.RowsAffected > 0 {
		return msg.RepeatedRoleErr
	}

	if err := db.Create(ur).Error; err != nil {
		return err
	}

	return nil
}
