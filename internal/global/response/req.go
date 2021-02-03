package response

import "time"

type LoginReq struct {
	Password string `json:"password"`
	Email    string `json:"email"`
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
