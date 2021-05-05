package service

import (
	"errors"
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/common/oss"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/store"
	"github.com/HaHadaxigua/melancholy/internal/file/utils"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/HaHadaxigua/melancholy/persistence"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
)

var FileSvc FileService

type FileService interface {
	UserRoot(uid int) (*msg.RspFolderList, error)                              // 列出用户的根目录
	FolderList(req *msg.ReqFolderListFilter) (*msg.RspFolderList, error)       // 列出文件夹
	FolderGetInfo(req *msg.ReqFolderGetInfo) (*msg.RspFolderListItem, error)   // 列出单个文件夹
	FolderCreate(req *msg.ReqFolderCreate) error                               // 创建文件夹
	FolderUpload(req *msg.ReqFolderUpdate) error                               // 上传文件夹
	FolderDelete(req *msg.ReqFolderDelete) error                               // 删除文件夹
	FolderPatchDelete(req *msg.ReqFolderPatchDelete) error                     // 文件夹批量删除
	FolderRDelete(req *msg.ReqFolderDelete) error                              // 递归删除文件夹
	FolderRPatchDelete(req *msg.ReqFolderPatchDelete) error                    // 递归的批量删除文件夹
	FolderInclude(req *msg.ReqFolderInclude) (*msg.RspFileSearchResult, error) // 列出给定文件夹下包含的内容

	FileSearch(req *msg.ReqFileSearch) (*msg.RspFileSearchResult, error)       // 文件搜索
	FileList(req *msg.ReqFileListFilter) (*msg.RspFileList, error)             // 列出文件
	FileUpload(req *msg.ReqFileUpload) error                                   // 上传文件，小文件上传
	FileCreate(req *msg.ReqFileCreate) (*msg.RspFileListItem, error)           // 创建文件
	FileDelete(req *msg.ReqFileDelete) error                                   // 删除文件
	FilePatchDelete(req *msg.ReqFilePatchDelete) error                         // 文件的批量删除
	FileSimpleDownload(req *msg.ReqFileDownload) (*msg.RspFileDownload, error) // 处理简单文件下载

	// 处理文件分片上传
	FileMultiCheck(req *msg.ReqFileMultiCheck) (*msg.RspFileMultiCheck, error)          // 检查文件上传情况
	FileMultiUpload(req *msg.ReqFileMultiUpload) (*msg.RspFileMultiUpload, error)       // 文件分片的上传
	FileMultiMerge(req *msg.ReqFileMultiMerge) (*msg.RspFileMultiMerge, error)          // 请求将文件分片进行合并
	FileMultiDownload(req *msg.ReqFileMultiDownload) (*msg.RspFIleMultiDownload, error) // 文件的分片下载

	// 文件夹和文件的整合方法
	DeleteInIntegration(req *msg.ReqDeleteInIntegration) error // 一个方法来处理文件夹和文件的删除方法

	// 针对特定类型的方法
	FindFileByType(req *msg.ReqFindFileByType) (*msg.RspFindFileByType, error) // 寻找当前用户的图片文件
	CreateDoc(req *msg.ReqDocFile) (*msg.RspCreateDocFile, error)              // 创建文本文件
	GetDocContent(req *msg.ReqDocFile) (string, error)                         // 获取文稿文件内容

	IsHashExisted(hash string) (bool, *model.File, error) // 文件hash是否已经存在
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
	folders, err := s.store.ListSubFolders(envir.RootFileID, uid)
	if err != nil {
		return nil, err
	}
	files, err := s.store.FileList(envir.RootFileID, uid)
	fileItems := FunctionalFile(files, buildFileItemRsp).([]*msg.RspFileListItem)
	var rsp msg.RspFolderList
	list := FunctionalFolder(folders, buildFolderItemRsp).([]*msg.RspFolderListItem)
	rsp.FolderItems = list
	rsp.FileItems = fileItems
	rsp.FolderTotal = len(list)
	rsp.FileTotal = len(fileItems)
	rsp.Total = rsp.FolderTotal + rsp.FileTotal
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

// FolderGetInfo 获取文件夹信息
func (s fileService) FolderGetInfo(req *msg.ReqFolderGetInfo) (*msg.RspFolderListItem, error) {
	folder, err := s.store.GetFolder(req.FolderID, req.UserID, false)
	if err != nil {
		return nil, err
	}
	return folder.ToFolderListItem(), nil
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
		parentFolder, err := s.store.GetFolder(req.ParentID, req.UserID, true)
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
		ID:      envir.RootFileID,
		OwnerID: req.UserID,
		Name:    envir.RootFileID,
	}); err != nil {
		return err
	}
	if err := s.store.FolderAppend(envir.RootFileID, folder); err != nil {
		return err
	}
	return nil
}

