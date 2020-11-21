package store

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/model/user"
)

//SaveExitLog 记录退出log
func SaveExitLog(el *user.ExitLog) error {
	db := GetConn()
	if err := db.Create(el).Error; err != nil {
		return err
	}
	return nil
}

//FindExitLog 寻找退出日志
func FindExitLog(token string) (*user.ExitLog, error) {

	fmt.Println("测试")


	db := GetConn()
	el := &user.ExitLog{}
	res := db.Model(el).Where("token = ?", token).Scan(el)
	if res.Error != nil {
		return nil, res.Error
	} else if res.RowsAffected < 1 {
		return nil, nil
	} else {
		return el, nil
	}
}
