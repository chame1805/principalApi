package infraestructure

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}


func NewRabbitMQService() *RabbitMQService {
	conn, err := amqp.Dial("amqp://chame:chame0104@52.206.28.216:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	q, err := ch.QueueDeclare(
		"ColaDeChame", // Nombre de la cola
		true,          // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	return &RabbitMQService{
		conn:    conn,
		channel: ch,
		queue:   q.Name,
	}
}


func (r *RabbitMQService) Publish(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.channel.PublishWithContext(ctx,
		"",        // Exchange
		r.queue,   // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message),
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", message)
	return nil
}

// Close cierra la conexión y el canal
func (r *RabbitMQService) Close() {
	r.channel.Close()
	r.conn.Close()
}
