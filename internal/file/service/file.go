package service

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/file/consts"
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"gorm.io/gorm"
)

var File FileService

type FileService interface {
	ListFileSpace(uid int) (*msg.RspFolderList, error)
	CreateFolder(req *msg.ReqFolderCreate) error
	UpdateFolder(req *msg.ReqFolderUpdate) error
	DeleteFolder(folderID string, userID int) error
}

type fileService struct {
	store store.FileStore
}

func NewFileService(conn *gorm.DB) *fileService {
	return &fileService{
		store.NewFolderStore(conn),
	}
}

/**
list folder's in root.
*/
func (s fileService) ListFileSpace(uid int) (*msg.RspFolderList, error) {
	folders, err := s.store.GetUserFolders(uid, consts.RootFileID)
	if err != nil {
		return nil, err
	}

	rsp := &msg.RspFolderList{}

	list := FunctionalFolderMap(folders, buildFolderRsp)
	switch list.(type) {
	case error:
		return nil, list.(error)
	case []*msg.RspFolderListItem:
		rsp.List = list.([]*msg.RspFolderListItem)
		return rsp, nil
	}
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
		_parentFolder, err := s.store.FindFolder(req.ParentID, true)
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

func (s fileService) UpdateFolder(req *msg.ReqFolderUpdate) error {
	if !tools.VerifyFileName(req.NewName) {
		return msg.ErrBadFilename
	}
	return s.store.UpdateFolder(req)
}

func (s fileService) DeleteFolder(folderID string, userID int) error {
	return s.store.DeleteFolder(folderID, userID)
}
