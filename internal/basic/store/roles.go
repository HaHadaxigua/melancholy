package store

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"gorm.io/gorm"
)

type RoleStore interface {
	InsertRole(r *model.Role) error
	ListRoles(filter *msg.ReqRoleFilter, withPermission bool) ([]*model.Role, int, error)
	Delete(rid int) error
	FindRole(rid int) (*model.Role, error)
	GetRoleByID(rid int, withPerms bool) (*model.Role, error)
	PermissionManager(rid, pid, operation int) error
}

type roleStore struct {
	db *gorm.DB
}

func NewRoleStore(db *gorm.DB) *roleStore {
	return &roleStore{
		db: db,
	}
}

func (s roleStore) InsertRole(r *model.Role) error {
	return s.db.Model(&model.Role{}).Create(r).Error
}

func (s roleStore) ListRoles(filter *msg.ReqRoleFilter, withPermission bool) ([]*model.Role, int, error) {
	query := s.db.Model(&model.Role{})
	var ret []*model.Role
	if withPermission {
		query = query.Preload("Permissions")
	}

	if filter.Fuzzy != "" {
		query = query.Where("name like ?", "%"+filter.Fuzzy+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Find(&ret).Error; err != nil {
		return nil, 0, err
	}
	return ret, int(total), nil
}

func (s roleStore) Delete(rid int) error {
	return s.db.Delete(&model.Role{ID: rid}).Error
}
func (s roleStore) FindRole(rid int) (*model.Role, error) {
	var role model.Role
	if err := s.db.Model(&model.Role{ID: rid}).Take(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s roleStore) GetRoleByID(rid int, withPerms bool) (*model.Role, error) {
	var role model.Role
	query := s.db.Model(&model.Role{}).Where("id = ?", rid)
	if withPerms {
		query = query.Preload("Permissions")
	}
	if err := query.Take(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
func (s roleStore) PermissionManager(rid, pid, operation int) error {
	query := s.db.Model(&model.Role{ID: rid}).Association("Permissions")
	switch operation {
	case consts.AppendPermission:
		return query.Append(&model.Permission{ID: pid})
	case consts.RemovePermission:
		return query.Delete(&model.Permission{ID: pid})
	default:
		return nil
	}
}
