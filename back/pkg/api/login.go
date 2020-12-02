package api

import (
	"errors"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/pkg/middleware"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1/admin"
	serviceu "github.com/HaHadaxigua/melancholy/pkg/service/v1/user"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	storeu "github.com/HaHadaxigua/melancholy/pkg/store/user"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// @Summary Login
// @Description 登录接口
// @Tags 基础接口
// @Accept json
// @Produce json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /login [get]
//Login
func Login(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "badRequest",
		})
		return
	}
	type auth struct {
		Email    string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	valid := validation.Validation{}
	a := auth{Email: email, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	status := msg.OK
	if ok {
		userId := storeu.CheckUserExist(email, password)
		if userId > -1 {
			token, err := tools.GenerateToken(email, password)
			if err != nil {
				status = msg.AuthCheckTokenErr
			} else {
				data["token"] = token
				status = msg.OK
				c.Set("user_id", userId)
			}
		} else {
			status = msg.UserNameOrPwdIncorrectlyErr
			status.Data = msg.UserNameOrPwdIncorrectlyErrorMsg
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	if status != msg.OK {
		c.JSON(http.StatusBadRequest, status)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": status.Code,
			"msg":  status.Message,
			"data": data,
		})
	}
}

//Register
func Register(c *gin.Context) {
	r := &msg.UserRequest{}
	if err := c.BindJSON(r); err != nil {
		e := msg.BadRequest
		e.Data = err.Error()
		c.JSON(http.StatusBadRequest, e)
		return
	}

	user, err := serviceu.CreateUser(r)
	if err != nil && errors.Is(err, msg.UserHasExistedErr){
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	// 赋予角色
	err = service.AddUserRoles(user.ID, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}


// Logout 退出登录
func Logout(c *gin.Context) {
	ah := middleware.AuthHeader{}

	if err := c.ShouldBindHeader(&ah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": msg.AuthAccessTokenIllegalErrorMsg,
		})
		return
	}

	userId := c.GetInt("user_id")
	// 写退出表

	exitReq := &ent.ExitLog{
		Token: ah.AccessToken,
		UserID: userId,
	}

	err := store.SaveExitLog(exitReq)
	if err != nil {
		e := msg.UserExitErr
		c.JSON(http.StatusBadRequest, e)
	}else{
		c.JSON(http.StatusOK, msg.Ok)
	}

}
