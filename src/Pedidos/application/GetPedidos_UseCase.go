package application

import (
	"Pedidos-Api/src/Pedidos/domain"
	"Pedidos-Api/src/Pedidos/domain/entities"
)

type GetPedidoUseCase struct {
	db domain.IPedido
}

func NewGetPedidoUseCase(db domain.IPedido) *GetPedidoUseCase {
	return &GetPedidoUseCase{
		db: db,
	}
}

func (uc *GetPedidoUseCase) Run(id int) (entities.Pedido, error) {
	return uc.db.FindByID(id)
}