// folderCreate 用于内部的创建文件夹 没有错误判断, 用于创建特殊的文件夹
func (s fileService) folderCreate(req *msg.ReqFolderCreate) error {
	folder := &model.Folder{
		ID:       envir.DocFolderID,
		Name:     req.FolderName,
		ParentID: req.ParentID,
		OwnerID:  req.UserID,
	}
	return s.store.FolderCreate(folder)
}

func (s fileService) FolderUpload(req *msg.ReqFolderUpdate) error {
	if !tools.VerifyFileName(req.NewName) {
		return msg.ErrBadFilename
	}
	return s.store.FolderUpdate(req)
}

// FolderDelete 删除文件夹 需要删除子文件夹
func (s fileService) FolderDelete(req *msg.ReqFolderDelete) error {
	// todo: 递归的找出当前文件夹的子文件夹
	if req.FolderID == envir.RootFileID {
		// 禁止删除根目录
		return nil
	}
	var (
		needToDeleteFolderID []string
		needToDeleteFileID   []string
		recursive            func(folderID string) error // 找出当前文件夹中的所有文件和文件夹
	)
	recursive = func(folderID string) error {
		folders, files, err := s.store.FolderInclude(req.FolderID, req.UserID)
		if err != nil {
			return err
		}
		folderIDs := folders.GetIDs()
		needToDeleteFolderID = append(needToDeleteFolderID, folderIDs...)
		needToDeleteFileID = append(needToDeleteFileID, files.GetIDs()...)
		for _, e := range folderIDs {
			recursive(e)
		}
		return nil
	}

	recursive(req.FolderID)

	return s.store.FolderDelete(req)
}

// FolderRDelete 递归删除文件夹
func (s fileService) FolderRDelete(req *msg.ReqFolderDelete) error {
	// todo: 递归的找出当前文件夹的子文件夹
	if req.FolderID == envir.RootFileID {
		// 禁止删除根目录
		return nil
	}
	var (
		needToDeleteFolderID []string
		needToDeleteFileID   []string
		recursive            func(folderID string) error // 找出当前文件夹中的所有文件和文件夹
	)
	needToDeleteFolderID = append(needToDeleteFolderID, req.FolderID)
	recursive = func(folderID string) error {
		folders, files, err := s.store.FolderInclude(folderID, req.UserID)
		if err != nil {
			return err
		}
		folderIDs := folders.GetIDs()
		needToDeleteFolderID = append(needToDeleteFolderID, folderIDs...)
		needToDeleteFileID = append(needToDeleteFileID, files.GetIDs()...)
		for _, e := range folderIDs {
			recursive(e)
		}
		return nil
	}

	err := recursive(req.FolderID)
	if err != nil {
		return err
	}

	if len(needToDeleteFolderID) != 0 {
		r := &msg.ReqFolderPatchDelete{
			FolderIDs: needToDeleteFolderID,
			UserID:    req.UserID,
		}
		if err := s.store.FolderPatchDelete(r); err != nil {
			return err
		}
	}

	if len(needToDeleteFileID) != 0 {
		r := &msg.ReqFilePatchDelete{
			FileIDs: needToDeleteFileID,
			UserID:  req.UserID,
		}
		if err := s.store.FilePatchDelete(r); err != nil {
			return err
		}
	}
	return nil
}

// FolderPatchDelete 文件夹批量删除
func (s fileService) FolderPatchDelete(req *msg.ReqFolderPatchDelete) error {
	return s.store.FolderPatchDelete(req)
}

// FolderRPatchDelete 批量删除文件夹包括其子内容
func (s fileService) FolderRPatchDelete(req *msg.ReqFolderPatchDelete) error {
	// todo 优化协程的并发处理
	for _, e := range req.FolderIDs {
		req := &msg.ReqFolderDelete{
			FolderID: e,
			UserID:   req.UserID,
		}
		if err := s.FolderRDelete(req); err != nil {
			return err
		}
	}
	return nil
}

