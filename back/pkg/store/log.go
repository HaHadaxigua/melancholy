package store

import (
	"github.com/HaHadaxigua/melancholy/pkg/model"
)

//SaveExitLog 记录退出log
func SaveExitLog(el *model.ExitLog) error {
	db := GetConn()
	if err := db.Create(el).Error; err != nil {
		return err
	}
	return nil
}

//FindExitLog 寻找退出日志
func FindExitLog(token string) (*model.ExitLog, error) {
	db := GetConn()
	el := &model.ExitLog{}
	res := db.Model(el).Where("token = ?", token).Scan(el)
	if res.Error != nil {
		return nil, res.Error
	} else if res.RowsAffected < 1 {
		return nil, nil
	} else {
		return el, nil
	}
}
