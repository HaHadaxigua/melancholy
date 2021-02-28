package service

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/common/oss"
	"github.com/HaHadaxigua/melancholy/internal/file/consts"
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	"gorm.io/gorm"
)

var FileSvc FileService

type FileService interface {
	ListFileSpace(uid int) (*msg.RspFolderList, error)
	CreateFolder(req *msg.ReqFolderCreate) error
	UpdateFolder(req *msg.ReqFolderUpdate) error
	DeleteFolder(folderID string, userID int) error

	UploadFile(req *msg.ReqFileUpload) error

	// test
	CreateFile(req *msg.ReqFileCreate) error
	DeleteFile(fileID string, userID int) error
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
	folders, err := s.store.FoldersList(uid, consts.RootFileID)
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
		_folder, err := s.store.FolderFind(req.ParentID, req.UserID)
		if err != nil {
			return err
		}
		_filteredFolders := FunctionalFolderFilter(_folder.Subs, func(r *model.Folder) bool {
			if r.Name == req.FolderName {
				return true
			}
			return false
		})
		if len(_filteredFolders) >= 1 {
			return msg.ErrFileHasExisted
		}

		return s.store.FolderAppend(req.ParentID, folder)
	}

	if err := s.store.FolderCreate(&model.Folder{
		ID:      consts.RootFileID,
		OwnerID: req.UserID,
		Name:    consts.RootFileID,
	}); err != nil {
		return err
	}
	if err := s.store.FolderAppend(consts.RootFileID, folder); err != nil {
		return err
	}
	return nil
}

func (s fileService) UpdateFolder(req *msg.ReqFolderUpdate) error {
	if !tools.VerifyFileName(req.NewName) {
		return msg.ErrBadFilename
	}
	return s.store.FolderUpdate(req)
}

func (s fileService) DeleteFolder(folderID string, userID int) error {
	return s.store.FolderDelete(folderID, userID)
}

func (s fileService) CreateFile(req *msg.ReqFileCreate) error {
	_, err := s.store.FolderFind(req.ParentID, req.UserID)
	if err != nil {
		return err
	}
	_files, err := s.store.FileList(req.ParentID, req.UserID)
	if err != nil {
		return err
	}
	files := FunctionalFileFilter(_files, func(f *model.File) bool {
		if f.Name == req.FileName {
			return true
		}
		return false
	})
	if len(files) > 0 {
		return response.BadRequest
	}

	fid, err := tools.SnowflakeId()
	if err != nil {
		return err
	}
	file := &model.File{
		ID:       fid,
		Name:     req.FileName,
		ParentID: req.ParentID,
		MD5:      "todo:生成md5",
	}
	return s.store.FileCreate(file)
}

func (s fileService) DeleteFile(fileID string, userID int) error {
	_file, err := s.store.FileFind(fileID)
	if err != nil {
		return err
	}
	folder, err := s.store.FolderFind(_file.ParentID, userID)
	if err != nil {
		return err
	}
	if folder.OwnerID != userID {
		return response.BadRequest
	}
	return s.store.FileDelete(fileID, folder.ID)
}

func (s fileService) UploadFile(req *msg.ReqFileUpload) error {
	fmt.Println(req.FileHeader.Filename)

	oss.UploadFile(req.FileHeader.Filename, nil)

	return nil
}
