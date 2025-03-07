package infraestructure

import (
	"principalApi/src/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetReservHospitalController struct {
	useCase *application.GetReservaUseCase
}

func NewGetReservHospitalController(useCase *application.GetReservaUseCase) *GetReservHospitalController {
	return &GetReservHospitalController{useCase: useCase}
}

func (rh *GetReservHospitalController) Execute(c *gin.Context) {
	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id") // Suponiendo que el ID se pasa en la URL como /reservas/:id
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Ejecutar el caso de uso con el ID
	reservas, err := rh.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con los datos obtenidos
	c.JSON(http.StatusOK, gin.H{"reservas": reservas})
}
