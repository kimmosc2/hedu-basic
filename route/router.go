package route

import (
	"github.com/gin-gonic/gin"
	"hedu-basic/controller"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()

	// 中间件
	{
		// SSL 转发中间件
		// engine.Use(middleware.TlsHandler())

	}

	official := engine.Group("/api/v1/official")
	{
		// change password service
		official.POST("/changePassword",controller.ChangePassWord)

	}
	return engine
}
