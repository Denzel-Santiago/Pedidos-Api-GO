package application

import (
	"Pedidos-Api/src/Pedidos/domain"
	"Pedidos-Api/src/Pedidos/domain/entities"
)

type CreatePedidoUseCase struct {
	db domain.IPedido
}

func NewCreatePedidoUseCase(db domain.IPedido) *CreatePedidoUseCase {
	return &CreatePedidoUseCase{
		db: db,
	}
}

func (uc *CreatePedidoUseCase) Run(pedido *entities.Pedido) (*entities.Pedido, error) {
	err := uc.db.Save(*pedido)
	return pedido, err
}