//  FolderInclude 列出给定文件夹下包含的内容, 包括文件和文件夹 todo: 整合UserRoot方法
func (s fileService) FolderInclude(req *msg.ReqFolderInclude) (*msg.RspFileSearchResult, error) {
	if req.FolderID == "" || req.FolderID == " " {
		req.FolderID = envir.RootFileID
	}
	curFolder, err := s.store.GetFolder(req.FolderID, req.UserID, false)
	if err != nil {
		return nil, err
	}
	folders, err := s.store.ListSubFolders(req.FolderID, req.UserID)
	if err != nil {
		return nil, err
	}
	files, err := s.store.FileList(req.FolderID, req.UserID)
	if err != nil {
		return nil, err
	}
	rsp := buildFileSearchResult(folders, files)
	rsp.ParentID = curFolder.ParentID // 当前文件夹的父文件夹

	msg.BuildRspFileSearchResultSort(req.SortedBy, req.Descending) // 处理排序
	sort.Sort(rsp)

	return rsp, nil
}

func (s fileService) FileCreate(req *msg.ReqFileCreate) (*msg.RspFileListItem, error) {
	if !req.Verify() {
		return nil, msg.ErrFileUnSupport
	}
	folder, err := s.store.GetFolder(req.ParentID, req.UserID, false)
	if err != nil {
		return nil, err
	}
	if folder == nil {
		return nil, gorm.ErrRecordNotFound
	}
	// 1. 列出当前文件夹下的文件
	_files, err := s.store.FileList(req.ParentID, req.UserID)
	if err != nil {
		return nil, err
	}
	// if existed repeated name file
	files := FunctionalFileFilter(_files, func(f *model.File) bool {
		if f.Name == req.FileName && f.Suffix == req.FileType {
			return true
		}
		return false
	})
	if len(files) > 0 {
		return nil, msg.ErrFileHasExisted
	}

	fid, err := tools.SnowflakeId()
	if err != nil {
		return nil, err
	}
	var suffix int
	if v, ok := envir.MapFileTypeToID[path.Ext(req.FileName)]; !ok {
		return nil, msg.ErrFileUnSupport
	} else {
		suffix = v
	}

	file := &model.File{
		ID:         fid,
		OwnerID:    req.UserID,
		ParentID:   req.ParentID, // create empty file don't need to generate file
		Name:       req.FileName,
		Suffix:     suffix,
		Size:       req.Size,
		Address:    req.Address,
		BucketName: req.BucketName,
		Ftype:      envir.GetFileType(suffix),
		Hash:       req.Hash,
	}
	if err = s.store.FileCreate(file); err != nil {
		return nil, err
	}
	return file.ToFileListItem(), nil
}

func (s fileService) FileDelete(req *msg.ReqFileDelete) error {
	_file, err := s.store.FileFind(req.FileID, req.UserID)
	if err != nil {
		return err
	}
	folder, err := s.store.GetFolder(_file.ParentID, req.UserID, false)
	if err != nil {
		return err
	}
	if folder.OwnerID != req.UserID {
		return response.BadRequest
	}
	return s.store.FileDelete(req.FileID, folder.ID)
}

func (s fileService) FilePatchDelete(req *msg.ReqFilePatchDelete) error {
	return s.store.FilePatchDelete(req)
}

// FileUpload todo 上传文件的处理
func (s fileService) FileUpload(req *msg.ReqFileUpload) error {
	hash := utils.CalcBytesHashInSHA(req.Data)
	var (
		bucketName string
		ossAddress string
	)
	location := utils.GetOSType()
	isExist, originFile, err := s.IsHashExisted(hash)
	if err != nil {
		return err
	}
	if isExist {
		bucketName = originFile.BucketName
		ossAddress = originFile.Address
	} else {
		bucketName, ossAddress = oss.BuildBucketNameAndAddress(req.UserID, req.FileHeader.Filename)
	}

	createFileReq := &msg.ReqFileCreate{
		ParentID:   req.ParentID,
		FileName:   req.FileHeader.Filename,
		FileType:   req.FileType,
		UserID:     req.UserID,
		Size:       int(req.FileHeader.Size),
		BucketName: bucketName,
		Address:    ossAddress,
		Hash:       hash,
	}

	if _, err := s.FileCreate(createFileReq); err != nil {
		return err
	}

	if !isExist {
		pm := persistence.NewResourceManager()
		go func() {
			err = pm.SaveSimpleFile(createFileReq.FileName, location, req.Data)
		}()

		go func() {
			err = oss.AliyunOss.UploadBytes(bucketName, req.FileHeader.Filename, req.Data)
		}()
	}

	if err != nil {
		return err
	}
	return nil
}

