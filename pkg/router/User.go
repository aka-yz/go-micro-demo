package router

import (
	"entry-task/pkg/api"
	"github.com/gin-gonic/gin"
)

// SetupUserApiRouter 启动用户信息路由
func SetupUserApiRouter(Router *gin.RouterGroup) {
	// 用户列表
	Router.POST("/list", api.GetUserList)
	// 编辑用户信息
	Router.POST("/edit", api.EditUserInfo)
	// 删除用户
	Router.POST("/del", api.DelUserInfo)
	// 新增用户
	Router.POST("/add", api.AddUserInfo)
}
