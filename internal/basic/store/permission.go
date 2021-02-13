/******
** @date : 2/11/2021 2:01 PM
** @author : zrx
** @description:
******/
package store

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"gorm.io/gorm"
)

type PermissionStore interface {
	InsertPermission(p *model.Permission) error
	ListPermission(r *msg.ReqPermissionFilter) ([]*model.Permission, int, error)
	FindPermission(pid int) (*model.Permission, error)

}

type permissionStore struct {
	db *gorm.DB
}

func NewPermissionStore(db *gorm.DB) *permissionStore {
	return &permissionStore{
		db: db,
	}
}

func (s permissionStore) InsertPermission(p *model.Permission) error {
	return s.db.Create(p).Error
}
func (s permissionStore) ListPermission(r *msg.ReqPermissionFilter) ([]*model.Permission, int, error) {
	var perms []*model.Permission
	query := s.db.Model(&model.Permission{})
	if r.Fuzzy != "" {
		query = query.Where("name like %?%", r.Fuzzy)
	}
	if r.Limit < 1 {
		r.Limit = 10
	}
	if r.Offset < 0 {
		r.Offset = 0
	}
	query = query.Offset(r.Offset).Limit(r.Limit)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Find(&perms).Error; err != nil {
		return nil, 0, err
	}
	return perms, int(total), nil
}
func (s permissionStore) FindPermission(pid int) (*model.Permission, error){
	var perm model.Permission
	if err := s.db.Model(&model.Permission{ID: pid}).Take(&perm).Error; err != nil {
		return nil,err
	}
	return &perm, nil
}
