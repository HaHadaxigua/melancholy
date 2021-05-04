package store

import (
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"gorm.io/gorm"
)

func buildBaseQuery(db *gorm.DB, req interface{}) *gorm.DB {
	switch v := req.(type) {
	case *msg.ReqFileSearch:
		if v.Start != nil {
			db = db.Where("start >= ?", v.Start)
		}
		if v.End != nil {
			db = db.Where("end <= ?", v.End)
		}
		if v.Limit < -1 || v.Limit > 100 {
			db = db.Limit(50)
		}
		if v.Offset < -1 {
			db = db.Offset(-1)
		}
	case *msg.ReqFindFileByType:
		if v.Limit < -1 || v.Limit > 100 {
			db = db.Limit(50)
		}
		if v.Offset < -1 {
			db = db.Offset(-1)
		}
	}
	return db
}
