package service

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"gorm.io/gorm"
)

var Role RoleService

type RoleService interface {
	NewRole(r *msg.ReqRoleCreate) error
	ListRoles(r *msg.ReqRoleListFilter, withPermission bool) (*msg.RspRoleList, error)
	DeleteRole(rid int) error
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

func (s roleService) ListRoles(r *msg.ReqRoleListFilter, withPermission bool) (*msg.RspRoleList, error) {
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
