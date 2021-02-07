package service

import (
	"context"
	"fmt"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
)

var FolderService IFolderService

type IFolderService interface {
	Create(r *msg.ReqFolderCreate) error
	ListFolder(uid, pid int) ([]*ent.Folder, error)
}

type folderService struct {
	folderStore store.IFolderStore
}

// desc a file tree
type FileTreeNode struct {
	Name  string          `json:"name"`
	Val   *ent.Folder     `json:"val"`
	Child []*FileTreeNode `json:"child"`
}

func NewFileTreeNode(folder *ent.Folder, folders []*ent.Folder) (node *FileTreeNode) {
	node = &FileTreeNode{
		Name:  folder.Name,
		Val:   folder,
		Child: make([]*FileTreeNode, 0),
	}
	for _, f := range folders {
		node.Child = append(node.Child, &FileTreeNode{
			Name:  f.Name,
			Val:   f,
			Child: nil,
		})
	}
	return
}

func NewFolderService(client *ent.Client, ctx context.Context) *folderService {
	return &folderService{
		folderStore: store.NewFolderStore(client, ctx),
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
func (fs *folderService) Create(r *msg.ReqFolderCreate) error {
	if !tools.VerifyFileName(r.Filename) {
		return response.InvalidParamsErr
	}
	f, err := NewFolder(r.UserID, r.ParentID, r.Filename)
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
// 文件表需要以字符和用户id作为主键
func (fs *folderService) ListFolder(uid, cid int) ([]*ent.Folder, error) {
	folder, err := fs.folderStore.GetFolderByID(uid)
	if err != nil {
		return nil, err
	}
	folders, err := fs.folderStore.GetFolderByUserID(uid, cid)
	root := NewFileTreeNode(folder, folders)

	fmt.Println(root)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

// fixme:
func (fs *folderService) ListRootFolder(uid int) (*ent.Folder, error) {
	res, err := fs.folderStore.GetRootFolder(uid)
	if err != nil {
		return nil, err
	}
	return res, nil
}
