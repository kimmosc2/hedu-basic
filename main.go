package main

import (
	"hedu-basic/route"
)

func main() {
	engine := route.NewRouter()
	engine.Run(":1234")
}
