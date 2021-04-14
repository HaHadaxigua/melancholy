/******
** @date : 2/11/2021 2:51 PM
** @author : zrx
** @description:
******/
package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouters(r gin.IRouter) {
	auth := r.Group("/auth")
	// open
	auth.POST("/login", login)
	auth.POST("/register", register) // 注册用户
}

func login(c *gin.Context) {
	req := &msg.ReqLogin{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.BadRequest)
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, response.NewErr(msg.ErrUserNameOrPwdWrong))
		return
	}
	type auth struct {
		Email    string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	valid := validation.Validation{}
	a := auth{Email: req.Email, Password: req.Password}
	ok, err := valid.Valid(&a)
	if ok {
		_user, err := service.User.GetUserByEmail(req.Email)
		if err != nil || _user == nil {
			c.JSON(http.StatusInternalServerError, response.NewErr(err))
			return
		}
		if _user.ID > 0 && _user.Status != consts.UserStatusBlocked {
			if !tools.VerifyPassword(_user.Password, req.Password, _user.Salt) {
				c.JSON(http.StatusOK, msg.ErrUserNameOrPwdWrong)
				return
			}
			token, err := tools.JwtGenerateToken(_user.ID, req.Email, req.Password, 2)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.NewErr(err))
				return
			}
			// 构造登陆的返回体
			rsp := &msg.RspLogin{
				User: &msg.UserInfo{
					ID:        _user.ID,
					Username:  _user.Username,
					Mobile:    _user.Mobile,
					Email:     _user.Email,
					Status:    _user.Status,
					Avatar:    _user.Avatar,
					CreatedAt: _user.CreatedAt,
					UpdatedAt: _user.UpdatedAt,
				},
				Token: token,
			}
			c.JSON(http.StatusOK, response.Ok(rsp))
			return
		} else {
			c.JSON(http.StatusOK, msg.ErrUserNotFound)
			return
		}
	}
	c.JSON(http.StatusBadRequest, err)
	return
}

// register
func register(c *gin.Context) {
	var r msg.ReqRegister

	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}

	user, err := service.User.CreateUser(&r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}

	c.JSON(http.StatusOK, response.Ok(user.ID))
}
