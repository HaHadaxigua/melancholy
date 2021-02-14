package store

import (
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FileStore interface {
	GetUserFolders(userID int, folderID string) ([]*model.Folder, error)
	FindFolder(folderID string, withSubs bool) (*model.Folder, error)
	CreateFolder(folder *model.Folder) error
	AppendFolder(parentID string, folder *model.Folder) error
	UpdateFolder(req *msg.ReqFolderUpdate) error

	ListFolders(folders []int, withSubFolders, withFiles bool) ([]*model.Folder, error)
	DeleteFolder(folderID string, ownerID int) error
}

type fileStore struct {
	db *gorm.DB
}

func NewFolderStore(conn *gorm.DB) *fileStore {
	return &fileStore{
		db: conn,
	}
}

/**
列出用户在某个文件夹下的所有文件和文件夹
*/
func (s fileStore) GetUserFolders(uid int, folderID string) ([]*model.Folder, error) {
	var folders []*model.Folder
	query := s.db.Model(&model.Folder{ID: folderID}).Preload("Files")
	if err := query.Select("folders.id, owner_id, name, created_at,updated_at,deleted_at").Where("owner_id = ?", uid).Association("Subs").Find(&folders); err != nil {
		return nil, err
	}
	return folders, nil
}

func (s fileStore) ListFolders(folders []int, withSubFolders, withFiles bool) ([]*model.Folder, error) {
	return nil, nil
}

func (s fileStore) FindFolder(folderID string, withSubs bool) (*model.Folder, error) {
	var folder model.Folder
	query := s.db.Model(&model.Folder{ID: folderID})
	if withSubs {
		query = query.Association("Subs").DB
	}
	if err := query.Take(&folder).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (s fileStore) CreateFolder(folder *model.Folder) error {
	return s.db.Create(folder).Error
}

func (s fileStore) AppendFolder(parentID string, folder *model.Folder) error {
	query := s.db.Model(&model.Folder{ID: parentID}).Association("Subs")
	if err := query.Append(folder); err != nil {
		return err
	}
	return nil
}

func (s fileStore) UpdateFolder(req *msg.ReqFolderUpdate) error {
	query := s.db.Model(&model.Folder{ID: req.FolderID}).Where("owner_id = ?", req.UserID)
	return query.Update("name", req.NewName).Error
}

func (s fileStore) DeleteFolder(folderID string, ownerID int) error {
	query := s.db
	return query.Select(clause.Associations).Where("owner_id = ?", ownerID).Delete(&model.Folder{ID: folderID}).Error
}
