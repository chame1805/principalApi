package application

import (
	"principalApi/src/domain"
	"fmt"
)

type UpdateReservaUseCase struct {
	db domain.IReservaRepocitory
}

func NewUpdateReservaUseCase(db domain.IReservaRepocitory) *UpdateReservaUseCase {
	return &UpdateReservaUseCase{db: db}
}

func (uc *UpdateReservaUseCase) Execute(id int, name string, fecha string, hora string, numeroPersonas int, service string) error {
	// Llamar al repositorio para actualizar la reserva
	err := uc.db.Update(id, name, fecha, hora, numeroPersonas, service)
	if err != nil {
		return fmt.Errorf("error al actualizar la reserva: %v", err)
	}
	return nil
}
