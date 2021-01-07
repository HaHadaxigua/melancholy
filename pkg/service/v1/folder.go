package v1

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

var FolderService IFolderService

type IFolderService interface {
	CreateFolder(r *msg.FolderRequest) error
	ListCurrentFolder(uid, pid int) ([]*ent.Folder, error)
	genPath(pid int, name string) (string, error)
}

type folderService struct {
	folderStore store.IFolderStore
}

func NewFolderService() *folderService{
	return &folderService{
		folderStore: store.FolderStore,
	}
}

func NewFolder(authorID, pid int, name string) (*ent.Folder, error){
	path, err := FolderService.genPath(pid, name)
	if err != nil {
		return nil, err
	}
	return &ent.Folder{
		Author: authorID,
		Parent: pid,
		Name:   name,
		Path:   path,
	}, nil
}

// CreateFolder
func(fs *folderService) CreateFolder(r *msg.FolderRequest) error {
	if !VerifyReq(r) {
		return nil
	}
	f, err  := NewFolder(r.Creator, r.ParentId, r.Name)
	if err != nil {
		return err
	}
	err = fs.folderStore.CreateFolder(f)
	if err != nil {
		return err
	}
	return nil
}

func(fs *folderService) ListCurrentFolder(uid, pid int) ([]*ent.Folder, error) {
	res, err := fs.folderStore.GetFolderByUserID(uid, 0)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func(fs *folderService) ListRootFolder(uid int) (*ent.Folder, error) {
	res, err := fs.folderStore.GetRootFolder(uid)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// fixme: failed
func(fs *folderService) genPath(pid int, name string) (string, error) {
	curFolder, err := fs.folderStore.GetFolderByID(pid)
	if err != nil {
		return "", err
	}
	subFolders, err := fs.folderStore.GetSubFolders(pid)
	if err != nil {
		return "", err
	}
	for _, sub := range subFolders {
		if sub.Name == name {
			return "", msg.FileRepeatErr
		}
	}
	res := fmt.Sprintf("%s/%s", curFolder.Path, name)
	return res, nil
}
