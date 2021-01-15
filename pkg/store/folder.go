package store

import (
	"context"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/folder"
)

var FolderStore *folderStore

type IFolderStore interface {
	CreateFolder(r *ent.Folder) error
	GetFolderByID(id int) (*ent.Folder, error)
	GetSubFolders(pid int) ([]*ent.Folder, error)
	GetFolderByUserID(uid, pid int) ([]*ent.Folder, error)
	GetRootFolder(uid int) (*ent.Folder, error)
}

type folderStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewFolderStore(client *ent.Client, ctx context.Context) *folderStore {
	return &folderStore{
		client: client,
		ctx:    ctx,
	}
}

func (fs *folderStore) CreateFolder(r *ent.Folder) error {
	r, err := fs.client.Folder.Create().
		SetName(r.Name).
		SetOwner(r.Owner).
		SetParent(r.Parent).
		SetPath(r.Path).
		Save(fs.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (fs *folderStore) GetFolderByID(id int) (*ent.Folder, error) {
	f, err := fs.client.Folder.Query().Where(folder.IDEQ(id)).Only(fs.ctx)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs *folderStore) GetSubFolders(pid int) ([]*ent.Folder, error) {
	res, err := fs.client.Folder.Query().Where(folder.IDEQ(pid)).QueryC().All(fs.ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (fs *folderStore) GetFolderByUserID(uid, pid int) ([]*ent.Folder, error) {
	f, err := fs.client.Folder.Query().Where(folder.OwnerEQ(uid), folder.ParentEQ(pid)).All(fs.ctx)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs *folderStore) GetRootFolder(uid int) (*ent.Folder, error) {
	res, err := fs.client.Folder.Query().Where(folder.OwnerEQ(uid), folder.ParentEQ(0)).QueryC().Only(fs.ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
