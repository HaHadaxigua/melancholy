package msg

type RspFriendItem struct {
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}

type RspFriendList struct {
	List  []*RspFriendItem `json:"list"`
	Total int              `json:"total"`
}
