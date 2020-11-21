package user

import (
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//logout 退出登录
func logout(c *gin.Context) {
	ah := middleware.AuthHeader{}

	if err := c.ShouldBindHeader(&ah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": msg.AuthAccessTokenIllegalErrorMsg,
		})
		return
	}

	userId := c.GetInt("user_id")
	// 写退出表

	exitLog := &model.ExitLog{
		Date:   time.Now(),
		UserID: userId,
		Token:  ah.AccessToken,
	}

	err := store.SaveExitLog(exitLog)
	if err != nil {
		e := msg.UserExitErr
		c.JSON(http.StatusBadRequest, e)
	}else{
		c.JSON(http.StatusOK, msg.Ok)
	}

}
