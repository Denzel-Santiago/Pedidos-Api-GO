package application

import (
	"Pedidos-Api/src/pedidos/domain"
	"Pedidos-Api/src/pedidos/domain/entities"
)

type GetPedidosUseCase struct {
	repo domain.PedidoRepository
}

func NewGetPedidosUseCase(repo domain.PedidoRepository) *GetPedidosUseCase {
	return &GetPedidosUseCase{repo: repo}
}

func (uc *GetPedidosUseCase) Execute(pedidoID int) (*entities.Pedido, error) {
	return uc.repo.FindByID(pedidoID)
}
