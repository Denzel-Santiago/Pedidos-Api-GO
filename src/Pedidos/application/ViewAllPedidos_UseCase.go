package application

import (
	"Pedidos-Api/src/Pedidos/domain"
	"Pedidos-Api/src/Pedidos/domain/entities"
)

type ViewAllPedidosUseCase struct {
	db domain.IPedido
}

func NewViewAllPedidosUseCase(db domain.IPedido) *ViewAllPedidosUseCase {
	return &ViewAllPedidosUseCase{
		db: db,
	}
}

func (uc *ViewAllPedidosUseCase) Run() ([]entities.Pedido, error) {
	return uc.db.GetAll()
}
