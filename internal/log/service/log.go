/******
** @date : 2/3/2021 12:39 AM
** @author : zrx
** @description:
******/
package service

var LogService logService

type ILogService interface {
}

type logService struct {
}
func NewLogService() *logService {
	return &logService{}
}

func (s logService) NewExitLog() error {
	return nil
}
