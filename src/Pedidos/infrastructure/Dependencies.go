package infrastructure

import (
	"Pedidos-Api/src/core"
	"Pedidos-Api/src/pedidos/application"
	"Pedidos-Api/src/pedidos/domain"
)

type Dependencies struct {
	DB                  *core.DB
	PedidoRepository    domain.PedidoRepository
	CreatePedidoUseCase *application.CreatePedidosUseCase
	GetPedidoUseCase    *application.GetPedidosUseCase
	PedidosHandler      *application.PedidosHandler
}

func SetupDependencies() (*Dependencies, error) {
	db, err := core.NewDB()
	if err != nil {
		return nil, err
	}

	pedidoRepo := domain.NewPedidoRepository(db.DB)
	createPedidoUseCase := application.NewCreatePedidosUseCase(pedidoRepo)
	getPedidoUseCase := application.NewGetPedidosUseCase(pedidoRepo)
	pedidosHandler := application.NewPedidosHandler(createPedidoUseCase, getPedidoUseCase)

	return &Dependencies{
		DB:                  db,
		PedidoRepository:    pedidoRepo,
		CreatePedidoUseCase: createPedidoUseCase,
		GetPedidoUseCase:    getPedidoUseCase,
		PedidosHandler:      pedidosHandler,
	}, nil
}