// FileSearch 文件的模糊搜索
func (s fileService) FileSearch(req *msg.ReqFileSearch) (*msg.RspFileSearchResult, error) {
	folders, files, err := s.store.FileSearch(req)
	if err != nil {
		return nil, err
	}
	res := buildFileSearchResult(folders, files)
	return res, nil
}

// ListFiles 列出指定文件夹中的文件
func (s fileService) FileList(req *msg.ReqFileListFilter) (*msg.RspFileList, error) {
	folder, err := s.store.GetFolder(req.FolderID, req.UserID, false)
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

// FileSimpleDownload 处理简单文件下载
func (s fileService) FileSimpleDownload(req *msg.ReqFileDownload) (*msg.RspFileDownload, error) {
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

// DeleteInIntegration 同时有文件和文件夹需要去删除
func (s fileService) DeleteInIntegration(req *msg.ReqDeleteInIntegration) error {
	var err error
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		if len(req.FileIDs) != 0 {
			r := &msg.ReqFilePatchDelete{
				FileIDs: req.FileIDs,
				UserID:  req.UserID,
			}
			err = s.FilePatchDelete(r)
		}
		wg.Done()
	}()

	go func() {
		if len(req.FolderIDs) != 0 {
			r := &msg.ReqFolderPatchDelete{
				FolderIDs: req.FolderIDs,
				UserID:    req.UserID,
			}
			err = s.FolderPatchDelete(r)
		}
		wg.Done()
	}()

	wg.Wait()

	return err
}

// FileMultiCheck 检查文件分分片的上传情况，返回一个分片列表
func (s fileService) FileMultiCheck(req *msg.ReqFileMultiCheck) (*msg.RspFileMultiCheck, error) {
	var rsp msg.RspFileMultiCheck
	hashPath := fmt.Sprintf("%s%s", conf.C.Application.TmpFile, req.Hash) // 以hash为文件名的文件夹
	if !utils.PathExists(hashPath) {                                      // 如果不存在该文件夹
		rsp.ChunkList = make([]string, 0)
		return &rsp, nil
	}
	var chunkList []string
	var state int
	files, err := ioutil.ReadDir(hashPath)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		fileName := f.Name()
		chunkList = append(chunkList, fileName)
		fileBaseName := strings.Split(fileName, ".")[0]
		if fileBaseName == req.Filename { // 如果存在一个文件名一致的文件，则说明已经上传完成了,也就不需要获取完整的分片列表了
			state = 1
			break
		}
	}
	rsp.ChunkList = chunkList
	rsp.State = state
	return &rsp, nil
}

// FileMultiUpload 分片文件的上传处理,当全部分片上传完成后，会进行合并,这里上传之后只是上传到了本地服务器中，还需要上传到对象存储中
func (s fileService) FileMultiUpload(req *msg.ReqFileMultiUpload) (*msg.RspFileMultiUpload, error) {
	hashPath := fmt.Sprintf("%s%s", conf.C.Application.TmpFile, req.Hash) // 以hash为文件名的文件夹
	// 不存在文件夹则进行创建
	if !utils.PathExists(hashPath) {
		os.Mkdir(hashPath, os.ModePerm)
	}
	// 将文件保存到对应路径
	if err := req.C.SaveUploadedFile(req.FileHeader, utils.GetMultiFileName(hashPath, req.ChunkID)); err != nil {
		return nil, err
	}
	// 读取文件夹下的文件分片，告诉前端已经当前文件的上传情况
	var chunkList []string
	files, err := ioutil.ReadDir(hashPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		filename := file.Name()
		if _, ok := envir.ExcludeFiles[filename]; ok {
			continue
		}
		chunkList = append(chunkList, filename)
	}
	var rsp msg.RspFileMultiUpload
	rsp.ChunkList = chunkList
	return &rsp, nil
}

