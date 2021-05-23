package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/HaHadaxigua/melancholy/internal/user/msg"
	"github.com/HaHadaxigua/melancholy/internal/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupUserRouters(r gin.IRouter) {
	// secured
	secured := r.Group("/s", middleware.Auth)


	// 好友相关api
	friend := secured.Group("/friend")
	friend.POST("/list", getFriendList)

}

// getFriendList 获取用户好友列表
func getFriendList(c *gin.Context) {
	var req msg.ReqFriendList
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)

	rsp, err := service.UserSvc.GetFriendList(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}
