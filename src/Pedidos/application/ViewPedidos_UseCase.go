package application

import (
	"Pedidos-Api/src/Pedidos/domain"
	"Pedidos-Api/src/Pedidos/domain/entities"
)

type ViewPedidoUseCase struct {
	db domain.IPedido
}

func NewViewPedidoUseCase(db domain.IPedido) *ViewPedidoUseCase {
	return &ViewPedidoUseCase{
		db: db,
	}
}

func (uc *ViewPedidoUseCase) Run(id int) (entities.Pedido, error) {
	return uc.db.FindByID(id)
}
