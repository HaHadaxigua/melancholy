/******
** @date : 2/11/2021 2:01 PM
** @author : zrx
** @description:
******/
package store

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"gorm.io/gorm"
)

type PermissionStore interface {
	InsertPermission(p *model.Permission) error
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
