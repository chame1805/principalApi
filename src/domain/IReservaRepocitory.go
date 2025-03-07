package domain

type IReservaRepocitory interface {
	Save(name string, fecha string, hora string, numeroPersonas int, service string)error
	GetReserva(id int) (map[string]interface{}, error) // Ahora recibe un ID
}
