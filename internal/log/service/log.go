/******
** @date : 2/3/2021 12:39 AM
** @author : zrx
** @description:
******/
package service

import "github.com/HaHadaxigua/melancholy/ent"

var LogService logService

type ILogService interface {
	NewExitLog(el *ent.ExitLog) error
}

type logService struct {
}

func NewLogService() *logService {
	return &logService{}
}

func (s logService) NewExitLog(el *ent.ExitLog) error {
	return nil
}
