package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	fConst "github.com/HaHadaxigua/melancholy/internal/file/envir"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func SetupUserRouters(r gin.IRouter) {
	secured := r.Group("/user", middleware.Auth)
	// user's
	userGroup := secured.Group("/u")
	userGroup.POST("/setInfo", setInfo)
	userGroup.POST("/setAvatar", setAvatar)
	userGroup.GET("/refreshUserInfo", refreshUserInfo)
	// friend's
}

// setInfo 设置用户信息接口
func setInfo(c *gin.Context) {
	var req msg.ReqSetUserInfo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.UserID = c.GetInt(consts.UserID)
	if err := service.User.SetUserInfo(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

// setAvatar 设置用户头像
func setAvatar(c *gin.Context) {
	// 这种方式将文件读入了内存，可能会导致内存爆掉
	file, header, err := c.Request.FormFile(fConst.FileUpload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	var req msg.ReqUpdateAvatar

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	req.FileHeader = header
	req.Data = data
	req.UserID = c.GetInt(consts.UserID)

	if err := service.User.UpdateAvatar(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.OK)
}

// refreshUserInfo 刷新用户信息
func refreshUserInfo(c *gin.Context) {
	_user, ok := c.Get(consts.User)
	if !ok {
		c.JSON(http.StatusBadRequest, response.NewErr(nil))
		return
	}
	user := _user.(*model.User)
	token := c.GetString(consts.CurrentToken)
	rsp := &msg.RspLogin{
		User: &msg.UserInfo{
			ID:                user.ID,
			Username:          user.Username,
			Mobile:            user.Mobile,
			Email:             user.Email,
			Status:            user.Status,
			Avatar:            user.Avatar,
			CreatedAt:         user.CreatedAt,
			UpdatedAt:         user.UpdatedAt,
			OssEndPoint:       user.OssEndPoint,
			OssAccessKey:      user.OssAccessKey,
			OssAccessSecret:   user.OssAccessSecret,
			CloudAccessKey:    user.CloudAccessKey,
			CloudAccessSecret: user.CloudAccessSecret,
		},
		Token: token,
	}
	c.JSON(http.StatusOK, response.Ok(rsp))
}
