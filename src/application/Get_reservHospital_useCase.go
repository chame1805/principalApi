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

func (uc *GetReservaUseCase) Execute(id int) (map[string]interface{}, error) {
	return uc.db.GetReserva(id) // Pasar el ID a la función GetReserva
}