// FileMultiMerge 文件合并操作
func (s fileService) FileMultiMerge(req *msg.ReqFileMultiMerge) (*msg.RspFileMultiMerge, error) {
	hashPath := utils.GetMultiFilePath(req.Hash)
	// 不存在文件夹说明请求错误
	if !utils.PathExists(hashPath) {
		return nil, msg.ErrMergeFileFailed
	}
	var (
		err error
		rsp msg.RspFileMultiMerge
	)
	rsp.Done = make(chan struct{}, 0)
	// 这里开启的协程是与此函数平级的，并不会因为函数的退出而退出
	go func() {
		// 进行文件合并
		err = utils.MergeFiles(hashPath, req.Filename)
		rsp.Result = err
		rsp.Done <- struct{}{}
	}()
	return &rsp, nil
}

// FileMultiDownload 处理文件的分片下载
func (s fileService) FileMultiDownload(req *msg.ReqFileMultiDownload) (*msg.RspFIleMultiDownload, error) {
	// 1。 判断本地有无下载好的文件
	// 2。 将文件进行分片的传输
	return nil, nil
}

// FindFileByType 根据文件类型查找文件
func (s fileService) FindFileByType(req *msg.ReqFindFileByType) (*msg.RspFindFileByType, error) {
	files, total, err := s.store.FindFileByType(req)
	if err != nil {
		return nil, err
	}
	var rsp msg.RspFindFileByType
	rsp.List = files.ToRspFindFileItemByType()
	rsp.Total = total
	return &rsp, nil
}

// CreateDoc 创建文本类型文件
func (s fileService) CreateDoc(req *msg.ReqDocFile) (*msg.RspCreateDocFile, error) {
	// 判断文档文件夹是否存在
	_, err := s.store.GetFolder(envir.DocFolderID, req.UserID, false)
	if err != nil {
		flag := errors.Is(err, gorm.ErrRecordNotFound)
		if flag {
			// 需要去创建文件夹
			folderReq := &msg.ReqFolderCreate{
				FolderName: envir.DocFolderID,
				ParentID:   envir.RootFileID,
				UserID:     req.UserID,
			}
			if err = s.folderCreate(folderReq); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 计算hash 判断文件是否已经上传过
	hash := utils.CalcStringHashInSHA(req.Content)
	isExist, originFile, err := s.IsHashExisted(hash)
	if err != nil {
		return nil, err
	}
	var (
		bucketName string
		ossAddress string
	)
	location := utils.GetOSType()
	if isExist {
		bucketName = originFile.BucketName
		ossAddress = originFile.Address
	} else {
		bucketName, ossAddress = oss.BuildBucketNameAndAddress(req.UserID, req.Name)
	}

	// 创建绑定的基本文件
	fileCreate := &msg.ReqFileCreate{
		ParentID: envir.DocFolderID,
		FileName: req.Name,
		FileType: envir.FileTypeTxtID,
		Size:     len(req.Content),
		Address:  ossAddress,
		UserID:   req.UserID,
		Hash:     hash,
	}

	createRsp, err := s.FileCreate(fileCreate)
	if err != nil {
		return nil, err
	}

	docFile := &model.DocFile{
		ID:      createRsp.ID,
		Content: req.Content,
	}
	// 保存记录
	if err = s.store.CreateDocFile(docFile); err != nil {
		return nil, err
	}

	if !isExist {
		// 将文件持久化
		pm := persistence.NewResourceManager()
		go func() {
			pm.SaveSimpleFile(req.Name, location, []byte(req.Content))
		}()

		// 将文件上传到oss
		go func() {
			oss.AliyunOss.UploadString(bucketName, req.Name, req.Content)
		}()
	}

	rsp := &msg.RspCreateDocFile{
		ID:       createRsp.ID,
		Name:     req.Name,
		Size:     createRsp.Size,
		Address:  createRsp.Address,
		Content:  req.Content,
		CreateAt: createRsp.CreatedAt,
		UpdateAt: createRsp.UpdatedAt,
	}
	return rsp, nil
}

// GetDocContent 根据文件id 找出文件内容
func (s fileService) GetDocContent(req *msg.ReqDocFile) (string, error) {
	_, err := s.store.FileFind(req.ID, req.UserID)
	if err != nil {
		return "", err
	}
	docFile, err := s.store.GetDocFile(req.ID)
	if err != nil {
		return "", err
	}
	return docFile.Content, nil
}

// IsHashExisted 文件hash是否存在
func (s fileService) IsHashExisted(hash string) (bool, *model.File, error) {
	var isExisted bool
	file, err := s.store.FindFileByHash(hash)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, err
		}
		isExisted = false
	} else {
		isExisted = true
	}
	return isExisted, file, nil
}
