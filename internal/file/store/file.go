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
	GetFolder(folderID string, withSub bool)(*model.Folder, error)
	FolderCreate(folder *model.Folder) error
	// FolderAppend 给目标文件夹添加文件夹
	FolderAppend(folderID string, folder *model.Folder) error
	FolderUpdate(req *msg.ReqFolderUpdate) error
	FolderDelete(req *msg.ReqFolderDelete) error

	FileFind(fileID string, userID int) (*model.File, error)
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

func(s fileStore) GetFolder(folderID string, withSub bool)(*model.Folder, error) {
	var folder model.Folder
	query := s.db.Model(&model.Folder{ID: folderID})
	if withSub {
		query = query.Preload("Subs")
	}
	if err := query.Scan(&folder).Error; err != nil {
		return nil, err
	}
	return &folder, nil
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


// FileList 列出一个文件夹中的所有文件
func (s fileStore) FileList(folderID string, userID int) ([]*model.File, error) {
	var files []*model.File
	query := s.db.Model(&model.File{})
	if err := query.Where("parent_id = ? and owner_id = ?", folderID, userID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (s fileStore) FileFind(fileID string, userID int) (*model.File, error) {
	var file model.File
	if err := s.db.Model(&model.File{ID: fileID,OwnerID: userID}).Take(&file).Error; err != nil {
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
