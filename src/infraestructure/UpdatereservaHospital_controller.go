package infraestructure

import (
	"principalApi/src/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateReservHospitalController struct {
	useCase *application.UpdateReservaUseCase
}

func NewUpdateReservHospitalController(useCase *application.UpdateReservaUseCase) *UpdateReservHospitalController {
	return &UpdateReservHospitalController{useCase: useCase}
}

func (rh *UpdateReservHospitalController) Execute(c *gin.Context) {
	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Crear un objeto para almacenar los datos de la reserva
	var reserva struct {
		Name           string `json:"name"`
		Fecha          string `json:"fecha"`
		Hora           string `json:"hora"`
		NumeroPersonas int    `json:"numeroPersonas"`
		Service        string `json:"service"`
	}

	// Obtener los datos de la solicitud
	if err := c.ShouldBindJSON(&reserva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al caso de uso para actualizar la reserva
	err = rh.useCase.Execute(id, reserva.Name, reserva.Fecha, reserva.Hora, reserva.NumeroPersonas, reserva.Service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reserva actualizada correctamente"})
}
