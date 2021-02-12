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
