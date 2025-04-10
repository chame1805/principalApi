package domain

type IReservaRepocitory interface {
	Save(name string, fecha string, hora string, numeroPersonas int, service string) error
	GetAllReservas() ([]map[string]interface{}, error)
	Delete(id int) error           // Método para eliminar la reserva
	Update(id int, name string, fecha string, hora string, numeroPersonas int, service string) error // Método para actualizar
}
