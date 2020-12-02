package store

import (
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/ent/exitlog"
)

//SaveExitLog 记录退出log
func SaveExitLog(el *ent.ExitLog) error {
	client := GetClient()
	ctx := GetCtx()

	_, err := client.ExitLog.Create().SetUserID(el.UserID).SetToken(el.Token).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

//FindExitLog 寻找退出日志
func FindExitLog(token string) (*ent.ExitLog, error) {
	client := GetClient()
	ctx := GetCtx()

	el, err := client.ExitLog.Query().Where(exitlog.TokenEQ(token)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return el, nil
}
