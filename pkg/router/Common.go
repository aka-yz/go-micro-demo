package router

import (
	"github.com/gin-gonic/gin"
)

// SetupCommonApiRouter 启动公共 Api 路由
func SetupCommonApiRouter(Router *gin.RouterGroup) {
	// 登录
	Router.POST("/user/login", api2.UserLogin)
	// 登出
	Router.POST("/user/logout", api2.UserLogout)
	// 个人信息详情
	Router.POST("/user/detail", api2.GetUserDetail)
	// 活动列表
	Router.POST("/activity/list", api2.GetActivityList)
	// 活动详情
	Router.POST("/activity/detail", api2.GetActivityDetail)
	// 参加活动
	Router.POST("/activity/join", api2.JoinActivity)
	// 退出活动
	Router.POST("/activity/exit", api2.ExitActivity)
	// 新增评论
	Router.POST("/comment/add", api2.AddComment)
}
