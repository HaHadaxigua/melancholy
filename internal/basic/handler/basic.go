package handler

import (
	"errors"
	"github.com/HaHadaxigua/melancholy/internal/basic/consts"
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/basic/msg"
	"github.com/HaHadaxigua/melancholy/internal/basic/service"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func SetupBasicRouters(r gin.IRouter) {
	secured := r.Group("/basic", middleware.JWT)

	role := secured.Group("/r")
	role.GET("/role", listRoles)
	role.POST("/role", createRole)
	role.DELETE("role/:id", deleteRole)
	role.PATCH("/appendPerm", appendPerm)
	role.PATCH("/removePerm", removePerm)

	user := secured.Group("/u")
	user.GET("/user", listUsers)
	user.PATCH("/appendRole", appendRole)
	user.PATCH("/removeRole", removeRole)

	perms := secured.Group("/p")
	perms.GET("/perm", listPerms)
	perms.POST("/perm", createPerm)

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
	if rsp, err := service.Role.ListRoles(req, true); err != nil {
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, response.NewErr(err))
			return
		}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, response.NewErr(err))
			return
		}
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func listPerms(c *gin.Context) {
	req := &msg.ReqPermissionFilter{}
	if err := c.BindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if rsp, err := service.Permission.ListPermission(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(rsp))
	}
}

func createPerm(c *gin.Context) {
	req := &msg.ReqPermissionCreate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.Permission.NewPermission(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func appendPerm(c *gin.Context) {
	req := &msg.ReqRolePermAssociation{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.Role.PermissionManager(req.RoleID, req.PermissionID, consts.AppendPermission); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, response.NewErr(err))
			return
		}
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func removePerm(c *gin.Context) {
	req := &msg.ReqRolePermAssociation{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.Role.PermissionManager(req.RoleID, req.PermissionID, consts.RemovePermission); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, response.NewErr(err))
			return
		}
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}
