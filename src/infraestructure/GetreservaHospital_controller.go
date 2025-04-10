package infraestructure

import (
	"principalApi/src/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetReservHospitalController struct {
	useCase *application.GetReservaUseCase
}

func NewGetReservHospitalController(useCase *application.GetReservaUseCase) *GetReservHospitalController {
	return &GetReservHospitalController{useCase: useCase}
}

func (rh *GetReservHospitalController) Execute(c *gin.Context) {
	// Ejecutar el caso de uso para obtener todas las reservas
	reservas, err := rh.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"reservas": reservas})
}
