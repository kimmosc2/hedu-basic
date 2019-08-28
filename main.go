package main

import (
	"hedu-basic/route"
)

func main() {
	engine := route.NewRouter()
	engine.Static("/","./static/")
	engine.Run(":1234")
}
