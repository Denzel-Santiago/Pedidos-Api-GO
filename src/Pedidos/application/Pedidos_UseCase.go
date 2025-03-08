package application

import (
	"Pedidos-Api/src/pedidos/domain"
	"Pedidos-Api/src/pedidos/domain/entities"
)

type CreatePedidosUseCase struct {
	repo domain.PedidoRepository
}

func NewCreatePedidosUseCase(repo domain.PedidoRepository) *CreatePedidosUseCase {
	return &CreatePedidosUseCase{repo: repo}
}

func (uc *CreatePedidosUseCase) Execute(pedido entities.Pedido) error {
	return uc.repo.Save(pedido)
}
