package api

import (
	model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1/user"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Login
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	type auth struct {
		Username string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	status := msg.OK
	if ok {
		isExist := CheckAuth(username, password)
		if isExist {
			token, err := tools.GenerateToken(username, password)
			if err != nil {
				status = msg.AuthCheckTokenErr
			} else {
				data["token"] = token
				status = msg.OK
			}
		} else {
			status = msg.UserNameOrPwdIncorrectlyErr
			status.Cause = "用户名或密码错误"
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

//CheckAuth 判断用户是否存在
func CheckAuth(username, password string) bool {
	var auth model.User
	db := store.GetConn()
	result := db.Model(model.User{}).Where("username= ?", username).First(&auth)
	if result.Error != nil {
		return false
	}
	if result.RowsAffected >= 1 {
		flag := tools.VerifyPassword(auth.Password, password+auth.Salt)
		return flag
	}

	return false
}
