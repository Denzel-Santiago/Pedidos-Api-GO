package application

import (
	"Pedidos-Api/src/Pedidos/domain"
)

type DeletePedidoUseCase struct {
	db domain.IPedido
}

func NewDeletePedidoUseCase(db domain.IPedido) *DeletePedidoUseCase {
	return &DeletePedidoUseCase{
		db: db,
	}
}

func (uc *DeletePedidoUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
