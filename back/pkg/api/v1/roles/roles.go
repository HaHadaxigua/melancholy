package roles

import (
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	service "github.com/HaHadaxigua/melancholy/pkg/service/v1/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetAllRoles all roles
func GetAllRoles(c *gin.Context) {
	roles, err := service.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

//AddRole 添加角色
func AddRole(c *gin.Context) {
	type roleReq struct {
		Name string `json:"name"`
	}
	req := &roleReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		e := msg.InvalidParamsErr
		e.Cause = err.Error()
		c.JSON(http.StatusBadRequest, e)
		return
	}
	err = service.AddRole(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, msg.OK)
}

//AddRoleToUser 给用户添加权限
func AddUserRoles(c *gin.Context){
	uid := c.GetInt("user_id")
	if uid < 0 {
		c.JSON(http.StatusBadRequest, msg.InvalidParamsErr)
		return
	}
	type roleReq struct {
		ID int `json:"id"`
	}
	req := &roleReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		e := msg.InvalidParamsErr
		e.Cause = err.Error()
		c.JSON(http.StatusBadRequest, e)
		return
	}

	err = service.AddUserRoles(uid, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, msg.OK)
}


// GetUserRoles 获取用户当前拥有的角色
func GetUserRoles(c *gin.Context){
	uid := c.GetInt("user_id")
	if uid < 0 {
		c.JSON(http.StatusBadRequest, msg.InvalidParamsErr)
		return
	}
	roles, err := service.GetRolesByUserID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, roles)
}