package infraestructure

import (
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
   "principalApi/src/application"
)

type CreateReservHospitalController struct {
    useCase *application.Create_reservHospital_useCase
}

func NewReservationHospitalController(useCase *application.Create_reservHospital_useCase) *CreateReservHospitalController {
    return &CreateReservHospitalController{useCase: useCase}
}

func (rh *CreateReservHospitalController) Execute(c *gin.Context) {
    var reserva struct {
        Name           string `json:"name"`
        Fecha          string `json:"fecha"`
        Hora           string `json:"hora"`
        NumeroPersonas int    `json:"numeroPersonas"`
        Service        string `json:"service"`
    }

    if err := c.ShouldBindJSON(&reserva); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    if reserva.NumeroPersonas <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "El número de personas debe ser mayor que cero"})
        return
    }

    
    if _, err := time.Parse("2006-01-02", reserva.Fecha); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha inválida, el formato debe ser YYYY-MM-DD"})
        return
    }

    // Llamar al caso de uso para procesar la reserva
    rh.useCase.Execute(reserva.Name, reserva.Fecha, reserva.Hora, reserva.NumeroPersonas, reserva.Service)

    c.JSON(http.StatusOK, gin.H{"message": "Reserva creada correctamente"})
}