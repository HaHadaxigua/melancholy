package v1

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

func NewFolder(aid, pid int, name string) (*ent.Folder, error){
	path, err := genPath(pid, name)
	if err != nil {
		return nil, err
	}
	return &ent.Folder{
		Author: aid,
		Parent: pid,
		Name:   name,
		Path:   path,
	}, nil
}

// CreateFolder
func CreateFolder(r *msg.FolderRequest) error {
	if !VerifyReq(r) {
		return nil
	}
	f, err  := NewFolder(r.Creator, r.ParentId, r.Name)
	if err != nil {
		return err
	}
	err = store.CreateFolder(f)
	if err != nil {
		return err
	}
	return nil
}

func ListCurrentFolder(uid, pid int) ([]*ent.Folder, error) {
	res, err := store.GetFolderByUserID(uid, 0)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ListRootFolder(uid int) (*ent.Folder, error) {
	res, err := store.GetRootFolder(uid)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func genPath(pid int, name string) (string, error) {
	curFolder, err := store.GetFolderByID(pid)
	if err != nil {
		return "", err
	}
	subFolders, err := store.GetSubFolders(pid)
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
