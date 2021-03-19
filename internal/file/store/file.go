package store

import (
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"gorm.io/gorm"
)

type FileStore interface {
	List(folderID string, userID int) ([]*model.Folder, error)
	FolderFind(folderID string, ownerID int) (*model.Folder, error)
	GetFolder(folderID string, withSub bool)(*model.Folder, error)
	FolderCreate(folder *model.Folder) error
	FolderAppend(parentID string, folder *model.Folder) error
	FolderUpdate(req *msg.ReqFolderUpdate) error
	FolderDelete(folderID string, ownerID int) error

	FileFind(fileID string) (*model.File, error)
	FileList(parentID string) ([]*model.File, error)
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

// List 列出用户在文件夹下的所有文件和子文件夹
func (s fileStore) List(folderID string, userID int) ([]*model.Folder, error) {
	var folders []*model.Folder
	query := s.db.Model(&model.Folder{ID: folderID}).Preload("Files")
	if err := query.Where("owner_id = ?", userID).
		Association("Subs").Find(&folders); err != nil {
		return nil, err
	}
	return folders, nil
}

func (s fileStore) FolderFind(folderID string, ownerID int) (*model.Folder, error) {
	var folder model.Folder
	query1 := s.db
	if err := query1.Where("owner_id = ? and id = ?", ownerID, folderID).Find(&folder).Error; err != nil {
		return nil, err
	}

	var subFolders []*model.Folder
	subQuery := s.db.Table("folder_sub").Select("sub_id").Where("folder_id = ?", folderID)
	if err := s.db.Model(&model.Folder{}).Where("id in (?)", subQuery).Find(&subFolders).Error; err != nil{
		return nil, err
	}
	folder.Subs = subFolders
	return &folder, nil
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

func (s fileStore) FolderAppend(parentID string, folder *model.Folder) error {
	query := s.db.Model(&model.Folder{ID: parentID}).Association("Subs")
	if err := query.Append(folder); err != nil {
		return err
	}
	return nil
}

func (s fileStore) FolderUpdate(req *msg.ReqFolderUpdate) error {
	query := s.db.Model(&model.Folder{ID: req.FolderID}).Where("owner_id = ?", req.UserID)
	return query.Update("name", req.NewName).Error
}

func (s fileStore) FolderDelete(folderID string, ownerID int) error {
	query := s.db
	return query.Model(&model.Folder{ID: folderID, OwnerID: ownerID}).Association("subs").Clear()
}

// 列出一个文件夹中的所有文件
func (s fileStore) FileList(folderID string) ([]*model.File, error) {
	var files []*model.File
	query := s.db.Model(&model.File{})
	if err := query.Where("parent_id = ?", folderID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (s fileStore) FileFind(fileID string) (*model.File, error) {
	var file model.File
	if err := s.db.Model(&model.File{ID: fileID}).Take(&file).Error; err != nil {
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
