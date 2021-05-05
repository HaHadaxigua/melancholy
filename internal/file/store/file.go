package store

import (
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FileStore interface {
	GetFolder(folderID string, userID int, withSub bool) (*model.Folder, error)    // 根据文件id找出文件夹
	ListSubFolders(folderID string, userID int) ([]*model.Folder, error)           // 列出当前文件夹下的所有子文件夹
	FolderInclude(folderID string, userID int) (model.Folders, model.Files, error) // 当前文件夹中包含的内容
	FolderFindByName(req *msg.ReqFileSearch) ([]*model.Folder, error)              // 根据名称找出文件夹
	FolderCreate(folder *model.Folder) error                                       // 创建文件夹
	FolderAppend(folderID string, folder *model.Folder) error                      // 给目标文件夹添加文件夹
	FolderUpdate(req *msg.ReqFolderUpdate) error                                   // 更新文件夹
	FolderDelete(req *msg.ReqFolderDelete) error                                   // 删除文件夹
	FolderPatchDelete(req *msg.ReqFolderPatchDelete) error                         // 批量删除文件夹

	FileSearch(req *msg.ReqFileSearch) ([]*model.Folder, []*model.File, error) // 根据给出的名字找出相应的文件夹或者是文件
	FileFind(fileID string, userID int) (*model.File, error)                   // 根据id查找文件
	FileFindByName(req *msg.ReqFileSearch) ([]*model.File, error)              // 根据名称找出文件
	FileList(parentID string, userID int) ([]*model.File, error)               // 列出一个文件夹下的所有文件
	FileCreate(file *model.File) error                                         // 创建文件
	FileDelete(fileID, parentID string) error                                  // 删除文件
	FilePatchDelete(req *msg.ReqFilePatchDelete) error                         // 批量删除文件
	FindFileByType(req *msg.ReqFindFileByType) (model.Files, int, error)       // 根据文件类型寻找文件
	FindFileByHash(hash string) (*model.File, error)                           // 通过文件hash来查找文件

	CreateDocFile(docFile *model.DocFile) error       // 创建文本类型的文件
	GetDocFile(fileID string) (*model.DocFile, error) // 通过文件id来找出文稿结构
}

type fileStore struct {
	db *gorm.DB
}

func NewFolderStore(conn *gorm.DB) *fileStore {
	return &fileStore{
		db: conn,
	}
}

//  ListSubFolders 列出用户在给定文件夹下的所有子文件夹
func (s fileStore) ListSubFolders(folderID string, userID int) ([]*model.Folder, error) {
	var folders []*model.Folder
	query := s.db.Model(&model.Folder{ID: folderID})
	if err := query.
		Where("owner_id = ? and parent_id = ?", userID, folderID).
		Preload("Subs").
		Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}

// GetFolder 获取指定文件夹
func (s fileStore) GetFolder(folderID string, userID int, withSub bool) (*model.Folder, error) {
	var folder model.Folder
	query := s.db.Model(&model.Folder{}).Where(" id = ? and owner_id = ?", folderID, userID)
	if withSub {
		query = query.Preload("Subs")
	}
	if err := query.Take(&folder).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

// FolderInclude 当前文件夹中包含的文件和文件夹
func (s fileStore) FolderInclude(folderID string, userID int) (model.Folders, model.Files, error) {
	folders, err := s.ListSubFolders(folderID, userID)
	if err != nil {
		return nil, nil, err
	}
	files, err := s.FileList(folderID, userID) // 当前文件夹中包含的文件
	if err != nil {
		return nil, nil, err
	}
	return folders, files, nil
}

// FolderFindByName 模糊搜索文件夹
func (s fileStore) FolderFindByName(req *msg.ReqFileSearch) ([]*model.Folder, error) {
	var files []*model.Folder
	query := s.db.Model(&model.Folder{})
	query = buildBaseQuery(query, req)
	if err := query.Where("name like ?", "%"+req.Fuzzy+"%").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (s fileStore) FolderCreate(folder *model.Folder) error {
	return s.db.Create(folder).Error
}

func (s fileStore) FolderAppend(folderID string, folder *model.Folder) error {
	query := s.db.Model(&model.Folder{ID: folderID}).Association("Subs")
	if err := query.Append(folder); err != nil {
		return err
	}
	return nil
}

func (s fileStore) FolderUpdate(req *msg.ReqFolderUpdate) error {
	query := s.db.Model(&model.Folder{ID: req.FolderID}).Where("owner_id = ?", req.UserID)
	return query.Update("name", req.NewName).Error
}

func (s fileStore) FolderDelete(req *msg.ReqFolderDelete) error {
	query := s.db
	return query.
		Select(clause.Associations).
		Delete(&model.Folder{ID: req.FolderID, OwnerID: req.UserID}).
		Error
}

// FolderPatchDelete 批量删除文件夹
func (s fileStore) FolderPatchDelete(req *msg.ReqFolderPatchDelete) error {
	query := s.db
	return query.Model(&model.Folder{}).Where("id in ? and owner_id = ?", req.FolderIDs, req.UserID).Select(clause.Associations).Delete(&model.Folder{}).Error
}

// FileFindByName 模糊搜索文件
func (s fileStore) FileFindByName(req *msg.ReqFileSearch) ([]*model.File, error) {
	var files []*model.File
	query := s.db.Model(&model.File{})
	query = buildBaseQuery(query, req)
	if req.Fuzzy != "" {
		query = query.Where("name like ?", "%"+req.Fuzzy+"%")
	}
	if err := query.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// FileList 列出一个文件夹中的所有文件
func (s fileStore) FileList(folderID string, userID int) ([]*model.File, error) {
	var files []*model.File
	query := s.db.Model(&model.File{})
	if err := query.Where("parent_id = ? and owner_id = ?", folderID, userID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// FileSearch 通过名称搜索文件，可能搜索出文件夹，也可能搜索出文件
func (s fileStore) FileSearch(req *msg.ReqFileSearch) ([]*model.Folder, []*model.File, error) {
	folders, err := s.FolderFindByName(req)
	if err != nil {
		return nil, nil, err
	}
	files, err := s.FileFindByName(req)
	if err != nil {
		return nil, nil, err
	}
	return folders, files, nil
}

func (s fileStore) FileFind(fileID string, userID int) (*model.File, error) {
	var file model.File
	if err := s.db.Model(&model.File{ID: fileID, OwnerID: userID}).Take(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (s fileStore) FileCreate(file *model.File) error {
	return s.db.Create(file).Error
}

func (s fileStore) FileDelete(fileID, parentID string) error {
	query := s.db
	if err := query.Model(&model.Folder{ID: parentID}).Association("Files").Delete(&model.File{ID: fileID}); err != nil {
		return err
	}
	if err := query.Model(&model.File{}).Delete(&model.File{ID: fileID}).Error; err != nil {
		return err
	}
	return nil
}

func (s fileStore) FilePatchDelete(req *msg.ReqFilePatchDelete) error {
	return s.db.Model(&model.File{}).
		Where("id in ? and owner_id = ?", req.FileIDs, req.UserID).
		Select(clause.Associations).
		Delete(&model.File{}).
		Error
}

// FindFileByType 根据文件类型寻找文件
func (s fileStore) FindFileByType(req *msg.ReqFindFileByType) (model.Files, int, error) {
	query := s.db.Model(&model.File{})
	query = buildBaseQuery(query, req)
	var files []*model.File
	query = query.Where("owner_id = ? and ftype = ?", req.UserID, req.FileType)
	if err := query.Find(&files).Error; err != nil {
		return nil, 0, err
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return files, int(total), nil
}

// FindFileByHash 根据文件Hash来查找文件
func (s fileStore) FindFileByHash(hash string) (*model.File, error) {
	query := s.db.Model(&model.File{})
	var file model.File
	if err := query.Where("hash = ?", hash).Take(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

// CreateDocFile 创建文本类型文件
func (s fileStore) CreateDocFile(docFile *model.DocFile) error {
	return s.db.Create(docFile).Error
}

//  GetDocContent 根据文稿id找出文稿内容
func (s fileStore) GetDocFile(fileID string) (*model.DocFile, error) {
	var docFile model.DocFile
	query := s.db.Model(&model.DocFile{})
	if err := query.Where("id = ?", fileID).Take(&docFile).Error; err != nil {
		return nil, err
	}
	return &docFile, nil
}
