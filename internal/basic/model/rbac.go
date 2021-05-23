/******
** @date : 2/9/2021 12:46 AM
** @author : zrx
** @description:
******/
package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       int    `json:"id" 	gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`

	OssEndPoint       string `form:"ossEndPoint" json:"ossEndPoint"`
	OssAccessKey      string `form:"ossAccessKey" json:"ossAccessKey"`
	OssAccessSecret   string `form:"ossAccessSecret" json:"ossAccessSecret"`
	CloudAccessKey    string `form:"cloudAccessKey" json:"cloudAccessKey"`
	CloudAccessSecret string `form:"cloudAccessSecret" json:"cloudAccessSecret"`

	Roles []*Role `json:"roles" gorm:"many2many:user_role"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (u User) TableName() string {
	return "users"
}

type Users []*User

// ToIDMap slice 转换为 id map
func (users Users) ToIDMap() map[int]*User {
	res := make(map[int]*User)
	for i := 0; i < len(users); i++ {
		res[users[i].ID] = users[i]
	}
	return res
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Users       []*User       `json:"users" gorm:"many2many:user_role"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:role_permission"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (r Role) TableName() string {
	return "roles"
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (p Permission) TableName() string {
	return "permissions"
}
