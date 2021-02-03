package handler

import (
	"errors"
	"github.com/HaHadaxigua/melancholy/ent"
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	logModule "github.com/HaHadaxigua/melancholy/internal/log"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SetupAuthRouters(r gin.IRouter) {
	// open
	r.POST("/login", Login)
	r.POST("/register", Register) // 注册用户
	r.GET("/logout", Logout)
}

// @Summary Login
// @Description 登录接口
// @Tags 基础接口
// @Accept json
// @Produce json
// @Param who query string true "人名"
// @Success 200 {string} string "{"response": "hello Razeen"}"
// @Failure 400 {string} string "{"response": "who are you"}"
// @Router /login [POST]
// Login
func Login(c *gin.Context) {
	req := &response.LoginReq{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.BadRequest)
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, response.UserNameOrPwdIncorrectlyErr)
		return
	}
	type auth struct {
		Email    string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	valid := validation.Validation{}
	a := auth{Email: req.Email, Password: req.Password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	status := response.OK
	if ok {
		userId := service.UserService.CheckUserExist(req.Email, req.Password)
		if userId > -1 {
			token, err := tools.JwtGenerateToken(userId, req.Email, req.Password, 2)
			if err != nil {
				status = response.AuthCheckTokenErr
			} else {
				data["token"] = token
				status = response.OK
				c.Set("user_id", userId)
			}
		} else {
			status = response.UserNameOrPwdIncorrectlyErr
			status.Data = response.UserNameOrPwdIncorrectlyErrorMsg
		}
	} else {
		for _, err := range valid.Errors {
			logrus.Println(err.Key, err.Message)
		}
	}

	if status != response.OK {
		c.JSON(http.StatusBadRequest, status)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": status.Code,
			"response":  status.Message,
			"data": data,
		})
	}
}

// Register
func Register(c *gin.Context) {
	r := &response.UserRequest{}
	if err := c.BindJSON(r); err != nil {
		e := response.BadRequest
		e.Data = err.Error()
		c.JSON(http.StatusBadRequest, e)
		return
	}

	user, err := service.UserService.CreateUser(r)
	if err != nil && errors.Is(err, response.UserHasExistedErr) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Logout
func Logout(c *gin.Context) {
	ah := middleware.AuthHeader{}

	if err := c.ShouldBindHeader(&ah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": response.AuthAccessTokenIllegalErrorMsg,
		})
		return
	}

	userId := c.GetInt("user_id")

	exitReq := &ent.ExitLog{
		Token:  ah.AccessToken,
		UserID: userId,
	}

	err := logModule.ModuleLog.LogService.NewExitLog(exitReq)
	if err != nil {
		e := response.UserExitErr
		c.JSON(http.StatusBadRequest, e)
	} else {
		c.JSON(http.StatusOK, response.Ok)
	}

}
