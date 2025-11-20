package controller

import (
	application "menu/src/Endpoint/Canciones/Application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCancionController struct {
	controller *application.GetAllCancionesUseCase
}

func NewGetAllController(app *application.GetAllCancionesUseCase)*GetCancionController{
	return &GetCancionController{
		controller: app,
	}
}

func (c *GetCancionController) GetAllCanciones(ctx *gin.Context) {
	canciones, err := c.controller.Execute()

	if len(canciones) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No hay canciones disponibles",
			"data":    []interface{}{},
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, canciones)
}