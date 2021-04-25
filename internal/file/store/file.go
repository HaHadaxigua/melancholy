package store

import (
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FileStore interface {
	// ListSubFolders 列出当前文件夹下的所有子文件夹
	ListSubFolders(folderID string, userID int) ([]*model.Folder, error)
	GetFolder(folderID string, userID int, withSub bool) (*model.Folder, error)
	// FolderFindByName 根据名称找出文件夹
	FolderFindByName(req *msg.ReqFileSearch) ([]*model.Folder, error)
	FolderCreate(folder *model.Folder) error
	// FolderAppend 给目标文件夹添加文件夹
	FolderAppend(folderID string, folder *model.Folder) error
	FolderUpdate(req *msg.ReqFolderUpdate) error
	FolderDelete(req *msg.ReqFolderDelete) error

	// FileSearch 根据给出的名字找出相应的文件夹或者是文件
	FileSearch(req *msg.ReqFileSearch) ([]*model.Folder, []*model.File, error)
	// FileFind 根据id查找文件
	FileFind(fileID string, userID int) (*model.File, error)
	// FileFindByName 根据名称找出文件
	FileFindByName(req *msg.ReqFileSearch) ([]*model.File, error)
	// FileList 列出一个文件夹下的所有文件
	FileList(parentID string, userID int) ([]*model.File, error)
	FileCreate(file *model.File) error
	FileDelete(fileID, parentID string) error
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
	query := s.db.Model(&model.Folder{ID: folderID}).Where("owner_id = ?", userID)
	if withSub {
		query = query.Preload("Subs")
	}
	if err := query.Scan(&folder).Error; err != nil {
		return nil, err
	}
	return &folder, nil
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
	return query.Select(clause.Associations).Delete(&model.Folder{ID: req.FolderID, OwnerID: req.UserID}).Error
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
