package model

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/user/msg"
	"time"
)

// Friend 强好友关系，添加好友需要对方同意
type Friend struct {
	ID        int       `json:"id"`        // 自增的friend关系ID
	From      int       `json:"from"`      // 来自A用户向B用户的好友申请
	To        int       `json:"to"`        //　B用户的ID
	Status    int       `json:"status"`    // 好友状态
	CreatedAt time.Time `json:"createdAt"` // 好友申请日期
	UpdatedAt time.Time `json:"updatedAt"` // 申请同意日期
	DeletedAt time.Time `json:"deletedAt"` // 删除会导致 不再存在好友关系
}

func (Friend) TableName() string {
	return "friends"
}

func (f Friend) ToRsp(user *model.User) *msg.RspFriendItem {
	return &msg.RspFriendItem{
		UserID:   user.ID,
		UserName: user.Username,
		Avatar:   user.Avatar,
	}
}

type Friends []*Friend

// ToRsp 构建返回体时 需要用户信息的map
func (friends Friends) ToRsp(mem map[int]*model.User, total int) *msg.RspFriendList {
	list := make([]*msg.RspFriendItem, len(friends))
	for i, e := range friends {
		list[i] = e.ToRsp(mem[e.To])
	}
	return &msg.RspFriendList{
		List:  list,
		Total: total,
	}
}

// GetFriendUserIDList 获取好友列表中的好友id
func (friends Friends) GetFriendUserIDList() []int {
	idList := make([]int, len(friends))
	for i := 0; i < len(friends); i++ {
		idList[i] = friends[i].To
	}
	return idList
}
