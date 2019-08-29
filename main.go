package main

import (
	"github.com/gin-gonic/gin"
	"hedu-basic/config"
	"hedu-basic/route"
)

func main() {
	engine := route.NewRouter()
	ginConfiginit(engine)
	// no ssl
	engine.Run(config.Appconf.AppPort)

	// use ssl
	// engine.RunTLS(config.Appconf.AppPort, config.Appconf.SSLPem, config.Appconf.SSLKey)
}

func ginConfiginit(engine *gin.Engine) {
	engine.Static("/", config.Appconf.StaticPath)
	gin.SetMode(config.Appconf.AppMode)
}
