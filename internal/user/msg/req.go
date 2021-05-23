package msg

// ReqFriendList
type ReqFriendList struct {
	FriendName string `form:"friendName" json:"friendName"` // 可以根据好友名称搜索好友
	Offset     int    `form:"offset" json:"offset"`
	Limit      int    `form:"limit" json:"limit"`
	UserID     int
}

// ReqAddFriend 添加用户请求
type ReqAddFriend struct {
	TargetUserID int `json:"targetUserID"` // 需要添加的用户id

	UserID int
}
