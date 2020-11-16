package api

import (
	model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store"
	"github.com/HaHadaxigua/melancholy/pkg/tools"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetAuth(c *gin.Context) {
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
	status := msg.InvalidParamsErr
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
			status = msg.AuthCheckTokenErr
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": status.Code,
		"msg":  status.Message,
		"data": data,
	})
}

// todo: 判断用户名和密码时需要进行重新操作
func CheckAuth(username, password string) bool {
	type User struct {
		ID       int
		Username string
		Password string
	}

	var auth User
	db := store.GetConn()
	db.Model(model.User{}).Select("id").Where(User{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
