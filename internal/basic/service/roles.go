package service

import (
	"context"
	"errors"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
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

func NewRoleService(client *ent.Client, ctx context.Context) *roleService {
	return &roleService{
		roleStore: store.NewRoleStore(client, ctx),
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
		return errors.New(response.RepeatedRoleMsg)
	}
	return nil
}

// AddRoleToUser 添加用户角色
func (rs *roleService) AddUserRoles(uID, rID int) error {
	err := rs.roleStore.AppendRoleToUser(rID, uID)
	if err != nil {
		if ent.IsNotFound(err) {
			return response.RoleNotFoundErr
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
