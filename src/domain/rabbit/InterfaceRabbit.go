package domain

type MessageBroker interface {
	Publish(message string) error
}
