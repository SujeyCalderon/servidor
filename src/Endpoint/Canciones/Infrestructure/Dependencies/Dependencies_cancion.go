package dependencies

import(
	application "menu/src/Endpoint/Canciones/Application"
	sql "menu/src/Endpoint/Canciones/Infrestructure/Sql"
	controllers "menu/src/Endpoint/Canciones/Infrestructure/Controller"
	routers "menu/src/Endpoint/Canciones/Infrestructure/Routers"
	"github.com/gin-gonic/gin"
)

func InitCancion(e *gin.Engine){
    ps, err := sql.NewMySQL()
    if err != nil {
        panic(err)
    }

    getAllCancionUseCase := application.NewGetAllCanciones(ps)
    getAllCancionController := controllers.NewGetAllController(getAllCancionUseCase)

    routers.RouterCancion(e, getAllCancionController)
}