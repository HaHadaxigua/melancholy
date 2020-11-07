package model

import "time"

import "gorm.io/gorm"

// 基本模型 嵌入到自己的结构体中去
type Model struct {
	ID        int64          `json:"Id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"deleted_at default:-"` // 逻辑删除
}
