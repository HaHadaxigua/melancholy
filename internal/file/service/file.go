package service

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/file/consts"
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"gorm.io/gorm"
)

var FileSvc FileService

type FileService interface {
	UserSpace(uid int) (*msg.RspFolderList, error)
	FolderList(req *msg.ReqFolderListFilter) (*msg.RspFolderList, error)
	FolderCreate(req *msg.ReqFolderCreate) error
	FolderUpload(req *msg.ReqFolderUpdate) error
	FolderDelete(folderID string, userID int) error

	FileList(req *msg.ReqFileListFilter) (*msg.RspFileList, error)
	FileUpload(req *msg.ReqFileUpload) error
	FileCreate(req *msg.ReqFileCreate) error
	FileDelete(fileID string, userID int) error
}

type fileService struct {
	store store.FileStore
}

func NewFileService(conn *gorm.DB) *fileService {
	return &fileService{
		store.NewFolderStore(conn),
	}
}

// ListFileSpace 列出用户的根文件夹
func (s fileService) UserSpace(uid int) (*msg.RspFolderList, error) {
	folders, err := s.store.List(consts.RootFileID, uid)
	if err != nil {
		return nil, err
	}

	var rsp  msg.RspFolderList
	list := FunctionalFolder(folders, buildFolderItemRsp).([]*msg.RspFolderListItem)
	rsp.List = list
	return &rsp, nil
}

//  ListFolders 列出文件夹
func (s fileService) FolderList(req *msg.ReqFolderListFilter) (*msg.RspFolderList, error) {
	folders, err := s.store.List(req.FolderID, req.UserID)
	if err != nil {
		return nil, err
	}

	var rsp msg.RspFolderList
	list := FunctionalFolder(folders, buildFolderItemRsp).([]*msg.RspFolderListItem)
	rsp.List= list
	return &rsp, nil
}

// FolderCreate 创建文件夹
func (s fileService) FolderCreate(req *msg.ReqFolderCreate) error {
	fid, err := tools.SnowflakeId()
	if err != nil {
		return err
	}
	folder := &model.Folder{
		ID:      fid,
		Name:    req.FolderName,
		OwnerID: req.UserID,
	}

	// if req ParentID is not empty, create folder in this folder
	if req.ParentID != "" {
		_folder, err := s.store.GetFolder(req.ParentID, true)
		if err != nil {
			return err
		}
		if _folder.OwnerID != req.UserID {
			return msg.ErrTargetFolderNotExist
		}
		// is exist repeated folder
		_filteredFolders := FunctionalFolderFilter(_folder.Subs, func(r *model.Folder) bool {
			if r.Name == req.FolderName  {
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

func (s fileService) FolderUpload(req *msg.ReqFolderUpdate) error {
	if !tools.VerifyFileName(req.NewName) {
		return msg.ErrBadFilename
	}
	return s.store.FolderUpdate(req)
}

func (s fileService) FolderDelete(folderID string, userID int) error {
	return s.store.FolderDelete(folderID, userID)
}

func (s fileService) FileCreate(req *msg.ReqFileCreate) error {
	folder, err := s.store.FolderFind(req.ParentID, req.UserID)
	if err != nil {
		return err
	}
	if folder == nil {
		return gorm.ErrRecordNotFound
	}
	_files, err := s.store.FileList(req.ParentID)
	if err != nil {
		return err
	}
	// if existed repeated name file
	files := FunctionalFileFilter(_files, func(f *model.File) bool {
		if f.Name == req.FileName && f.Suffix == req.FileType{
			return true
		}
		return false
	})
	if len(files) > 0 {
		return msg.ErrFileHasExisted
	}

	fid, err := tools.SnowflakeId()
	if err != nil {
		return err
	}
	file := &model.File{
		ID:       fid,
		Name:     req.FileName,
		ParentID: req.ParentID, // create empty file don't need to generate file
	}
	return s.store.FileCreate(file)
}

func (s fileService) FileDelete(fileID string, userID int) error {
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

func (s fileService) FileUpload(req *msg.ReqFileUpload) error {
	fmt.Println(req.FileHeader.Filename)
	return nil
}

// ListFiles 列出指定文件夹中的文件
func(s fileService) FileList(req *msg.ReqFileListFilter) (*msg.RspFileList, error){
	folder, err := s.store.GetFolder(req.FolderID, false)
	if err != nil {
		return nil, err
	}
	if folder == nil {
		return nil, gorm.ErrRecordNotFound
	}

	var rsp msg.RspFileList
	files, err := s.store.FileList(req.FolderID)
	if err != nil {
		return nil, err
	}
	items := FunctionalFile(files, buildFileItemRsp).([]*msg.RspFileListItem)
	rsp.List = items
	return &rsp, nil
}
