package model

import "time"

//LoginLog 记录下token日志，可以用来限制登录端
type LoginLog struct {
	Date   time.Time `json:"date"`
	UserID int       `json:"userId"`
	Token  string    `json:"token"`
}

func (e *LoginLog) TableName() string {
	return "login_log"
}

//ExitLog token黑名单 阻止退出后token任然有效
type ExitLog struct {
	Date   time.Time `json:"date"`
	UserID int       `json:"userId"`
	Token  string    `json:"token"`
}

func (e *ExitLog) TableName() string {
	return "exit_log"
}