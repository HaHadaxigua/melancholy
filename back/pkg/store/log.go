package store

import (
	"context"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/exitlog"
)

var ExitLogStore IExitLogStore

type IExitLogStore interface {
	SaveExitLog(el *ent.ExitLog) error
	GetExitLog(token string) (*ent.ExitLog, error)
}

type exitLogStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewExitLogStore(client *ent.Client, ctx context.Context) *exitLogStore {
	return &exitLogStore{
		client: client,
		ctx:    ctx,
	}
}

func (es *exitLogStore) SaveExitLog(el *ent.ExitLog) error {
	_, err := es.client.ExitLog.Create().SetUserID(el.UserID).SetToken(el.Token).Save(es.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (es *exitLogStore) GetExitLog(token string) (*ent.ExitLog, error) {
	el, err := es.client.ExitLog.Query().Where(exitlog.TokenEQ(token)).Only(es.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return el, nil
}
