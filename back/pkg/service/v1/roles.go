package v1

import (
	"errors"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	store2 "github.com/HaHadaxigua/melancholy/pkg/store"
)

//  GetAllRoles 获取所有的角色
func GetAllRoles() ([]*ent.Role, error){
	roles, err := store2.GetAllRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// AddRole 添加用角色
func AddRole(name string) error{
	role, err := store2.CreateRole(name)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.New(msg.RepeatedRoleMsg)
	}
	return nil
}

// AddRoleToUser 添加用户角色
func AddUserRoles(uID, rID int) error {
	err := store2.AddUserRoles(rID, uID)
	if err != nil {
		if ent.IsNotFound(err) {
			return msg.RoleNotFoundErr
		}
		return err
	}
	return nil
}

// GetRoleByUserID 根据用户id寻找角色
func GetRolesByUserID(uID int) ([]*ent.Role, error){
	res, err := store2.GetRolesByUserID(uID)
	if err != nil {
		return nil, err
	}
	return res , nil
}
