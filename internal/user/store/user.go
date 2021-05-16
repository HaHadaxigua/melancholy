package store

import (
	"gorm.io/gorm"
)

type UserStore interface {
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{
		db,
	}
}
