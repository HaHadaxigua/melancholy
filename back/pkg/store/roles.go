package store

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/role"
	"github.com/HaHadaxigua/melancholy/ent/user"
)

func CreateRole(name string) (*ent.Role, error) {
	client := GetClient()
	ctx := GetCtx()

	r, err := client.Role.Create().SetName(name).SetStatus("0").Save(ctx)
	if err != nil {
		return nil, err
	}

	return r, err
}

func ListRoles() ([]*ent.Role, error) {
	client := GetClient()
	ctx := GetCtx()
	roles, err := client.Role.Query().Where(role.DeletedAtIsNil()).All(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func ListRolesByUserID(uID int) ([]*ent.Role, error) {
	client := GetClient()
	ctx := GetCtx()

	roles, err := client.User.Query().Where(user.IDEQ(uID)).QueryRoles().All(ctx)
	if err != nil {
		return []*ent.Role{}, err
	}
	return roles, nil
}

// append role to user
func AppendRoleToUser(roleID, userID int) error {
	client := GetClient()
	ctx := GetCtx()
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