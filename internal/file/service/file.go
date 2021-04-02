package service

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/common/oss"
	"github.com/HaHadaxigua/melancholy/internal/file/consts"
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"gorm.io/gorm"
)

var FileSvc FileService

type FileService interface {
	UserRoot(uid int) (*msg.RspFolderList, error)                        // 列出用户的根目录
	FolderList(req *msg.ReqFolderListFilter) (*msg.RspFolderList, error) // 列出文件夹
	FolderCreate(req *msg.ReqFolderCreate) error                         // 创建文件夹
	FolderUpload(req *msg.ReqFolderUpdate) error                         // 上传文件夹
	FolderDelete(req *msg.ReqFolderDelete) error                         // 删除文件夹

	FileList(req *msg.ReqFileListFilter) (*msg.RspFileList, error) // 列出文件
	FileUpload(req *msg.ReqFileUpload) error                       // 上传文件
	FileCreate(req *msg.ReqFileCreate) error                       // 创建文件
	FileDelete(fileID string, userID int) error                    // 删除文件

	FileDownload(req *msg.ReqFileDownload) (*msg.RspFileDownload, error) // 下载文件,需要流式处理
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
func (s fileService) UserRoot(uid int) (*msg.RspFolderList, error) {
	folders, err := s.store.ListSubFolders(consts.RootFileID, uid)
	if err != nil {
		return nil, err
	}
	files, err := s.store.FileList(consts.RootFileID, uid)
	fileItems := FunctionalFile(files, buildFileItemRsp).([]*msg.RspFileListItem)
	var rsp msg.RspFolderList
	list := FunctionalFolder(folders, buildFolderItemRsp).([]*msg.RspFolderListItem)
	rsp.FolderItems = list
	rsp.FileItems = fileItems
	return &rsp, nil
}

//  ListFolders 列出文件夹
func (s fileService) FolderList(req *msg.ReqFolderListFilter) (*msg.RspFolderList, error) {
	folders, err := s.store.ListSubFolders(req.FolderID, req.UserID)
	if err != nil {
		return nil, err
	}

	var rsp msg.RspFolderList
	list := FunctionalFolder(folders, buildFolderItemRsp).([]*msg.RspFolderListItem)
	rsp.FolderItems = list
	return &rsp, nil
}

// FolderCreate 创建文件夹
func (s fileService) FolderCreate(req *msg.ReqFolderCreate) error {
	// 1. 生成文件夹ID
	fid, err := tools.SnowflakeId()
	if err != nil {
		return err
	}
	folder := &model.Folder{
		ID:      fid,
		Name:    req.FolderName,
		OwnerID: req.UserID,
	}

	// 2. 如果请求创建的父文件夹ID存在
	if req.ParentID != "" {
		// 获取父文件夹
		parentFolder, err := s.store.GetFolder(req.ParentID, true)
		if err != nil {
			return err
		}
		if parentFolder.OwnerID != req.UserID {
			return msg.ErrTargetFolderNotExist
		}
		// is exist repeated folder
		// 判断是否有重名的文件夹
		_filteredFolders := FunctionalFolderFilter(parentFolder.Subs, func(r *model.Folder) bool {
			if r.Name == req.FolderName {
				return true
			}
			return false
		})
		if len(_filteredFolders) > 0 { // 存在重名的文件夹
			return msg.ErrFileHasExisted
		}
		return s.store.FolderAppend(req.ParentID, folder) // 添加文件夹
	}

	// 3. 否则在用户根目录创建文件夹
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

func (s fileService) FolderDelete(req *msg.ReqFolderDelete) error {
	return s.store.FolderDelete(req)
}

func (s fileService) FileCreate(req *msg.ReqFileCreate) error {
	if !req.Verify() {
		return msg.ErrFileUnSupport
	}
	folder, err := s.store.GetFolder(req.ParentID, false)
	if err != nil {
		return err
	}
	if folder == nil {
		return gorm.ErrRecordNotFound
	}
	// 1. 列出当前文件夹下的文件
	_files, err := s.store.FileList(req.ParentID, req.UserID)
	if err != nil {
		return err
	}
	// if existed repeated name file
	files := FunctionalFileFilter(_files, func(f *model.File) bool {
		if f.Name == req.FileName && f.Suffix == req.FileType {
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
	_file, err := s.store.FileFind(fileID, userID)
	if err != nil {
		return err
	}
	folder, err := s.store.GetFolder(_file.ParentID, false)
	if err != nil {
		return err
	}
	if folder.OwnerID != userID {
		return response.BadRequest
	}
	return s.store.FileDelete(fileID, folder.ID)
}

// FileUpload todo 上传文件的处理
func (s fileService) FileUpload(req *msg.ReqFileUpload) error {
	createFileReq := &msg.ReqFileCreate{
		ParentID: req.ParentID,
		FileName: req.FileHeader.Filename,
		FileType: req.FileType,
		UserID:   req.UserID,
	}

	if err := s.FileCreate(createFileReq); err != nil {
		return err
	}

	if err := oss.AliyunOss.UploadBytes(fmt.Sprint(req.UserID), req.FileHeader.Filename, req.Data); err != nil {
		return err
	}

	return nil
}

// ListFiles 列出指定文件夹中的文件
func (s fileService) FileList(req *msg.ReqFileListFilter) (*msg.RspFileList, error) {
	folder, err := s.store.GetFolder(req.FolderID, false)
	if err != nil {
		return nil, err
	}
	if folder == nil {
		return nil, gorm.ErrRecordNotFound
	}

	var rsp msg.RspFileList
	files, err := s.store.FileList(req.FolderID, req.UserID)
	if err != nil {
		return nil, err
	}
	items := FunctionalFile(files, buildFileItemRsp).([]*msg.RspFileListItem)
	rsp.List = items
	return &rsp, nil
}

// DownloadFile 处理文件下载 fixme: 完成文件下载的部分
func (s fileService) FileDownload(req *msg.ReqFileDownload) (*msg.RspFileDownload, error) {
	logicFile, err := s.store.FileFind(req.FileID, req.UserID)
	if err != nil {
		return nil, err
	}
	buf, err := oss.AliyunOss.DownloadFileByStream(logicFile.BucketName, logicFile.ObjectName)
	if err != nil {
		return nil, err
	}
	ret := make([]byte, buf.Len())
	buf.Read(ret)
	var rsp msg.RspFileDownload
	rsp.Content = ret
	rsp.FileName = logicFile.Name
	return &rsp, nil
}
