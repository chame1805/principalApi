package application

import (
	"fmt"
	domain "principalApi/src/domain/rabbit"
)


type SendMessageUseCase struct {
	rabbitMQSvc domain.MessageBroker
}

func NewSendMessageUseCase(rabbitMQSvc domain.MessageBroker) *SendMessageUseCase {
	return &SendMessageUseCase{rabbitMQSvc: rabbitMQSvc}
}

func (uc *SendMessageUseCase)Execute(message string)  {
	err := uc.rabbitMQSvc.Publish(message)
	if err != nil {
		fmt.Println("Error al enviar mensaje a RabbitMQ:", err)
	}
	}