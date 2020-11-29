package admin

import (
	model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	store "github.com/HaHadaxigua/melancholy/pkg/store/user"
)

//  GetAllRoles 获取所有的角色
func GetAllRoles() ([]*model.Role, error){
	roles, err := store.GetAllRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// AddRole 添加用角色
func AddRole(name string) error{
	err := store.CreateRoles(name)
	if err != nil {
		return err
	}
	return nil
}

//AddRoleToUser 添加用户角色
func AddUserRoles(uID, rID int) error {
	err := store.AddUserRoles(rID, uID)
	if err != nil {
		return err
	}
	return nil
}