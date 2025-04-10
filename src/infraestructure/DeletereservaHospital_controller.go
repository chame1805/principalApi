package infraestructure

import (
	"principalApi/src/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteReservHospitalController struct {
	useCase *application.DeleteReservaUseCase
}

func NewDeleteReservHospitalController(useCase *application.DeleteReservaUseCase) *DeleteReservHospitalController {
	return &DeleteReservHospitalController{useCase: useCase}
}

func (rh *DeleteReservHospitalController) Execute(c *gin.Context) {
	
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Ejecutar el caso de uso para eliminar la reserva
	err = rh.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reserva eliminada correctamente"})
}
