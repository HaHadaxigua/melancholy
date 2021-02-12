package handler

import (
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
	role.POST("/role", createRole)
	role.DELETE("role/:id", deleteRole)
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
	c.JSON(http.StatusOK, response.OK)
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
	c.JSON(http.StatusOK, response.OK)
}
