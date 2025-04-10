package application

import (
	"principalApi/src/domain"
	"fmt"
)

type DeleteReservaUseCase struct {
	db domain.IReservaRepocitory
}

func NewDeleteReservaUseCase(db domain.IReservaRepocitory) *DeleteReservaUseCase {
	return &DeleteReservaUseCase{db: db}
}

func (uc *DeleteReservaUseCase) Execute(id int) error {
	// Llamar al repositorio para eliminar la reserva
	err := uc.db.Delete(id)
	if err != nil {
		return fmt.Errorf("error al eliminar la reserva: %v", err)
	}
	return nil
}
