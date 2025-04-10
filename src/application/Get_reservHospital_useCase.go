package application

import (
	"principalApi/src/domain"
)

type GetReservaUseCase struct {
	db domain.IReservaRepocitory	
}

func NewGetReservaUseCase(db domain.IReservaRepocitory) *GetReservaUseCase {
	return &GetReservaUseCase{db: db}
}

// Modificamos el Execute para obtener todas las reservas.
func (uc *GetReservaUseCase) Execute() ([]map[string]interface{}, error) {
	return uc.db.GetAllReservas() // Llamamos a GetAllReservas en vez de GetReserva
}
