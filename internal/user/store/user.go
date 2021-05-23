package store

import (
	"github.com/HaHadaxigua/melancholy/internal/user/model"
	"github.com/HaHadaxigua/melancholy/internal/user/msg"
	"gorm.io/gorm"
)

type UserStore interface {
	FindOnesFriend(req *msg.ReqFriendList) ([]*model.Friend, int, error) // 通过给出用户id找出他的好友
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{
		db,
	}
}

// FindOnesFriend 通过给出用户id找出他的好友
func (store userStore) FindOnesFriend(req *msg.ReqFriendList) ([]*model.Friend, int, error) {
	var res []*model.Friend
	query := store.db.Model(&model.Friend{})
	var count int64
	query = query.Where("status = 1").Where(`to =　? or from = ? `, req.UserID, req.UserID)
	if req.Offset >= 1 {
		query.Offset(req.Offset)
	} else {
		query.Offset(-1)
	}
	if req.Limit >= 1 && req.Limit <= 20 {
		query.Limit(req.Limit)
	} else {
		query.Limit(20)
	}
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, int(count), nil
}

// AddFriend 添加好友
func (store userStore) AddFriend(req *msg.ReqAddFriend) error {
	query := store.db
	friend := &model.Friend{
		From:   req.UserID,
		To:     req.TargetUserID,
		Status: 0,
	}
	return query.Create(friend).Error
}

// SetFriendStatus 设置好友状态
func (store userStore) SetFriendStatus(friendID, status int) error {
	query := store.db.Model(&model.Friend{})
	if err := query.Where("id = ?", friendID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
