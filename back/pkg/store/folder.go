package store

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/folder"
)

func CreateFolder(r *ent.Folder) error {
	client := GetClient()
	ctx := GetCtx()
	r, err := client.Folder.Create().
		SetName(r.Name).
		SetAuthor(r.Author).
		SetParent(r.Parent).
		SetPath(r.Path).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetFolderByID(id int) (*ent.Folder, error) {
	client := GetClient()
	ctx := GetCtx()

	f, err := client.Folder.Query().Where(folder.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetSubFolders(pid int) ([]*ent.Folder, error) {
	client := GetClient()
	ctx := GetCtx()

	res, err := client.Folder.Query().Where(folder.IDEQ(pid)).QueryC().All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetFolderByUserID(uid, pid int) ([]*ent.Folder, error) {
	client := GetClient()
	ctx := GetCtx()

	f, err := client.Folder.Query().Where(folder.AuthorEQ(uid), folder.ParentEQ(pid)).All(ctx)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetRootFolder(uid int)(*ent.Folder, error){
	client := GetClient()
	ctx := GetCtx()

	res, err := client.Folder.Query().Where(folder.AuthorEQ(uid), folder.ParentEQ(0)).QueryC().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}