package routers

import(
	"github.com/gin-gonic/gin"
	controllers "menu/src/Endpoint/Canciones/Infrestructure/Controller"
)

func RouterCancion(g *gin.Engine,
	GetCancionController *controllers.GetCancionController,){
		router := g.Group("/api/v1/cancion")

		router.GET("/getAll", GetCancionController.GetAllCanciones)
}