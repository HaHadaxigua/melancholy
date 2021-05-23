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
)

func SetupBasicRouters(r gin.IRouter) {
	secured := r.Group("/basic", middleware.Auth)

	role := secured.Group("/r")
	role.POST("/role/list", listRoles)
	role.POST("/role/create", createRole)
	role.POST("/role/delete", deleteRole)
	role.POST("/appendPerm", appendPerm)
	role.POST("/removePerm", removePerm)

	user := secured.Group("/u")
	user.POST("/user", listUsers)
	user.POST("/appendRole", appendRole)
	user.POST("/removeRole", removeRole)

	perms := secured.Group("/p")
	perms.POST("/perm/list", listPerms)
	perms.POST("/perm/create", createPerm)
	perms.POST("/perm/delete", deletePerm)

}

func createRole(c *gin.Context) {
	var req msg.ReqRoleCreate
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.Role.NewRole(&req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func deleteRole(c *gin.Context) {
	var req msg.ReqRoleDelete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if req.RoleID <= 0 {
		c.JSON(http.StatusBadRequest, response.NewErr(nil))
		return
	}

	if err := service.Role.DeleteRole(req.RoleID); err != nil {

		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func listRoles(c *gin.Context) {
	var req msg.ReqRoleFilter
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if rsp, err := service.Role.ListRoles(&req, true); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(rsp))
		return
	}
}

func listUsers(c *gin.Context) {
	var req msg.ReqUserFilter
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if rsp, err := service.User.ListUsers(&req, true); err != nil {
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

// deletePerm 删除权限
func deletePerm(c *gin.Context) {
	var req msg.ReqPermissionDelete
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	if err := service.Permission.DeletePermission(req.PermissionID); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
}

// appendPerm 给角色添加权限
func appendPerm(c *gin.Context) {
	var req msg.ReqRolePermAssociation
	if err := c.BindJSON(&req); err != nil {
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

// removePerm 给角色移除权限
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
