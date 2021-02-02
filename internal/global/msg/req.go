package msg

import "time"

type LoginReq struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}


type CreateFolderReq struct {
	Name string `json:"name"`
	ParentID int `json:"parentID"`
}

type FolderRequest struct {
	Creator  int    `json:"creator"`
	ParentId int    `json:"parentId"`
	Filename string `json:"filename"`
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
