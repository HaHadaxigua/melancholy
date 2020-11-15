package user

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Register
func Register(c *gin.Context) {
	r := &msg.UserRequest{}
	c.BindJSON(r)
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
