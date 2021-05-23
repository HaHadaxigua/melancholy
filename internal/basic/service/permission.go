/******
** @date : 2/11/2021 1:59 PM
** @author : zrx
** @description:
******/
package service

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/store"
	"gorm.io/gorm"
)

var Permission PermissionService

type PermissionService interface {
	NewPermission(r *msg.ReqPermissionCreate) error
	ListPermission(r *msg.ReqPermissionFilter) (*msg.RspPermList, error)
	FindPermission(pid int) (*model.Permission, error)
	DeletePermission(pid int) error
}

type permissionService struct {
	store store.PermissionStore
}

func NewPermissionService(db *gorm.DB) *permissionService {
	return &permissionService{
		store: store.NewPermissionStore(db),
	}
}
func (s permissionService) NewPermission(r *msg.ReqPermissionCreate) error {
	p := &model.Permission{
		Name: r.PermissionName,
	}
	return s.store.InsertPermission(p)
}
func (s permissionService) ListPermission(r *msg.ReqPermissionFilter) (*msg.RspPermList, error) {
	var rsp msg.RspPermList
	perms, count, err := s.store.ListPermission(r)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &rsp, nil
	}

	rsp = msg.RspPermList{
		List:  (FunctionalPermissionMap(perms, buildPermRsp)).([]*msg.RspPermListItem),
		Total: count,
	}
	return &rsp, nil
}

func (s permissionService) FindPermission(pid int) (*model.Permission, error) {
	return s.store.FindPermission(pid)
}

// DeletePermission 删除权限
func (s permissionService) DeletePermission(pid int) error {
	return s.store.DeletePermission(pid)
}
