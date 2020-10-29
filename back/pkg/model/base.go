package model

import "time"

import "gorm.io/gorm"

// 基本模型 嵌入到自己的结构体中去
type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"` // 逻辑删除
}
