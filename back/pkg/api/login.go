package api

import (
	//model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1/user"
	store "github.com/HaHadaxigua/melancholy/pkg/store/user"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

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
		userId := store.CheckUserExist(email, password)
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
			status.Cause = msg.UserNameOrPwdIncorrectlyErrorMsg
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
		e.Cause = err.Error()
		c.JSON(http.StatusBadRequest, e)
		return
	}
	user, err := service.CreateUser(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
