package store

import (
	"context"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/role"
	"github.com/HaHadaxigua/melancholy/ent/user"
)



type IRoleStore interface {
	CreateRole(name string) (*ent.Role, error)
	ListRoles() ([]*ent.Role, error)
	ListRolesByUserID(uID int) ([]*ent.Role, error)
	AppendRoleToUser(roleID, userID int) error
}

type roleStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewRoleStore(client *ent.Client, ctx context.Context) *roleStore {
	return &roleStore{
		client: client,
		ctx:    ctx,
	}
}

func (rs *roleStore) CreateRole(name string) (*ent.Role, error) {
	r, err := rs.client.Role.Create().SetName(name).SetStatus("0").Save(rs.ctx)
	if err != nil {
		return nil, err
	}

	return r, err
}

func (rs *roleStore) ListRoles() ([]*ent.Role, error) {
	roles, err := rs.client.Role.Query().Where(role.DeletedAtIsNil()).All(rs.ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (rs *roleStore) ListRolesByUserID(uID int) ([]*ent.Role, error) {
	roles, err := rs.client.User.Query().Where(user.IDEQ(uID)).QueryRoles().All(rs.ctx)
	if err != nil {
		return []*ent.Role{}, err
	}
	return roles, nil
}

// append role to user
func (rs *roleStore) AppendRoleToUser(roleID, userID int) error {
	_, err := rs.client.Role.Query().Where(role.IDEQ(roleID)).Only(rs.ctx)
	if err != nil {
		return err
	}
	_, err = rs.client.User.UpdateOneID(userID).AddRoleIDs(roleID).Save(rs.ctx)
	if err != nil {
		return err
	}
	return nil
}
