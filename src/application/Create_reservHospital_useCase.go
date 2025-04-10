package application

import (
	"fmt"
	"principalApi/src/domain"
)

// Estructura del caso de uso
type Create_reservHospital_useCase struct {
	db                domain.IReservaRepocitory
	sendMessageUseCase *SendMessageUseCase
}

// Constructor que recibe la base de datos y el caso de uso para enviar mensajes
func NewReservaUseCase(db domain.IReservaRepocitory, sendMessageUseCase *SendMessageUseCase) *Create_reservHospital_useCase {
	return &Create_reservHospital_useCase{
		db:                db,
		sendMessageUseCase: sendMessageUseCase, // Se usa correctamente el parámetro recibido
	}
}

// Método para ejecutar la creación de la reserva
func (uc *Create_reservHospital_useCase) Execute(name string, fecha string, hora string, numeroPersonas int, service string) {
	err := uc.db.Save(name, fecha, hora, numeroPersonas, service)
	if err != nil {
		fmt.Println("Error al guardar la reserva:", err)
		return
	}

	// Enviar mensaje
	message := fmt.Sprintf("Reserva creada: Nombre=%s, Fecha=%s, Hora=%s, Personas=%d, Servicio=%s", name, fecha, hora, numeroPersonas, service)
	uc.sendMessageUseCase.Execute(message) // Se usa el campo correcto
}
