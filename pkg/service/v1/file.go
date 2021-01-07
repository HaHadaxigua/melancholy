package v1

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
)


// VerifyReq 验证请求合法性
/**
	To verify req is legal

 */
func VerifyReq(r *msg.FolderRequest) bool {
	if r.Creator <= 0 || r.Name == "" || r.Name == " " || r.ParentId < 0 {
		return false
	}
	return true
}

