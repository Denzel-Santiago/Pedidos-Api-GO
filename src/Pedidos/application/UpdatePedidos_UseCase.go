package application

import (
	"Pedidos-Api/src/Pedidos/domain"
	"Pedidos-Api/src/Pedidos/domain/entities"
)

type UpdatePedidoUseCase struct {
	db domain.IPedido
}

func NewUpdatePedidoUseCase(db domain.IPedido) *UpdatePedidoUseCase {
	return &UpdatePedidoUseCase{
		db: db,
	}
}

func (uc *UpdatePedidoUseCase) Run(id int, pedido entities.Pedido) error {
	return uc.db.Update(id, pedido)
}
