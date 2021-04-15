package store

import (
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"gorm.io/gorm"
)

func buildBaseQuery(db *gorm.DB, req *msg.ReqFileSearch) *gorm.DB {
	if req.Start != nil {
		db = db.Where("start >= ?", req.Start)
	}
	if req.End != nil {
		db = db.Where("end <= ?", req.End)
	}
	if req.Limit < -1 || req.Limit > 100 {
		db = db.Limit(50)
	}
	if req.Offset < -1 {
		db = db.Offset(-1)
	}
	return db
}
