package infraestructure

import (
	"principalApi/src/application"
	"principalApi/src/infraestructure/service"

	"github.com/gin-gonic/gin"
)

func SetReservaHospitalRoutes(r *gin.Engine) {
	rs := NewMySQLReserva() // Instancia del repositorio
	rabbitMQ := infraestructure.NewRabbitMQService() // Instancia de RabbitMQ

	// Pasamos también rabbitMQ al caso de uso
	CreateReservHospitalController := NewReservationHospitalController(application.NewReservaUseCase(rs, rabbitMQ))
	GetReservHospitalController := NewGetReservHospitalController(application.NewGetReservaUseCase(rs))

	r.POST("/reservas", CreateReservHospitalController.Execute)
	r.GET("/reservas/:id", GetReservHospitalController.Execute)
}
