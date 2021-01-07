package v1

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/pkg/store"
)

var ExitLogService IExitLogService

type IExitLogService interface {
	SaveExitLog(el *ent.ExitLog) error
}

type exitLogService struct {
	exitLogStore IExitLogService
}

func NewExitLogService() *exitLogService {
	return &exitLogService{
		exitLogStore: store.ExitLogStore,
	}
}

func (es *exitLogService) SaveExitLog(el *ent.ExitLog) error {
	return es.SaveExitLog(el)
}
