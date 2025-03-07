package application

import (
	"fmt"
	"principalApi/src/domain"
	infraestructure "principalApi/src/infraestructure/service"
	//"principalApi/src/infraestructure"
	// "principalApi/src/infraestructure" // Corrección aquí
)

type Create_reservHospital_useCase struct {
	db          domain.IReservaRepocitory
	rabbitMQSvc *infraestructure.RabbitMQService // Corrección aquí
}

// Constructor que recibe la base de datos y el servicio de mensajería
func NewReservaUseCase(db domain.IReservaRepocitory, rabbitMQSvc *infraestructure.RabbitMQService) *Create_reservHospital_useCase { // Corrección aquí
	return &Create_reservHospital_useCase{db: db, rabbitMQSvc: rabbitMQSvc}
}

func (uc *Create_reservHospital_useCase) Execute(name string, fecha string, hora string, numeroPersonas int, service string) {
	err := uc.db.Save(name, fecha, hora, numeroPersonas, service)
	if err != nil {
		fmt.Println("Error al guardar la reserva:", err)
		return
	}

	// Enviar mensaje a RabbitMQ
	message := fmt.Sprintf("Reserva creada: Nombre=%s, Fecha=%s, Hora=%s, Personas=%d, Servicio=%s", name, fecha, hora, numeroPersonas, service)
	err = uc.rabbitMQSvc.SendMessage(message)
	if err != nil {
		fmt.Println("Error al enviar mensaje a RabbitMQ:", err)
	}
}
