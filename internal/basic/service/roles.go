package service

import (
	"errors"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/global/msg"
)

var RoleService IRoleService

type IRoleService interface {
	GetAllRoles() ([]*ent.Role, error)
	AddRole(name string) error
	AddUserRoles(uID, rID int) error
	GetRolesByUserID(uID int) ([]*ent.Role, error)
}

type roleService struct {
	roleStore store.IRoleStore
}

func NewRoleService() *roleService {
	return &roleService{
		roleStore: store.RoleStore,
	}
}

//  GetAllRoles 获取所有的角色
func (rs *roleService) GetAllRoles() ([]*ent.Role, error) {
	roles, err := rs.roleStore.ListRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// AddRole 添加用角色
func (rs *roleService) AddRole(name string) error {
	role, err := rs.roleStore.CreateRole(name)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.New(msg.RepeatedRoleMsg)
	}
	return nil
}

// AddRoleToUser 添加用户角色
func (rs *roleService) AddUserRoles(uID, rID int) error {
	err := rs.roleStore.AppendRoleToUser(rID, uID)
	if err != nil {
		if ent.IsNotFound(err) {
			return msg.RoleNotFoundErr
		}
		return err
	}
	return nil
}

// GetRoleByUserID 根据用户id寻找角色
func (rs *roleService) GetRolesByUserID(uID int) ([]*ent.Role, error) {
	res, err := rs.roleStore.ListRolesByUserID(uID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
