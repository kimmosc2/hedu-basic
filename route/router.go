package route

import (
	"github.com/gin-gonic/gin"
	"hedu-basic/controller"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()

	official := engine.Group("/api/v1/official")
	{
		// change password service
		official.POST("/changePassword",controller.ChangePassWord)

	}
	return engine
}
