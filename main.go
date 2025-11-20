package main

import (
	"menu/src/Middleware"

	"github.com/gin-gonic/gin"
	dependencies "menu/src/Endpoint/Canciones/Infrestructure/Dependencies"
)

func main(){
	sos := gin.Default()
	sos.Use(Middleware.Cors())
	dependencies.InitCancion(sos)
	sos.Run(":8080")
}