package infraestructure

import (
	"principalApi/src/application"
	"principalApi/src/infraestructure/service"

	"github.com/gin-gonic/gin"
)

func SetReservaHospitalRoutes(r *gin.Engine) {
	rs := NewMySQLReserva() // Instancia del repositorio
	rabbitMQ := infraestructure.NewRabbitMQService() // Instancia de RabbitMQ
	sendMessageUseCase := application.NewSendMessageUseCase(rabbitMQ)

	// Pasamos también rabbitMQ al caso de uso
	CreateReservHospitalController := NewReservationHospitalController(application.NewReservaUseCase(rs, sendMessageUseCase))
	GetReservHospitalController := NewGetReservHospitalController(application.NewGetReservaUseCase(rs))
	UpdateReservHospitalController := NewUpdateReservHospitalController(application.NewUpdateReservaUseCase(rs))
	DeleteReservHospitalController := NewDeleteReservHospitalController(application.NewDeleteReservaUseCase(rs))	

	r.POST("/reservas", CreateReservHospitalController.Execute)
	r.GET("/reservas", GetReservHospitalController.Execute)
	r.PUT("/reservas/:id", UpdateReservHospitalController.Execute)
	r.DELETE("/reservas/:id", DeleteReservHospitalController.Execute)
}
