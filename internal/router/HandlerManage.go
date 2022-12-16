package router

import (
	"github.com/aka-yz/go-micro-core/extension"
	"github.com/aka-yz/go-micro-core/providers/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-micro-demo/internal/api"
	"time"
)

type HandlerManagerImpl struct {
	UserApi api.UserApi `inject:""`
}

func (h *HandlerManagerImpl) HandlerList() []*extension.GinHandlerRegister {
	return []*extension.GinHandlerRegister{
		{
			HttpMethod:   constants.GinMethodGet,
			RelativePath: "/user/list",
			Handlers: []gin.HandlerFunc{
				h.UserApi.GetUserList,
			},
		},
		// 如果有后续其他的，继续如上追加在这儿就好了
	}
}

func (h *HandlerManagerImpl) MiddlewareList() []gin.HandlerFunc {
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     constants.AllowedOrigins,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     constants.AllowedHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "...."
		// },
		MaxAge: 12 * time.Hour,
	})
	return []gin.HandlerFunc{gin.Logger(), gin.Recovery(), corsConfig}
}

func (h *HandlerManagerImpl) InjectName() string {
	return constants.HandlerInjectName
}
