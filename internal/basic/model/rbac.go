/******
** @date : 2/9/2021 12:46 AM
** @author : zrx
** @description:
******/
package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" 	gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Status   int    `json:"status"`

	Roles []*Role `json:"roles" gorm:"many2many:user_role"`
	gorm.Model
}

func (u User) TableName() string {
	return "users"
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Users       []*User       `json:"users" gorm:"many2many:user_role"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:role_permission"`
	gorm.Model
}

func (r Role) TableName() string {
	return "roles"
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p Permission) TableName() string {
	return "permissions"
}
