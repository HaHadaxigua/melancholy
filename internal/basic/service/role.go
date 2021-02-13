package service

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"gorm.io/gorm"
)

var Role RoleService

type RoleService interface {
	NewRole(r *msg.ReqRoleCreate) error
	ListRoles(r *msg.ReqRoleFilter, withPermission bool) (*msg.RspRoleList, error)
	DeleteRole(rid int) error
	GetRoleByID(rid int, withPerms bool) (*model.Role, error)
	PermissionManager(rid, pid, operation int) error
}

type roleService struct {
	store store.RoleStore
}

func NewRoleService(db *gorm.DB) *roleService {
	return &roleService{
		store: store.NewRoleStore(db),
	}
}

func (s roleService) NewRole(r *msg.ReqRoleCreate) error {
	role := &model.Role{
		Name: r.RoleName,
	}
	return s.store.InsertRole(role)
}

func (s roleService) ListRoles(r *msg.ReqRoleFilter, withPermission bool) (*msg.RspRoleList, error) {
	rsp := &msg.RspRoleList{}

	roles, total, err := s.store.ListRoles(r, withPermission)
	if err != nil {
		return nil, err
	}

	rsp.Total = total
	rsp.List = (FunctionalRoleMap(roles, buildRoleRsp)).([]*msg.RspRoleListItem)

	return rsp, nil
}

func (s roleService) DeleteRole(rid int) error {
	return s.store.Delete(rid)
}

func (s roleService) GetRoleByID(rid int, withPerms bool) (*model.Role, error) {
	return s.store.GetRoleByID(rid, withPerms)
}

func (s roleService) PermissionManager(rid, pid, operation int) error {
	role, err := s.GetRoleByID(rid, true)
	if err != nil {
		return err
	}
	_, err = Permission.FindPermission(pid)
	if err != nil {
		return err
	}

	perms := FunctionalPermissionFilter(role.Permissions, func(p *model.Permission) bool {
		if p.ID == pid {
			return true
		}
		return false
	})

	switch operation {
	case consts.AppendPermission:
		if len(perms) > 0 {
			return nil
		}
	case consts.RemovePermission:
		if len(perms) < 1 {
			return nil
		}
	}
	return s.store.PermissionManager(rid, pid, operation)
}
