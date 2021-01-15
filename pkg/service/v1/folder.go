package v1

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/pkg/common"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

var FolderService IFolderService

type IFolderService interface {
	CreateFolder(r *msg.FolderRequest) error
	ListFolder(uid, pid int) ([]*ent.Folder, error)
}

type folderService struct {
	folderStore store.IFolderStore
}

func NewFolderService() *folderService {
	return &folderService{
		folderStore: store.FolderStore,
	}
}

func NewFolder(authorID, pid int, name string) (*ent.Folder, error) {
	return &ent.Folder{
		Owner:  authorID,
		Parent: pid,
		Name:   name,
	}, nil
}

// CreateFolder
func (fs *folderService) CreateFolder(r *msg.FolderRequest) error {
	if !common.VerifyFileName(r.Filename) {
		return msg.InvalidParamsErr
	}
	f, err := NewFolder(r.Creator, r.ParentId, r.Filename)
	if err != nil {
		return err
	}
	err = fs.folderStore.CreateFolder(f)
	if err != nil {
		return err
	}
	return nil
}

// fixme: generate tree struct
func (fs *folderService) ListFolder(uid, cid int) ([]*ent.Folder, error) {
	res, err := fs.folderStore.GetFolderByUserID(uid, 0)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// fixme:
func (fs *folderService) ListRootFolder(uid int) (*ent.Folder, error) {
	res, err := fs.folderStore.GetRootFolder(uid)
	if err != nil {
		return nil, err
	}
	return res, nil
}
