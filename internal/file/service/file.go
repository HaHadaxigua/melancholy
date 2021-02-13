package service

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/file/consts"
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"gorm.io/gorm"
)

var File FileService

type FileService interface {
	ListUserFolders(uid int) (*msg.RspFileList, error)
	CreateFolder(req *msg.ReqFolderCreate) error
}

type fileService struct {
	store store.FileStore
}

func NewFileService(conn *gorm.DB) *fileService {
	return &fileService{
		store.NewFolderStore(conn),
	}
}

func (s fileService) ListUserFolders(uid int) (*msg.RspFileList, error) {
	folders, err := s.store.GetUserFolders(uid)
	if err != nil {
		return nil, err
	}

	fmt.Println(folders)

	return nil, nil
}

func (s fileService) CreateFolder(req *msg.ReqFolderCreate) error {
	fid, err := tools.SnowflakeId()
	if err != nil {
		return err
	}
	folder := &model.Folder{
		ID:      fid,
		Name:    req.FolderName,
		OwnerID: req.UserID,
	}

	if req.ParentID != "" {
		_parentFolder, err := s.store.FindFolder(req.ParentID)
		if err != nil {
			return err
		}
		_filteredFolders := FunctionalFolderFilter(_parentFolder.Subs, func(r *model.Folder) bool {
			if r.Name == req.FolderName {
				return true
			}
			return false
		})
		if len(_filteredFolders) >= 1 {
			return msg.ErrFileHasExisted
		}

		return s.store.AppendFolder(req.ParentID, folder)
	}

	if err := s.store.CreateFolder(&model.Folder{
		ID:      consts.RootFileID,
		OwnerID: req.UserID,
		Name:    consts.RootFileID,
	}); err != nil {
		return err
	}
	if err := s.store.AppendFolder(consts.RootFileID, folder); err != nil {
		return err
	}
	return nil
}
