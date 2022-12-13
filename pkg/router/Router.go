package router

import (
	filter2 "entry-task/pkg/filter"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// SetupRouter 启动路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//gin.SetMode(gin.ReleaseMode)

	pprof.Register(r)

	// 初始化拦截器
	r.Use(filter2.CorsHandler())

	commonApiGroup := r.Group("/")
	userApiGroup := r.Group("/user/manage")
	activityApiGroup := r.Group("/activity/manage")
	classificationApiGroup := r.Group("/classification/manage")
	// 造数据
	createData := r.Group("/create/data")

	// user
	SetupUserApiRouter(userApiGroup)
	// common
	SetupCommonApiRouter(commonApiGroup)
	// activity
	SetupActivityApiRouter(activityApiGroup)
	// 造数据
	SetupCreateDataRouter(createData)
	// classification
	SetupClassificationApiRouter(classificationApiGroup)

	return r
}
