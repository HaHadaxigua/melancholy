package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupBasicRouters(r gin.IRouter) {
	secured := r.Group("/basic", middleware.JWT)

	role := secured.Group("/r")
	role.GET("/role", listRoles)
	role.POST("/role", createRole)
	role.DELETE("role/:id", deleteRole)

	user := secured.Group("/u")
	user.GET("/user", listUsers)
	user.PATCH("/appendRole", appendRole)
	user.PATCH("/removeRole", removeRole)
}

func createRole(c *gin.Context) {
	req := &msg.ReqRoleCreate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.Role.NewRole(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func deleteRole(c *gin.Context) {
	_rid := c.Param("id")
	rid, err := strconv.Atoi(_rid)
	if err != nil || rid <= 0 {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}

	if err := service.Role.DeleteRole(rid); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func listRoles(c *gin.Context) {
	req := &msg.ReqRoleFilter{}
	if err := c.BindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if rsp, err := service.Role.ListRoles(req, false); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(rsp))
		return
	}
}

func listUsers(c *gin.Context) {
	req := &msg.ReqUserFilter{}
	if err := c.BindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if rsp, err := service.User.ListUsers(req, true); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(rsp))
		return
	}
}

func appendRole(c *gin.Context) {
	req := &msg.ReqUserRoleAssociation{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.User.RoleManager(req.UserID, req.RoleID, consts.AppendRole); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func removeRole(c *gin.Context) {
	req := &msg.ReqUserRoleAssociation{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.User.RoleManager(req.UserID, req.RoleID, consts.RemoveRole); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}
