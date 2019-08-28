package main

import (
	"github.com/gin-gonic/gin"
	"hedu-basic/config"
	"hedu-basic/route"
)



func main() {
	engine := route.NewRouter()
	ginConfiginit(engine)
	engine.Run(config.Appconf.AppPort)
}

func ginConfiginit(engine *gin.Engine)  {
	engine.Static("/",config.Appconf.StaticPath)
	gin.SetMode(config.Appconf.AppMode)
}