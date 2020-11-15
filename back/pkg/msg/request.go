package msg

import "time"

//DirRequest 文件夹请求
type FolderRequest struct {
	Creator  int64  `json:"creator"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
}

//NewDirRequest 构造请求
func NewDirRequest(c, p int64, n string) *FolderRequest {
	return &FolderRequest{
		Creator:  c,
		Name:     n,
		ParentId: p,
	}
}

//VideoRequest 视频请求
type VideoRequest struct {
	Fid       string
	FType     int
	FSize     int
	CreatedAt time.Time
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
