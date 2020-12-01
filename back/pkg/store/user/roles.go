package user

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/role"
	"github.com/HaHadaxigua/melancholy/ent/user"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

// CreateRoles 创建用户
func CreateRole(name string) (*ent.Role, error) {
	client := store.GetClient()
	ctx := store.GetCtx()

	r, err := client.Role.Create().SetName(name).SetStatus("0").Save(ctx)
	if err != nil {
		return nil, err
	}

	return r, err
}

// GetAllRoles 获取所有的角色
func GetAllRoles() ([]*ent.Role, error) {
	client := store.GetClient()
	ctx := store.GetCtx()
	roles, err := client.Role.Query().Where(role.DeletedAtIsNil()).All(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// AddUserRoles 添加角色给用户
func AddUserRoles(roleID, userID int) error {
	client := store.GetClient()
	ctx := store.GetCtx()
	// todo: 判断角色是否真实存在
	_, err := client.Role.Query().Where(role.IDEQ(roleID)).Only(ctx)
	if err != nil {
		return err
	}
	_, err = client.User.UpdateOneID(userID).AddRoleIDs(roleID).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetRolesByUserID(uID int) ([]*ent.Role, error) {
	client := store.GetClient()
	ctx := store.GetCtx()

	roles, err := client.User.Query().Where(user.IDEQ(uID)).QueryRoles().All(ctx)
	if err != nil {
		return []*ent.Role{}, err
	}
	return roles, nil
}
